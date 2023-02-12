//go:generate go run testdata/gen/mus.go -slice $ARG
package musgo

import (
	"errors"
	"math"
	"reflect"
	"testing"

	"github.com/ymz-ncnk/musgo/errs"
	"github.com/ymz-ncnk/musgo/testdata"
	tdmg "github.com/ymz-ncnk/musgo/testdata/musgen"
)

func TestGeneratedSliceAliasCode(t *testing.T) {

	t.Run("Simple slice", func(t *testing.T) {
		for _, val := range []tdmg.StrSliceAlias{
			{"hello", "world"},
			{"", "tekejt"},
			{
				"hello", "hello", "hello", "hello", "hlo", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "h", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",

				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "helo", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
				"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello",
			},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Slice of pointers", func(t *testing.T) {
		var (
			n = math.MaxFloat64
			m = 0.0
			k = -0.19283
		)
		for _, val := range []tdmg.FloatPtrSliceAlias{
			{&n, &m, &k},
			{nil, nil, nil},
			{nil, &k, nil},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Slice of aliases", func(t *testing.T) {
		for _, val := range []tdmg.IntAliasSliceAlias{
			{1, 0, -21923},
			{12828, -18128, 19292},
			{math.MaxInt, math.MinInt, 19292},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Slice of pointer aliases", func(t *testing.T) {
		var (
			n tdmg.Uint64Alias = 181727361
			m tdmg.Uint64Alias = math.MaxUint64
			k tdmg.Uint64Alias = 0
		)
		for _, val := range []tdmg.Uint64PtrAliasSliceAlias{
			{&n, &m, &k},
			{nil, nil, nil},
			{&k, nil, &k},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Slice of slices", func(t *testing.T) {
		for _, val := range []tdmg.BoolSliceSliceAlias{
			{
				{true, false},
				{true, true, true},
				{
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, false, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, false, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
					true, true, true, true, true, true, true, true, true, true,
				},
			},
			{
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},

				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
				{true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true}, {true},
			},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Unmarshal slice of slices with empty buf", func(t *testing.T) {
		var val tdmg.BoolSliceSliceAlias
		_, err := val.Unmarshal([]byte{})
		if err != errs.ErrSmallBuf {
			t.Errorf("unexpected '%v' error", err)
		}
	})

	t.Run("Slice of arrays", func(t *testing.T) {
		for _, val := range []tdmg.ByteArraySliceAlias{
			{
				[2]byte{0x01, 0x60},
				[2]byte{0x50, math.MaxInt8},
			},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Slice of pointer arrays", func(t *testing.T) {
		for _, val := range []tdmg.FloatArrayPtrSliceAlias{
			{
				&[2]float32{102.1826, -123.12},
				&[2]float32{10, 0.2836},
			},
			{nil, nil, nil},
			{&[2]float32{100, -0.2}, nil, nil},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Slice of maps", func(t *testing.T) {
		for _, val := range []tdmg.IntStrMapSliceAlias{
			{
				map[int16]string{345: "io"},
				map[int16]string{19292: "poll"},
			},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Slice of pointer maps", func(t *testing.T) {
		for _, val := range []tdmg.Uint32Int32MapPtrSliceAlias{
			{
				&map[uint32]int32{345: 82},
				&map[uint32]int32{19292: 893837},
			},
			{nil},
			{nil, &map[uint32]int32{10: 2}},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Slice of CustomType-s", func(t *testing.T) {
		for _, val := range []tdmg.StructTypeSliceAlias{
			{
				tdmg.SimpleStructType{Int: 8},
				tdmg.SimpleStructType{Int: 23429},
			},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Slice of pointer CustomType-s", func(t *testing.T) {
		for _, val := range []tdmg.StructTypePtrSliceAlias{
			{
				&tdmg.SimpleStructType{Int: 8},
				&tdmg.SimpleStructType{Int: 23429},
			},
			{nil, nil},
			{&tdmg.SimpleStructType{Int: -2}, nil},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Tricky slice", func(t *testing.T) {
		var (
			getVal = func(k *[]tdmg.StringAlias,
				m map[*[]tdmg.StringAlias]map[tdmg.SimpleStructType][]int) (
				map[tdmg.SimpleStructType][]int, bool) {
				for mk, mvl := range m {
					if reflect.DeepEqual(*k, *mk) {
						return mvl, true
					}
				}
				return nil, false
			}
		)
		val := tdmg.TrickySliceAlias{
			{
				{
					&[]tdmg.StringAlias{"s"}: {
						tdmg.SimpleStructType{Int: 1}: []int{23},
					},
				},
			},
			{
				{
					&[]tdmg.StringAlias{"t"}: {
						tdmg.SimpleStructType{Int: -1123}: []int{63532},
					},
				},
			},
		}
		aval, err := testdata.ExecGeneratedCode(val)
		if err != nil {
			t.Error(err)
		}
		sval := *aval.(*tdmg.TrickySliceAlias)
		for i, va := range val {
			for j, vam := range va {
				for k, vamm := range vam {
					avamm, pst := getVal(k, sval[i][j])
					if !pst {
						t.Errorf("Marshal - %v, result of Unmarshal - %v", val, aval)
					}
					for k2, v2 := range vamm {
						if !reflect.DeepEqual(v2, avamm[k2]) {
							t.Errorf("Marshal - %v, result of Unmarshal - %v", val, aval)
						}
					}
				}
			}
		}
	})

	t.Run("Buffer ends", func(t *testing.T) {
		var val tdmg.StrSliceAlias
		err := testdata.TestBufferEnds(&val, []byte{})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Length is too big", func(t *testing.T) {
		var val tdmg.StrSliceAlias
		buf := []byte{12, 6, 1, 1, 1}
		err := testdata.TestBufferEnds(&val, buf)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Slice elems validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidUintSliceAlias{
			{2, 11},
		} {
			_, err := testdata.ExecGeneratedCode(val)
			unmarshalErr := errors.Unwrap(err)
			if sliceElemErr, ok := unmarshalErr.(*errs.SliceError); ok {
				if sliceElemErr.Cause() != tdmg.ErrBiggerThanTen {
					t.Errorf("wrong error cause '%v'", sliceElemErr.Cause())
				}
			} else {
				t.Errorf("wrong error '%v'", unmarshalErr)
			}
		}
	})

	t.Run("Slice pointer elems validation", func(t *testing.T) {
		var (
			s  = "str"
			s1 = &s
		)
		for _, val := range []tdmg.ValidPtrStringSliceAlias{
			{s1, nil},
		} {
			_, err := testdata.ExecGeneratedCode(val)
			unmarshalErr := errors.Unwrap(err)
			if sliceElemErr, ok := unmarshalErr.(*errs.SliceError); ok {
				if sliceElemErr.Index() != 1 {
					t.Errorf("wrong error index '%v'", sliceElemErr.Index())
				}
				if sliceElemErr.Cause() != tdmg.ErrNil {
					t.Errorf("wrong error cause '%v'", sliceElemErr.Cause())
				}
			} else {
				t.Errorf("wrong error '%v'", unmarshalErr)
			}
		}
	})

	t.Run("Slice validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidUintSliceAlias{
			{8, 6},
		} {
			_, err := testdata.ExecGeneratedCode(val)
			unmarshalErr := errors.Unwrap(err)
			if unmarshalErr != tdmg.ErrSliceSumBiggerThanTen {
				t.Errorf("wrong error '%v'", unmarshalErr)
			}
		}
	})

	t.Run("Slice length validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidUintSliceAlias{
			{1, 0, 0, 0},
		} {
			_, err := testdata.ExecGeneratedCode(val)
			unmarshalErr := errors.Unwrap(err)
			if unmarshalErr != errs.ErrMaxLengthExceeded {
				t.Errorf("wrong error '%v'", unmarshalErr)
			}
		}
	})

}
