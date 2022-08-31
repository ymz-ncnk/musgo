package musgo

// //go:generate go run testdata/make/musable.go -intraw $ARG
// package musgo

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/ymz-ncnk/musgo/errs"
// 	mgtd "github.com/ymz-ncnk/musgo/testdata/musgen"
// 	utils "github.com/ymz-ncnk/musgo/testdata/musgen/utils"
// )

// func TestGeneratedIntRawEncodingCode(t *testing.T) {
// 	var err error
// 	// uint64
// 	{
// 		typeName := mgtd.Uint64RawAliasTypeDesc.Name
// 		for _, num := range []mgtd.Uint64RawAlias{151492823822937, 113191817} {
// 			initVal := mgtd.Uint64RawAlias(num)
// 			zeroVal := mgtd.Uint64RawAlias(0)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// uint32
// 	{
// 		typeName := mgtd.Uint32RawAliasTypeDesc.Name
// 		for _, num := range []mgtd.Uint32RawAlias{94383726, 1} {
// 			initVal := mgtd.Uint32RawAlias(num)
// 			zeroVal := mgtd.Uint32RawAlias(0)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// uint16
// 	{
// 		typeName := mgtd.Uint16RawAliasTypeDesc.Name
// 		for _, num := range []mgtd.Uint16RawAlias{929, 11111} {
// 			initVal := mgtd.Uint16RawAlias(num)
// 			zeroVal := mgtd.Uint16RawAlias(0)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// uint8
// 	{
// 		typeName := mgtd.Uint8RawAliasTypeDesc.Name
// 		for _, num := range []mgtd.Uint8RawAlias{255, 0} {
// 			initVal := mgtd.Uint8RawAlias(num)
// 			zeroVal := mgtd.Uint8RawAlias(0)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// uint
// 	{
// 		typeName := mgtd.UintRawAliasTypeDesc.Name
// 		for _, num := range []mgtd.UintRawAlias{545349, 11514141414} {
// 			initVal := mgtd.UintRawAlias(num)
// 			zeroVal := mgtd.UintRawAlias(0)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// int64
// 	{
// 		typeName := mgtd.Int64RawAliasTypeDesc.Name
// 		for _, num := range []mgtd.Int64RawAlias{-2937437292, 17620} {
// 			initVal := mgtd.Int64RawAlias(num)
// 			zeroVal := mgtd.Int64RawAlias(0)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// int32
// 	{
// 		typeName := mgtd.Int32RawAliasTypeDesc.Name
// 		for _, num := range []mgtd.Int32RawAlias{-17253, 0} {
// 			initVal := mgtd.Int32RawAlias(num)
// 			zeroVal := mgtd.Int32RawAlias(0)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// int16
// 	{
// 		typeName := mgtd.Int16RawAliasTypeDesc.Name
// 		for _, num := range []mgtd.Int16RawAlias{-55, -3} {
// 			initVal := mgtd.Int16RawAlias(num)
// 			zeroVal := mgtd.Int16RawAlias(0)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// int8
// 	{
// 		typeName := mgtd.Int8RawAliasTypeDesc.Name
// 		for _, num := range []mgtd.Int8RawAlias{-1, 43} {
// 			initVal := mgtd.Int8RawAlias(num)
// 			zeroVal := mgtd.Int8RawAlias(0)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// int
// 	{
// 		typeName := mgtd.IntRawAliasTypeDesc.Name
// 		for _, num := range []mgtd.IntRawAlias{-726, 10283763525} {
// 			initVal := mgtd.IntRawAlias(num)
// 			zeroVal := mgtd.IntRawAlias(0)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// array of ints
// 	{
// 		typeName := mgtd.IntRawArrayAliasTypeDesc.Name
// 		for _, num := range [][3]int{{1, 0, -21923}, {12828, -18128, 19292}} {
// 			initVal := mgtd.IntRawArrayAlias(num)
// 			zeroVal := mgtd.IntRawArrayAlias(*new([3]int))
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// map uint16 int32
// 	{
// 		typeName := mgtd.Uint16Int32RawMapAliasTypeDesc.Name
// 		for _, m := range []map[uint16]int32{{1: -1, 18: 1827374}} {
// 			initVal := mgtd.Uint16Int32RawMapAlias(m)
// 			zeroVal := mgtd.Uint16Int32RawMapAlias(*new(map[uint16]int32))
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}

// 	// array of int double pointers
// 	{
// 		typeName := mgtd.IntRawPtrPtrPtrAliasSliceAliasTypeDesc.Name
// 		var foo int = 181727361
// 		foo1 := &foo
// 		foo2 := &foo1
// 		foo3 := &foo2
// 		var bar int = 28283839844
// 		bar1 := &bar
// 		bar2 := &bar1
// 		bar3 := &bar2
// 		var car int = 0
// 		car1 := &car
// 		car2 := &car1
// 		car3 := &car2
// 		for _, num := range [][]***int{{foo3, bar3, car3}} {
// 			initVal := mgtd.IntRawPtrPtrPtrAliasSliceAlias(num)
// 			zeroVal := mgtd.IntRawPtrPtrPtrAliasSliceAlias(*new([]***int))
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}

// 	// struct
// 	{
// 		typeName := mgtd.RawStructTypeDesc.Name
// 		var n int = 2
// 		n1 := &n
// 		var j int16 = -87
// 		var mr map[*int16]int8 = map[*int16]int8{&j: 3}
// 		for _, num := range []mgtd.RawStructType{
// 			{
// 				UintRaw:      9,
// 				UintRaw16:    95,
// 				IntRawPtrPtr: &n1,
// 				MapRawPtr:    &mr,
// 			},
// 		} {
// 			initVal := num
// 			zeroVal := *new(mgtd.RawStructType)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if initVal.UintRaw != zeroVal.UintRaw ||
// 					initVal.UintRaw16 != zeroVal.UintRaw16 ||
// 					**initVal.IntRawPtrPtr != **zeroVal.IntRawPtrPtr {
// 					// !reflect.DeepEqual(initVal.MapRawPtr, zeroVal.MapRawPtr) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// int32 small buf
// 	{
// 		arr := []byte{25, 255, 21}
// 		val := mgtd.Int32RawAlias(0)
// 		_, err := val.Unmarshal(arr)
// 		if err != errs.ErrSmallBuf {
// 			t.Error("too small buf is ok")
// 		}
// 	}
// 	// valid int32
// 	{
// 		typeName := mgtd.ValidInt32RawTypeDesc.Name
// 		for _, num := range []mgtd.ValidInt32RawAlias{-173} {
// 			initVal := mgtd.ValidInt32RawAlias(num)
// 			zeroVal := mgtd.ValidInt32RawAlias(0)
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error(err)
// 			}
// 			genErr, ok := err.(*utils.GenError)
// 			if !ok {
// 				t.Error("wrong gen error")
// 			}
// 			if genErr.Cause() != mgtd.ErrNegative {
// 				t.Error("not ErrNegative error")
// 			}
// 		}
// 	}

// }
