# Musgo

Musgo is a Go code generator for binary MUS format with validation support.
Generated code converts data to and from MUS format. More info about it and
about the format you can find at "https://github.com/ymz-ncnk/musgen".

# How to use

First, you should download and install Go, version 1.4 or later.

Create in your home directory `foo` folder with the following structure:

```
foo/
 |‒‒‒make/
 |    |‒‒‒musable.go
 |‒‒‒validators/
 |    |‒‒‒validators.go
 |‒‒‒foo.go
```

__foo.go__
```go
//go:generate go run make/musable.go
package foo

type Foo struct {
	boo int `mus:"validators.BiggerThanTen"` // private fields are supported
	// too, while unmarshalling will be checked with BiggerThanTen validator
	zoo []int `mus:",,validators.BiggerThanTen"` // every element will be checked
	// with BiggerThanTen validator
	Bar MyString // alias types are supported too
	Car bool     `mus:"-"` // this field will be skiped
}

type MyString string
```

__validators/validators.go__
```go
package validators

import "errors"

var ErrBiggerThanTen error = errors.New("bigger then ten")

func BiggerThanTen(n int) error {
  if n > 10 {
    return ErrBiggerThanTen
  }
  return nil
}
```

__make/musable.go__
```go
// +build ignore

package main

import (
	"foo"
	"reflect"

	"github.com/ymz-ncnk/musgo"
)

func main() {
	musGo, err := musgo.New()
	if err != nil {
		panic(err)
	}
	// You should "Generate" for all involved custom types.
	unsafe := false // to generate safe code
	var myStr foo.MyString
	// Alias types don't support tags, so to set up validators we use
	// GenerateAlias method.
	maxLength := 5 // restricts length of MyString values to 5 characters
	err = musGo.GenerateAlias(reflect.TypeOf(myStr), unsafe, "", maxLength, "",
		"")
	if err != nil {
		panic(err)
	}
	// reflect.Type could be created without any variable.
	err = musGo.Generate(reflect.TypeOf((*foo.Foo)(nil)).Elem(), unsafe)
	if err != nil {
		panic(err)
	}
}
```

Run from the comamnd line:
```bash
$ cd ~/foo
$ go mod init foo
$ go get github.com/ymz-ncnk/musgo
$ go generate
```

Now you can see `Foo.musgen.go` and `MyString.musgen.go` files in the `foo` 
folder. Pay attention to the location of the generated files. The data type and 
the code generated for it must be in the same package.
Let's write some tests. Create `foo_test.go` file.

```
foo/
 |‒‒‒...
 |‒‒‒foo_test.go
```

__foo_test.go__
```go
package foo

import (
	"foo/validators"
	"reflect"
	"testing"

	"github.com/ymz-ncnk/musgo/errs"
)

func TestFooSerialization(t *testing.T) {
	foo := Foo{
		zoo: []int{4, 2},
		boo: 5,
		Bar: MyString("hello"),
		Car: true,
	}
	buf := make([]byte, foo.SizeMUS())
	foo.MarshalMUS(buf)

	afoo := Foo{}
	_, err := afoo.UnmarshalMUS(buf)
	if err != nil {
		t.Error(err)
	}
	foo.Car = false
	if !reflect.DeepEqual(foo, afoo) {
		t.Error("something went wrong")
	}
}

func TestFooValidation(t *testing.T) {
	// test simple validator
	{
		foo := Foo{
			boo: 11,
			zoo: []int{1, 2},
			Bar: "hello",
		}
		buf := make([]byte, foo.SizeMUS())
		foo.MarshalMUS(buf)

		afoo := Foo{}
		_, err := afoo.UnmarshalMUS(buf)
		if err == nil {
			t.Error("validation doesn't work")
		}
		fieldErr, ok := err.(errs.FieldError)
		if !ok {
			t.Error("wrong field error")
		}
		if fieldErr.FieldName() != "boo" {
			t.Error("wrong field error fieldName")
		}
		if fieldErr.Cause() != validators.ErrBiggerThanTen {
			t.Error("wrong error")
		}
	}
	// test element validator
	{
		foo := Foo{
			boo: 3,
			zoo: []int{1, 12, 2},
			Bar: "hello",
		}
		buf := make([]byte, foo.SizeMUS())
		foo.MarshalMUS(buf)

		afoo := Foo{}
		_, err := afoo.UnmarshalMUS(buf)
		if err == nil {
			t.Error("validation doesn't work")
		}
		fieldErr, ok := err.(errs.FieldError)
		if !ok {
			t.Error("wrong field error")
		}
		if fieldErr.FieldName() != "zoo" {
			t.Error("wrong field error fieldName")
		}
		sliceErr, ok := fieldErr.Cause().(errs.SliceError)
		if !ok {
			t.Error("wrong slice error")
		}
		if sliceErr.Index() != 1 {
			t.Error("wrong slice error index")
		}
		if sliceErr.Cause() != validators.ErrBiggerThanTen {
			t.Error("wrong error")
		}
	}
	// test max length
	{
		foo := Foo{
			boo: 8,
			zoo: []int{1, 2},
			Bar: "hello world",
		}
		buf := make([]byte, foo.SizeMUS())
		foo.MarshalMUS(buf)

		afoo := Foo{}
		_, err := afoo.UnmarshalMUS(buf)
		if err == nil {
			t.Error("validation doesn't work")
		}
		fieldErr, ok := err.(errs.FieldError)
		if !ok {
			t.Error("wrong field error")
		}
		if fieldErr.FieldName() != "Bar" {
			t.Error("wrong field error fieldName")
		}
		if fieldErr.Cause() != errs.ErrMaxLengthExceeded {
			t.Error("wrong error")
		}
	}
}
```

More advanced usage you can find at https://github.com/ymz-ncnk/musgotest.

When encoding multiple values, it is impractical to create a new buffer each 
time. It takes too long. Instead, you can use the same buffer for each Marshal:
```go
...
buf := make([]byte, FixedLength)
for foo := range foos {
  if foo.Size() > len(buf) {
    return errors.New("buf is too small")
  }
  i := foo.MarshalMUS(buf)
  err = handle(buf[:i])
  ...
}
```

To gain more performance, `recover()` function can be used:
```go
...
defer func() {
  if r := recover(); r != nil {
    return errors.New("buf is too small")
  }
}()
buf := make([]byte, FixedLength)
for foo := range foos {
  i := foo.MarshalMUS(buf)
  err = handle(buf[:i])
  ...
}
```
It will intercept every panic, so use it with careful.

# Supported Types

Supports following types:
  - uint64
  - uint32
  - uint18
  - uint8
  - uint
  - int64
  - int32
  - int18
  - int8
  - int
  - bool
  - byte
  - string
  - array
  - slice
  - map
  - struct
  - alias

Pointers are supported as well. But aliases to pointer types are not, Go
doesn't allow methods for such types.

# Private fields

You could encode and decode private fields too.

# Unsafe code

You could generate fast unsafe code. Read more about it at 
"https://github.com/ymz-ncnk/musgen".

# Validation

For every structure field you can set up validators using the
`mus:"Validator,MaxLength,ElemValidator,KeyValidator"` tag , where:
- Validator - it's a name of the function which will validate current 
  field.
- MaxLength - if field has a string, array, slice or map type, MaxLength 
  will restrict its length. Should be positive number.
- ElemValidator - it's a name of the function which will validate field's 
  elements, if field type is an array, slice of map.
- KeyValidator - it's a name of the function which will validate field's keys,
  if field type is a map.

All tag items, except MaxLength, should have the format like 
"package.FunctionName" or "FunctionName".

Decoding(and encoding) is performed in order, from the first field to the last 
one. That's why, it will stop with validation error on the first not valid 
field.

For alias type, you can set up validators with help of MusGo.GenerateAlias()
method.

## Validators

Validator is a function with the following signature `func (value Type) error`,
where `Type` is a type of the value to which the validator is applied.

A few examples,

```go
// Validator for the field.
type Foo struct {
  Field string `mus:"StrValidator"`
}

func StrValidator(str string) errorr {...}

// ElemValidator for the slice field.
type Bar struct {
  Field []string `mus:",,StrValidator"`
}

// KeyValidator for the map field.
type Zoo struct {
  Field map[int]string `mus:",,,StrValidator"`
}

// Validator for the field of a custom type.
type Far struct {
  Field Foo `mus:FooValidator`
}

func FooValidator(foo *Foo) error {...} // validators for custom types receive
// a pointer as argument

// Validator for the alias field.
type Ror []string

type Pac struct {
  Field Ror `mus:RorValidator`  // you can't set MaxLength or 
  // ElemValidator here, they should be applied for the Ror type.
}

func RorValidator(ror *Ror) error {...} // has pointer as argument too
```

## Errors

Often validation errors are wrapped by one of the predefined error 
(from `errs` package):
- FieldError - happens when a field validation failed. Contains field name and 
  cause.
- SliceError - happens when a validation of a slice element failed. Contains
  element index and cause.
- ArrayError - happens when a validation of an array element failed. Contains
  element index and cause.  
- MapKeyError - happens when a validation of a map key failed. Contains key and
  cause.
- MapValueError - happens when validation of a map value failed. Contains key,
  value and cause.

# Benchmarks

[github.com/alecthomas/go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks)