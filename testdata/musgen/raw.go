package musgen

import "github.com/ymz-ncnk/musgen"

type Uint64RawAlias uint64

var Uint64RawAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Uint64RawAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "uint64",
		Alias: "Uint64RawAlias", Encoding: "raw"}},
}

type Uint32RawAlias uint32

var Uint32RawAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Uint32RawAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "uint32",
		Alias: "Uint32RawAlias", Encoding: "raw"}},
}

type Uint16RawAlias uint16

var Uint16RawAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Uint16RawAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "uint16",
		Alias: "Uint16RawAlias", Encoding: "raw"}},
}

type Uint8RawAlias uint8

var Uint8RawAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Uint8RawAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "uint8",
		Alias: "Uint8RawAlias", Encoding: "raw"}},
}

type UintRawAlias uint

var UintRawAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "UintRawAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "uint",
		Alias: "UintRawAlias", Encoding: "raw"}},
}

type Int64RawAlias int64

var Int64RawAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Int64RawAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "int64",
		Alias: "Int64RawAlias", Encoding: "raw"}},
}

type Int32RawAlias int32

var Int32RawAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Int32RawAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "int32",
		Alias: "Int32RawAlias", Encoding: "raw"}},
}

type Int16RawAlias int16

var Int16RawAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Int16RawAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "int16",
		Alias: "Int16RawAlias", Encoding: "raw"}},
}

type Int8RawAlias int8

var Int8RawAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Int8RawAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "int8",
		Alias: "Int8RawAlias", Encoding: "raw"}},
}

type IntRawAlias int

var IntRawAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "IntRawAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "int",
		Alias: "IntRawAlias", Encoding: "raw"}},
}

type IntRawArrayAlias [3]int

var IntRawArrayAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "IntRawArrayAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[3]int",
		Alias: "IntRawArrayAlias", ElemEncoding: "raw"}},
}

type Uint16Int32RawMapAlias map[uint16]int32

var Uint16Int32RawMapAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "Uint16Int32RawMapAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "map-0[uint16]-0int32",
		Alias: "Uint16Int32RawMapAlias", KeyEncoding: "raw", ElemEncoding: "raw"}},
}

type IntRawPtrPtrPtrAliasSliceAlias []***int

var IntRawPtrPtrPtrAliasSliceAliasTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "IntRawPtrPtrPtrAliasSliceAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "[]***int",
		Alias: "IntRawPtrPtrPtrAliasSliceAlias", ElemEncoding: "raw"}},
}

type RawStructType struct {
	UintRaw      uint             `mus:"BiggerThanTenUint#raw"`
	UintRaw16    uint             `mus:"#raw"`
	IntRawPtrPtr **int            `mus:"#raw"`
	MapRawPtr    *map[*int16]int8 `mus:",,#raw,BiggerThanTenInt16Ptr#raw"`
	SliceRaw     []uint           `mus:",,BiggerThanTenUint#raw"`
}

var RawStructTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "RawStructType",
	Fields: []musgen.FieldDesc{
		{Name: "UintRaw", Type: "uint", Validator: "BiggerThanTenUint",
			Encoding: "raw"},
		{Name: "UintRaw16", Type: "uint", Encoding: "raw"},
		{Name: "IntRawPtrPtr", Type: "**int", Encoding: "raw"},
		{Name: "MapRawPtr", Type: "*map-0[*int16]-0int8", ElemEncoding: "raw",
			KeyValidator: "BiggerThanTenInt16Ptr", KeyEncoding: "raw"},
		{Name: "SliceRaw", Type: "[]uint", ElemValidator: "BiggerThanTenUint",
			ElemEncoding: "raw"},
	},
}

type ValidInt32RawAlias int32

var ValidInt32RawTypeDesc musgen.TypeDesc = musgen.TypeDesc{
	Package: "musgen",
	Name:    "ValidInt32RawAlias",
	Fields: []musgen.FieldDesc{{Name: "", Type: "int32",
		Alias: "ValidInt32RawAlias", Validator: "PositiveValidInt32AliasRaw",
		Encoding: "raw"}},
}
