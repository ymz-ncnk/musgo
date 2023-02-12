//go:generate go run testdata/gen/mus.go -struct $ARG
package musgo

import (
	"math"
	"testing"

	"github.com/ymz-ncnk/musgo/errs"
	"github.com/ymz-ncnk/musgo/testdata"
	tdmg "github.com/ymz-ncnk/musgo/testdata/musgen"
)

func TestGeneratedStructCode(t *testing.T) {

	makeValidStrucType := func() tdmg.ValidStructType {
		var (
			str  = "str"
			str1 = &str
		)
		return tdmg.ValidStructType{
			StringPtr: str1,
			SlicePtr:  &[]uint16{},
			ArrayPtr:  &[2]int32{},
			MapPtr:    &map[string]int{},
			StructPtr: &tdmg.SimpleStructType{},
		}
	}

	t.Run("Struct", func(t *testing.T) {
		var (
			un  uint = 51
			un1      = &un
			un2      = &un1

			n  int           = 43
			n1               = &n
			n2               = &n1
			na tdmg.IntAlias = 5

			str  = "strte"
			str1 = &str
			str2 = &str1

			b  byte = 0x09
			b1      = &b
			b2      = &b1

			bl  bool = true
			bl1      = &bl
			bl2      = &bl1

			sl  []uint16 = []uint16{1123, 0, math.MaxUint16}
			sl1          = &sl
			sl2          = &sl1

			ar  [2]int32 = [2]int32{123123123, math.MinInt32}
			ar1          = &ar
			ar2          = &ar1

			m map[string]int = map[string]int{"world": 20, "string": math.MaxInt,
				"small": math.MinInt}
			m1 = &m
			m2 = &m1

			st  tdmg.SimpleStructType = tdmg.SimpleStructType{Int: 99}
			st1                       = &st
			st2                       = &st1
		)
		for _, val := range []tdmg.StructType{
			{
				Uint:          12,
				Uint16:        12032,
				Uint32:        123912913,
				Uint64:        1231233333333333333,
				UintPtr:       &un,
				UintPtrPtrPtr: &un2,

				Int:          -1123,
				Int16:        0,
				Int32:        1231231231,
				Int64:        -12309823487234892,
				IntPtr:       &n,
				IntPtrPtrPtr: &n2,
				IntAliasPtr:  &na,

				String:          "some antoherlkjalkja setlwkethwleh lkjwelkjwelr",
				StringPtr:       &str,
				StringPtrPtrPtr: &str2,

				Byte:          0x10,
				BytePtr:       &b,
				BytePtrPtrPtr: &b2,

				Bool:          true,
				BoolPtr:       &bl,
				BoolPtrPtrPtr: &bl2,

				Slice:          []uint{123, 345},
				SlicePtr:       &sl,
				SlicePtrPtrPtr: &sl2,

				Array:          [2]int{-123, -345},
				ArrayPtr:       &ar,
				ArrayPtrPtrPtr: &ar2,

				Map:          map[string]int{"hello": 5},
				MapPtr:       &m,
				MapPtrPtrPtr: &m2,

				Struct:          tdmg.SimpleStructType{Int: 28},
				StructPtr:       &st,
				StructPtrPtrPtr: &st2,

				Tricky: [2]map[[2]tdmg.IntAlias]map[tdmg.StringAlias][2]string{
					{[2]tdmg.IntAlias{1, 2}: {"some": [2]string{"s", "r"}}},
					{[2]tdmg.IntAlias{150, 2222}: {"an": [2]string{"erer", "qqqpq"}}},
				},
			},
			{
				Uint:          12,
				Uint16:        12032,
				Uint32:        123912913,
				Uint64:        1231233333333333333,
				UintPtr:       nil,
				UintPtrPtrPtr: nil,

				Int:          -1123,
				Int16:        0,
				Int32:        1231231231,
				Int64:        -12309823487234892,
				IntPtr:       nil,
				IntPtrPtrPtr: nil,
				IntAliasPtr:  nil,

				String:          "some antoherlkjalkja setlwkethwleh lkjwelkjwelr",
				StringPtr:       nil,
				StringPtrPtrPtr: nil,

				Byte:          0x10,
				BytePtr:       nil,
				BytePtrPtrPtr: nil,

				Bool:          true,
				BoolPtr:       nil,
				BoolPtrPtrPtr: nil,

				Slice:          []uint{123, 345},
				SlicePtr:       nil,
				SlicePtrPtrPtr: nil,

				Array:          [2]int{-123, -345},
				ArrayPtr:       nil,
				ArrayPtrPtrPtr: nil,

				Map:          map[string]int{"hello": 5},
				MapPtr:       nil,
				MapPtrPtrPtr: nil,

				Struct:          tdmg.SimpleStructType{Int: 28},
				StructPtr:       nil,
				StructPtrPtrPtr: nil,

				Tricky: [2]map[[2]tdmg.IntAlias]map[tdmg.StringAlias][2]string{
					{[2]tdmg.IntAlias{1, 2}: {"slkjer": [2]string{"s", "r"}}},
					{[2]tdmg.IntAlias{150, 2222}: {"an": [2]string{"ake", "qqqpq"}}},
				},
			},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}

	})

	t.Run("Buffer ends", func(t *testing.T) {
		{
			var val tdmg.StructType
			err := testdata.TestBufferEnds(&val, []byte{})
			if err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Struct uint field validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.Uint64 = 120
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "Uint64", tdmg.ErrBiggerThanTen, t)
			}
		}
	})

	t.Run("Struct int field validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.Int8 = 19
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "Int8", tdmg.ErrBiggerThanTen, t)
			}
		}
	})

	t.Run("Struct string field validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.String = "hello w"
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "String", tdmg.ErrNotEmptyString, t)
			}
		}
	})

	t.Run("Struct pointer string field validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.StringPtr = nil
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "StringPtr", tdmg.ErrNil, t)
			}
		}
	})

	t.Run("Struct byte field validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.Byte = 21
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "Byte", tdmg.ErrBiggerThanTen, t)
			}
		}
	})

	t.Run("Struct bool field validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.Bool = true
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "Bool", tdmg.ErrPositiveBool, t)
			}
		}
	})

	t.Run("Struct slice field validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.Slice = []uint{7, 4}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "Slice",
					tdmg.ErrSliceSumBiggerThanTen,
					t)
			}
		}
	})

	t.Run("Struct slice field elem validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.Slice = []uint{5, 11}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestSliceElemErr(err, "Slice", 1, tdmg.ErrBiggerThanTen, t)
			}
		}
	})

	t.Run("Struct pointer slice field validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.SlicePtr = &[]uint16{8, 8}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "SlicePtr",
					tdmg.ErrSliceSumBiggerThanTen,
					t)
			}
		}
	})

	t.Run("Struct pointer slice field elem validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.SlicePtr = &[]uint16{8, 11, 8}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestSliceElemErr(err, "SlicePtr", 1, tdmg.ErrBiggerThanTen, t)
			}
		}
	})

	t.Run("Struct array field validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.Array = [2]int{7, 4}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "Array",
					tdmg.ErrArraySumBiggerThanTen,
					t)
			}
		}
	})

	t.Run("Struct array field elem validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.Array = [2]int{5, 11}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestSliceElemErr(err, "Array", 1, tdmg.ErrBiggerThanTen, t)
			}
		}
	})

	t.Run("Struct pointer array field validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.ArrayPtr = &[2]int32{4, 8}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "ArrayPtr",
					tdmg.ErrArraySumBiggerThanTen,
					t)
			}
		}
	})

	t.Run("Struct pointer array field elem validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.ArrayPtr = &[2]int32{12, 8}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestSliceElemErr(err, "ArrayPtr", 0, tdmg.ErrBiggerThanTen, t)
			}
		}
	})

	t.Run("Struct map field validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.Map = map[string]int{"str1": 5, "str2": 10}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "Map", tdmg.ErrMapSumBiggerThanTen,
					t)
			}
		}
	})

	t.Run("Struct map field key validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.Map = map[string]int{"": 5, "hello": 9, "rt": 1}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestMapKeyErr(err, "Map", "hello", tdmg.ErrStrIsHello, t)
			}
		}
	})

	t.Run("Struct map field value validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.Map = map[string]int{"": 5, "not hello": 10, "rt": 88}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestMapValueErr(err, "Map", "rt", 88, tdmg.ErrBiggerThanTen, t)
			}
		}
	})

	t.Run("Struct pointer map field validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.MapPtr = &map[string]int{"str1": 5, "str2": 10}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "MapPtr",
					tdmg.ErrMapSumBiggerThanTen,
					t)
			}
		}
	})

	t.Run("Struct pointer map field key validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.MapPtr = &map[string]int{"": 5, "hello": 10, "rt": 1}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestMapKeyErr(err, "MapPtr", "hello", tdmg.ErrStrIsHello, t)
			}
		}
	})

	t.Run("Struct pointer map field value validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.MapPtr = &map[string]int{"": 5, "not hello": 10, "rt": 88}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestMapValueErr(err, "MapPtr", "rt", 88,
					tdmg.ErrBiggerThanTen,
					t)
			}
		}
	})

	t.Run("Struct validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.Struct = tdmg.SimpleStructType{Int: 15}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "Struct", tdmg.ErrSimpleStructType,
					t)
			}
		}
	})

	t.Run("Struct validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.StructPtr = &tdmg.SimpleStructType{Int: 15}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "StructPtr", tdmg.ErrSimpleStructType,
					t)
			}
		}
	})

	t.Run("Struct string field length validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.String = "qwertyuiopa"
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "String", errs.ErrMaxLengthExceeded,
					t)
			}
		}
	})

	t.Run("Struct slice field length validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.Slice = []uint{1, 1, 1, 0, 0, 0, 0, 0, 0}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "Slice", errs.ErrMaxLengthExceeded,
					t)
			}
		}
	})

	t.Run("Struct map field length validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.Map = map[string]int{"str1": 0, "str2": 1, "str3": 2, "str4": 3,
					"str5": -1, "str6": 0}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "Map", errs.ErrMaxLengthExceeded, t)
			}
		}
	})

	t.Run("Struct pointer map field length validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStructType{
			func() tdmg.ValidStructType {
				st := makeValidStrucType()
				st.MapPtr = &map[string]int{"str1": 0, "str2": 0, "str3": -10,
					"str4": 1, "str5": -1, "str6": -2}
				return st
			}(),
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				testdata.TestFieldErrAndCause(err, "MapPtr", errs.ErrMaxLengthExceeded,
					t)
			}
		}
	})

	t.Run("Struct without fields", func(t *testing.T) {
		for _, val := range []tdmg.FieldlessStructType{
			{Number: 5, Slice: []int{1, 2}},
		} {
			aval, err := testdata.ExecGeneratedCode(val)
			if err != nil {
				t.Error("validation didn't work")
			} else {
				sval := aval.(*tdmg.FieldlessStructType)
				if sval.Number != 0 {
					t.Error("there is a content for FieldlessStructType")
				}
				if sval.Slice != nil {
					t.Error("there is a content for FieldlessStructType")
				}
			}
		}
	})

}
