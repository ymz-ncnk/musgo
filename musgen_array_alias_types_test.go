package musgo

import (
	"errors"
	"math"
	"testing"

	"github.com/ymz-ncnk/muserrs"
	"github.com/ymz-ncnk/musgo/v2/testdata"
	tdmg "github.com/ymz-ncnk/musgo/v2/testdata/musgen"
)

func TestGeneratedArrayAliasCode(t *testing.T) {

	t.Run("Simple array", func(t *testing.T) {
		for _, val := range []tdmg.StrArrayAlias{
			{"hello", "world", ")(*(*&*&^&^$#"},
			{"", "", ""},
			{"some", "", "12342g2"}} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of pointers", func(t *testing.T) {
		var (
			f1 = -math.MaxFloat64
			f2 = 0.0
			f3 = math.MaxFloat64
			f4 = 11.8476301
			f5 = math.SmallestNonzeroFloat64
			f6 = 2342342.0
		)
		for _, val := range []tdmg.FloatPtrArrayAlias{
			{&f1, &f2, &f3},
			{&f4, &f5, &f6},
			{nil, nil, nil},
			{nil, &f3, nil},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of aliases", func(t *testing.T) {
		for _, val := range []tdmg.IntAliasArrayAlias{
			{math.MinInt, 0, math.MaxInt},
			{12828, -18128, 19292}} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of pointer aliases", func(t *testing.T) {
		var (
			n1 tdmg.Uint64Alias = 181727361
			n2 tdmg.Uint64Alias = math.MaxUint64
			n3 tdmg.Uint64Alias = 0
		)
		for _, val := range []tdmg.Uint64PtrAliasArrayAlias{
			{&n1, &n2, &n3},
			{nil, nil, nil},
			{&n3, nil, nil},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of slices", func(t *testing.T) {
		for _, val := range []tdmg.Int64SliceArrayAlias{
			{
				{123, 0, math.MinInt64},
				{0, 0, 0},
				{-1, -2, -12342},
			},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of pointer slices", func(t *testing.T) {
		var (
			slice1 []uint16 = []uint16{math.MaxUint16, 0, 12}
			slice2 []uint16 = []uint16{0, 0, 0}
			slice3 []uint16 = []uint16{657, 23, 981}
		)
		for _, val := range []tdmg.Uint16SlicePtrArrayAlias{
			{&slice1, &slice2, &slice3},
			{nil, nil, nil},
			{nil, nil, &slice3},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of arrays", func(t *testing.T) {
		for _, val := range []tdmg.BoolArrayArrayAlias{
			{
				{true, false},
				{true, true, true},
			},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of arrays", func(t *testing.T) {
		var (
			arr1 = [2]byte{0x08, math.MaxInt8}
			arr2 = [2]byte{0x00, 0x09}
		)
		for _, val := range []tdmg.BytePtrArrayPtrArrayAlias{
			{&arr1, &arr2},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of maps with keys - arrays", func(t *testing.T) {
		var (
			map1 = map[int]string{1: "some1", math.MaxInt: "some2", 10: "another"}
			map2 = map[int]string{math.MinInt: "some3"}
			map3 = map[int]string{}
		)
		for _, val := range []tdmg.IntStrMapArrayAlias{
			{map1, map2, map3},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of pointer maps", func(t *testing.T) {
		var (
			map1 = map[uint32]int32{math.MaxUint32: math.MaxInt32, 123: -123}
			map2 = map[uint32]int32{7: 1, 120383: -12323423}
			map3 = map[uint32]int32{0: math.MinInt32, 1: 0}
			map4 = map[uint32]int32{7: 1, 120383: -12323423}
		)
		for _, val := range []tdmg.Uint32Int32MapArrayAlias{
			{&map1, &map2, &map3, &map4},
			{nil, nil, nil, nil},
			{nil, &map3, nil, &map3},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of CustomType-s", func(t *testing.T) {
		for _, val := range []tdmg.StructTypeArrayAlias{
			{
				tdmg.SimpleStructType{Int: 17},
				tdmg.SimpleStructType{Int: -123},
			},
			{
				tdmg.SimpleStructType{Int: 11289719871},
				tdmg.SimpleStructType{},
			},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of pointer CustomType-s", func(t *testing.T) {
		for _, val := range []tdmg.StructTypePtrArrayAlias{
			{
				&tdmg.SimpleStructType{Int: math.MinInt},
				&tdmg.SimpleStructType{Int: -123},
				&tdmg.SimpleStructType{Int: -123},
			},
			{
				&tdmg.SimpleStructType{Int: math.MaxInt},
				&tdmg.SimpleStructType{},
				&tdmg.SimpleStructType{Int: -182},
			},
			{nil, nil, nil},
			{nil, &tdmg.SimpleStructType{}, nil},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Tricky array alias", func(t *testing.T) {
		for _, val := range []tdmg.TrickyArrayAlias{
			{
				{
					{
						{
							tdmg.SimpleStructType{Int: 1}: {1},
						},
						{
							tdmg.SimpleStructType{Int: 3}: {2},
						},
						{
							tdmg.SimpleStructType{Int: 1}: {9},
						},
					},
					{
						{
							tdmg.SimpleStructType{Int: 0}: {0},
						},
						{
							tdmg.SimpleStructType{Int: 0}: {233},
						},
						{
							tdmg.SimpleStructType{Int: 112}: {90},
						},
					},
				},
				{
					{
						{
							tdmg.SimpleStructType{Int: 1}: {4},
						},
						{
							tdmg.SimpleStructType{Int: 3}: {2},
						},
						{
							tdmg.SimpleStructType{Int: 1}: {9},
						},
					},
				},
			}} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Zero length array", func(t *testing.T) {
		for _, val := range []tdmg.StrZeroLengthArrayAlias{
			{},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Buffer ends", func(t *testing.T) {
		var val tdmg.StrArrayAlias
		if err := testdata.TestBufferEnds(&val, []byte{}); err != nil {
			t.Error(err)
		}
	})

	t.Run("Array of triple pointer ints", func(t *testing.T) {
		var (
			n  int = 181727361
			n1     = &n
			n2     = &n1
			n3     = &n2
			m  int = -28283839844
			m1     = &m
			m2     = &m1
			m3     = &m2
			c  int = 0
			c1     = &c
			c2     = &c1
			c3     = &c2
		)
		for _, val := range []tdmg.IntPtrPtrPtrAliasArrayAlias{
			{n3, m3, c3},
			{nil, nil, nil},
			{nil, nil, c3},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array elems validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidIntArrayAlias{
			{2, 11},
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				unmarshalErr := errors.Unwrap(err)
				if arrErr, ok := unmarshalErr.(*muserrs.ArrayError); ok {
					if arrErr.Cause() != tdmg.ErrBiggerThanTen {
						t.Errorf("wrong error cause '%v'", arrErr.Cause())
					}
				} else {
					t.Errorf("wrong error '%v'", unmarshalErr)
				}
			}
		}
	})

	t.Run("Array pointer elems validation", func(t *testing.T) {
		var (
			n  int = 64
			n1     = &n
		)
		for _, val := range []tdmg.ValidPtrIntArrayAlias{
			{n1, nil},
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				unmarshalErr := errors.Unwrap(err)
				if arrErr, ok := unmarshalErr.(*muserrs.ArrayError); ok {
					if arrErr.Index() != 1 {
						t.Errorf("wrong error index '%v'", arrErr.Index())
					}
					if arrErr.Cause() != tdmg.ErrNil {
						t.Errorf("wrong error cause '%v'", arrErr.Cause())
					}
				} else {
					t.Errorf("wrong error '%v'", unmarshalErr)
				}
			}
		}
	})

	t.Run("Array validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidIntArrayAlias{
			{8, 6},
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				unmarshalErr := errors.Unwrap(err)
				if unmarshalErr != tdmg.ErrArraySumBiggerThanTen {
					t.Errorf("wrong error '%v'", unmarshalErr)
				}
			}
		}
	})

}
