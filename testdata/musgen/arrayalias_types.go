package musgen

import (
	musgenmod "github.com/ymz-ncnk/musgen"
)

type StrArrayAlias [3]string

var StrArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "StrArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:  "[3]string",
		Alias: "StrArrayAlias",
	}},
}

type FloatPtrArrayAlias [3]*float64

var FloatPtrArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "FloatPtrArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:  "[3]*float64",
		Alias: "FloatPtrArrayAlias",
	}},
}

type IntAliasArrayAlias [3]IntAlias

var IntAliasArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "IntAliasArrayAlias",
	Fields: []musgenmod.FieldDesc{{Name: "", Type: "[3]IntAlias",
		Alias: "IntAliasArrayAlias"}},
}

type Uint64PtrAliasArrayAlias [3]*Uint64Alias

var Uint64PtrAliasArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "Uint64PtrAliasArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:  "[3]*Uint64Alias",
		Alias: "Uint64PtrAliasArrayAlias",
	}},
}

type Int64SliceArrayAlias [3][]int64

var Int64SliceArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "Int64SliceArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:  "[3][]int64",
		Alias: "Int64SliceArrayAlias",
	}},
}

type Uint16SlicePtrArrayAlias [3]*[]uint16

var Uint16SlicePtrArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "Uint16SlicePtrArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:  "[3]*[]uint16",
		Alias: "Uint16SlicePtrArrayAlias",
	}},
}

type BoolArrayArrayAlias [3][3]bool

var BoolArrayArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "BoolArrayArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:  "[3][3]bool",
		Alias: "BoolArrayArrayAlias",
	}},
}

type BytePtrArrayPtrArrayAlias [2]*[2]byte

var BytePtrArrayPtrArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "BytePtrArrayPtrArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:  "[2]*[2]uint8",
		Alias: "BytePtrArrayPtrArrayAlias",
	}},
}

type IntStrMapArrayAlias [3]map[int]string

var IntStrMapArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "IntStrMapArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:  "[3]map-0[int]-0string",
		Alias: "IntStrMapArrayAlias",
	}},
}

type Uint32Int32MapArrayAlias [4]*map[uint32]int32

var Uint32Int32MapArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "Uint32Int32MapArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:  "[4]*map-0[uint32]-0int32",
		Alias: "Uint32Int32MapArrayAlias",
	}},
}

type StructTypeArrayAlias [3]SimpleStructType

var StructTypeArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "StructTypeArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:  "[3]SimpleStructType",
		Alias: "StructTypeArrayAlias",
	}},
}

type StructTypePtrArrayAlias [3]*SimpleStructType

var StructTypePtrArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "StructTypePtrArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:  "[3]*SimpleStructType",
		Alias: "StructTypePtrArrayAlias",
	}},
}

type TrickyArrayAlias [2][][3]map[SimpleStructType][1]int

var TrickyArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "TrickyArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:  "[2][][3]map-0[SimpleStructType]-0[1]int",
		Alias: "TrickyArrayAlias",
	}},
}

type StrZeroLengthArrayAlias [0]string

var StrZeroLengthArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "StrZeroLengthArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:  "[0]string",
		Alias: "StrZeroLengthArrayAlias",
	}},
}

type IntPtrPtrPtrAliasArrayAlias [3]***int

var IntPtrPtrPtrAliasArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "IntPtrPtrPtrAliasArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:  "[3]***int",
		Alias: "IntPtrPtrPtrAliasArrayAlias",
	}},
}

type ValidIntArrayAlias [2]int

var ValidIntArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "ValidIntArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:          "[2]int",
		Validator:     "ValidIntArrayAliasSumBiggerThanTen",
		ElemValidator: "BiggerThanTenInt", Alias: "ValidIntArrayAlias",
	}},
}

type ValidPtrIntArrayAlias [2]*int

var ValidPtrIntArrayAliasTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "ValidPtrIntArrayAlias",
	Fields: []musgenmod.FieldDesc{{
		Type:          "[2]*int",
		ElemValidator: "NotNilInt", Alias: "ValidPtrIntArrayAlias",
	}},
}
