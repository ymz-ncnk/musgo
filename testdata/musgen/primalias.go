package musgen

import "github.com/ymz-ncnk/musgen"

type Uint64Alias uint64

var Uint64AliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Uint64Alias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "uint64", MaxLength: 9,
		Alias: "Uint64Alias"}},
}

type Uint32Alias uint32

var Uint32AliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Uint32Alias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "uint32", MaxLength: 5,
		Alias: "Uint32Alias"}},
}

type Uint16Alias uint16

var Uint16AliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Uint16Alias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "uint16", MaxLength: 3,
		Alias: "Uint16Alias"}},
}

type Uint8Alias uint8

var Uint8AliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Uint8Alias",
	Fields:  []musgen.FieldDesc{{Name: "", Type: "uint8", Alias: "Uint8Alias"}},
}

type UintAlias uint

var UintAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "UintAlias",
	Fields:  []musgen.FieldDesc{{Name: "", Type: "uint", Alias: "UintAlias"}},
}

type Int64Alias int64

var Int64AliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Int64Alias",
	Fields:  []musgen.FieldDesc{{Name: "", Type: "int64", Alias: "Int64Alias"}},
}

type Int32Alias int32

var Int32AliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Int32Alias",
	Fields:  []musgen.FieldDesc{{Name: "", Type: "int32", Alias: "Int32Alias"}},
}

type Int16Alias int16

var Int16AliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Int16Alias",
	Fields:  []musgen.FieldDesc{{Name: "", Type: "int16", Alias: "Int16Alias"}},
}

type Int8Alias int8

var Int8AliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Int8Alias",
	Fields:  []musgen.FieldDesc{{Name: "", Type: "int8", Alias: "Int8Alias"}},
}

type IntAlias int

var IntAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "IntAlias",
	Fields:  []musgen.FieldDesc{{Name: "", Type: "int", Alias: "IntAlias"}},
}

type Float64Alias float64

var Float64AliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Float64Alias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "float64",
		Alias: "Float64Alias"}},
}

type Float32Alias float64

var Float32AliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Float32Alias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "float32",
		Alias: "Float32Alias"}},
}

type BoolAlias bool

var BoolAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "BoolAlias",
	Fields:  []musgen.FieldDesc{{Name: "", Type: "bool", Alias: "BoolAlias"}},
}

type ByteAlias byte

var ByteAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "ByteAlias",
	Fields:  []musgen.FieldDesc{{Name: "", Type: "uint8", Alias: "ByteAlias"}},
}

type StringAlias string

var StringAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "StringAlias",
	Fields:  []musgen.FieldDesc{{Name: "", Type: "string", Alias: "StringAlias"}},
}
