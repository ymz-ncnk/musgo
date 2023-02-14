package musgen

import musgen_mod "github.com/ymz-ncnk/musgen/v2"

type Uint64Alias uint64

var Uint64AliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "Uint64Alias",
	Fields: []musgen_mod.FieldDesc{{
		Type:      "uint64",
		MaxLength: 9,
		Alias:     "Uint64Alias",
	}},
}

type Uint32Alias uint32

var Uint32AliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "Uint32Alias",
	Fields: []musgen_mod.FieldDesc{{
		Type:      "uint32",
		MaxLength: 5,
		Alias:     "Uint32Alias",
	}},
}

type Uint16Alias uint16

var Uint16AliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "Uint16Alias",
	Fields: []musgen_mod.FieldDesc{{
		Type:      "uint16",
		MaxLength: 3,
		Alias:     "Uint16Alias",
	}},
}

type Uint8Alias uint8

var Uint8AliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "Uint8Alias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "uint8",
		Alias: "Uint8Alias",
	}},
}

type UintAlias uint

var UintAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "UintAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "uint",
		Alias: "UintAlias",
	}},
}

type Int64Alias int64

var Int64AliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "Int64Alias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "int64",
		Alias: "Int64Alias",
	}},
}

type Int32Alias int32

var Int32AliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "Int32Alias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "int32",
		Alias: "Int32Alias",
	}},
}

type Int16Alias int16

var Int16AliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "Int16Alias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "int16",
		Alias: "Int16Alias",
	}},
}

type Int8Alias int8

var Int8AliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "Int8Alias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "int8",
		Alias: "Int8Alias",
	}},
}

type IntAlias int

var IntAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "IntAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "int",
		Alias: "IntAlias",
	}},
}

type Float64Alias float64

var Float64AliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "Float64Alias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "float64",
		Alias: "Float64Alias",
	}},
}

type Float32Alias float32

var Float32AliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "Float32Alias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "float32",
		Alias: "Float32Alias",
	}},
}

type BoolAlias bool

var BoolAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "BoolAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "bool",
		Alias: "BoolAlias",
	}},
}

type ByteAlias byte

var ByteAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "ByteAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "uint8",
		Alias: "ByteAlias",
	}},
}

type StringAlias string

var StringAliasTypeDesc musgen_mod.TypeDesc = musgen_mod.TypeDesc{
	Package: "musgen",
	Name:    "StringAlias",
	Fields: []musgen_mod.FieldDesc{{
		Type:  "string",
		Alias: "StringAlias",
	}},
}
