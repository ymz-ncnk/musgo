package musgen

import musgen_mod "github.com/ymz-ncnk/musgen/v2"

type StrSliceAlias []string

var StrSliceAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "StrSliceAlias",
	Fields: []musgen_mod.FieldDesc{{Name: "", Type: "[]string",
		Alias: "StrSliceAlias"}},
}

type FloatPtrSliceAlias []*float64

var FloatPtrSliceAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "FloatPtrSliceAlias",
	Fields: []musgen_mod.FieldDesc{{Name: "", Type: "[]*float64",
		Alias: "FloatPtrSliceAlias"}},
}

type IntAliasSliceAlias []IntAlias

var IntAliasSliceAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "IntAliasSliceAlias",
	Fields: []musgen_mod.FieldDesc{{Name: "", Type: "[]IntAlias",
		Alias: "IntAliasSliceAlias"}},
}

type Uint64PtrAliasSliceAlias []*Uint64Alias

var Uint64PtrAliasSliceAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "Uint64PtrAliasSliceAlias",
	Fields: []musgen_mod.FieldDesc{{Name: "", Type: "[]*Uint64Alias",
		Alias: "Uint64PtrAliasSliceAlias"}},
}

type BoolSliceSliceAlias [][]bool

var BoolSliceSliceAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "BoolSliceSliceAlias",
	Fields: []musgen_mod.FieldDesc{{Name: "", Type: "[][]bool",
		Alias: "BoolSliceSliceAlias"}},
}

type ByteArraySliceAlias [][2]byte

var ByteArraySliceAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "ByteArraySliceAlias",
	Fields: []musgen_mod.FieldDesc{{Name: "", Type: "[][2]byte",
		Alias: "ByteArraySliceAlias"}},
}

type FloatArrayPtrSliceAlias []*[2]float32

var FloatArrayPtrSliceAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "FloatArrayPtrSliceAlias",
	Fields: []musgen_mod.FieldDesc{{Name: "", Type: "[]*[2]float32",
		Alias: "FloatArrayPtrSliceAlias"}},
}

type IntStrMapSliceAlias []map[int16]string

var IntStrMapSliceAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "IntStrMapSliceAlias",
	Fields: []musgen_mod.FieldDesc{{Name: "", Type: "[]map-0[int16]-0string",
		Alias: "IntStrMapSliceAlias"}},
}

type Uint32Int32MapPtrSliceAlias []*map[uint32]int32

var Uint32Int32MapPtrSliceAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "Uint32Int32MapPtrSliceAlias",
	Fields: []musgen_mod.FieldDesc{{Name: "", Type: "[]*map-0[uint32]-0int32",
		Alias: "Uint32Int32MapPtrSliceAlias"}},
}

type StructTypeSliceAlias []SimpleStructType

var StructTypeSliceAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "StructTypeSliceAlias",
	Fields: []musgen_mod.FieldDesc{{Name: "", Type: "[]SimpleStructType",
		Alias: "StructTypeSliceAlias"}},
}

type StructTypePtrSliceAlias []*SimpleStructType

var StructTypePtrSliceAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "StructTypePtrSliceAlias",
	Fields: []musgen_mod.FieldDesc{{Name: "", Type: "[]*SimpleStructType",
		Alias: "StructTypePtrSliceAlias"}},
}

type TrickySliceAlias [][2]map[*[]StringAlias]map[SimpleStructType][]int

var TrickySliceAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "TrickySliceAlias",
	Fields: []musgen_mod.FieldDesc{
		{
			Type:  "[][2]map-1[*[]StringAlias]-1map-0[SimpleStructType]-0[]int",
			Alias: "TrickySliceAlias",
		},
	},
}

type ValidUintSliceAlias []uint

var ValidUintSliceAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "ValidUintSliceAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:          "[]uint",
		Validator:     "ValidUintSliceAliasSumBiggerThanTen",
		MaxLength:     3,
		ElemValidator: "BiggerThanTenUint",
		Alias:         "ValidUintSliceAlias"}},
}

type ValidPtrStringSliceAlias []*string

var ValidPtrStringSliceAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "ValidPtrStringSliceAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:          "[]*string",
		ElemValidator: "NotNilString",
		Alias:         "ValidPtrStringSliceAlias"}},
}
