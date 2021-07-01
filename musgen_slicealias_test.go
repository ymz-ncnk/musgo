package musgo

// //go:generate go run testdata/make/musable.go -slice $ARG
// package musgo

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/ymz-ncnk/musgo/errs"
// 	mgtd "github.com/ymz-ncnk/musgo/testdata/musgen"
// 	utils "github.com/ymz-ncnk/musgo/testdata/musgen/utils"
// )

// func TestGeneratedSliceAliasCode(t *testing.T) {
// 	var err error
// 	// simple slice
// 	{
// 		typeName := mgtd.StrSliceAliasTypeDesc.Name
// 		for _, num := range [][]string{{"hello", "world"}, {"", "tekejt"}} {
// 			initVal := mgtd.StrSliceAlias(num)
// 			zeroVal := mgtd.StrSliceAlias(*new([]string))
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
// 	// slice of pointers
// 	{
// 		typeName := mgtd.FloatPtrSliceAliasTypeDesc.Name
// 		foo := 191.10
// 		bar := 0.0
// 		car := -0.19283
// 		for _, num := range [][]*float64{{&foo, &bar, &car}} {
// 			initVal := mgtd.FloatPtrSliceAlias(num)
// 			zeroVal := mgtd.FloatPtrSliceAlias(*new([]*float64))
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
// 	// slice of alias
// 	{
// 		typeName := mgtd.IntAliasSliceAliasTypeDesc.Name
// 		for _, num := range [][]mgtd.IntAlias{
// 			{1, 0, -21923},
// 			{12828, -18128, 19292},
// 		} {
// 			initVal := mgtd.IntAliasSliceAlias(num)
// 			zeroVal := mgtd.IntAliasSliceAlias(*new([]mgtd.IntAlias))
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
// 	// slice of alias pointers
// 	{
// 		typeName := mgtd.Uint64PtrAliasSliceAliasTypeDesc.Name
// 		var foo mgtd.Uint64Alias = 181727361
// 		var bar mgtd.Uint64Alias = 28283839844
// 		var car mgtd.Uint64Alias = 0
// 		for _, num := range [][]*mgtd.Uint64Alias{{&foo, &bar, &car}} {
// 			initVal := mgtd.Uint64PtrAliasSliceAlias(num)
// 			zeroVal := mgtd.Uint64PtrAliasSliceAlias(*new([]*mgtd.Uint64Alias))
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
// 	// slice of arrays
// 	{
// 		typeName := mgtd.BoolSliceSliceAliasTypeDesc.Name
// 		for _, num := range [][][]bool{{{true, false}, {true, true, true}}} {
// 			initVal := mgtd.BoolSliceSliceAlias(num)
// 			zeroVal := mgtd.BoolSliceSliceAlias(*new([][]bool))
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
// 	// slice of arrays
// 	{
// 		typeName := mgtd.ByteArraySliceAliasTypeDesc.Name
// 		for _, num := range [][][2]byte{{
// 			[2]byte{0x01, 0x60},
// 			[2]byte{0x50, 0x78},
// 		}} {
// 			initVal := mgtd.ByteArraySliceAlias(num)
// 			zeroVal := mgtd.ByteArraySliceAlias(*new([][2]byte))
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
// 	// slice of arrays pointers
// 	{
// 		typeName := mgtd.FloatArrayPtrSliceAliasTypeDesc.Name
// 		for _, num := range [][]*[2]float32{{
// 			&[2]float32{102.1826, -123.12},
// 			&[2]float32{10, 0.2836},
// 		}} {
// 			initVal := mgtd.FloatArrayPtrSliceAlias(num)
// 			zeroVal := mgtd.FloatArrayPtrSliceAlias(*new([]*[2]float32))
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
// 	// slice of maps
// 	{
// 		typeName := mgtd.IntStrMapSliceAliasTypeDesc.Name
// 		for _, num := range [][]map[int16]string{{
// 			map[int16]string{345: "io"},
// 			map[int16]string{19292: "poll"},
// 		}} {
// 			initVal := mgtd.IntStrMapSliceAlias(num)
// 			zeroVal := mgtd.IntStrMapSliceAlias(*new([]map[int16]string))
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
// 	// slice of maps pointers
// 	{
// 		typeName := mgtd.Uint32Int32MapPtrSliceAliasTypeDesc.Name
// 		for _, num := range [][]*map[uint32]int32{{
// 			&map[uint32]int32{345: 82},
// 			&map[uint32]int32{19292: 893837},
// 		}} {
// 			initVal := mgtd.Uint32Int32MapPtrSliceAlias(num)
// 			zeroVal := mgtd.Uint32Int32MapPtrSliceAlias(*new([]*map[uint32]int32))
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
// 	// slice of CustomType
// 	{
// 		typeName := mgtd.StructTypeSliceAliasTypeDesc.Name
// 		for _, num := range [][]mgtd.SimpleStructType{{
// 			mgtd.SimpleStructType{Int: 8},
// 			mgtd.SimpleStructType{Int: 23429},
// 		}} {
// 			initVal := mgtd.StructTypeSliceAlias(num)
// 			zeroVal := mgtd.StructTypeSliceAlias(*new([]mgtd.SimpleStructType))
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
// 	// slice of CustomType pointers
// 	{
// 		typeName := mgtd.StructTypePtrSliceAliasTypeDesc.Name
// 		for _, num := range [][]*mgtd.SimpleStructType{{
// 			&mgtd.SimpleStructType{Int: 8},
// 			&mgtd.SimpleStructType{Int: 23429},
// 		}} {
// 			initVal := mgtd.StructTypePtrSliceAlias(num)
// 			zeroVal := mgtd.StructTypePtrSliceAlias(*new([]*mgtd.SimpleStructType))
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
// 	// tricky slice type
// 	{
// 		getVal := func(k *[]mgtd.StringAlias,
// 			m map[*[]mgtd.StringAlias]map[mgtd.SimpleStructType][]int) (map[mgtd.SimpleStructType][]int, bool) {
// 			// ka := toSlice(k)
// 			for mk, mvl := range m {
// 				if reflect.DeepEqual(*k, *mk) {
// 					return mvl, true
// 				}
// 			}
// 			return nil, false
// 		}

// 		typeName := mgtd.TrickySliceAliasTypeDesc.Name
// 		num := [][2]map[*[]mgtd.StringAlias]map[mgtd.SimpleStructType][]int{
// 			{
// 				{
// 					&[]mgtd.StringAlias{"s"}: {mgtd.SimpleStructType{Int: 1}: []int{23}},
// 				},
// 			},
// 			{
// 				{
// 					&[]mgtd.StringAlias{"t"}: {mgtd.SimpleStructType{Int: -1123}: []int{63532}},
// 				},
// 			},
// 		}

// 		initVal := mgtd.TrickySliceAlias(num)
// 		zeroVal := mgtd.TrickySliceAlias(
// 			*new([][2]map[*[]mgtd.StringAlias]map[mgtd.SimpleStructType][]int))
// 		err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 		if err != nil {
// 			t.Error(err)
// 		} else {
// 			// not works when map's key is pointer
// 			// if !reflect.DeepEqual(initVal, zeroVal) {
// 			// 	t.Errorf(MuErrMsg, typeName)
// 			// }
// 			for i, as := range initVal {
// 				for j, a := range as {
// 					for k, im := range a {
// 						zm, pst := getVal(k, zeroVal[i][j])
// 						if !pst {
// 							t.Errorf(utils.MuErrMsg, typeName)
// 						}
// 						for kk, imm := range im {
// 							if !reflect.DeepEqual(imm, zm[kk]) {
// 								t.Errorf(utils.MuErrMsg, typeName)
// 							}
// 						}
// 					}

// 				}
// 			}
// 		}
// 	}
// 	// slice, buffer ends
// 	{
// 		var val mgtd.StrSliceAlias
// 		err := utils.TestBufferEnds(&val, []byte{})
// 		if err != nil {
// 			t.Error(err)
// 		}
// 	}
// 	// slice, length is too big
// 	{
// 		var val mgtd.StrSliceAlias
// 		buf := []byte{12, 6, 1, 1, 1}
// 		err := utils.TestBufferEnds(&val, buf)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 	}
// 	// invalid elem slice
// 	{
// 		typeName := mgtd.ValidUintSliceAliasTypeDesc.Name
// 		for _, num := range [][]uint{{2, 11}} {
// 			initVal := mgtd.ValidUintSliceAlias(num)
// 			zeroVal := mgtd.ValidUintSliceAlias(*new([]uint))
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error(err)
// 			}
// 			genErr, ok := err.(*utils.GenError)
// 			if !ok {
// 				t.Error("wrong gen error")
// 			}
// 			arrayErr, ok := genErr.Cause().(errs.ArrayError)
// 			if !ok {
// 				t.Error("wrong array error")
// 			}
// 			if arrayErr.Cause() != mgtd.ErrBiggerThanTen {
// 				t.Error("wrong cause")
// 			}
// 		}
// 	}
// 	// invalid slice
// 	{
// 		typeName := mgtd.ValidUintSliceAliasTypeDesc.Name
// 		for _, num := range [][]uint{{8, 6}} {
// 			initVal := mgtd.ValidUintSliceAlias(num)
// 			zeroVal := mgtd.ValidUintSliceAlias(*new([]uint))
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error(err)
// 			}
// 			genErr, ok := err.(*utils.GenError)
// 			if !ok {
// 				t.Error("wrong gen error")
// 			}
// 			if genErr.Cause() != mgtd.ErrSliceSumBiggerThanTen {
// 				t.Error("wrong cause")
// 			}
// 		}
// 	}
// 	// invalid MaxLength slice
// 	{
// 		typeName := mgtd.ValidUintSliceAliasTypeDesc.Name
// 		for _, num := range [][]uint{{1, 0, 0, 0}} {
// 			initVal := mgtd.ValidUintSliceAlias(num)
// 			zeroVal := mgtd.ValidUintSliceAlias(*new([]uint))
// 			err = utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error(err)
// 			}
// 			genErr, ok := err.(*utils.GenError)
// 			if !ok {
// 				t.Error("wrong gen error")
// 			}
// 			if genErr.Cause() != errs.ErrMaxLengthExceeded {
// 				t.Error("wrong MaxLength cause")
// 			}
// 		}
// 	}
// }
