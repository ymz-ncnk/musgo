package musgen

import "github.com/ymz-ncnk/musgen"

type StrArrayAlias [3]string

var StrArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "StrArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[3]string",
		Alias: "StrArrayAlias"}},
}

type FloatPtrArrayAlias [3]*float64

var FloatPtrArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "FloatPtrArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[3]*float64",
		Alias: "FloatPtrArrayAlias"}},
}

type FloatNilPtrArrayAlias [3]*float32

var FloatNilPtrArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "FloatNilPtrArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[3]*float32",
		Alias: "FloatNilPtrArrayAlias"}},
}

type Int16NilPtrArrayAlias [3]*int16

var Int16NilPtrArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Int16NilPtrArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[3]*int16",
		Alias: "Int16NilPtrArrayAlias"}},
}

type IntAliasArrayAlias [3]IntAlias

var IntAliasArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "IntAliasArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[3]IntAlias",
		Alias: "IntAliasArrayAlias"}},
}

type Uint64PtrAliasArrayAlias [3]*Uint64Alias

var Uint64PtrAliasArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Uint64PtrAliasArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[3]*Uint64Alias",
		Alias: "Uint64PtrAliasArrayAlias"}},
}

type Int64SliceArrayAlias [3][]int64

var Int64SliceArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Int64SliceArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[3][]int64",
		Alias: "Int64SliceArrayAlias"}},
}

type Uint16SlicePtrArrayAlias [3]*[]uint16

var Uint16SlicePtrArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Uint16SlicePtrArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[3]*[]uint16",
		Alias: "Uint16SlicePtrArrayAlias"}},
}

type BoolArrayArrayAlias [3][3]bool

var BoolArrayArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "BoolArrayArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[3][3]bool",
		Alias: "BoolArrayArrayAlias"}},
}

type BytePtrArrayPtrArrayAlias [2]*[2]byte

var BytePtrArrayPtrArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "BytePtrArrayPtrArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[2]*[2]uint8",
		Alias: "BytePtrArrayPtrArrayAlias"}},
}

type IntStrMapArrayAlias [3]map[int]string

var IntStrMapArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "IntStrMapArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[3]map-0[int]-0string",
		Alias: "IntStrMapArrayAlias"}},
}

type Uint32Int32MapArrayAlias [4]*map[uint32]int32

var Uint32Int32MapArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Uint32Int32MapArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[4]*map-0[uint32]-0int32",
		Alias: "Uint32Int32MapArrayAlias"}},
}

type StructTypeArrayAlias [3]SimpleStructType

var StructTypeArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "StructTypeArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[3]SimpleStructType",
		Alias: "StructTypeArrayAlias"}},
}

type StructTypePtrArrayAlias [3]*SimpleStructType

var StructTypePtrArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "StructTypePtrArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[3]*SimpleStructType",
		Alias: "StructTypePtrArrayAlias"}},
}

type TrickyArrayAlias [2][][3]map[SimpleStructType][1]int

var TrickyArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "TrickyArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[2][][3]map-0[SimpleStructType]-0[1]int",
		Alias: "TrickyArrayAlias"}},
}

type StrZeroLengthArrayAlias [0]string

var StrZeroLengthArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "StrZeroLengthArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[0]string",
		Alias: "StrZeroLengthArrayAlias"}},
}

type IntPtrPtrPtrAliasArrayAlias [3]***int

var IntPtrPtrPtrAliasArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "IntPtrPtrPtrAliasArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[3]***int",
		Alias: "IntPtrPtrPtrAliasArrayAlias"}},
}

type ValidIntArrayAlias [2]int

var ValidIntArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "ValidIntArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[2]int",
		Validator:     "ValidIntArrayAliasSumBiggerThanTen",
		ElemValidator: "BiggerThanTenInt", Alias: "ValidIntArrayAlias"}},
}
