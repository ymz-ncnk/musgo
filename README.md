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
  zoo []int `mus:",3,validators.BiggerThanTen"`
  Bar MyString
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
  var myStr foo.MyString
  err = musGo.Generate(reflect.TypeOf(myStr), false)
  if err != nil {
    panic(err)
  }
  // Variable creation can be skipped.
  err = musGo.Generate(reflect.TypeOf(*foo.Foo(nil)).Elem(), false)
  if err != nil {
    panic(err)
  }
}
```

Run from the comamnd line:
```
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
foo
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
    Bar: MyString("hello world"),
    zoo: []int{4, 2}, // private fields are supported too
  }
  buf := make([]byte, foo.SizeMUS())
  foo.MarshalMUS(buf)

  afoo := Foo{}
  _, err := afoo.UnmarshalMUS(buf)
  if err != nil {
    t.Error(err)
  }
  if !reflect.DeepEqual(foo, afoo) {
    t.Error("something went wrong")
  }
}

func TestFooValidation(t *testing.T) {
  // test max length
  {
    foo := Foo{
      Bar: MyString("hello world"),
      zoo: []int{9, -2, 0, 5},
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
    if fieldErr.Cause() != errs.ErrMaxLengthExceeded {
      t.Error("wrong error")
    }
  }
  // test element validator
  {
    foo := Foo{
      Bar: MyString("hello world"),
      zoo: []int{8, 12},
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
  
Decoding(and encoding) is performed in order, from the first field to the last 
one. That's why, it will stop with validation error on the first not valid 
field.

For alias type, you can set up validators with help of MusGo.GenerateAlias()
method.

# Benchmarks

[github.com/alecthomas/go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks)