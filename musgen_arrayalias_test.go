package musgo

// //go:generate go run testdata/make/musable.go -array $ARG
// package musgo

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/ymz-ncnk/musgo/errs"
// 	mgtd "github.com/ymz-ncnk/musgo/testdata/musgen"
// )

// func TestGeneratedArrayAliasCode(t *testing.T) {
// 	var err error
// 	// simple array
// 	{
// 		typeName := mgtd.StrArrayAliasTypeDesc.Name
// 		for _, num := range [][3]string{{"hello", "world"}, {"", "tekejt"}} {
// 			initVal := mgtd.StrArrayAlias(num)
// 			zeroVal := mgtd.StrArrayAlias(*new([3]string))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(mgtd.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// array of pointers
// 	{
// 		typeName := mgtd.FloatPtrArrayAliasTypeDesc.Name
// 		foo := 191.10
// 		bar := 0.0
// 		car := -0.19283
// 		for _, num := range [][3]*float64{{&foo, &bar, &car}} {
// 			initVal := mgtd.FloatPtrArrayAlias(num)
// 			zeroVal := mgtd.FloatPtrArrayAlias(*new([3]*float64))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(mgtd.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// array of alias
// 	{
// 		typeName := mgtd.IntAliasArrayAliasTypeDesc.Name
// 		for _, num := range [][3]mgtd.IntAlias{{1, 0, -21923}, {12828, -18128, 19292}} {
// 			initVal := mgtd.IntAliasArrayAlias(num)
// 			zeroVal := mgtd.IntAliasArrayAlias(*new([3]mgtd.IntAlias))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(mgtd.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// array of alias pointers
// 	{
// 		typeName := mgtd.Uint64PtrAliasArrayAliasTypeDesc.Name
// 		var foo mgtd.Uint64Alias = 181727361
// 		var bar mgtd.Uint64Alias = 28283839844
// 		var car mgtd.Uint64Alias = 0
// 		for _, num := range [][3]*mgtd.Uint64Alias{{&foo, &bar, &car}} {
// 			initVal := mgtd.Uint64PtrAliasArrayAlias(num)
// 			zeroVal := mgtd.Uint64PtrAliasArrayAlias(*new([3]*mgtd.Uint64Alias))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(mgtd.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// array of slices
// 	{
// 		typeName := mgtd.Int64SliceArrayAliasTypeDesc.Name
// 		for _, num := range [][3][]int64{{{123, 0, 12342}, {0, 0, 0}, {-1, -2, -12342}}} {
// 			initVal := mgtd.Int64SliceArrayAlias(num)
// 			zeroVal := mgtd.Int64SliceArrayAlias(*new([3][]int64))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(mgtd.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// array of slices pointers
// 	{
// 		typeName := mgtd.Uint16SlicePtrArrayAliasTypeDesc.Name
// 		var slice1 []uint16 = []uint16{7788, 0, 12}
// 		var slice2 []uint16 = []uint16{0, 0, 0}
// 		var slice3 []uint16 = []uint16{657, 23, 981}
// 		for _, num := range [][3]*[]uint16{{&slice1, &slice2, &slice3}} {
// 			initVal := mgtd.Uint16SlicePtrArrayAlias(num)
// 			zeroVal := mgtd.Uint16SlicePtrArrayAlias(*new([3]*[]uint16))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(mgtd.MuErrMsg, typeName)
// 				}

// 			}
// 		}
// 	}
// 	// array of arrays
// 	{
// 		typeName := mgtd.BoolArrayArrayAliasTypeDesc.Name
// 		for _, num := range [][3][3]bool{{{true, false}, {true, true, true}}} {
// 			initVal := mgtd.BoolArrayArrayAlias(num)
// 			zeroVal := mgtd.BoolArrayArrayAlias(*new([3][3]bool))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(mgtd.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// array of arrays pointers
// 	{
// 		typeName := mgtd.BytePtrArrayPtrArrayAliasTypeDesc.Name
// 		arr1 := [2]byte{0x08, 0x55}
// 		arr2 := [2]byte{0x00, 0x09}
// 		for _, num := range [][2]*[2]byte{{&arr1, &arr2}} {
// 			initVal := mgtd.BytePtrArrayPtrArrayAlias(num)
// 			zeroVal := mgtd.BytePtrArrayPtrArrayAlias(*new([2]*[2]byte))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(mgtd.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// array of maps with keys - arrays
// 	{
// 		typeName := mgtd.IntStrMapArrayAliasTypeDesc.Name
// 		map1 := map[int]string{1: "some1", 876: "some2"}
// 		map2 := map[int]string{2: "some3"}
// 		map3 := map[int]string{}
// 		a := [3]map[int]string{map1, map2, map3}
// 		num := a
// 		initVal := mgtd.IntStrMapArrayAlias(num)
// 		zeroVal := mgtd.IntStrMapArrayAlias(*new([3]map[int]string))
// 		err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 		if err != nil {
// 			t.Error(err)
// 		} else {
// 			if !reflect.DeepEqual(initVal, zeroVal) {
// 				t.Errorf(mgtd.MuErrMsg, typeName)
// 			}
// 		}
// 	}
// 	// array of maps pointers
// 	{
// 		typeName := mgtd.Uint32Int32MapArrayAliasTypeDesc.Name
// 		map1 := map[uint32]int32{1927: 0, 123: -123}
// 		map2 := map[uint32]int32{7: 1, 120383: -12323423}
// 		map3 := map[uint32]int32{0: 0, 1: 0}
// 		map4 := map[uint32]int32{7: 1, 120383: -12323423}
// 		num := [4]*map[uint32]int32{&map1, &map2, &map3, &map4}
// 		initVal := mgtd.Uint32Int32MapArrayAlias(num)
// 		zeroVal := mgtd.Uint32Int32MapArrayAlias(*new([4]*map[uint32]int32))
// 		err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 		if err != nil {
// 			t.Error(err)
// 		} else {
// 			if !reflect.DeepEqual(initVal, zeroVal) {
// 				t.Errorf(mgtd.MuErrMsg, typeName)
// 			}
// 		}
// 	}
// 	// array of CustomType
// 	{
// 		typeName := mgtd.StructTypeArrayAliasTypeDesc.Name
// 		for _, num := range [][3]mgtd.SimpleStructType{
// 			{
// 				mgtd.SimpleStructType{Int: 17},
// 				mgtd.SimpleStructType{Int: -123},
// 			},
// 			{
// 				mgtd.SimpleStructType{Int: 11289719871},
// 				mgtd.SimpleStructType{},
// 			},
// 		} {
// 			initVal := mgtd.StructTypeArrayAlias(num)
// 			zeroVal := mgtd.StructTypeArrayAlias(*new([3]mgtd.SimpleStructType))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(mgtd.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// array of CustomType Ptr
// 	{
// 		typeName := mgtd.StructTypePtrArrayAliasTypeDesc.Name
// 		for _, num := range [][3]*mgtd.SimpleStructType{
// 			{
// 				&mgtd.SimpleStructType{Int: -1289371},
// 				&mgtd.SimpleStructType{Int: -123},
// 				&mgtd.SimpleStructType{Int: -123},
// 			},
// 			{
// 				&mgtd.SimpleStructType{Int: 999999},
// 				&mgtd.SimpleStructType{},
// 				&mgtd.SimpleStructType{Int: -182},
// 			},
// 		} {
// 			initVal := mgtd.StructTypePtrArrayAlias(num)
// 			zeroVal := mgtd.StructTypePtrArrayAlias(*new([3]*mgtd.SimpleStructType))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(mgtd.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// tricky array alias
// 	{
// 		typeName := mgtd.TrickyArrayAliasTypeDesc.Name
// 		for _, num := range [][2][][3]map[mgtd.SimpleStructType][1]int{
// 			{
// 				{
// 					{
// 						{
// 							mgtd.SimpleStructType{Int: 1}: {1},
// 						},
// 						{
// 							mgtd.SimpleStructType{Int: 3}: {2},
// 						},
// 						{
// 							mgtd.SimpleStructType{Int: 1}: {9},
// 						},
// 					},
// 					{
// 						{
// 							mgtd.SimpleStructType{Int: 0}: {0},
// 						},
// 						{
// 							mgtd.SimpleStructType{Int: 0}: {233},
// 						},
// 						{
// 							mgtd.SimpleStructType{Int: 112}: {90},
// 						},
// 					},
// 				},
// 				{
// 					{{mgtd.SimpleStructType{Int: 1}: {4}}, {mgtd.SimpleStructType{Int: 3}: {2}},
// 						{mgtd.SimpleStructType{Int: 1}: {9}}},
// 				},
// 			}} {

// 			initVal := mgtd.TrickyArrayAlias(num)
// 			zeroVal := mgtd.TrickyArrayAlias(*new([2][][3]map[mgtd.SimpleStructType][1]int))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(mgtd.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// zero length array
// 	{
// 		typeName := mgtd.StrZeroLengthArrayAliasTypeDesc.Name
// 		for _, num := range [][0]string{
// 			{},
// 		} {
// 			initVal := mgtd.StrZeroLengthArrayAlias(num)
// 			zeroVal := mgtd.StrZeroLengthArrayAlias(*new([0]string))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				for i, ct := range initVal {
// 					if ct != zeroVal[i] {
// 						t.Errorf(mgtd.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// array, buffer ends
// 	{
// 		var val mgtd.StrArrayAlias
// 		err := mgtd.TestBufferEnds(&val, []byte{})
// 		if err != nil {
// 			t.Error(err)
// 		}
// 	}
// 	// array of int double pointers
// 	{
// 		typeName := mgtd.IntPtrPtrPtrAliasArrayAliasTypeDesc.Name
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
// 		for _, num := range [][3]***int{{foo3, bar3, car3}} {
// 			initVal := mgtd.IntPtrPtrPtrAliasArrayAlias(num)
// 			zeroVal := mgtd.IntPtrPtrPtrAliasArrayAlias(*new([3]***int))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if !reflect.DeepEqual(initVal, zeroVal) {
// 					t.Errorf(mgtd.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// invalid elem array
// 	{
// 		typeName := mgtd.ValidIntArrayAliasTypeDesc.Name
// 		for _, num := range [][2]int{{2, 11}} {
// 			initVal := mgtd.ValidIntArrayAlias(num)
// 			zeroVal := mgtd.ValidIntArrayAlias(*new([2]int))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error(err)
// 			}
// 			genErr, ok := err.(*mgtd.GenError)
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
// 	// invalid array
// 	{
// 		typeName := mgtd.ValidIntArrayAliasTypeDesc.Name
// 		for _, num := range [][2]int{{8, 6}} {
// 			initVal := mgtd.ValidIntArrayAlias(num)
// 			zeroVal := mgtd.ValidIntArrayAlias(*new([2]int))
// 			err = mgtd.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error(err)
// 			}
// 			genErr, ok := err.(*mgtd.GenError)
// 			if !ok {
// 				t.Error("wrong gen error")
// 			}
// 			if genErr.Cause() != mgtd.ErrArraySumBiggerThanTen {
// 				t.Error("wrong cause")
// 			}
// 		}
// 	}
// }
