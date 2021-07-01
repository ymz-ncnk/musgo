package musgo

// //go:generate go run testdata/make/musable.go -struct $ARG
// package musgo

// import (
// 	"fmt"
// 	"reflect"
// 	"testing"

// 	"github.com/ymz-ncnk/musgo/errs"
// 	mgtd "github.com/ymz-ncnk/musgo/testdata/musgen"
// 	utils "github.com/ymz-ncnk/musgo/testdata/musgen/utils"
// )

// func TestGeneratedStructCode(t *testing.T) {
// 	var err error
// 	// struct
// 	{
// 		typeName := mgtd.StructTypeDesc.Name
// 		var un uint = 51
// 		un1 := &un
// 		un2 := &un1
// 		var n int = 43
// 		n1 := &n
// 		n2 := &n1
// 		var str = "strte"
// 		str1 := &str
// 		str2 := &str1
// 		var b byte = 0x09
// 		b1 := &b
// 		b2 := &b1
// 		var bl bool = true
// 		bl1 := &bl
// 		bl2 := &bl1
// 		var sl []uint16 = []uint16{1123, 0}
// 		sl1 := &sl
// 		sl2 := &sl1
// 		var ar [2]int32 = [2]int32{123123123, 10}
// 		ar1 := &ar
// 		ar2 := &ar1
// 		var m map[string]int = map[string]int{"world": 20}
// 		m1 := &m
// 		m2 := &m1
// 		var st mgtd.SimpleStructType = mgtd.SimpleStructType{Int: 99}
// 		st1 := &st
// 		st2 := &st1
// 		for _, num := range []mgtd.StructType{
// 			{
// 				Uint:          12,
// 				Uint16:        12032,
// 				Uint32:        123912913,
// 				Uint64:        1231233333333333333,
// 				UintPtr:       &un,
// 				UintPtrPtrPtr: &un2,

// 				Int:          -1123,
// 				Int16:        0,
// 				Int32:        1231231231,
// 				Int64:        -12309823487234892,
// 				IntPtr:       &n,
// 				IntPtrPtrPtr: &n2,

// 				String:          "some antoherlkjalkja setlwkethwleh lkjwelkjwelr",
// 				StringPtr:       &str,
// 				StringPtrPtrPtr: &str2,

// 				Byte:          0x10,
// 				BytePtr:       &b,
// 				BytePtrPtrPtr: &b2,
// 				Bool:          true,
// 				BoolPtr:       &bl,
// 				BoolPtrPtrPtr: &bl2,

// 				Slice:          []uint{123, 345},
// 				SlicePtr:       &sl,
// 				SlicePtrPtrPtr: &sl2,

// 				Array:          [2]int{-123, -345},
// 				ArrayPtr:       &ar,
// 				ArrayPtrPtrPtr: &ar2,

// 				Map:          map[string]int{"hello": 5},
// 				MapPtr:       &m,
// 				MapPtrPtrPtr: &m2,

// 				Struct:          mgtd.SimpleStructType{Int: 28},
// 				StructPtr:       &st,
// 				StructPtrPtrPtr: &st2,

// 				Tricky: [2]map[[2]mgtd.IntAlias]map[mgtd.StringAlias][2]string{
// 					{[2]mgtd.IntAlias{1, 2}: {"some": [2]string{"s", "r"}}},
// 					{[2]mgtd.IntAlias{150, 2222}: {"an": [2]string{"erer", "qqqpq"}}},
// 				},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.StructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if initVal.Uint != zeroVal.Uint ||
// 					initVal.Uint16 != zeroVal.Uint16 ||
// 					initVal.Uint32 != zeroVal.Uint32 ||
// 					initVal.Uint64 != zeroVal.Uint64 ||
// 					*initVal.UintPtr != *zeroVal.UintPtr ||

// 					initVal.Int != zeroVal.Int ||
// 					initVal.Int16 != zeroVal.Int16 ||
// 					initVal.Int32 != zeroVal.Int32 ||
// 					initVal.Int64 != zeroVal.Int64 ||
// 					*initVal.IntPtr != *zeroVal.IntPtr ||

// 					initVal.String != zeroVal.String ||
// 					*initVal.StringPtr != *zeroVal.StringPtr ||
// 					initVal.Byte != zeroVal.Byte ||
// 					*initVal.BytePtr != *zeroVal.BytePtr ||
// 					initVal.Bool != zeroVal.Bool ||
// 					*initVal.BoolPtr != *zeroVal.BoolPtr ||

// 					!reflect.DeepEqual(initVal.Slice, zeroVal.Slice) ||
// 					!reflect.DeepEqual(initVal.SlicePtr, zeroVal.SlicePtr) ||

// 					!reflect.DeepEqual(initVal.Array, zeroVal.Array) ||
// 					!reflect.DeepEqual(initVal.ArrayPtr, zeroVal.ArrayPtr) ||

// 					!reflect.DeepEqual(initVal.Map, zeroVal.Map) ||
// 					!reflect.DeepEqual(initVal.MapPtr, zeroVal.MapPtr) ||

// 					!reflect.DeepEqual(initVal.Struct, zeroVal.Struct) ||
// 					!reflect.DeepEqual(initVal.StructPtr, zeroVal.StructPtr) ||

// 					!reflect.DeepEqual(initVal.Tricky, zeroVal.Tricky) {
// 					t.Error(fmt.Errorf(utils.MuErrMsg, typeName))
// 				}
// 			}
// 		}
// 	}
// 	// struct, buffer ends
// 	{
// 		var val mgtd.StructType
// 		err := utils.TestBufferEnds(&val, []byte{})
// 		if err != nil {
// 			t.Error(err)
// 		}
// 	}
// }

// // Validation
// func TestGeneratedValidStructCode(t *testing.T) {
// 	var err error
// 	// uint validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				Uint64:    120,
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "Uint64", mgtd.ErrBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// int validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				Int8:      19,
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "Int8", mgtd.ErrBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// string validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				String:    "hello w",
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "String", mgtd.ErrNotEmptyString)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// byte validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				Byte:      21,
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "Byte", mgtd.ErrBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// byte validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				Bool:      true,
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "Bool", mgtd.ErrPositiveBool)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// slice validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				Slice:     []uint{7, 4},
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "Slice", mgtd.ErrSliceSumBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// slice elem validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				Slice:     []uint{5, 11},
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testSliceElemErr(err, "Slice", 1, mgtd.ErrBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// slice ptr validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				SlicePtr:  &[]uint16{8, 8},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "SlicePtr", mgtd.ErrSliceSumBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// slice ptr elem validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				SlicePtr:  &[]uint16{8, 11, 8},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testSliceElemErr(err, "SlicePtr", 1, mgtd.ErrBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// array validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				Array:     [2]int{7, 4},
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "Array", mgtd.ErrArraySumBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// array elem validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				Array:     [2]int{5, 11},
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testSliceElemErr(err, "Array", 1, mgtd.ErrBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// array ptr validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{4, 8},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "ArrayPtr", mgtd.ErrArraySumBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// array ptr elem validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{12, 8},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testArrayElemErr(err, "ArrayPtr", 0, mgtd.ErrBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// map validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				Map:       map[string]int{"str1": 5, "str2": 10},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "Map", mgtd.ErrMapSumBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// map key validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				Map:       map[string]int{"": 5, "hello": 9, "rt": 1},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testMapKeyErr(err, "Map", "hello", mgtd.ErrStrIsHello)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// map value validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				Map:       map[string]int{"": 5, "not hello": 10, "rt": 88},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testMapValueErr(err, "Map", "rt", 88, mgtd.ErrBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// map ptr validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{"str1": 5, "str2": 10},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "MapPtr", mgtd.ErrMapSumBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// map ptr key validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{"": 5, "hello": 10, "rt": 1},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testMapKeyErr(err, "MapPtr", "hello", mgtd.ErrStrIsHello)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// map ptr value validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{"": 5, "not hello": 10, "rt": 88},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testMapValueErr(err, "MapPtr", "rt", 88, mgtd.ErrBiggerThanTen)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// struct validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				Struct:    mgtd.SimpleStructType{Int: 15},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "Struct", mgtd.ErrSimpleStructType)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// struct ptr validator
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{Int: 15},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "StructPtr", mgtd.ErrSimpleStructType)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// struct string maxLength
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				String:    "qwertyuiopa",
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "String", errs.ErrMaxLengthExceeded)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// struct slice maxLength
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				Slice:     []uint{1, 1, 1, 0, 0, 0, 0, 0, 0},
// 				SlicePtr:  &[]uint16{},
// 				ArrayPtr:  &[2]int32{},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "Slice", errs.ErrMaxLengthExceeded)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// struct map maxLength
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				SlicePtr: &[]uint16{},
// 				ArrayPtr: &[2]int32{},
// 				Map: map[string]int{"str1": 0, "str2": 1, "str3": 2, "str4": 3,
// 					"str5": -1, "str6": 0},
// 				MapPtr:    &map[string]int{},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "Map", errs.ErrMaxLengthExceeded)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// struct map ptr maxLength
// 	{
// 		typeName := mgtd.ValidStructTypeDesc.Name
// 		for _, num := range []mgtd.ValidStructType{
// 			{
// 				SlicePtr: &[]uint16{},
// 				ArrayPtr: &[2]int32{},
// 				MapPtr: &map[string]int{"str1": 0, "str2": 0, "str3": -10, "str4": 1,
// 					"str5": -1, "str6": -2},
// 				StructPtr: &mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.ValidStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				err := testValidErr(err, "MapPtr", errs.ErrMaxLengthExceeded)
// 				if err != nil {
// 					t.Error(err)
// 				}
// 			}
// 		}
// 	}
// 	// struct without fields
// 	{
// 		typeName := mgtd.FieldlessStructTypeDesc.Name
// 		for _, num := range []mgtd.FieldlessStructType{
// 			{Number: 5, Slice: []int{1, 2}},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.FieldlessStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error("validation didn't work")
// 			} else {
// 				if zeroVal.Number != 0 {
// 					t.Error("there is a content for FieldlessStructType")
// 				}
// 				if zeroVal.Slice != nil {
// 					t.Error("there is a content for FieldlessStructType")
// 				}
// 			}
// 		}
// 	}
// }

// // -----------------------------------------------------------------------------

// func testValidErr(err error, requiredField string,
// 	requiredCause error) error {
// 	cause, err := fieldErrorCause(err, requiredField, requiredCause)
// 	if err != nil {
// 		return err
// 	}
// 	if cause != requiredCause {
// 		return fmt.Errorf("wrong cause: %v, should be: %v", cause,
// 			requiredCause)
// 	}
// 	return nil
// }

// func testSliceElemErr(err error, requiredField string, requiredElem int,
// 	requiredCause error) error {
// 	fieldErrCause, err := fieldErrorCause(err, requiredField, requiredCause)
// 	if err != nil {
// 		return err
// 	}
// 	sliceErr, ok := fieldErrCause.(errs.SliceError)
// 	if !ok {
// 		return fmt.Errorf("received not SliceError: %v", sliceErr.Cause())
// 	}
// 	if sliceErr.Index() != requiredElem {
// 		return fmt.Errorf("failed wrong element: %v, should be %v", sliceErr.Index(),
// 			requiredElem)
// 	}
// 	if sliceErr.Cause() != requiredCause {
// 		return fmt.Errorf("wrong cause: %v, should be: %v", sliceErr.Cause(),
// 			requiredCause)
// 	}
// 	return nil
// }

// func testArrayElemErr(err error, requiredField string, requiredElem int,
// 	requiredCause error) error {
// 	fieldErrCause, err := fieldErrorCause(err, requiredField, requiredCause)
// 	if err != nil {
// 		return err
// 	}
// 	arrayErr, ok := fieldErrCause.(errs.ArrayError)
// 	if !ok {
// 		return fmt.Errorf("received not arrayError: %v", arrayErr.Cause())
// 	}
// 	if arrayErr.Index() != requiredElem {
// 		return fmt.Errorf("failed wrong element: %v, should be %v", arrayErr.Index(),
// 			requiredElem)
// 	}
// 	if arrayErr.Cause() != requiredCause {
// 		return fmt.Errorf("wrong cause: %v, should be: %v", arrayErr.Cause(),
// 			requiredCause)
// 	}
// 	return nil
// }

// func testMapKeyErr(err error, requiredField string, requiredKey string,
// 	requiredCause error) error {
// 	fieldErrCause, err := fieldErrorCause(err, requiredField, requiredCause)
// 	if err != nil {
// 		return err
// 	}
// 	mapKeyErr, ok := fieldErrCause.(*errs.MapKeyError)
// 	if !ok {
// 		return fmt.Errorf("received not mapKeyError: %v", mapKeyErr.Cause())
// 	}
// 	if mapKeyErr.Key() != requiredKey {
// 		return fmt.Errorf("failed wrong key: %v, should be %v", mapKeyErr.Key(),
// 			requiredKey)
// 	}
// 	if mapKeyErr.Cause() != requiredCause {
// 		return fmt.Errorf("wrong cause: %v, should be: %v", mapKeyErr.Cause(),
// 			requiredCause)
// 	}
// 	return nil
// }

// func testMapValueErr(err error, requiredField string, requiredKey string,
// 	requiredValue int, requiredCause error) error {
// 	fieldErrCause, err := fieldErrorCause(err, requiredField, requiredCause)
// 	if err != nil {
// 		return err
// 	}
// 	mapValueErr, ok := fieldErrCause.(*errs.MapValueError)
// 	if !ok {
// 		return fmt.Errorf("received not mapKeyError: %v", mapValueErr.Cause())
// 	}
// 	if mapValueErr.Key() != requiredKey {
// 		return fmt.Errorf("failed wrong key: %v, should be %v", mapValueErr.Key(),
// 			requiredKey)
// 	}
// 	if mapValueErr.Key() != requiredKey {
// 		return fmt.Errorf("failed wrong key: %v, should be %v", mapValueErr.Key(),
// 			requiredKey)
// 	}
// 	if mapValueErr.Value() != requiredValue {
// 		return fmt.Errorf("failed wrong value: %v, should be %v", mapValueErr.Value(),
// 			requiredValue)
// 	}
// 	if mapValueErr.Cause() != requiredCause {
// 		return fmt.Errorf("wrong cause: %v, should be: %v", mapValueErr.Cause(),
// 			requiredCause)
// 	}
// 	return nil
// }

// func fieldErrorCause(occuredErr error, requiredField string,
// 	requiredCause error) (cause error, err error) {
// 	unmarshalErr, ok := occuredErr.(*utils.GenError)
// 	if !ok {
// 		return nil, fmt.Errorf("received unknown error: %v", occuredErr)
// 	}
// 	var fieldErr errs.FieldError
// 	fieldErr, ok = unmarshalErr.Cause().(errs.FieldError)
// 	if !ok {
// 		return nil, fmt.Errorf("received not FieldError: %v", unmarshalErr.Cause())
// 	}
// 	if fieldErr.FieldName() != requiredField {
// 		return nil, fmt.Errorf("failed wrong field: %v, should be: %v",
// 			fieldErr.FieldName(), requiredField)
// 	}
// 	return fieldErr.Cause(), nil
// }
