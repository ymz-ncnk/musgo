package musgen

import "github.com/ymz-ncnk/musgen"

type StrSliceAlias []string

var StrSliceAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "StrSliceAlias",
	Fields:  []musgen.FieldDesc{{Name: "", Type: "[]string", Alias: "StrSliceAlias"}},
}

type FloatPtrSliceAlias []*float64

var FloatPtrSliceAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "FloatPtrSliceAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[]*float64",
		Alias: "FloatPtrSliceAlias"}},
}

type IntAliasSliceAlias []IntAlias

var IntAliasSliceAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "IntAliasSliceAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[]IntAlias",
		Alias: "IntAliasSliceAlias"}},
}

type Uint64PtrAliasSliceAlias []*Uint64Alias

var Uint64PtrAliasSliceAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Uint64PtrAliasSliceAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[]*Uint64Alias",
		Alias: "Uint64PtrAliasSliceAlias"}},
}

type BoolSliceSliceAlias [][]bool

var BoolSliceSliceAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "BoolSliceSliceAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[][]bool",
		Alias: "BoolSliceSliceAlias"}},
}

type ByteArraySliceAlias [][2]byte

var ByteArraySliceAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "ByteArraySliceAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[][2]byte",
		Alias: "ByteArraySliceAlias"}},
}

type FloatArrayPtrSliceAlias []*[2]float32

var FloatArrayPtrSliceAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "FloatArrayPtrSliceAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[]*[2]float32",
		Alias: "FloatArrayPtrSliceAlias"}},
}

type IntStrMapSliceAlias []map[int16]string

var IntStrMapSliceAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "IntStrMapSliceAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[]map-0[int16]-0string",
		Alias: "IntStrMapSliceAlias"}},
}

type Uint32Int32MapPtrSliceAlias []*map[uint32]int32

var Uint32Int32MapPtrSliceAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Uint32Int32MapPtrSliceAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[]*map-0[uint32]-0int32",
		Alias: "Uint32Int32MapPtrSliceAlias"}},
}

type StructTypeSliceAlias []SimpleStructType

var StructTypeSliceAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "StructTypeSliceAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[]SimpleStructType",
		Alias: "StructTypeSliceAlias"}},
}

type StructTypePtrSliceAlias []*SimpleStructType

var StructTypePtrSliceAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "StructTypePtrSliceAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[]*SimpleStructType",
		Alias: "StructTypePtrSliceAlias"}},
}

type TrickySliceAlias [][2]map[*[]StringAlias]map[SimpleStructType][]int

var TrickySliceAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "TrickySliceAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[][2]map-1[*[]StringAlias]-1map-0[SimpleStructType]-0[]int",
		Alias: "TrickySliceAlias"}},
}

type ValidUintSliceAlias []uint

var ValidUintSliceAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "ValidUintSliceAlias",
	Fields: []musgen.FieldDesc{{
		Name:          "",
		Type:          "[]uint",
		Validator:     "ValidUintSliceAliasSumBiggerThanTen",
		MaxLength:     3,
		ElemValidator: "BiggerThanTenUint",
		Alias:         "ValidUintSliceAlias"}},
}
