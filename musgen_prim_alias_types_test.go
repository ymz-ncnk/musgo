//go:generate go run testdata/gen/mus.go -prim $ARG
package musgo

import (
	"math"
	"testing"

	"github.com/ymz-ncnk/musgo/errs"
	"github.com/ymz-ncnk/musgo/testdata"
	tdmg "github.com/ymz-ncnk/musgo/testdata/musgen"
)

func TestPrimAlias(t *testing.T) {

	t.Run("Uint64", func(t *testing.T) {
		for _, val := range []tdmg.Uint64Alias{
			0,
			11555598373362,
			math.MaxUint64,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Uint64 incorrect encoding", func(t *testing.T) {
		arr := []byte{255, 255, 255}
		val := tdmg.Uint64Alias(0)
		_, err := val.Unmarshal(arr)
		if err != errs.ErrSmallBuf {
			t.Error("incorrect encoding is ok")
		}
	})

	t.Run("Uint64 last byte is too big, should be 0 or 1", func(t *testing.T) {
		// 131 - 10000011
		// 128 - 10000000
		arr := []byte{131, 128, 128, 128, 128, 128, 128, 128, 128, 2}
		val := tdmg.Uint64Alias(0)
		_, err := val.Unmarshal(arr)
		if err != errs.ErrOverflow {
			t.Error("incorrect encoding is ok")
		}
	})

	t.Run("Uint64 encoding is too long", func(t *testing.T) {
		arr := []byte{255, 255, 255, 255, 255, 225, 255, 255, 255, 2}
		val := tdmg.Uint64Alias(0)
		_, err := val.Unmarshal(arr)
		if err != errs.ErrOverflow {
			t.Error("too long encoding is ok")
		}
	})

	t.Run("Uint32", func(t *testing.T) {
		for _, val := range []tdmg.Uint32Alias{
			0,
			9937362,
			math.MaxUint32,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Uinte32 last byte is too big, should be less or equal to 15",
		func(t *testing.T) {
			// 131 - 10000011
			// 128 - 10000000
			arr := []byte{131, 128, 128, 128, 16}
			val := tdmg.Uint32Alias(0)
			_, err := val.Unmarshal(arr)
			if err != errs.ErrOverflow {
				t.Error("incorrect encoding is ok")
			}
		})

	t.Run("Uint32 encoding is too long", func(t *testing.T) {
		arr := []byte{255, 255, 255, 255, 255, 2}
		val := tdmg.Uint32Alias(0)
		_, err := val.Unmarshal(arr)
		if err != errs.ErrOverflow {
			t.Error("too long encoding is ok")
		}
	})

	t.Run("Uint16", func(t *testing.T) {
		for _, val := range []tdmg.Uint16Alias{
			0,
			34262,
			math.MaxUint16,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Uint16 last byte is too big, should be less or equal to 3",
		func(t *testing.T) {
			// 131 - 10000011
			// 128 - 10000000
			arr := []byte{131, 128, 4}
			val := tdmg.Uint16Alias(0)
			_, err := val.Unmarshal(arr)
			if err != errs.ErrOverflow {
				t.Error("incorrect encoding is ok")
			}
		})

	t.Run("Uint16 encoding is too long", func(t *testing.T) {
		arr := []byte{255, 255, 255, 2}
		val := tdmg.Uint16Alias(0)
		_, err := val.Unmarshal(arr)
		if err != errs.ErrOverflow {
			t.Error("too long encoding is ok")
		}
	})

	t.Run("Uint8", func(t *testing.T) {
		for _, val := range []tdmg.Uint8Alias{
			0,
			99,
			math.MaxUint8,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Uint", func(t *testing.T) {
		for _, val := range []tdmg.UintAlias{
			0,
			1234943,
			math.MaxUint,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Uint, buffer ends", func(t *testing.T) {
		var val tdmg.IntAlias
		err := testdata.TestBufferEnds(&val, []byte{})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Int64", func(t *testing.T) {
		for _, val := range []tdmg.Int64Alias{
			math.MinInt64,
			0,
			922337203623425,
			math.MaxInt64,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Int64 incorrect encoding", func(t *testing.T) {
		arr := []byte{255, 255, 255}
		val := tdmg.Int64Alias(0)
		_, err := val.Unmarshal(arr)
		if err != errs.ErrSmallBuf {
			t.Error("incorrect encoding is ok")
		}
	})

	t.Run("Int64 encoding is too long", func(t *testing.T) {
		arr := []byte{255, 255, 255, 255, 255, 225, 255, 255, 255, 2}
		val := tdmg.Int64Alias(0)
		_, err := val.Unmarshal(arr)
		if err != errs.ErrOverflow {
			t.Error("too long encoding is ok")
		}
	})

	t.Run("Int32", func(t *testing.T) {
		for _, val := range []tdmg.Int32Alias{
			math.MinInt32,
			0,
			21474000,
			math.MaxInt32,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Int32 encoding is too long", func(t *testing.T) {
		arr := []byte{255, 255, 255, 255, 255, 2}
		val := tdmg.Int32Alias(0)
		_, err := val.Unmarshal(arr)
		if err != errs.ErrOverflow {
			t.Error("too long encoding is ok")
		}
	})

	t.Run("Int16", func(t *testing.T) {
		for _, val := range []tdmg.Int16Alias{
			math.MinInt16,
			0,
			1134,
			math.MaxInt16,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Int16 encoding is too long", func(t *testing.T) {
		arr := []byte{255, 255, 255, 2}
		val := tdmg.Int16Alias(0)
		_, err := val.Unmarshal(arr)
		if err != errs.ErrOverflow {
			t.Error("too long encoding is ok")
		}
	})

	t.Run("Int8", func(t *testing.T) {
		for _, val := range []tdmg.Int8Alias{
			math.MinInt8,
			0,
			123,
			math.MaxInt8,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Int", func(t *testing.T) {
		for _, val := range []tdmg.IntAlias{
			math.MinInt,
			0,
			1209748,
			math.MaxInt,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Int, buffer ends", func(t *testing.T) {
		var val tdmg.UintAlias
		err := testdata.TestBufferEnds(&val, []byte{})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Float64", func(t *testing.T) {
		for _, val := range []tdmg.Float64Alias{
			-math.MaxFloat64,
			0,
			math.SmallestNonzeroFloat64,
			12093.1827,
			math.MaxFloat64,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Float32", func(t *testing.T) {
		for _, val := range []tdmg.Float32Alias{
			-math.MaxFloat32,
			0,
			math.SmallestNonzeroFloat32,
			1093.11,
			math.MaxFloat32,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Bool", func(t *testing.T) {
		for _, val := range []tdmg.BoolAlias{
			true,
			false,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Bool, buffer ends", func(t *testing.T) {
		var val tdmg.BoolAlias
		err := testdata.TestBufferEnds(&val, []byte{})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Bool wrong byte", func(t *testing.T) {
		arr := []byte{3}
		val := tdmg.BoolAlias(false)
		_, err := val.Unmarshal(arr)
		if err != errs.ErrWrongByte {
			t.Error("wrong byte is ok")
		}
	})

	t.Run("Byte", func(t *testing.T) {
		for _, val := range []tdmg.ByteAlias{
			0,
			38,
			235,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Byte, buffer ends", func(t *testing.T) {
		var val tdmg.ByteAlias
		err := testdata.TestBufferEnds(&val, []byte{})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("String", func(t *testing.T) {
		for _, val := range []tdmg.StringAlias{
			"hello world",
			"",
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("String, too small buf", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				if r != errs.ErrSmallBuf {
					t.Error("unexpected panic msg")
				}
			}
		}()
		val := tdmg.StringAlias("Hello world")
		buf := make([]byte, 6)
		val.Marshal(buf)
		t.Error("expected to panic")
	})

	t.Run("String, length is too big", func(t *testing.T) {
		arr := []byte{12, 1, 1}
		val := tdmg.StringAlias("")
		_, err := val.Unmarshal(arr)
		if err != errs.ErrSmallBuf {
			t.Error("too short buf is ok")
		}
	})

	t.Run("String, negative length", func(t *testing.T) {
		arr := []byte{1, 12, 12}
		val := tdmg.StringAlias("")
		_, err := val.Unmarshal(arr)
		if err != errs.ErrNegativeLength {
			t.Error("negative length in byte string is ok")
		}
	})

	t.Run("String, buffer ends", func(t *testing.T) {
		val := tdmg.StringAlias("")
		err := testdata.TestBufferEnds(&val, []byte{})
		if err != nil {
			t.Error(err)
		}
	})

}
