package musgen

import musgenmod "github.com/ymz-ncnk/musgen"

type SimpleStructType struct {
	Int int
}

var SimpleStructTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "SimpleStructType",
	Fields: []musgenmod.FieldDesc{
		{Name: "Int", Type: "int"},
	},
}

// `-, validator, elemValidator, maxLength`

type StructType struct {
	Uint          uint
	Uint16        uint16
	Uint32        uint32
	Uint64        uint64
	UintPtr       *uint
	UintPtrPtrPtr ***uint

	Int          int
	Int8         int8
	Int16        int16
	Int32        int32
	Int64        int64
	IntPtr       *int
	IntPtrPtrPtr ***int

	String          string
	StringPtr       *string
	StringPtrPtrPtr ***string
	Byte            byte
	BytePtr         *byte
	BytePtrPtrPtr   ***byte
	Bool            bool
	BoolPtr         *bool
	BoolPtrPtrPtr   ***bool

	Slice          []uint
	SlicePtr       *[]uint16
	SlicePtrPtrPtr ***[]uint16

	Array          [2]int
	ArrayPtr       *[2]int32
	ArrayPtrPtrPtr ***[2]int32

	Map          map[string]int
	MapPtr       *map[string]int
	MapPtrPtrPtr ***map[string]int

	Struct          SimpleStructType
	StructPtr       *SimpleStructType
	StructPtrPtrPtr ***SimpleStructType

	Tricky [2]map[[2]IntAlias]map[StringAlias][2]string
}

var StructTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "StructType",
	Fields: []musgenmod.FieldDesc{
		{Name: "Uint", Type: "uint"},
		{Name: "Uint16", Type: "uint16"},
		{Name: "Uint32", Type: "uint32"},
		{Name: "Uint64", Type: "uint64"},
		{Name: "UintPtr", Type: "*uint"},
		{Name: "UintPtrPtrPtr", Type: "***uint"},

		{Name: "Int", Type: "int"},
		{Name: "Int8", Type: "int8"},
		{Name: "Int16", Type: "int16"},
		{Name: "Int32", Type: "int32"},
		{Name: "Int64", Type: "int64"},
		{Name: "IntPtr", Type: "*int"},
		{Name: "IntPtrPtrPtr", Type: "***int"},

		{Name: "String", Type: "string"},
		{Name: "StringPtr", Type: "*string"},
		{Name: "StringPtrPtrPtr", Type: "***string"},

		{Name: "Byte", Type: "uint8"},
		{Name: "BytePtr", Type: "*uint8"},
		{Name: "BytePtrPtrPtr", Type: "***uint8"},
		{Name: "Bool", Type: "bool"},
		{Name: "BoolPtr", Type: "*bool"},
		{Name: "BoolPtrPtrPtr", Type: "***bool"},

		{Name: "Slice", Type: "[]uint"},
		{Name: "SlicePtr", Type: "*[]uint16"},
		{Name: "SlicePtrPtrPtr", Type: "***[]uint16"},

		{Name: "Array", Type: "[2]int"},
		{Name: "ArrayPtr", Type: "*[2]int32"},
		{Name: "ArrayPtrPtrPtr", Type: "***[2]int32"},

		{Name: "Map", Type: "map-0[string]-0int"},
		{Name: "MapPtr", Type: "*map-0[string]-0int"},
		{Name: "MapPtrPtrPtr", Type: "***map-0[string]-0int"},

		{Name: "Struct", Type: "SimpleStructType"},
		{Name: "StructPtr", Type: "*SimpleStructType"},
		{Name: "StructPtrPtrPtr", Type: "***SimpleStructType"},

		{Name: "Tricky", Type: "[2]map-1[[2]IntAlias]-1map-0[StringAlias]-0[2]string"},
	},
}

type ValidStructType struct {
	Uint64 uint64

	Int8 int8

	String string
	Byte   byte
	Bool   bool

	Slice    []uint
	SlicePtr *[]uint16

	Array    [2]int
	ArrayPtr *[2]int32

	Map    map[string]int
	MapPtr *map[string]int

	Struct    SimpleStructType
	StructPtr *SimpleStructType
}

var ValidStructTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "ValidStructType",
	Fields: []musgenmod.FieldDesc{
		{Name: "Uint64", Type: "uint64", Validator: "BiggerThanTenUint64"},

		{Name: "Int8", Type: "int8", Validator: "BiggerThanTenInt8"},

		{Name: "String", Type: "string", Validator: "NotEmptyString", MaxLength: 10},
		{Name: "Byte", Type: "uint8", Validator: "BiggerThanTenByte"},
		{Name: "Bool", Type: "bool", Validator: "PositiveBool"},

		{Name: "Slice", Type: "[]uint", ElemValidator: "BiggerThanTenUint",
			Validator: "UintSliceSumBiggerThanTen", MaxLength: 8},
		{Name: "SlicePtr", Type: "*[]uint16", ElemValidator: "BiggerThanTenUint16",
			Validator: "Uint16SlicePtrSumBiggerThanTen"},

		{Name: "Array", Type: "[2]int", ElemValidator: "BiggerThanTenInt",
			Validator: "IntArraySumBiggerThanTen"},
		{Name: "ArrayPtr", Type: "*[2]int32", ElemValidator: "BiggerThanTenInt32",
			Validator: "Int32ArrayPtrSumBiggerThanTen"},

		{Name: "Map", Type: "map-0[string]-0int", KeyValidator: "StrIsHello",
			ElemValidator: "BiggerThanTenInt", Validator: "MapSumBiggerThanTen",
			MaxLength: 5},
		{Name: "MapPtr", Type: "*map-0[string]-0int", KeyValidator: "StrIsHello",
			ElemValidator: "BiggerThanTenInt", Validator: "MapPtrSumBiggerThanTen",
			MaxLength: 4},

		{Name: "Struct", Type: "SimpleStructType",
			Validator: "ValidSimpleStructType"},
		{Name: "StructPtr", Type: "*SimpleStructType",
			Validator: "ValidSimpleStructPtrType"},
	},
}

type FieldlessStructType struct {
	Number int
	Slice  []int
}

var FieldlessStructTypeDesc musgenmod.TypeDesc = musgenmod.TypeDesc{
	Package: "musgen",
	Name:    "FieldlessStructType",
	Fields:  []musgenmod.FieldDesc{},
}
