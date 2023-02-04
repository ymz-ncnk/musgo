package musgen

import musgenmod "github.com/ymz-ncnk/serialization/musgen"

type StrSliceAlias []string

var StrSliceAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "StrSliceAlias",
	Fields: []musgenmod.FieldDesc{{Name: "", Type: "[]string",
		Alias: "StrSliceAlias"}},
}

type FloatPtrSliceAlias []*float64

var FloatPtrSliceAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "FloatPtrSliceAlias",
	Fields: []musgenmod.FieldDesc{{Name: "", Type: "[]*float64",
		Alias: "FloatPtrSliceAlias"}},
}

type IntAliasSliceAlias []IntAlias

var IntAliasSliceAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "IntAliasSliceAlias",
	Fields: []musgenmod.FieldDesc{{Name: "", Type: "[]IntAlias",
		Alias: "IntAliasSliceAlias"}},
}

type Uint64PtrAliasSliceAlias []*Uint64Alias

var Uint64PtrAliasSliceAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "Uint64PtrAliasSliceAlias",
	Fields: []musgenmod.FieldDesc{{Name: "", Type: "[]*Uint64Alias",
		Alias: "Uint64PtrAliasSliceAlias"}},
}

type BoolSliceSliceAlias [][]bool

var BoolSliceSliceAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "BoolSliceSliceAlias",
	Fields: []musgenmod.FieldDesc{{Name: "", Type: "[][]bool",
		Alias: "BoolSliceSliceAlias"}},
}

type ByteArraySliceAlias [][2]byte

var ByteArraySliceAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "ByteArraySliceAlias",
	Fields: []musgenmod.FieldDesc{{Name: "", Type: "[][2]byte",
		Alias: "ByteArraySliceAlias"}},
}

type FloatArrayPtrSliceAlias []*[2]float32

var FloatArrayPtrSliceAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "FloatArrayPtrSliceAlias",
	Fields: []musgenmod.FieldDesc{{Name: "", Type: "[]*[2]float32",
		Alias: "FloatArrayPtrSliceAlias"}},
}

type IntStrMapSliceAlias []map[int16]string

var IntStrMapSliceAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "IntStrMapSliceAlias",
	Fields: []musgenmod.FieldDesc{{Name: "", Type: "[]map-0[int16]-0string",
		Alias: "IntStrMapSliceAlias"}},
}

type Uint32Int32MapPtrSliceAlias []*map[uint32]int32

var Uint32Int32MapPtrSliceAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "Uint32Int32MapPtrSliceAlias",
	Fields: []musgenmod.FieldDesc{{Name: "", Type: "[]*map-0[uint32]-0int32",
		Alias: "Uint32Int32MapPtrSliceAlias"}},
}

type StructTypeSliceAlias []SimpleStructType

var StructTypeSliceAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "StructTypeSliceAlias",
	Fields: []musgenmod.FieldDesc{{Name: "", Type: "[]SimpleStructType",
		Alias: "StructTypeSliceAlias"}},
}

type StructTypePtrSliceAlias []*SimpleStructType

var StructTypePtrSliceAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "StructTypePtrSliceAlias",
	Fields: []musgenmod.FieldDesc{{Name: "", Type: "[]*SimpleStructType",
		Alias: "StructTypePtrSliceAlias"}},
}

type TrickySliceAlias [][2]map[*[]StringAlias]map[SimpleStructType][]int

var TrickySliceAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "TrickySliceAlias",
	Fields: []musgenmod.FieldDesc{
		{
			Type:  "[][2]map-1[*[]StringAlias]-1map-0[SimpleStructType]-0[]int",
			Alias: "TrickySliceAlias",
		},
	},
}

type ValidUintSliceAlias []uint

var ValidUintSliceAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "ValidUintSliceAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:          "[]uint",
		Validator:     "ValidUintSliceAliasSumBiggerThanTen",
		MaxLength:     3,
		ElemValidator: "BiggerThanTenUint",
		Alias:         "ValidUintSliceAlias"}},
}
