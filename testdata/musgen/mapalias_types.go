package musgen

import musgen_mod "github.com/ymz-ncnk/musgen/v2"

type StrIntMapAlias map[string]int

var StrIntMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "StrIntMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "map-0[string]-0int",
		Alias: "StrIntMapAlias",
	}},
}

type StrPtrIntPtrMapAlias map[*string]*int

var StrPtrIntPtrMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "StrPtrIntPtrMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "map-0[*string]-0*int",
		Alias: "StrPtrIntPtrMapAlias",
	}},
}

type StrAliasIntAliasMapAlias map[StringAlias]IntAlias

var StrAliasIntAliasMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "StrAliasIntAliasMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "map-0[StringAlias]-0IntAlias",
		Alias: "StrAliasIntAliasMapAlias",
	}},
}

type StrAliasPtrIntAliasPtrMapAlias map[*StringAlias]*IntAlias

var StrAliasPtrIntAliasPtrMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "StrAliasPtrIntAliasPtrMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "map-0[*StringAlias]-0*IntAlias",
		Alias: "StrAliasPtrIntAliasPtrMapAlias",
	}},
}

type BoolInt16SliceMapAlias map[bool][]int16

var BoolInt16SliceMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "BoolInt16SliceMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "map-0[bool]-0[]int16",
		Alias: "BoolInt16SliceMapAlias",
	}},
}

type ByteUint16SlicePtrMapAlias map[byte]*[]uint16

var ByteUint16SlicePtrMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "ByteUint16SlicePtrMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "map-0[byte]-0*[]uint16",
		Alias: "ByteUint16SlicePtrMapAlias",
	}},
}

type Int32Float64ArrayMapAlias map[int32][2]float64

var Int32Float64ArrayMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "Int32Float64ArrayMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "map-0[int32]-0[2]float64",
		Alias: "Int32Float64ArrayMapAlias",
	}},
}

type Float32Uint32ArrayPtrMapAlias map[float32]*[2]uint32

var Float32Uint32ArrayPtrMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "Float32Uint32ArrayPtrMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "map-0[float32]-0*[2]uint32",
		Alias: "Float32Uint32ArrayPtrMapAlias",
	}},
}

type FloatByteBoolMapMapAlias map[float32]map[byte]bool

var FloatByteBoolMapMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "FloatByteBoolMapMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "map-1[float32]-1map-0[byte]-0bool",
		Alias: "FloatByteBoolMapMapAlias",
	}},
}

type UintIntStringMapPtrMapAlias map[uint16]*map[int]string

var UintIntStringMapPtrMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "UintIntStringMapPtrMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "map-1[uint16]-1*map-0[int]-0string",
		Alias: "UintIntStringMapPtrMapAlias",
	}},
}

type StructTypeStructTypeMapAlias map[SimpleStructType]SimpleStructType

var StructTypeStructTypeMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "StructTypeStructTypeMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "map-0[SimpleStructType]-0SimpleStructType",
		Alias: "StructTypeStructTypeMapAlias",
	}},
}

type StructTypePtrStructTypePtrMapAlias map[*SimpleStructType]*SimpleStructType

var StructTypePtrStructTypePtrMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "StructTypePtrStructTypePtrMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "map-0[*SimpleStructType]-0*SimpleStructType",
		Alias: "StructTypePtrStructTypePtrMapAlias",
	}},
}

type TrickyMapAlias map[[2]StringAlias]map[*[]SimpleStructType]map[*map[int]string][2]int

var TrickyMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "TrickyMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "map-3[[2]StringAlias]-3map-2[*[]SimpleStructType]-2map-1[*map-0[int]-0string]-1[2]int",
		Alias: "TrickyMapAlias",
	}},
}

type ValidStringIntMapAlias map[string]int

var ValidStringIntMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "ValidStringIntMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:          "map-0[string]-0int",
		Validator:     "ValidStringIntMapAliasSumBiggerThanTen",
		MaxLength:     3,
		ElemValidator: "BiggerThanTenInt",
		KeyValidator:  "StrIsHello",
		Alias:         "ValidStringIntMapAlias"}},
}

type ValidPtrIntPtrIntMapAlias map[*int]*int

var ValidPtrIntPtrIntMapAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "ValidPtrIntPtrIntMapAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:          "map-0[*int]-0*int",
		ElemValidator: "NotNilInt",
		KeyValidator:  "NotNilInt",
		Alias:         "ValidPtrIntPtrIntMapAlias"}},
}
