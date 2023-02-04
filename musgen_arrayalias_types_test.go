//go:generate go run testdata/gen/mus.go -array $ARG
package musgo

import (
	"errors"
	"testing"

	"github.com/ymz-ncnk/serialization/musgo/errs"
	"github.com/ymz-ncnk/serialization/musgo/testdata"
	tdmg "github.com/ymz-ncnk/serialization/musgo/testdata/musgen"
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

	t.Run("Srray of pointers", func(t *testing.T) {
		var (
			f1 = 191.10
			f2 = 0.0
			f3 = -0.19283
			f4 = 1.0
			f5 = -181928.0 // min value
			f6 = 2342342.0 // max value
		)
		for _, val := range []tdmg.FloatPtrArrayAlias{
			{&f1, &f2, &f3},
			{&f4, &f5, &f6}} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of aliases", func(t *testing.T) {
		for _, val := range []tdmg.IntAliasArrayAlias{
			{1, 0, -21923},
			{12828, -18128, 19292}} { // mix,max value
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of pointer aliases", func(t *testing.T) {
		var (
			foo tdmg.Uint64Alias = 181727361
			bar tdmg.Uint64Alias = 28283839844
			car tdmg.Uint64Alias = 0
		)
		for _, val := range []tdmg.Uint64PtrAliasArrayAlias{
			{&foo, &bar, &car},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of slices", func(t *testing.T) {
		for _, val := range []tdmg.Int64SliceArrayAlias{
			{
				{123, 0, 12342},
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
			slice1 []uint16 = []uint16{7788, 0, 12}
			slice2 []uint16 = []uint16{0, 0, 0}
			slice3 []uint16 = []uint16{657, 23, 981}
		)
		for _, val := range []tdmg.Uint16SlicePtrArrayAlias{
			{
				&slice1,
				&slice2,
				&slice3,
			},
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
			arr1 = [2]byte{0x08, 0x55}
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
			map1 = map[int]string{1: "some1", 876: "some2"}
			map2 = map[int]string{2: "some3"}
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
			map1 = map[uint32]int32{1927: 0, 123: -123}
			map2 = map[uint32]int32{7: 1, 120383: -12323423}
			map3 = map[uint32]int32{0: 0, 1: 0}
			map4 = map[uint32]int32{7: 1, 120383: -12323423}
		)
		for _, val := range []tdmg.Uint32Int32MapArrayAlias{
			{&map1, &map2, &map3, &map4},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array of CustomType", func(t *testing.T) {
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

	t.Run("Array of pointer CustomType", func(t *testing.T) {
		for _, val := range []tdmg.StructTypePtrArrayAlias{
			{
				&tdmg.SimpleStructType{Int: -1289371},
				&tdmg.SimpleStructType{Int: -123},
				&tdmg.SimpleStructType{Int: -123},
			},
			{
				&tdmg.SimpleStructType{Int: 999999},
				&tdmg.SimpleStructType{},
				&tdmg.SimpleStructType{Int: -182},
			},
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

	t.Run("Array of double pointer ints", func(t *testing.T) {
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
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Array elem validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidIntArrayAlias{{2, 11}} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				unmarshalErr := errors.Unwrap(err)
				if arrErr, ok := unmarshalErr.(*errs.ArrayError); ok {
					if arrErr.Cause() != tdmg.ErrBiggerThanTen {
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
