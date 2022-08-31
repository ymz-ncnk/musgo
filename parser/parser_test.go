package parser

import (
	"fmt"
	"io"
	"math/big"
	"reflect"
	"testing"

	"github.com/ymz-ncnk/musgen"
)

var currPkg = "parser"

func TestParseSimpleType(t *testing.T) {
	var v int
	_, err := Parse(reflect.TypeOf(v))
	if shouldFail(err) {
		t.Error("required to fail")
	}
}

func TestParsePtrTypeAlias(t *testing.T) {
	var vp *int
	_, err := Parse(reflect.TypeOf(vp))
	if shouldFail(err) {
		t.Error("required to fail")
	}

	type IntPtrAlias *int
	var v IntPtrAlias
	_, err = Parse(reflect.TypeOf(v))
	if shouldFail(err) {
		t.Error("required to fail")
	}

	type Struct struct{}
	type StructAlias *Struct
	var sap StructAlias
	_, err = Parse(reflect.TypeOf(sap))
	if shouldFail(err) {
		t.Error("required to fail")
	}

	var sp *Struct
	_, err = Parse(reflect.TypeOf(sp))
	if shouldFail(err) {
		t.Error("required to fail")
	}
}

func TestParseInterfaceTypeAlias(t *testing.T) {
	type InterfaceAlias io.Reader
	var v InterfaceAlias
	_, err := Parse(reflect.TypeOf(v))
	if err == nil {
		t.Error("required to fail")
	}
	if err.Error() != "type is nil" {
		t.Error("wrong error")
	}

	type InterfacePtrAlias **io.Reader
	var vp InterfacePtrAlias
	_, err = Parse(reflect.TypeOf(vp))
	if shouldFail(err) {
		t.Error("required to fail")
	}
}

func TestParseFromAnotherPkg(t *testing.T) {
	type MyStruct struct {
		big **big.Int
	}
	var v MyStruct
	td, err := Parse(reflect.TypeOf(v))
	if err != nil {
		t.Error(err)
	}
	etd := musgen.TypeDesc{
		Package: currPkg,
		Name:    "MyStruct",
		Fields: []musgen.FieldDesc{{
			Name: "big",
			Type: "**big.Int",
		}},
	}
	if !reflect.DeepEqual(td, etd) {
		t.Fail()
	}
}

func TestParsePrimitiveTypeAlias(t *testing.T) {
	type BoolAlias bool
	var v BoolAlias
	td, err := Parse(reflect.TypeOf(v))
	if err != nil {
		t.Error(err)
	}
	etd := musgen.TypeDesc{
		Package: currPkg,
		Name:    "BoolAlias",
		Fields: []musgen.FieldDesc{{
			Name:      "",
			Type:      "bool",
			MaxLength: 0,
			Alias:     "BoolAlias",
		}},
	}
	if !reflect.DeepEqual(td, etd) {
		t.Error("BoolAlias")
	}

	type Uint64Alias uint64
	var vu64 Uint64Alias
	td, err = Parse(reflect.TypeOf(vu64))
	if err != nil {
		t.Error(err)
	}
	etd = musgen.TypeDesc{
		Package: currPkg,
		Name:    "Uint64Alias",
		Fields: []musgen.FieldDesc{{
			Name:  "",
			Type:  "uint64",
			Alias: "Uint64Alias",
		}},
	}
	if !reflect.DeepEqual(td, etd) {
		t.Error("Uint64Alias")
	}

	type Uint32Alias uint32
	var vu32 Uint32Alias
	td, err = Parse(reflect.TypeOf(vu32))
	if err != nil {
		t.Error(err)
	}
	etd = musgen.TypeDesc{
		Package: currPkg,
		Name:    "Uint32Alias",
		Fields: []musgen.FieldDesc{{
			Name:  "",
			Type:  "uint32",
			Alias: "Uint32Alias",
		}},
	}
	if !reflect.DeepEqual(td, etd) {
		t.Error("Uint32Alias")
	}

	type Uint16Alias uint16
	var vu16 Uint16Alias
	td, err = Parse(reflect.TypeOf(vu16))
	if err != nil {
		t.Error(err)
	}
	etd = musgen.TypeDesc{
		Package: currPkg,
		Name:    "Uint16Alias",
		Fields: []musgen.FieldDesc{{
			Name:  "",
			Type:  "uint16",
			Alias: "Uint16Alias",
		}},
	}
	if !reflect.DeepEqual(td, etd) {
		t.Error("Uint16Alias")
	}

	type Uint8Alias uint8
	var vu8 Uint8Alias
	td, err = Parse(reflect.TypeOf(vu8))
	if err != nil {
		t.Error(err)
	}
	etd = musgen.TypeDesc{
		Package: currPkg,
		Name:    "Uint8Alias",
		Fields: []musgen.FieldDesc{{
			Name:  "",
			Type:  "uint8",
			Alias: "Uint8Alias",
		}},
	}
	if !reflect.DeepEqual(td, etd) {
		t.Error("Uint8Alias")
	}

	type UintAlias uint
	var vu UintAlias
	td, err = Parse(reflect.TypeOf(vu))
	if err != nil {
		t.Error(err)
	}
	etd = musgen.TypeDesc{
		Package: currPkg,
		Name:    "UintAlias",
		Fields: []musgen.FieldDesc{{
			Name:  "",
			Type:  "uint",
			Alias: "UintAlias",
		}},
	}
	if !reflect.DeepEqual(td, etd) {
		t.Error("UintAlias")
	}
}

func TestParseArrayAlias(t *testing.T) {
	type ArrayAlias [3]int
	var v ArrayAlias
	td, err := Parse(reflect.TypeOf(v))
	if err != nil {
		t.Error(err)
	}
	etd := musgen.TypeDesc{
		Package: currPkg,
		Name:    "ArrayAlias",
		Fields: []musgen.FieldDesc{{
			Name:      "",
			Type:      "[3]int",
			MaxLength: 0,
			Alias:     "ArrayAlias",
		}},
	}
	if !reflect.DeepEqual(td, etd) {
		t.Fail()
	}
}

func TestParseSliceAlias(t *testing.T) {
	type SliceAlias []string
	var v SliceAlias
	td, err := Parse(reflect.TypeOf(v))
	if err != nil {
		t.Error(err)
	}
	etd := musgen.TypeDesc{
		Package: currPkg,
		Name:    "SliceAlias",
		Fields: []musgen.FieldDesc{{
			Name:      "",
			Type:      "[]string",
			MaxLength: 0,
			Alias:     "SliceAlias",
		}},
	}
	if !reflect.DeepEqual(td, etd) {
		t.Fail()
	}
}
func TestParseMapAlias(t *testing.T) {
	type MapAlias map[int32]float64
	var v MapAlias
	td, err := Parse(reflect.TypeOf(v))
	if err != nil {
		t.Error(err)
	}
	etd := musgen.TypeDesc{
		Package: currPkg,
		Name:    "MapAlias",
		Fields: []musgen.FieldDesc{{
			Name:      "",
			Type:      "map-0[int32]-0float64",
			MaxLength: 0,
			Alias:     "MapAlias",
		}},
	}
	if !reflect.DeepEqual(td, etd) {
		t.Fail()
	}
}

func TestParseStruct(t *testing.T) {
	type Struct struct {
		R int
		T float64
	}
	var v Struct
	td, err := Parse(reflect.TypeOf(v))
	if err != nil {
		t.Error(err)
	}
	etd := musgen.TypeDesc{
		Package: currPkg,
		Name:    "Struct",
		Fields: []musgen.FieldDesc{
			{
				Name:      "R",
				Type:      "int",
				MaxLength: 0,
				Alias:     "",
			},
			{
				Name:      "T",
				Type:      "float64",
				MaxLength: 0,
				Alias:     "",
			},
		},
	}
	if !reflect.DeepEqual(td, etd) {
		t.Fail()
	}
}

func TestParseStructAlias(t *testing.T) {
	type Struct struct {
		ui uint8
		F  float32
	}
	type StructAlias Struct
	var v StructAlias
	td, err := Parse(reflect.TypeOf(v))
	if err != nil {
		t.Error(err)
	}
	etd := musgen.TypeDesc{
		Package: currPkg,
		Name:    "StructAlias",
		Fields: []musgen.FieldDesc{
			{
				Name:      "ui",
				Type:      "uint8",
				MaxLength: 0,
				Alias:     "",
			},
			{
				Name:      "F",
				Type:      "float32",
				MaxLength: 0,
				Alias:     "",
			},
		},
	}
	if !reflect.DeepEqual(td, etd) {
		t.Fail()
	}
}

func TestParseNestedMapAlias(t *testing.T) {
	type NestedMapAlias map[*map[int64]map[uint8]uint32]**map[string]map[int]**string
	var v NestedMapAlias
	td, err := Parse(reflect.TypeOf(v))
	if err != nil {
		t.Error(err)
	}
	etd := musgen.TypeDesc{
		Package: currPkg,
		Name:    "NestedMapAlias",
		Fields: []musgen.FieldDesc{{
			Name:      "",
			Type:      "map-4[*map-1[int64]-1map-0[uint8]-0uint32]-4**map-3[string]-3map-2[int]-2**string",
			MaxLength: 0,
			Alias:     "NestedMapAlias",
		}},
	}
	if !reflect.DeepEqual(td, etd) {
		t.Fail()
	}
}

func TestParseTrickyStructAlias(t *testing.T) {
	type St struct {
		ui uint8
	}
	type ArrayAlias [2]int
	type Struct struct {
		b     bool
		Array [3]*[]**[3]map[int]string
		St    St
		Stp   *St
		MSt   map[St]map[ArrayAlias][]*St
		Aa    ArrayAlias
	}
	type FirstStructAlias Struct
	type SecondStructAlias FirstStructAlias
	var v SecondStructAlias
	td, err := Parse(reflect.TypeOf(v))
	if err != nil {
		t.Error(err)
	}
	etd := musgen.TypeDesc{
		Package: currPkg,
		Name:    "SecondStructAlias",
		Fields: []musgen.FieldDesc{
			{
				Name:      "b",
				Type:      "bool",
				MaxLength: 0,
				Alias:     "",
			},
			{
				Name:      "Array",
				Type:      "[3]*[]**[3]map-0[int]-0string",
				MaxLength: 0,
				Alias:     "",
			},
			{
				Name:      "St",
				Type:      "St",
				MaxLength: 0,
				Alias:     "",
			},
			{
				Name:      "Stp",
				Type:      "*St",
				MaxLength: 0,
				Alias:     "",
			},
			{
				Name:      "MSt",
				Type:      "map-1[St]-1map-0[ArrayAlias]-0[]*St",
				MaxLength: 0,
				Alias:     "",
			},
			{
				Name:      "Aa",
				Type:      "ArrayAlias",
				MaxLength: 0,
				Alias:     "",
			},
		},
	}
	if !reflect.DeepEqual(td, etd) {
		t.Fail()
	}
}

func TestParseStructWithTags(t *testing.T) {
	// test skip
	{
		// Note, we can't set skip flag on alias.
		type Struct struct {
			myUint uint8 `mus:"-"`
		}
		var v Struct
		td, err := Parse(reflect.TypeOf(v))
		if err != nil {
			t.Error(err)
		}
		etd := musgen.TypeDesc{
			Package: "parser",
			Name:    "Struct",
			Fields:  []musgen.FieldDesc{},
		}
		if !reflect.DeepEqual(td, etd) {
			t.Error("parse skip tag failed")
		}
	}
	// test invalid skip
	{
		type Struct struct {
			myUint uint8 `mus:"-,validator"`
		}
		var v Struct
		_, err := Parse(reflect.TypeOf(v))
		if err == nil {
			t.Error("invalid tag is ok")
		}
		if err.Error() != fmt.Sprintf(InvalidTagFormatErrMsg, "myUint") {
			t.Error("wrong error")
		}
	}
	// test validator string
	{
		type Struct struct {
			myStr string `mus:"validator"`
		}
		var v Struct
		td, err := Parse(reflect.TypeOf(v))
		if err != nil {
			t.Error(err)
		}
		etd := musgen.TypeDesc{
			Package: "parser",
			Name:    "Struct",
			Fields: []musgen.FieldDesc{
				{
					Name:      "myStr",
					Type:      "string",
					Validator: "validator",
				},
			},
		}
		if !reflect.DeepEqual(td, etd) {
			t.Error("parse skip tag failed")
		}
	}
	// TODO test validator for other types
	// test maxLength string
	{
		type Struct struct {
			myStr string `mus:",5"`
		}
		var v Struct
		td, err := Parse(reflect.TypeOf(v))
		if err != nil {
			t.Error(err)
		}
		etd := musgen.TypeDesc{
			Package: "parser",
			Name:    "Struct",
			Fields: []musgen.FieldDesc{
				{
					Name:      "myStr",
					Type:      "string",
					MaxLength: 5,
				},
			},
		}
		if !reflect.DeepEqual(td, etd) {
			t.Error("parse skip tag failed")
		}
	}
	// test maxLength array
	{
		type Struct struct {
			myStrArr [2]string `mus:",5"`
		}
		var v Struct
		_, err := Parse(reflect.TypeOf(v))
		if err == nil {
			t.Error("invalid tag is ok")
		}
		if err.Error() != fmt.Sprintf(InvalidTagOwnMaxLengthErrMsg, "myStrArr") {
			t.Error("wrong error")
		}
	}
	// test maxLength slice
	{
		type Struct struct {
			myStrSlice []string `mus:",10"`
		}
		var v Struct
		td, err := Parse(reflect.TypeOf(v))
		if err != nil {
			t.Error(err)
		}
		etd := musgen.TypeDesc{
			Package: "parser",
			Name:    "Struct",
			Fields: []musgen.FieldDesc{
				{
					Name:      "myStrSlice",
					Type:      "[]string",
					MaxLength: 10,
				},
			},
		}
		if !reflect.DeepEqual(td, etd) {
			t.Error("parse tag with maxLength failed")
		}
	}
	// test maxLength map
	{
		type Struct struct {
			myMap map[string]int `mus:",-1"`
		}
		var v Struct
		_, err := Parse(reflect.TypeOf(v))
		if err == nil {
			t.Error("invalid tag is ok")
		}
		if err.Error() != fmt.Sprintf(InvalidTagMaxLengthErrMsg, "myMap") {
			t.Error("wrong error")
		}
	}
	// test elemValidator string
	{
		type Struct struct {
			myStr string `mus:",,elemValidator"`
		}
		var v Struct
		_, err := Parse(reflect.TypeOf(v))
		if err == nil {
			t.Error("invalid tag is ok")
		}
		if err.Error() != fmt.Sprintf(InvalidTagOwnElemValidatorErrMsg, "myStr") {
			t.Error("wrong error")
		}
	}
	// test elemValidator array
	{
		type Struct struct {
			myStrArr [2]string `mus:",,elemValidator"`
		}
		var v Struct
		td, err := Parse(reflect.TypeOf(v))

		if err != nil {
			t.Error(err)
		}
		etd := musgen.TypeDesc{
			Package: "parser",
			Name:    "Struct",
			Fields: []musgen.FieldDesc{
				{
					Name:          "myStrArr",
					Type:          "[2]string",
					ElemValidator: "elemValidator",
				},
			},
		}
		if !reflect.DeepEqual(td, etd) {
			t.Error("parse tag with elemValidator failed")
		}
	}
	// test elemValidator array
	{
		type Struct struct {
			myStrSlice []string `mus:",,elemValidator"`
		}
		var v Struct
		td, err := Parse(reflect.TypeOf(v))
		if err != nil {
			t.Error(err)
		}
		etd := musgen.TypeDesc{
			Package: "parser",
			Name:    "Struct",
			Fields: []musgen.FieldDesc{
				{
					Name:          "myStrSlice",
					Type:          "[]string",
					ElemValidator: "elemValidator",
				},
			},
		}
		if !reflect.DeepEqual(td, etd) {
			t.Error("parse tag with elemValidator failed")
		}
	}
	// test elemValidator map
	{
		type Struct struct {
			myStrMap map[string]int `mus:",,elemValidator"`
		}
		var v Struct
		td, err := Parse(reflect.TypeOf(v))
		if err != nil {
			t.Error(err)
		}
		etd := musgen.TypeDesc{
			Package: "parser",
			Name:    "Struct",
			Fields: []musgen.FieldDesc{
				{
					Name:          "myStrMap",
					Type:          "map-0[string]-0int",
					ElemValidator: "elemValidator",
				},
			},
		}
		if !reflect.DeepEqual(td, etd) {
			t.Error("parse tag with elemValidator failed")
		}
	}
	// test keyValidator string
	{
		type Struct struct {
			myStr string `mus:",,,keyValidator"`
		}
		var v Struct
		_, err := Parse(reflect.TypeOf(v))
		if err == nil {
			t.Error("invalid tag is ok")
		}
		if err.Error() != fmt.Sprintf(InvalidTagOwnKeyValidatorErrMsg, "myStr") {
			t.Error("wrong error")
		}
	}
	// test keyValidator array
	{
		type Struct struct {
			myStrArr [2]string `mus:",,,keyValidator"`
		}
		var v Struct
		_, err := Parse(reflect.TypeOf(v))
		if err == nil {
			t.Error("invalid tag is ok")
		}
		if err.Error() != fmt.Sprintf(InvalidTagOwnKeyValidatorErrMsg, "myStrArr") {
			t.Error("wrong error")
		}
	}
	// test keyValidator slice
	{
		type Struct struct {
			myStrSlice []string `mus:",,,keyValidator"`
		}
		var v Struct
		_, err := Parse(reflect.TypeOf(v))
		if err == nil {
			t.Error("invalid tag is ok")
		}
		if err.Error() != fmt.Sprintf(InvalidTagOwnKeyValidatorErrMsg, "myStrSlice") {
			t.Error("wrong error")
		}
	}
	// test keyValidator map
	{
		type Struct struct {
			myStrMap map[string]int `mus:",,,keyValidator"`
		}
		var v Struct
		td, err := Parse(reflect.TypeOf(v))
		if err != nil {
			t.Error(err)
		}
		etd := musgen.TypeDesc{
			Package: "parser",
			Name:    "Struct",
			Fields: []musgen.FieldDesc{
				{
					Name:         "myStrMap",
					Type:         "map-0[string]-0int",
					KeyValidator: "keyValidator",
				},
			},
		}
		if !reflect.DeepEqual(td, etd) {
			t.Error("parse tag with elemValidator failed")
		}
	}
	// encoding
	{
		type Struct struct {
			i map[bool]string `mus:"valid#enc,5,valid1#enc1,valid2#enc2"`
		}
		var v Struct
		td, err := Parse(reflect.TypeOf(v))
		if err != nil {
			t.Error(err)
		}
		etd := musgen.TypeDesc{
			Package: "parser",
			Name:    "Struct",
			Fields: []musgen.FieldDesc{
				{
					Name:          "i",
					Type:          "map-0[bool]-0string",
					Validator:     "valid",
					Encoding:      "enc",
					MaxLength:     5,
					ElemValidator: "valid1",
					ElemEncoding:  "enc1",
					KeyValidator:  "valid2",
					KeyEncoding:   "enc2",
				},
			},
		}
		if !reflect.DeepEqual(td, etd) {
			t.Error("parse tag with encoding failed")
		}
	}
	// encoding not valid
	{
		type Struct struct {
			i int `mus:",,#enc1"`
		}
		var v Struct
		_, err := Parse(reflect.TypeOf(v))
		if err == nil {
			t.Error("nil error")
		} else {
			if err.Error() !=
				fmt.Errorf(InvalidTagOwnElemValidatorErrMsg, "i").Error() {
				t.Error("wrong error")
			}
		}
	}

	// test maxLenght on string alias
	{
		type MyString string
		type Struct struct {
			myStr MyString `mus:",3"`
		}
		var v Struct
		_, err := Parse(reflect.TypeOf(v))
		if err == nil {
			t.Error("maxLength for string alias is ok")
		}
		if err.Error() != fmt.Errorf(InvalidTagOwnMaxLengthErrMsg, "myStr").Error() {
			t.Error("wrong error")
		}
	}
	// test elemValidator on array alias
	{
		type MyArray [2]int
		type Struct struct {
			myArr MyArray `mus:",,elemValidator"`
		}
		var v Struct
		_, err := Parse(reflect.TypeOf(v))
		if err == nil {
			t.Error("elemValidator for array alias is ok")
		}
		if err.Error() != fmt.Errorf(InvalidTagOwnElemValidatorErrMsg, "myArr").Error() {
			t.Error("wrong error")
		}
	}
	// test keyValidator on map alias
	{
		type MyMap map[int]int
		type Struct struct {
			myMap MyMap `mus:",,,keyValidator"`
		}
		var v Struct
		_, err := Parse(reflect.TypeOf(v))
		if err == nil {
			t.Error("keyValidator for map alias is ok")
		}
		if err.Error() != fmt.Errorf(InvalidTagOwnKeyValidatorErrMsg, "myMap").Error() {
			t.Error("wrong error")
		}
	}
}

func shouldFail(err error) bool {
	if err == nil {
		return true
	}
	if _, ok := err.(NotSupportedTypeError); !ok {
		return true
	}
	return false
}
