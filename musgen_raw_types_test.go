//go:generate go run testdata/gen/mus.go -intraw $ARG
package musgo

import (
	"errors"
	"testing"

	"github.com/ymz-ncnk/serialization/musgo/errs"
	"github.com/ymz-ncnk/serialization/musgo/testdata"
	tdmg "github.com/ymz-ncnk/serialization/musgo/testdata/musgen"
)

func TestRawEncoding(t *testing.T) {

	t.Run("Uint64Raw", func(t *testing.T) {
		for _, val := range []tdmg.Uint64RawAlias{
			151492823822937,
			113191817,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Uint32Raw", func(t *testing.T) {
		for _, val := range []tdmg.Uint32RawAlias{
			94383726,
			1,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Uint16Raw", func(t *testing.T) {
		for _, val := range []tdmg.Uint16RawAlias{
			929,
			11111,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Uint8Raw", func(t *testing.T) {
		for _, val := range []tdmg.Uint8RawAlias{
			255,
			0,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("UintRaw", func(t *testing.T) {
		for _, val := range []tdmg.UintRawAlias{
			545349,
			11514141414,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Int64Raw", func(t *testing.T) {
		for _, val := range []tdmg.Int64RawAlias{
			-2937437292,
			17620,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Int32Raw", func(t *testing.T) {
		for _, val := range []tdmg.Int32RawAlias{
			-17253,
			0,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Int16Raw", func(t *testing.T) {
		for _, val := range []tdmg.Int16RawAlias{
			-55,
			-3,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Int8Raw", func(t *testing.T) {
		for _, val := range []tdmg.Int8RawAlias{
			-1,
			43,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("IntRaw", func(t *testing.T) {
		for _, val := range []tdmg.IntRawAlias{
			-726,
			10283763525,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Float64Raw", func(t *testing.T) {
		for _, val := range []tdmg.Float64RawAlias{
			0.151492823822937,
			0.113191817,
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Float32Raw", func(t *testing.T) {
		for _, val := range []tdmg.Float32RawAlias{
			0.151492823822937,
			0.113191817,
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
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map with raw keys/values", func(t *testing.T) {
		for _, val := range []tdmg.Uint16Int32RawMapAlias{
			{1: -1, 18: 1827374},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array with raw triple pointers", func(t *testing.T) {
		var (
			n  float64 = 1.181727361
			n1         = &n
			n2         = &n1
			n3         = &n2

			m  float64 = 0.28283839844
			m1         = &m
			m2         = &m1
			m3         = &m2

			k  float64 = 0
			k1         = &k
			k2         = &k1
			k3         = &k2
		)
		for _, val := range []tdmg.Float64RawPtrPtrPtrAliasSliceAlias{
			{n3, m3, k3},
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

			k float64 = -87.21938

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
