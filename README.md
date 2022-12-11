# MusGo
MusGo is an extremely fast serializer based on code generation. It supports 
validation, different encodings, aliases, pointers, and private fields.

# Why we need another serializer?
1. With MusGo you can encode/decode your data really fast.
2. Encoded values take up so little space because the serialization format is 
  very simple.
3. Moreover, with MusGo invalid data decodes almost instantly, see the 
  Validation section.

# Binary serialization format
[github.com/ymz-ncnk/musgen](https://github.com/ymz-ncnk/musgen)

# Tests
The generated code is well tested (to run tests read the instructions in the 
[musgo_test.go](musgo_test.go) file). Test coverage is about 80%.

# Benchmarks
[github.com/alecthomas/go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks)

# How to use
First, you should download and install Go, version 1.4 or later.

Create in your home directory a `foo` folder with the following structure:
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
  num int `mus:"validators.Positive"` // private fields are supported
  // too, will be checked with BiggerThanTen validator while unmarshalling 
  arr []int `mus:",,validators.Positive"` // every element will be checked
  // with BiggerThanTen validator
  Alias StringAlias // alias types are supported too
  Bool bool     `mus:"-"` // this field will be skiped
}

type StringAlias string
```

__validators/validators.go__
```go
package validators

import "errors"

var ErrNegative error = errors.New("negative")

func Positive(n int) error {
  if n < 0 {
    return ErrNegative
  }
  return nil
}
```

__make/musable.go__
```go
//go:build ignore

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
  var alias foo.StringAlias
  // Alias types don't support tags, so to set up validators we use
  // GenerateAlias method.
  maxLength := 5 // restricts length of StringAlias values to 5 characters
  err = musGo.GenerateAlias(reflect.TypeOf(alias), unsafe, "", "", maxLength,
    "", "", "", "")
  if err != nil {
    panic(err)
  }
  // reflect.Type could be created without the explicit variable.
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

Now you can see `Foo.musgen.go` and `StringAlias.musgen.go` files in the `foo` 
folder. Pay attention to the location of the generated files. The data type and 
the code generated for it must be in the same package. Let's write some tests.
Create a `foo_test.go` file:
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
		num:   5,
		arr:   []int{4, 2},
		Alias: StringAlias("hello"),
		Bool:  true,
	}
	buf := make([]byte, foo.SizeMUS())
	foo.MarshalMUS(buf)

	afoo := Foo{}
	_, err := afoo.UnmarshalMUS(buf)
	if err != nil {
		t.Error(err)
	}
	foo.Bool = false
	if !reflect.DeepEqual(foo, afoo) {
		t.Error("something went wrong")
	}
}

func TestFooValidation(t *testing.T) {
	// test simple validator
	{
		foo := Foo{
			num:   -11,
			arr:   []int{1, 2},
			Alias: "hello",
		}
		buf := make([]byte, foo.SizeMUS())
		foo.MarshalMUS(buf)

		afoo := Foo{}
		_, err := afoo.UnmarshalMUS(buf)
		if err == nil {
			t.Error("validation error expected")
		}
		if fieldErr, ok := err.(errs.FieldError); ok {
			if fieldErr.FieldName() != "num" {
				t.Error("wrong FieldError's FieldName")
			}
			if fieldErr.Cause() != validators.ErrNegative {
				t.Error("wrong FieldError's Cause")
			}
		} else {
			t.Error("not FiledError")
		}
	}
	// test element validator
	{
		foo := Foo{
			num:   3,
			arr:   []int{1, -12, 2},
			Alias: "hello",
		}
		buf := make([]byte, foo.SizeMUS())
		foo.MarshalMUS(buf)

		afoo := Foo{}
		_, err := afoo.UnmarshalMUS(buf)
		if err == nil {
			t.Error("validation error expected")
		}
		if fieldErr, ok := err.(errs.FieldError); ok {
			if fieldErr.FieldName() != "arr" {
				t.Error("wrong FieldError's FieldName")
			}
			if sliceErr, ok := fieldErr.Cause().(errs.SliceError); ok {
				if sliceErr.Index() != 1 {
					t.Error("wrong SliceError's Index")
				}
				if sliceErr.Cause() != validators.ErrNegative {
					t.Error("wrong SliceError's Cause")
				}
			} else {
				t.Error("not SliceError")
			}
		} else {
			t.Error("not FiledError")
		}
	}
	// test max length
	{
		foo := Foo{
			num:   8,
			arr:   []int{1, 2},
			Alias: "hello world",
		}
		buf := make([]byte, foo.SizeMUS())
		foo.MarshalMUS(buf)

		afoo := Foo{}
		_, err := afoo.UnmarshalMUS(buf)
		if err == nil {
			t.Error("validation error expected")
		}
		if fieldErr, ok := err.(errs.FieldError); ok {
			if fieldErr.FieldName() != "Alias" {
				t.Error("wrong FieldError's FieldName")
			}
			if fieldErr.Cause() != errs.ErrMaxLengthExceeded {
				t.Error("wrong FieldError's Cause")
			}
		} else {
			t.Error("not FiledError")
		}
	}
}
```
More advanced usage you can find at https://github.com/ymz-ncnk/musgotry.

When encoding multiple values, it is impractical to create a new buffer each 
time, it takes too long. Instead, you can use the same buffer for each Marshal:
```go
...
buf := make([]byte, FixedLength)
for foo := range foos {
  if foo.Size() > len(buf) {
    return errors.New("buf is too small")
  }
  i = foo.MarshalMUS(buf)
  err = handle(buf[:i])
  ...
}
```

To gain more performance, the `recover()` function can be used:
```go
...
defer func() {
  if r := recover(); r != nil {
    return errors.New("buf is too small")
  }
}()
buf := make([]byte, FixedLength)
for _, foo := range foos {
  i = foo.MarshalMUS(buf)
  err = handle(buf[:i])
  ...
}
```
It will intercept every panic, so use it with careful.

# Supported Types
Supports following types:
  - `bool`
  - `byte`
  - `int`
  - `int8`
  - `int16`
  - `int32`
  - `int64`
  - `uint`
  - `uint8`
  - `uint16`
  - `uint32`
  - `uint64`
  - `float32`
  - `float64`
  - `string`
  - `array`
  - `slice`
  - `map`
  - `struct`
  - `alias`

Pointers are supported as well. But aliases to pointer types are not, Go
doesn't allow methods for such types.

# Private fields
You could encode and decode private fields too.

# Unsafe code
You could generate fast unsafe code. Read more about it at 
[github.com/ymz-ncnk/musgen](https://github.com/ymz-ncnk/musgen).

# Validation
For every structure field you can set up validators using the
`mus:"Validator,MaxLength,ElemValidator,KeyValidator"` tag , where:
- Validator - it's a name of the function that will validate the current 
  field.
- MaxLength - if the field is a string, array, slice or map, MaxLength will 
  restrict its length. Must be a positive number.
- ElemValidator - it's a name of the function that will validate field
  elements, if the field is an array, slice or map.
- KeyValidator - it's a name of the function that will validate field keys,
  if the field is a map.

All tag items, except MaxLength, must have the "package.FunctionName" or 
"FunctionName" format.

Decoding(and encoding) is performed in order, from the first field to the last 
one. That's why, it will stop with a validation error on the first not valid 
field. There is no practical reason for decoding the rest of the structure when 
we already know that it is not valid.

For an alias type, you can set up validators with help of the 
`MusGo.GenerateAlias()` method.

## Validators
Validator is a function with the following signature `func (value Type) error`,
where `Type` is a type of the value to which the validator is applied.

A few examples:
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

// Validator for the field of a custom pointer type.
type Far struct {
  Field *Foo `mus:FooValidator`
}

func FooValidator(foo *Foo) error {...}

// Validator for the alias field.
type Ror []string

type Pac struct {
  Field Ror `mus:RorValidator`  // you can't set MaxLength or 
  // ElemValidator here, they should be applied for the Ror type.
}

func RorValidator(ror Ror) error {...}
```

## Errors
Often validation errors are wrapped by one of the predefined error 
(from the `errs` package):
- FieldError - happens when field validation failed. Contains the field name
  and cause.
- SliceError - happens when validation of the slice element failed. Contains 
  the element index and cause.
- ArrayError - happens when validation of the array element failed. Contains 
  the element index and cause.
- MapKeyError - happens when validation of the map key failed. Contains the 
  key and cause.
- MapValueError - happens when validation of the map value failed. Contains 
  the key, value and cause.

# Encodings
All `uint`, `int` and `float` types support `Varint` and `Raw` encodings. By 
default `Varint` is used. You can choose `Raw` encoding using the `#raw` in
`mus:"Validator#raw,MaxLength,ElemValidator#raw,KeyValidator#raw"` tag. 

For example:
```go
// Set up Raw encoding without validator for the field.
type Foo struct {
  Field uint64 `mus:"#raw"`
}

// Set up validator and Raw encoding for the field.
type Foo struct {
  Field uint64 `mus:"Positive#raw"`
}
```

`Raw` encoding has better speed and worse size. Only on large numbers 
(> 2^48 in uint representation) it has same or lesser size as `Varint`.

For an alias type, you can set up encoding with help of the 
`MusGo.GenerateAlias()` method.

# DotMusGo
By default generated files create a mess in your folder. If you don't like this
try [github.com/ymz-ncnk/dotmusgo](https://github.com/ymz-ncnk/dotmusgo).

# Single number serialization
If all you want is to serialize a single number you can use:
- [github.com/ymz-ncnk/musgo_int](https://github.com/ymz-ncnk/musgo_int)
- [github.com/ymz-ncnk/musgo_int8](https://github.com/ymz-ncnk/musgo_int8)
- [github.com/ymz-ncnk/musgo_int16](https://github.com/ymz-ncnk/musgo_int16)
- [github.com/ymz-ncnk/musgo_int32](https://github.com/ymz-ncnk/musgo_int32)
- [github.com/ymz-ncnk/musgo_int64](https://github.com/ymz-ncnk/musgo_int64)
- [github.com/ymz-ncnk/musgo_uint](https://github.com/ymz-ncnk/musgo_uint)
- [github.com/ymz-ncnk/musgo_uint8](https://github.com/ymz-ncnk/musgo_uint8)
- [github.com/ymz-ncnk/musgo_uint16](https://github.com/ymz-ncnk/musgo_uint16)
- [github.com/ymz-ncnk/musgo_uint32](https://github.com/ymz-ncnk/musgo_uint32)
- [github.com/ymz-ncnk/musgo_uint64](https://github.com/ymz-ncnk/musgo_uint64)
- [github.com/ymz-ncnk/musgo_float32](https://github.com/ymz-ncnk/musgo_float32)
- [github.com/ymz-ncnk/musgo_float64](https://github.com/ymz-ncnk/musgo_float64)