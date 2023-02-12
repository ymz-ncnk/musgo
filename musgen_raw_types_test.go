//go:generate go run testdata/gen/mus.go -intraw $ARG
package musgo

import (
	"errors"
	"math"
	"testing"

	"github.com/ymz-ncnk/musgo/errs"
	"github.com/ymz-ncnk/musgo/testdata"
	tdmg "github.com/ymz-ncnk/musgo/testdata/musgen"
)

func TestRawEncoding(t *testing.T) {

	t.Run("Uint64Raw", func(t *testing.T) {
		for _, val := range []tdmg.Uint64RawAlias{
			0,
			151492823822937,
			113191817,
			math.MaxUint64,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Uint32Raw", func(t *testing.T) {
		for _, val := range []tdmg.Uint32RawAlias{
			0,
			1,
			94383726,
			math.MaxUint32,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Uint16Raw", func(t *testing.T) {
		for _, val := range []tdmg.Uint16RawAlias{
			0,
			929,
			11111,
			math.MaxUint16,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Uint8Raw", func(t *testing.T) {
		for _, val := range []tdmg.Uint8RawAlias{
			0,
			34,
			math.MaxInt8,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("UintRaw", func(t *testing.T) {
		for _, val := range []tdmg.UintRawAlias{
			0,
			545349,
			11514141414,
			math.MaxUint,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Int64Raw", func(t *testing.T) {
		for _, val := range []tdmg.Int64RawAlias{
			math.MinInt64,
			-2937437292,
			0,
			17620,
			math.MaxInt64,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Int32Raw", func(t *testing.T) {
		for _, val := range []tdmg.Int32RawAlias{
			math.MinInt32,
			-17253,
			0,
			math.MaxInt32,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Int16Raw", func(t *testing.T) {
		for _, val := range []tdmg.Int16RawAlias{
			math.MinInt16,
			-55,
			-3,
			0,
			math.MaxInt16,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Int8Raw", func(t *testing.T) {
		for _, val := range []tdmg.Int8RawAlias{
			math.MinInt8,
			-1,
			0,
			43,
			math.MaxInt8,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("IntRaw", func(t *testing.T) {
		for _, val := range []tdmg.IntRawAlias{
			math.MinInt,
			-726,
			0,
			10283763525,
			math.MaxInt,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Float64Raw", func(t *testing.T) {
		for _, val := range []tdmg.Float64RawAlias{
			-math.MaxFloat64,
			0,
			0.151492823822937,
			0.113191817,
			math.MaxFloat64,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Float32Raw", func(t *testing.T) {
		for _, val := range []tdmg.Float32RawAlias{
			-math.MaxFloat32,
			0.151492823822937,
			0.113191817,
			0,
			math.MaxFloat32,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array with raw elems", func(t *testing.T) {
		for _, val := range []tdmg.IntRawArrayAlias{
			{1, 0, -21923},
			{12828, -18128, 19292},
			{math.MaxInt, math.MinInt, 10102},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map with raw keys/values", func(t *testing.T) {
		for _, val := range []tdmg.Uint16Int32RawMapAlias{
			{1: -1, 18: 1827374},
			{0: math.MinInt32, math.MaxUint16: math.MaxInt32},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array with raw triple pointers", func(t *testing.T) {
		var (
			n  float64 = math.SmallestNonzeroFloat64
			n1         = &n
			n2         = &n1
			n3         = &n2

			m  float64 = -0.28283839844
			m1         = &m
			m2         = &m1
			m3         = &m2

			k  float64 = 0
			k1         = &k
			k2         = &k1
			k3         = &k2

			p  float64 = -math.MaxFloat64
			p1         = &p
			p2         = &p1
			p3         = &p2
		)
		for _, val := range []tdmg.Float64RawPtrPtrPtrAliasSliceAlias{
			{n3, m3, k3, p3},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("StructRaw", func(t *testing.T) {
		var (
			n  int = 2
			n1     = &n

			m  int16 = 1
			m1       = &m

			k float64 = -math.MaxFloat64

			p map[*float64]*int16 = map[*float64]*int16{&k: m1}
		)
		for _, val := range []tdmg.RawStructType{
			{
				UintRaw:      9,
				Float32Raw:   95,
				IntRawPtrPtr: &n1,
				MapRawPtr:    &p,
			},
		} {
			aval, err := testdata.ExecGeneratedCode(val)
			if err != nil {
				t.Error(err)
			} else {
				sval := aval.(*tdmg.RawStructType)
				if val.UintRaw != sval.UintRaw ||
					val.Float32Raw != sval.Float32Raw ||
					**val.IntRawPtrPtr != **sval.IntRawPtrPtr {
					// TODO
					// !reflect.DeepEqual(initVal.MapRawPtr, zeroVal.MapRawPtr) {
					t.Errorf("Marshal - %v, result of Unmarshal - %v", val, sval)
				}
			}
		}
	})

	t.Run("Int32Raw small buf", func(t *testing.T) {
		arr := []byte{25, 255, 21}
		val := tdmg.Int32RawAlias(0)
		_, err := val.Unmarshal(arr)
		if err != errs.ErrSmallBuf {
			t.Error("too small buf is ok")
		}
	})

	t.Run("Int32 validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidInt32RawAlias{
			-173,
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				unmarshalErr := errors.Unwrap(err)
				if unmarshalErr != tdmg.ErrNegative {
					t.Errorf("wrong error '%v'", unmarshalErr)
				}
			}
		}
	})

}
