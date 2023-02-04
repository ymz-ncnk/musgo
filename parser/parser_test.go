package parser

import (
	"io"
	"math/big"
	"reflect"
	"testing"
)

func TestParsePrimitiveType(t *testing.T) {
	var (
		v    int
		want = NewUnsupportedTypeError("int")
	)
	_, _, err := Parse(reflect.TypeOf(v))
	if err == nil || err.Error() != want.Error() {
		t.Errorf("want '%v', actual '%v'", want, err)
	}
}

func TestParsePtrType(t *testing.T) {

	t.Run("Unsupported simple pointer type", func(t *testing.T) {
		var (
			v    *int
			want = NewUnsupportedTypeError("*int")
		)
		_, _, err := Parse(reflect.TypeOf(v))
		if err == nil || err.Error() != want.Error() {
			t.Errorf("want '%v', actual '%v'", want, err)
		}
	})

	t.Run("Unsupported simple pointer type alias", func(t *testing.T) {
		type IntPtrAlias *int
		var (
			v    IntPtrAlias
			want = NewUnsupportedTypeError("parser.IntPtrAlias")
		)
		_, _, err := Parse(reflect.TypeOf(v))
		if err == nil || err.Error() != want.Error() {
			t.Errorf("want '%v', actual '%v'", want, err)
		}
	})

	t.Run("Unsupported atruct pointer", func(t *testing.T) {
		type Struct struct{}
		var (
			v    *Struct
			want = NewUnsupportedTypeError("*parser.Struct")
		)
		_, _, err := Parse(reflect.TypeOf(v))
		if err == nil || err.Error() != want.Error() {
			t.Errorf("want '%v', actual '%v'", want, err)
		}
	})

	t.Run("Unsupported struct pointer alias", func(t *testing.T) {
		type Struct struct{}
		type StructAlias *Struct
		var (
			v    StructAlias
			want = NewUnsupportedTypeError("parser.StructAlias")
		)
		_, _, err := Parse(reflect.TypeOf(v))
		if err == nil || err.Error() != want.Error() {
			t.Errorf("want '%v', actual '%v'", want, err)
		}
	})

}

func TestParseInterfaceAlias(t *testing.T) {

	t.Run("Interface alias as nil type", func(t *testing.T) {
		type InterfaceAlias io.Reader
		var (
			v    InterfaceAlias
			want = NewUnsupportedTypeError("nil")
		)
		_, _, err := Parse(reflect.TypeOf(v))
		if err == nil || err.Error() != want.Error() {
			t.Errorf("want '%v', actual '%v'", want, err)
		}
	})

	t.Run("Unsupported interface pointer alias", func(t *testing.T) {
		type InterfacePtrAlias **io.Reader
		var (
			v    InterfacePtrAlias
			want = NewUnsupportedTypeError("parser.InterfacePtrAlias")
		)
		_, _, err := Parse(reflect.TypeOf(v))
		if err == nil || err.Error() != want.Error() {
			t.Errorf("want '%v', actual '%v'", want, err)
		}
	})

}

func TestParseTypeFromAnotherPkg(t *testing.T) {
	type MyStruct struct {
		big **big.Int
	}
	var (
		v           MyStruct
		wantAliasOf = ""
		wantFields  = []string{"**big.Int"}
	)
	test(reflect.TypeOf(v), wantAliasOf, wantFields, t)
	_ = v.big
}

func TestParsePrimitiveTypeAlias(t *testing.T) {

	t.Run("Bool alias", func(t *testing.T) {
		type BoolAlias bool
		var (
			v           BoolAlias
			wantAliasOf          = "bool"
			wantFields  []string = nil
		)
		test(reflect.TypeOf(v), wantAliasOf, wantFields, t)
	})

	t.Run("Uint64 alias", func(t *testing.T) {
		type Uint64Alias uint64
		var (
			v           Uint64Alias
			wantAliasOf          = "uint64"
			wantFields  []string = nil
		)
		test(reflect.TypeOf(v), wantAliasOf, wantFields, t)
	})

	t.Run("Uint32 alias", func(t *testing.T) {
		type Uint32Alias uint32
		var (
			v           Uint32Alias
			wantAliasOf          = "uint32"
			wantFields  []string = nil
		)
		test(reflect.TypeOf(v), wantAliasOf, wantFields, t)
	})

	t.Run("Uint16 alias", func(t *testing.T) {
		type Uint16Alias uint16
		var (
			v           Uint16Alias
			wantAliasOf          = "uint16"
			wantFields  []string = nil
		)
		test(reflect.TypeOf(v), wantAliasOf, wantFields, t)
	})

	t.Run("Uint8 alias", func(t *testing.T) {
		type Uint8Alias uint8
		var (
			v           Uint8Alias
			wantAliasOf          = "uint8"
			wantFields  []string = nil
		)
		test(reflect.TypeOf(v), wantAliasOf, wantFields, t)
	})

	t.Run("Uint alias", func(t *testing.T) {
		type UintAlias uint
		var (
			v           UintAlias
			wantAliasOf          = "uint"
			wantFields  []string = nil
		)
		test(reflect.TypeOf(v), wantAliasOf, wantFields, t)
	})

}

func TestParseArrayAlias(t *testing.T) {
	type ArrayAlias [3]int
	var (
		v           ArrayAlias
		wantAliasOf          = "[3]int"
		wantFields  []string = nil
	)
	test(reflect.TypeOf(v), wantAliasOf, wantFields, t)
}

func TestParseSliceAlias(t *testing.T) {
	type SliceAlias []string
	var (
		v           SliceAlias
		wantAliasOf          = "[]string"
		wantFields  []string = nil
	)
	test(reflect.TypeOf(v), wantAliasOf, wantFields, t)
}

func TestParseMapAlias(t *testing.T) {
	type MapAlias map[int32]float64
	var (
		v           MapAlias
		wantAliasOf          = "map-0[int32]-0float64"
		wantFields  []string = nil
	)
	test(reflect.TypeOf(v), wantAliasOf, wantFields, t)
}

func TestParseStruct(t *testing.T) {
	type Struct struct {
		R int
		T float64
	}
	var (
		v           Struct
		wantAliasOf          = ""
		wantFields  []string = []string{"int", "float64"}
	)
	test(reflect.TypeOf(v), wantAliasOf, wantFields, t)
}

func TestParseStructAlias(t *testing.T) {
	type Struct struct {
		uint8Field uint8
		F          float32
	}
	type StructAlias Struct
	var (
		v           StructAlias
		wantAliasOf          = ""
		wantFields  []string = []string{"uint8", "float32"}
	)
	test(reflect.TypeOf(v), wantAliasOf, wantFields, t)
	_ = v.uint8Field
}

func TestParseNestedMapAlias(t *testing.T) {
	type NestedMapAlias map[*map[int64]map[uint8]uint32]**map[string]map[int]**string
	var (
		v           NestedMapAlias
		wantAliasOf          = "map-4[*map-1[int64]-1map-0[uint8]-0uint32]-4**map-3[string]-3map-2[int]-2**string"
		wantFields  []string = nil
	)
	test(reflect.TypeOf(v), wantAliasOf, wantFields, t)
}

func TestParseTrickyStructAlias(t *testing.T) {
	type AnotherStruct struct {
		uint8Field uint8
	}
	type ArrayAlias [2]int
	type Struct struct {
		boolField       bool
		ArrayField      [3]*[]**[3]map[int]string
		StructField     AnotherStruct
		StructPtrField  *AnotherStruct
		MapField        map[AnotherStruct]map[ArrayAlias][]*AnotherStruct
		ArrayAliasField ArrayAlias
	}
	type FirstStructAlias Struct
	type SecondStructAlias FirstStructAlias
	var (
		v           SecondStructAlias
		wantAliasOf          = ""
		wantFields  []string = []string{
			"bool",
			"[3]*[]**[3]map-0[int]-0string",
			"AnotherStruct",
			"*AnotherStruct",
			"map-1[AnotherStruct]-1map-0[ArrayAlias]-0[]*AnotherStruct",
			"ArrayAlias",
		}
	)
	test(reflect.TypeOf(v), wantAliasOf, wantFields, t)
	_ = v.boolField
	_ = v.StructField.uint8Field
}

// TODO move to tdesc_builder
// func TestParseStructWithTags(t *testing.T) {

// 	// t.Run("Skip flag", func(t *testing.T) {
// 	// 	// Note, we can't set skip flag on alias.
// 	// 	type Struct struct {
// 	// 		myUint uint8 `mus:"-"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = gen4type.TypeDesc{
// 	// 			Package: "parser",
// 	// 			Name:    "Struct",
// 	// 			Fields:  []gen4type.FieldDesc{},
// 	// 		}
// 	// 	)
// 	// 	td, err := Parse(reflect.TypeOf(v))
// 	// 	if err != nil {
// 	// 		t.Error(err)
// 	// 	}
// 	// 	if !reflect.DeepEqual(td, want) {
// 	// 		t.Errorf("want '%v', actual '%v'", td, want)
// 	// 	}
// 	// 	_ = v.myUint
// 	// })

// 	// t.Run("Invalid skip flag", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		uintField uint8 `mus:"-,validator"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = NewInvalidTagFormatError("uintField")
// 	// 	)
// 	// 	_, err := Parse(reflect.TypeOf(v))
// 	// 	if err == nil {
// 	// 		t.Error("invalid tag is ok")
// 	// 	}
// 	// 	if err == nil || err.Error() != want.Error() {
// 	// 		t.Errorf("actual '%v', want '%v'", err, want)
// 	// 	}
// 	// 	_ = v.uintField
// 	// })

// 	// t.Run("Validator of string field", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		strField string `mus:"validator"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = gen4type.TypeDesc{
// 	// 			Package: "parser",
// 	// 			Name:    "Struct",
// 	// 			Fields: []gen4type.FieldDesc{
// 	// 				{
// 	// 					Name: "strField",
// 	// 					Type: "string",
// 	// 					// Validator: "validator",
// 	// 				},
// 	// 			},
// 	// 		}
// 	// 	)
// 	// 	td, err := Parse(reflect.TypeOf(v))
// 	// 	if err != nil {
// 	// 		t.Error(err)
// 	// 	}
// 	// 	if !reflect.DeepEqual(td, want) {
// 	// 		t.Errorf("want '%v', actual '%v'", td, want)
// 	// 	}
// 	// 	_ = v.strField
// 	// })
// 	// TODO test validator for other types

// 	// t.Run("MaxLength of string field", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		strField string `mus:",5"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = gen4type.TypeDesc{
// 	// 			Package: "parser",
// 	// 			Name:    "Struct",
// 	// 			Fields: []gen4type.FieldDesc{
// 	// 				{
// 	// 					Name:      "strField",
// 	// 					Type:      "string",
// 	// 					MaxLength: 5,
// 	// 				},
// 	// 			},
// 	// 		}
// 	// 	)
// 	// 	td, err := Parse(reflect.TypeOf(v))
// 	// 	if err != nil {
// 	// 		t.Error(err)
// 	// 	}
// 	// 	if !reflect.DeepEqual(td, want) {
// 	// 		t.Errorf("want '%v', actual '%v'", td, want)
// 	// 	}
// 	// 	_ = v.strField
// 	// })

// 	// t.Run("MaxLength of array field", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		arrField [2]string `mus:",5"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = NewUnsupportedMaxLengthTagError("arrField")
// 	// 	)
// 	// 	_, err := Parse(reflect.TypeOf(v))
// 	// 	if err == nil {
// 	// 		t.Error("invalid tag is ok")
// 	// 	}
// 	// 	if err == nil || err.Error() != want.Error() {
// 	// 		t.Errorf("actual '%v', want '%v'", err, want)
// 	// 	}
// 	// 	_ = v.arrField
// 	// })

// 	// t.Run("MaxLength of slice field", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		sliceField []string `mus:",10"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = gen4type.TypeDesc{
// 	// 			Package: "parser",
// 	// 			Name:    "Struct",
// 	// 			Fields: []gen4type.FieldDesc{
// 	// 				{
// 	// 					Name:      "sliceField",
// 	// 					Type:      "[]string",
// 	// 					MaxLength: 10,
// 	// 				},
// 	// 			},
// 	// 		}
// 	// 	)
// 	// 	td, err := Parse(reflect.TypeOf(v))
// 	// 	if err != nil {
// 	// 		t.Error(err)
// 	// 	}
// 	// 	if !reflect.DeepEqual(td, want) {
// 	// 		t.Errorf("actual '%v', want '%v'", td, want)
// 	// 	}
// 	// 	_ = v.sliceField
// 	// })

// 	// t.Run("Invalid MaxLength of map field", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		mapField map[string]int `mus:",-1"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = NewInvalidMaxLengthTagError("mapField")
// 	// 	)
// 	// 	_, err := Parse(reflect.TypeOf(v))
// 	// 	if err == nil {
// 	// 		t.Error("invalid tag is ok")
// 	// 	}
// 	// 	if err == nil || err.Error() != want.Error() {
// 	// 		t.Errorf("actual '%v', want '%v'", err, want)
// 	// 	}
// 	// 	_ = v.mapField
// 	// })

// 	// t.Run("Unsupported ElemValidator of string field", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		strField string `mus:",,elemValidator"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = NewUnsupportedElemValidatorTagError("strField")
// 	// 	)
// 	// 	_, err := Parse(reflect.TypeOf(v))
// 	// 	if err == nil {
// 	// 		t.Error("invalid tag is ok")
// 	// 	}
// 	// 	if err == nil || err.Error() != want.Error() {
// 	// 		t.Errorf("actual '%v', want '%v'", err, want)
// 	// 	}
// 	// 	_ = v.strField
// 	// })

// 	// t.Run("ElemValidator of array field", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		arrField [2]string `mus:",,elemValidator"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = gen4type.TypeDesc{
// 	// 			Package: "parser",
// 	// 			Name:    "Struct",
// 	// 			Fields: []gen4type.FieldDesc{
// 	// 				{
// 	// 					Name:          "arrField",
// 	// 					Type:          "[2]string",
// 	// 					ElemValidator: "elemValidator",
// 	// 				},
// 	// 			},
// 	// 		}
// 	// 	)
// 	// 	td, err := Parse(reflect.TypeOf(v))
// 	// 	if err != nil {
// 	// 		t.Error(err)
// 	// 	}
// 	// 	if !reflect.DeepEqual(td, want) {
// 	// 		t.Errorf("actual '%v', want '%v'", td, want)
// 	// 	}
// 	// 	_ = v.arrField
// 	// })

// 	// t.Run("ElemValidator of slice field", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		sliceField []string `mus:",,elemValidator"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = gen4type.TypeDesc{
// 	// 			Package: "parser",
// 	// 			Name:    "Struct",
// 	// 			Fields: []gen4type.FieldDesc{
// 	// 				{
// 	// 					Name:          "sliceField",
// 	// 					Type:          "[]string",
// 	// 					ElemValidator: "elemValidator",
// 	// 				},
// 	// 			},
// 	// 		}
// 	// 	)
// 	// 	td, err := Parse(reflect.TypeOf(v))
// 	// 	if err != nil {
// 	// 		t.Error(err)
// 	// 	}
// 	// 	if !reflect.DeepEqual(td, want) {
// 	// 		t.Errorf("actual '%v', want '%v'", td, want)
// 	// 	}
// 	// 	_ = v.sliceField
// 	// })

// 	// t.Run("ElemValidator of map field", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		mapField map[string]int `mus:",,elemValidator"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = gen4type.TypeDesc{
// 	// 			Package: "parser",
// 	// 			Name:    "Struct",
// 	// 			Fields: []gen4type.FieldDesc{
// 	// 				{
// 	// 					Name:          "mapField",
// 	// 					Type:          "map-0[string]-0int",
// 	// 					ElemValidator: "elemValidator",
// 	// 				},
// 	// 			},
// 	// 		}
// 	// 	)
// 	// 	td, err := Parse(reflect.TypeOf(v))
// 	// 	if err != nil {
// 	// 		t.Error(err)
// 	// 	}
// 	// 	if !reflect.DeepEqual(td, want) {
// 	// 		t.Errorf("actual '%v', want '%v'", td, want)
// 	// 	}
// 	// 	_ = v.mapField
// 	// })

// 	// t.Run("Unsupported KeyValidator of string field", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		strField string `mus:",,,keyValidator"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = NewUnsupportedKeyValidatorTagError("strField")
// 	// 	)
// 	// 	_, err := Parse(reflect.TypeOf(v))
// 	// 	if err == nil {
// 	// 		t.Error("invalid tag is ok")
// 	// 	}
// 	// 	if err == nil || err.Error() != want.Error() {
// 	// 		t.Errorf("actual '%v', want '%v'", err, want)
// 	// 	}
// 	// 	_ = v.strField
// 	// })

// 	// t.Run("Unsupported KeyValidator of array field", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		arrField [2]string `mus:",,,keyValidator"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = NewUnsupportedKeyValidatorTagError("arrField")
// 	// 	)
// 	// 	_, err := Parse(reflect.TypeOf(v))
// 	// 	if err == nil {
// 	// 		t.Error("invalid tag is ok")
// 	// 	}
// 	// 	if err == nil || err.Error() != want.Error() {
// 	// 		t.Errorf("actual '%v', want '%v'", err, want)
// 	// 	}
// 	// 	_ = v.arrField
// 	// })

// 	// t.Run("Unsupported KeyValidator of slice field", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		sliceFiedl []string `mus:",,,keyValidator"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = NewUnsupportedKeyValidatorTagError("sliceFiedl")
// 	// 	)
// 	// 	_, err := Parse(reflect.TypeOf(v))
// 	// 	if err == nil {
// 	// 		t.Error("invalid tag is ok")
// 	// 	}
// 	// 	if err == nil || err.Error() != want.Error() {
// 	// 		t.Errorf("actual '%v', want '%v'", err, want)
// 	// 	}
// 	// 	_ = v.sliceFiedl
// 	// })

// 	// t.Run("KeyValidator of map field", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		mapField map[string]int `mus:",,,keyValidator"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = gen4type.TypeDesc{
// 	// 			Package: "parser",
// 	// 			Name:    "Struct",
// 	// 			Fields: []gen4type.FieldDesc{
// 	// 				{
// 	// 					Name:         "mapField",
// 	// 					Type:         "map-0[string]-0int",
// 	// 					KeyValidator: "keyValidator",
// 	// 				},
// 	// 			},
// 	// 		}
// 	// 	)
// 	// 	td, err := Parse(reflect.TypeOf(v))
// 	// 	if err != nil {
// 	// 		t.Error(err)
// 	// 	}
// 	// 	if !reflect.DeepEqual(td, want) {
// 	// 		t.Errorf("actual '%v', want '%v'", td, want)
// 	// 	}
// 	// 	_ = v.mapField
// 	// })

// 	// t.Run("Encoding", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		field map[bool]string `mus:"valid#enc,5,valid1#enc1,valid2#enc2"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = gen4type.TypeDesc{
// 	// 			Package: "parser",
// 	// 			Name:    "Struct",
// 	// 			Fields: []gen4type.FieldDesc{
// 	// 				{
// 	// 					Name:          "field",
// 	// 					Type:          "map-0[bool]-0string",
// 	// 					Validator:     "valid",
// 	// 					Encoding:      "enc",
// 	// 					MaxLength:     5,
// 	// 					ElemValidator: "valid1",
// 	// 					ElemEncoding:  "enc1",
// 	// 					KeyValidator:  "valid2",
// 	// 					KeyEncoding:   "enc2",
// 	// 				},
// 	// 			},
// 	// 		}
// 	// 	)
// 	// 	td, err := Parse(reflect.TypeOf(v))
// 	// 	if err != nil {
// 	// 		t.Error(err)
// 	// 	}
// 	// 	if !reflect.DeepEqual(td, want) {
// 	// 		t.Errorf("actual '%v', want '%v'", td, want)
// 	// 	}
// 	// 	_ = v.field
// 	// })

// 	// t.Run("Unsupported ElemEncoding of int field", func(t *testing.T) {
// 	// 	type Struct struct {
// 	// 		field int `mus:",,#enc1"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = NewUnsupportedElemValidatorTagError("field")
// 	// 	)
// 	// 	_, err := Parse(reflect.TypeOf(v))
// 	// 	if err == nil || err.Error() != want.Error() {
// 	// 		t.Errorf("actual '%v', want '%v'", err, want)
// 	// 	}
// 	// 	_ = v.field
// 	// })

// 	// t.Run("Unsupported MaxLength of string alias field", func(t *testing.T) {
// 	// 	type MyString string
// 	// 	type Struct struct {
// 	// 		aliasField MyString `mus:",3"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = NewUnsupportedMaxLengthTagError("aliasField")
// 	// 	)
// 	// 	_, err := Parse(reflect.TypeOf(v))
// 	// 	if err == nil {
// 	// 		t.Error("invalid tag is ok")
// 	// 	}
// 	// 	if err == nil || err.Error() != want.Error() {
// 	// 		t.Errorf("actual '%v', want '%v'", err, want)
// 	// 	}
// 	// 	_ = v.aliasField
// 	// })

// 	// t.Run("Unsupported ElemValidator of array alias field", func(t *testing.T) {
// 	// 	type MyArray [2]int
// 	// 	type Struct struct {
// 	// 		aliasField MyArray `mus:",,elemValidator"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = NewUnsupportedElemValidatorTagError("aliasField")
// 	// 	)
// 	// 	_, err := Parse(reflect.TypeOf(v))
// 	// 	if err == nil {
// 	// 		t.Error("invalid tag is ok")
// 	// 	}
// 	// 	if err == nil || err.Error() != want.Error() {
// 	// 		t.Errorf("actual '%v', want '%v'", err, want)
// 	// 	}
// 	// 	_ = v.aliasField
// 	// })

// 	// t.Run("Unsupported KeyValidator of map alias field", func(t *testing.T) {
// 	// 	type MyMap map[int]int
// 	// 	type Struct struct {
// 	// 		aliasField MyMap `mus:",,,keyValidator"`
// 	// 	}
// 	// 	var (
// 	// 		v    Struct
// 	// 		want = NewUnsupportedKeyValidatorTagError("aliasField")
// 	// 	)
// 	// 	_, err := Parse(reflect.TypeOf(v))
// 	// 	if err == nil {
// 	// 		t.Error("invalid tag is ok")
// 	// 	}
// 	// 	if err == nil || err.Error() != want.Error() {
// 	// 		t.Errorf("actual '%v', want '%v'", err, want)
// 	// 	}
// 	// 	_ = v.aliasField
// 	// })
// }

func test(tp reflect.Type, wantAliasOf string, wantFields []string,
	t *testing.T) {
	aliasOf, fields, err := Parse(tp)
	if err != nil {
		t.Error(err)
	}
	if aliasOf != wantAliasOf {
		t.Errorf("want '%v', actual '%v'", wantAliasOf, aliasOf)
	}
	if !reflect.DeepEqual(fields, wantFields) {
		t.Errorf("want '%v', actual '%v'", wantFields, fields)
	}
}
