package musgo

// //go:generate go run testdata/make/musable.go -map $ARG
// package musgo

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/ymz-ncnk/musgo/errs"
// 	mgtd "github.com/ymz-ncnk/musgo/testdata/musgen"
// 	utils "github.com/ymz-ncnk/musgo/testdata/musgen/utils"
// )

// func TestGeneratedMapAliasCode(t *testing.T) {
// 	// map's key can't be function, map or slice
// 	// simple map
// 	{
// 		typeName := mgtd.StrIntMapAliasTypeDesc.Name
// 		for _, num := range []map[string]int{
// 			{"some": 15},
// 			{"yeuryrywuww": -15141},
// 			{"": 0},
// 		} {
// 			initVal := mgtd.StrIntMapAlias(num)
// 			zeroVal := mgtd.StrIntMapAlias(*new(map[string]int))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				for k, v := range initVal {
// 					if v != zeroVal[k] {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// map of pointers
// 	{
// 		typeName := mgtd.StrPtrIntPtrMapAliasTypeDesc.Name
// 		htpg := "htpg"
// 		negInt := -17263
// 		emptyStr := ""
// 		zero := 0
// 		for _, num := range []map[*string]*int{
// 			{&htpg: &negInt},
// 			{&emptyStr: &zero},
// 		} {
// 			initVal := mgtd.StrPtrIntPtrMapAlias(num)
// 			zeroVal := mgtd.StrPtrIntPtrMapAlias(*new(map[*string]*int))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if len(initVal) != len(zeroVal) {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 				initMap := make(map[string]int)
// 				zeroMap := make(map[string]int)
// 				for k, v := range initVal {
// 					initMap[*k] = *v
// 				}
// 				for k, v := range zeroVal {
// 					zeroMap[*k] = *v
// 				}
// 				for k, v := range initMap {
// 					if zeroMap[k] != v {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// map of aliases
// 	{
// 		typeName := mgtd.StrAliasIntAliasMapAliasTypeDesc.Name
// 		for _, num := range []map[mgtd.StringAlias]mgtd.IntAlias{
// 			{"sky": 626262},
// 			{"town": -29283},
// 			{"": 0},
// 		} {
// 			initVal := mgtd.StrAliasIntAliasMapAlias(num)
// 			zeroVal := mgtd.StrAliasIntAliasMapAlias(*new(map[mgtd.StringAlias]mgtd.IntAlias))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				for k, v := range initVal {
// 					if v != zeroVal[k] {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// map of pointers to aliases
// 	{
// 		getVal := func(k *mgtd.StringAlias, m map[*mgtd.StringAlias]*mgtd.IntAlias) (*mgtd.IntAlias,
// 			bool) {
// 			var result mgtd.IntAlias = 0
// 			for mk, mvl := range m {
// 				if *k == *mk {
// 					return mvl, true
// 				}
// 			}
// 			return &result, false
// 		}
// 		typeName := mgtd.StrAliasPtrIntAliasPtrMapAliasTypeDesc.Name
// 		var str1 mgtd.StringAlias = "sky"
// 		var str2 mgtd.StringAlias = "town"
// 		var str3 mgtd.StringAlias = ""
// 		var num1 mgtd.IntAlias = 8239283422
// 		var num2 mgtd.IntAlias = 10192838
// 		var num3 mgtd.IntAlias = 0
// 		for _, num := range []map[*mgtd.StringAlias]*mgtd.IntAlias{
// 			{&str1: &num1},
// 			{&str2: &num2},
// 			{&str3: &num3},
// 		} {
// 			initVal := mgtd.StrAliasPtrIntAliasPtrMapAlias(num)
// 			zeroVal := mgtd.StrAliasPtrIntAliasPtrMapAlias(*new(map[*mgtd.StringAlias]*mgtd.IntAlias))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				for k, v := range initVal {
// 					zv, pst := getVal(k, zeroVal)
// 					if !pst {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 					if *v != *zv {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// map of slices
// 	{
// 		typeName := mgtd.BoolInt16SliceMapAliasTypeDesc.Name
// 		for _, num := range []map[bool][]int16{
// 			{true: {123, 434, 0}},
// 			{false: {-12, -123, 123}},
// 		} {
// 			initVal := mgtd.BoolInt16SliceMapAlias(num)
// 			zeroVal := mgtd.BoolInt16SliceMapAlias(*new(map[bool][]int16))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				for k, s := range initVal {
// 					for i, v := range s {
// 						if v != zeroVal[k][i] {
// 							t.Errorf(utils.MuErrMsg, typeName)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// map of pointers to slices
// 	{
// 		typeName := mgtd.ByteUint16SlicePtrMapAliasTypeDesc.Name
// 		for _, num := range []map[byte]*[]uint16{
// 			{0x50: &[]uint16{98, 333, 0}},
// 			{0x8: &[]uint16{12, 1123, 123}},
// 			{0x9: &[]uint16{}},
// 		} {
// 			initVal := mgtd.ByteUint16SlicePtrMapAlias(num)
// 			zeroVal := mgtd.ByteUint16SlicePtrMapAlias(*new(map[byte]*[]uint16))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				for k, s := range initVal {
// 					for i, v := range *s {
// 						if v != (*zeroVal[k])[i] {
// 							t.Errorf(utils.MuErrMsg, typeName)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// map of arrays
// 	{
// 		typeName := mgtd.Int32Float64ArrayMapAliasTypeDesc.Name
// 		for _, num := range []map[int32][2]float64{
// 			{19283: [2]float64{0.19, 434.937}},
// 			{0: [2]float64{}},
// 			{-12847: {-12, -123.3}},
// 		} {
// 			initVal := mgtd.Int32Float64ArrayMapAlias(num)
// 			zeroVal := mgtd.Int32Float64ArrayMapAlias(*new(map[int32][2]float64))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				for k, s := range initVal {
// 					for i, v := range s {
// 						if v != zeroVal[k][i] {
// 							t.Errorf(utils.MuErrMsg, typeName)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// map of pointers to arrays
// 	{
// 		typeName := mgtd.Float32Uint32ArrayPtrMapAliasTypeDesc.Name
// 		for _, num := range []map[float32]*[2]uint32{
// 			{19283.12: &[2]uint32{234, 434}},
// 			{0: &[2]uint32{}},
// 			{-12847: &[2]uint32{3, 9374}},
// 		} {
// 			initVal := mgtd.Float32Uint32ArrayPtrMapAlias(num)
// 			zeroVal := mgtd.Float32Uint32ArrayPtrMapAlias(*new(map[float32]*[2]uint32))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				for k, s := range initVal {
// 					for i, v := range s {
// 						if v != (*zeroVal[k])[i] {
// 							t.Errorf(utils.MuErrMsg, typeName)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// map of maps
// 	{
// 		typeName := mgtd.FloatByteBoolMapMapAliasTypeDesc.Name
// 		for _, num := range []map[float32]map[byte]bool{
// 			{56.22: {0x03: true, 0x20: true}},
// 			{0: {0x00: false}},
// 			{-123.34: {0x05: false, 0x23: true}},
// 		} {
// 			initVal := mgtd.FloatByteBoolMapMapAlias(num)
// 			zeroVal := mgtd.FloatByteBoolMapMapAlias(*new(map[float32]map[byte]bool))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				for k, v := range initVal {
// 					zv, pst := zeroVal[k]
// 					if !pst {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 					for sk, sv := range v {
// 						if sv != zv[sk] {
// 							t.Errorf(utils.MuErrMsg, typeName)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// map of *maps
// 	{
// 		typeName := mgtd.UintIntStringMapPtrMapAliasTypeDesc.Name
// 		for _, num := range []map[uint16]*map[int]string{
// 			{56: {-123554: "true", 12837: "city"}},
// 			{0: {0: ""}},
// 			{1918: {12918923131211111: "bird", -1: "&&&92-2=!"}},
// 		} {
// 			initVal := mgtd.UintIntStringMapPtrMapAlias(num)
// 			zeroVal := mgtd.UintIntStringMapPtrMapAlias(*new(map[uint16]*map[int]string))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				for k, v := range initVal {
// 					for ik, iv := range *v {
// 						if iv != (*zeroVal[k])[ik] {
// 							t.Errorf(utils.MuErrMsg, typeName)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// map of CustomType
// 	{
// 		typeName := mgtd.StructTypeStructTypeMapAliasTypeDesc.Name
// 		for _, num := range []map[mgtd.SimpleStructType]mgtd.SimpleStructType{
// 			{
// 				mgtd.SimpleStructType{Int: 12}: mgtd.SimpleStructType{Int: 2383},
// 			},
// 			{
// 				mgtd.SimpleStructType{Int: -12}: mgtd.SimpleStructType{Int: 111},
// 			},
// 		} {
// 			initVal := mgtd.StructTypeStructTypeMapAlias(num)
// 			zeroVal := mgtd.StructTypeStructTypeMapAlias(
// 				*new(map[mgtd.SimpleStructType]mgtd.SimpleStructType))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				for k, v := range initVal {
// 					if v != zeroVal[k] {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// map of pointers to CustomType
// 	{
// 		getVal := func(k *mgtd.SimpleStructType,
// 			m map[*mgtd.SimpleStructType]*mgtd.SimpleStructType) (*mgtd.SimpleStructType, bool) {
// 			for mk, mvl := range m {
// 				if *k == *mk {
// 					return mvl, true
// 				}
// 			}
// 			return nil, false
// 		}

// 		typeName := mgtd.StructTypePtrStructTypePtrMapAliasTypeDesc.Name
// 		for _, num := range []map[*mgtd.SimpleStructType]*mgtd.SimpleStructType{
// 			{
// 				&mgtd.SimpleStructType{Int: 12}: &mgtd.SimpleStructType{Int: 2383},
// 			},
// 			{
// 				&mgtd.SimpleStructType{Int: -12}: &mgtd.SimpleStructType{Int: 111},
// 			},
// 		} {
// 			initVal := mgtd.StructTypePtrStructTypePtrMapAlias(num)
// 			zeroVal := mgtd.StructTypePtrStructTypePtrMapAlias(
// 				*new(map[*mgtd.SimpleStructType]*mgtd.SimpleStructType))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				for k, v := range initVal {
// 					zv, pst := getVal(k, zeroVal)
// 					if !pst {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 					if *v != *zv {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// tricky map
// 	{
// 		getVal := func(k *[]mgtd.SimpleStructType,
// 			m map[*[]mgtd.SimpleStructType]map[*map[int]string][2]int) (map[*map[int]string][2]int, bool) {
// 			for mk, mvl := range m {
// 				if reflect.DeepEqual(*k, *mk) {
// 					return mvl, true
// 				}
// 			}
// 			return nil, false
// 		}
// 		getValByMap := func(k *map[int]string,
// 			m map[*map[int]string][2]int) ([2]int, bool) {
// 			for mk, mvl := range m {
// 				if reflect.DeepEqual(*k, *mk) {
// 					return mvl, true
// 				}
// 			}
// 			return [2]int{}, false
// 		}

// 		typeName := mgtd.TrickyMapAliasTypeDesc.Name
// 		for _, num := range []map[[2]mgtd.StringAlias]map[*[]mgtd.SimpleStructType]map[*map[int]string][2]int{
// 			{
// 				[2]mgtd.StringAlias{"s", "k"}: {
// 					&[]mgtd.SimpleStructType{
// 						{Int: 1},
// 						{Int: 0},
// 					}: {
// 						&map[int]string{4: "y", 6: "o", -34: "pp"}: [2]int{
// 							9,
// 							-9,
// 						},
// 					},
// 				},
// 			},
// 		} {
// 			initVal := mgtd.TrickyMapAlias(num)
// 			zeroVal := mgtd.TrickyMapAlias(*new(map[[2]mgtd.StringAlias]map[*[]mgtd.SimpleStructType]map[*map[int]string][2]int))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				for k, v := range initVal {
// 					zv := zeroVal[k]
// 					for kv, vv := range v {
// 						zvv, pst := getVal(kv, zv)
// 						if !pst {
// 							t.Errorf(utils.MuErrMsg, typeName)
// 						}
// 						for kvv, vvv := range vv {
// 							zvvv, pst := getValByMap(kvv, zvv)
// 							if !pst {
// 								t.Errorf(utils.MuErrMsg, typeName)
// 							}
// 							if vvv != zvvv {
// 								t.Errorf(utils.MuErrMsg, typeName)
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// map, buffer ends
// 	{
// 		var val mgtd.StrIntMapAlias
// 		err := utils.TestBufferEnds(&val, []byte{})
// 		if err != nil {
// 			t.Error(err)
// 		}
// 	}
// 	// map, length is too big
// 	{
// 		var val mgtd.StrIntMapAlias
// 		buf := []byte{12, 6, 1, 1, 1, 12}
// 		err := utils.TestBufferEnds(&val, buf)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 	}
// 	// There are no valid tests. They are made for the StructType. Here we are
// 	// checking only ValidStringIntMapAlias type generation.
// 	// invalid elem map
// 	{
// 		typeName := mgtd.ValidStringIntMapAliasTypeDesc.Name
// 		for _, num := range []map[string]int{{"one": 1, "eleven": 11}} {
// 			initVal := mgtd.ValidStringIntMapAlias(num)
// 			zeroVal := mgtd.ValidStringIntMapAlias(*new(map[string]int))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error(err)
// 			} else {
// 				genErr, ok := err.(*utils.GenError)
// 				if !ok {
// 					t.Error("wrong gen error")
// 				}
// 				mapValueErr, ok := genErr.Cause().(*errs.MapValueError)
// 				if !ok {
// 					t.Error("wrong map value error")
// 				}
// 				if mapValueErr.Cause() != mgtd.ErrBiggerThanTen {
// 					t.Error("wrong cause")
// 				}
// 			}
// 		}
// 	}
// 	// invalid key map
// 	{
// 		typeName := mgtd.ValidStringIntMapAliasTypeDesc.Name
// 		for _, num := range []map[string]int{{"one": 1, "hello": 0}} {
// 			initVal := mgtd.ValidStringIntMapAlias(num)
// 			zeroVal := mgtd.ValidStringIntMapAlias(*new(map[string]int))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error(err)
// 			} else {
// 				genErr, ok := err.(*utils.GenError)
// 				if !ok {
// 					t.Error("wrong gen error")
// 				}
// 				mapKeyErr, ok := genErr.Cause().(*errs.MapKeyError)
// 				if !ok {
// 					t.Error("wrong map key error")
// 				}
// 				if mapKeyErr.Cause() != mgtd.ErrStrIsHello {
// 					t.Error("wrong cause")
// 				}
// 			}
// 		}
// 	}
// 	// invalid map
// 	{
// 		typeName := mgtd.ValidStringIntMapAliasTypeDesc.Name
// 		for _, num := range []map[string]int{{"seven": 7, "six": 6}} {
// 			initVal := mgtd.ValidStringIntMapAlias(num)
// 			zeroVal := mgtd.ValidStringIntMapAlias(*new(map[string]int))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error(err)
// 			} else {
// 				genErr, ok := err.(*utils.GenError)
// 				if !ok {
// 					t.Error("wrong gen error")
// 				}
// 				if genErr.Cause() != mgtd.ErrMapSumBiggerThanTen {
// 					t.Error("wrong map error")
// 				}
// 			}
// 		}
// 	}
// 	// invalid MaxLength map
// 	{
// 		typeName := mgtd.ValidStringIntMapAliasTypeDesc.Name
// 		for _, num := range []map[string]int{{"one": 1, "zero": 0, "two": 1,
// 			"three": 3}} {
// 			initVal := mgtd.ValidStringIntMapAlias(num)
// 			zeroVal := mgtd.ValidStringIntMapAlias(*new(map[string]int))
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err == nil {
// 				t.Error(err)
// 			} else {
// 				genErr, ok := err.(*utils.GenError)
// 				if !ok {
// 					t.Error("wrong gen error")
// 				}
// 				if genErr.Cause() != errs.ErrMaxLengthExceeded {
// 					t.Error("wrong MaxLength cause")
// 				}
// 			}
// 		}
// 	}
// }
