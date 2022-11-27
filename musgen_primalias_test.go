package musgo

// //go:generate go run testdata/make/musable.go -prim $ARG
// package musgo

// import (
// 	"testing"

// 	"github.com/ymz-ncnk/musgo/errs"
// 	mgtd "github.com/ymz-ncnk/musgo/testdata/musgen"
// 	utils "github.com/ymz-ncnk/musgo/testdata/musgen/utils"
// )

// // Can't do this. Can't generate methods for such type, error - invalid receiver
// //  'typename' (pointer or interface type).
// // type Uint64PtrAlias *uint64

// func TestPrimAlias(t *testing.T) {
// 	// uint
// 	{
// 		typeName := mgtd.Uint64AliasTypeDesc.Name
// 		{
// 			for _, num := range []uint64{11555598373362, 0} {
// 				initVal := mgtd.Uint64Alias(num)
// 				zeroVal := mgtd.Uint64Alias(0)
// 				err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 				if err != nil {
// 					t.Error(err)
// 				} else {
// 					if initVal != zeroVal {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 		// uint64 incorrect encoding
// 		{
// 			arr := []byte{255, 255, 255}
// 			val := mgtd.Uint64Alias(0)
// 			_, err := val.Unmarshal(arr)
// 			if err != errs.ErrSmallBuf {
// 				t.Error("incorrect encoding is ok")
// 			}
// 		}
// 		// uint64 last byte is too big, should be less or equal to 1
// 		{
// 			// 131 - 10000011
// 			// 128 - 10000000
// 			arr := []byte{131, 128, 128, 128, 128, 128, 128, 128, 128, 2}
// 			val := mgtd.Uint64Alias(0)
// 			_, err := val.Unmarshal(arr)
// 			if err != errs.ErrOverflow {
// 				t.Error("incorrect encoding is ok")
// 			}
// 		}
// 		// uint64 encoding is to long
// 		{
// 			arr := []byte{255, 255, 255, 255, 255, 225, 255, 255, 255, 2}
// 			val := mgtd.Uint64Alias(0)
// 			_, err := val.Unmarshal(arr)
// 			if err != errs.ErrOverflow {
// 				t.Error("too long encoding is ok")
// 			}
// 		}
// 	}
// 	{
// 		typeName := mgtd.Uint32AliasTypeDesc.Name
// 		{
// 			for _, num := range []uint32{9937362, 0} {
// 				initVal := mgtd.Uint32Alias(num)
// 				zeroVal := mgtd.Uint32Alias(0)
// 				err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 				if err != nil {
// 					t.Error(err)
// 				} else {
// 					if initVal != zeroVal {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 		// uint32 last byte is too big, should be less or equal to 15
// 		{
// 			// 131 - 10000011
// 			// 128 - 10000000
// 			arr := []byte{131, 128, 128, 128, 16}
// 			val := mgtd.Uint32Alias(0)
// 			_, err := val.Unmarshal(arr)
// 			if err != errs.ErrOverflow {
// 				t.Error("incorrect encoding is ok")
// 			}
// 		}
// 		// uint32 encoding is to long
// 		{
// 			arr := []byte{255, 255, 255, 255, 255, 2}
// 			val := mgtd.Uint32Alias(0)
// 			_, err := val.Unmarshal(arr)
// 			if err != errs.ErrOverflow {
// 				t.Error("too long encoding is ok")
// 			}
// 		}
// 	}
// 	{
// 		typeName := mgtd.Uint16AliasTypeDesc.Name
// 		{
// 			for _, num := range []uint16{34262, 0} {
// 				initVal := mgtd.Uint16Alias(num)
// 				zeroVal := mgtd.Uint16Alias(0)
// 				err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 				if err != nil {
// 					t.Error(err)
// 				} else {
// 					if initVal != zeroVal {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 		// uint16 last byte is too big, should be less or equal to 3
// 		{
// 			// 131 - 10000011
// 			// 128 - 10000000
// 			arr := []byte{131, 128, 4}
// 			val := mgtd.Uint16Alias(0)
// 			_, err := val.Unmarshal(arr)
// 			if err != errs.ErrOverflow {
// 				t.Error("incorrect encoding is ok")
// 			}
// 		}
// 		// uint16 encoding is to long
// 		{
// 			arr := []byte{255, 255, 255, 2}
// 			val := mgtd.Uint16Alias(0)
// 			_, err := val.Unmarshal(arr)
// 			if err != errs.ErrOverflow {
// 				t.Error("too long encoding is ok")
// 			}
// 		}
// 	}
// 	{
// 		typeName := mgtd.Uint8AliasTypeDesc.Name
// 		for _, num := range []uint8{99, 0, 35} {
// 			initVal := mgtd.Uint8Alias(num)
// 			zeroVal := mgtd.Uint8Alias(0)
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if initVal != zeroVal {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	{
// 		typeName := mgtd.UintAliasTypeDesc.Name
// 		for _, num := range []uint{182373, 12123131231321} {
// 			initVal := mgtd.UintAlias(num)
// 			zeroVal := mgtd.UintAlias(0)
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if initVal != zeroVal {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// int
// 	{
// 		typeName := mgtd.Int64AliasTypeDesc.Name
// 		{
// 			for _, num := range []int64{-2761259239023909999, 0, 3425219181718387236} {
// 				initVal := mgtd.Int64Alias(num)
// 				zeroVal := mgtd.Int64Alias(0)
// 				err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 				if err != nil {
// 					t.Error(err)
// 				} else {
// 					if initVal != zeroVal {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 		// int64 incorrect encoding
// 		{
// 			arr := []byte{255, 255, 255}
// 			val := mgtd.Int64Alias(0)
// 			_, err := val.Unmarshal(arr)
// 			if err != errs.ErrSmallBuf {
// 				t.Error("incorrect encoding is ok")
// 			}
// 		}
// 		// int64 encoding is to long
// 		{
// 			arr := []byte{255, 255, 255, 255, 255, 225, 255, 255, 255, 2}
// 			val := mgtd.Int64Alias(0)
// 			_, err := val.Unmarshal(arr)
// 			if err != errs.ErrOverflow {
// 				t.Error("too long encoding is ok")
// 			}
// 		}
// 	}
// 	{
// 		typeName := mgtd.Int32AliasTypeDesc.Name
// 		{
// 			for _, num := range []int32{-1265126289, 0, 342521918} {
// 				initVal := mgtd.Int32Alias(num)
// 				zeroVal := mgtd.Int32Alias(0)
// 				err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 				if err != nil {
// 					t.Error(err)
// 				} else {
// 					if initVal != zeroVal {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 		// int32 encoding is to long
// 		{
// 			arr := []byte{255, 255, 255, 255, 255, 2}
// 			val := mgtd.Int32Alias(0)
// 			_, err := val.Unmarshal(arr)
// 			if err != errs.ErrOverflow {
// 				t.Error("too long encoding is ok")
// 			}
// 		}
// 	}
// 	{
// 		typeName := mgtd.Int16AliasTypeDesc.Name
// 		{
// 			for _, num := range []int16{-28373, 0, 3736} {
// 				initVal := mgtd.Int16Alias(num)
// 				zeroVal := mgtd.Int16Alias(0)
// 				err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 				if err != nil {
// 					t.Error(err)
// 				} else {
// 					if initVal != zeroVal {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 		// int16 encoding is to long
// 		{
// 			arr := []byte{255, 255, 255, 2}
// 			val := mgtd.Int16Alias(0)
// 			_, err := val.Unmarshal(arr)
// 			if err != errs.ErrOverflow {
// 				t.Error("too long encoding is ok")
// 			}
// 		}
// 	}
// 	{
// 		typeName := mgtd.Int8AliasTypeDesc.Name
// 		for _, num := range []int8{-2, 0, 37} {
// 			initVal := mgtd.Int8Alias(num)
// 			zeroVal := mgtd.Int8Alias(0)
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if initVal != zeroVal {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	{
// 		typeName := mgtd.IntAliasTypeDesc.Name
// 		for _, num := range []int{-2982384774636, 0, 22342348} {
// 			initVal := mgtd.IntAlias(num)
// 			zeroVal := mgtd.IntAlias(0)
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if initVal != zeroVal {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	{
// 		typeName := mgtd.Float64AliasTypeDesc.Name
// 		for _, num := range []float64{-261144411.3523, 0, 8272345.2028365} {
// 			initVal := mgtd.Float64Alias(num)
// 			zeroVal := mgtd.Float64Alias(0)
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if initVal != zeroVal {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	{
// 		typeName := mgtd.Float32AliasTypeDesc.Name
// 		for _, num := range []float32{-2374892.161515141, 0, 299773.28} {
// 			initVal := mgtd.Float32Alias(num)
// 			zeroVal := mgtd.Float32Alias(0)
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if initVal != zeroVal {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// bool
// 	{
// 		typeName := mgtd.BoolAliasTypeDesc.Name
// 		{
// 			for _, num := range []bool{true, false} {
// 				initVal := mgtd.BoolAlias(num)
// 				zeroVal := mgtd.BoolAlias(false)
// 				err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 				if err != nil {
// 					t.Error(err)
// 				} else {
// 					if initVal != zeroVal {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 		{
// 			arr := []byte{3}
// 			val := mgtd.BoolAlias(false)
// 			_, err := val.Unmarshal(arr)
// 			if err != errs.ErrWrongByte {
// 				t.Error("wrong byte bool is ok")
// 			}
// 		}
// 	}
// 	// byte
// 	{
// 		typeName := mgtd.ByteAliasTypeDesc.Name
// 		for _, num := range []byte{38, 235} {
// 			initVal := mgtd.ByteAlias(num)
// 			zeroVal := mgtd.ByteAlias(0x00)
// 			err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 			if err != nil {
// 				t.Error(err)
// 			} else {
// 				if initVal != zeroVal {
// 					t.Errorf(utils.MuErrMsg, typeName)
// 				}
// 			}
// 		}
// 	}
// 	// string
// 	{
// 		typeName := mgtd.StringAliasTypeDesc.Name
// 		{
// 			for _, num := range []string{"hello world", ""} {
// 				initVal := mgtd.StringAlias(num)
// 				zeroVal := mgtd.StringAlias("")
// 				err := utils.ExecGeneratedCode(initVal, &zeroVal, typeName)
// 				if err != nil {
// 					t.Error(err)
// 				} else {
// 					if initVal != zeroVal {
// 						t.Errorf(utils.MuErrMsg, typeName)
// 					}
// 				}
// 			}
// 		}
// 		// string, too small buf
// 		{
// 			defer func() {
// 				if r := recover(); r != nil {
// 					if r != errs.ErrSmallBuf {
// 						t.Error("unexpected panic msg")
// 					}
// 				}
// 			}()
// 			val := mgtd.StringAlias("Hello world")
// 			buf := make([]byte, 6)
// 			val.Marshal(buf)
// 			t.Error("expected to panic")
// 		}
// 		// string, length is too big
// 		{
// 			arr := []byte{12, 1, 1}
// 			val := mgtd.StringAlias("")
// 			_, err := val.Unmarshal(arr)
// 			if err != errs.ErrSmallBuf {
// 				t.Error("too short buf is ok")
// 			}
// 		}
// 		// string, negative length
// 		{
// 			arr := []byte{1, 12, 12}
// 			val := mgtd.StringAlias("")
// 			_, err := val.Unmarshal(arr)
// 			if err != errs.ErrNegativeLength {
// 				t.Error("negative length in byte string is ok")
// 			}
// 		}
// 		// uint, buffer ends
// 		{
// 			var val mgtd.IntAlias
// 			err := utils.TestBufferEnds(&val, []byte{})
// 			if err != nil {
// 				t.Error(err)
// 			}
// 		}
// 		// int, buffer ends
// 		{
// 			var val mgtd.UintAlias
// 			err := utils.TestBufferEnds(&val, []byte{})
// 			if err != nil {
// 				t.Error(err)
// 			}
// 		}
// 		// string, buffer ends
// 		{
// 			val := mgtd.StringAlias("")
// 			err := utils.TestBufferEnds(&val, []byte{})
// 			if err != nil {
// 				t.Error(err)
// 			}
// 		}
// 		// byte, buffer ends
// 		{
// 			var val mgtd.ByteAlias
// 			err := utils.TestBufferEnds(&val, []byte{})
// 			if err != nil {
// 				t.Error(err)
// 			}
// 		}
// 		// bool, buffer ends
// 		{
// 			var val mgtd.BoolAlias
// 			err := utils.TestBufferEnds(&val, []byte{})
// 			if err != nil {
// 				t.Error(err)
// 			}
// 		}
// 		// bool, buffer ends
// 		{
// 			var val mgtd.BoolAlias
// 			err := utils.TestBufferEnds(&val, []byte{})
// 			if err != nil {
// 				t.Error(err)
// 			}
// 		}
// 	}
// }
