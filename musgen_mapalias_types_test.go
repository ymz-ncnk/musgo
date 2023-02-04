//go:generate go run testdata/gen/mus.go -map $ARG
package musgo

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ymz-ncnk/serialization/musgo/errs"
	"github.com/ymz-ncnk/serialization/musgo/testdata"
	tdmg "github.com/ymz-ncnk/serialization/musgo/testdata/musgen"
)

func TestGeneratedMapAliasCode(t *testing.T) {

	// map's key can't be function, map or slice
	t.Run("Simple map", func(t *testing.T) {
		for _, val := range []tdmg.StrIntMapAlias{
			{"some": 15},
			{"yeuryrywuww": -15141},
			{"": 0},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map of pointers", func(t *testing.T) {
		var (
			htpg     = "htpg"
			negInt   = -17263
			emptyStr = ""
			zero     = 0
		)
		for _, val := range []tdmg.StrPtrIntPtrMapAlias{
			{&htpg: &negInt},
			{&emptyStr: &zero},
		} {
			aval, err := testdata.ExecGeneratedCode(val)
			if err != nil {
				t.Error(err)
			}
			var (
				m1 = make(map[string]int)
				m2 = make(map[string]int)
			)
			for k, v := range val {
				m1[*k] = *v
			}
			for k, v := range *aval.(*tdmg.StrPtrIntPtrMapAlias) {
				m2[*k] = *v
			}
			for k, v := range m1 {
				if m2[k] != v {
					t.Errorf("Marshal - %v, result of Unmarshal - %v", m1, m2)
				}
			}
		}
	})

	t.Run("Map of aliases", func(t *testing.T) {
		for _, val := range []tdmg.StrAliasIntAliasMapAlias{
			{"sky": 626262},
			{"town": -29283},
			{"": 0},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map of pointer aliases", func(t *testing.T) {
		var (
			str1 tdmg.StringAlias = "1"
			str2 tdmg.StringAlias = "'sow'"
			str3 tdmg.StringAlias = ""
			num1 tdmg.IntAlias    = 8239283422
			num2 tdmg.IntAlias    = 10192838
			num3 tdmg.IntAlias    = 0
		)
		for _, val := range []tdmg.StrAliasPtrIntAliasPtrMapAlias{
			{&str1: &num1},
			{&str2: &num2},
			{&str3: &num3},
		} {
			aval, err := testdata.ExecGeneratedCode(val)
			if err != nil {
				t.Error(err)
			}
			var (
				m1 = make(map[tdmg.StringAlias]tdmg.IntAlias)
				m2 = make(map[tdmg.StringAlias]tdmg.IntAlias)
			)
			for k, v := range val {
				m1[*k] = *v
			}
			for k, v := range *aval.(*tdmg.StrAliasPtrIntAliasPtrMapAlias) {
				m2[*k] = *v
			}
			for k, v := range m1 {
				if m2[k] != v {
					t.Errorf("Marshal - %v, result of Unmarshal - %v", m1, m2)
				}
			}
		}
	})

	t.Run("Map of slices", func(t *testing.T) {
		for _, val := range []tdmg.BoolInt16SliceMapAlias{
			{true: {123, 434, 0}},
			{false: {-12, -123, 123}},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map of pointer slices", func(t *testing.T) {
		for _, val := range []tdmg.ByteUint16SlicePtrMapAlias{
			{0x50: &[]uint16{98, 333, 0}},
			{0x8: &[]uint16{12, 1123, 123}},
			{0x9: &[]uint16{}},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map of arrays", func(t *testing.T) {
		for _, val := range []tdmg.Int32Float64ArrayMapAlias{
			{19283: [2]float64{0.19, 434.937}},
			{0: [2]float64{}},
			{-12847: {-12, -123.3}},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map of pointer arrays", func(t *testing.T) {
		for _, val := range []tdmg.Float32Uint32ArrayPtrMapAlias{
			{19283.12: &[2]uint32{234, 434}},
			{0: &[2]uint32{}},
			{-12847: &[2]uint32{3, 9374}},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map of maps", func(t *testing.T) {
		for _, val := range []tdmg.FloatByteBoolMapMapAlias{
			{56.22: {0x03: true, 0x20: true}},
			{0: {0x00: false}},
			{-123.34: {0x05: false, 0x23: true}},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map of pointer maps", func(t *testing.T) {
		for _, val := range []tdmg.UintIntStringMapPtrMapAlias{
			{56: {-123554: "true", 12837: "city"}},
			{0: {0: ""}},
			{1918: {12918923131211111: "bird", -1: "&&&92-2=!"}},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map of CustomType", func(t *testing.T) {
		for _, val := range []tdmg.StructTypeStructTypeMapAlias{
			{
				tdmg.SimpleStructType{Int: 12}: tdmg.SimpleStructType{Int: 2383},
			},
			{
				tdmg.SimpleStructType{Int: -12}: tdmg.SimpleStructType{Int: 111},
			},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map of pointer CustomType", func(t *testing.T) {
		for _, val := range []tdmg.StructTypePtrStructTypePtrMapAlias{
			{
				&tdmg.SimpleStructType{Int: 12}: &tdmg.SimpleStructType{Int: 2383},
			},
			{
				&tdmg.SimpleStructType{Int: -12}: &tdmg.SimpleStructType{Int: 111},
			},
		} {
			aval, err := testdata.ExecGeneratedCode(val)
			if err != nil {
				t.Error(err)
			}
			var (
				m1 = make(map[tdmg.SimpleStructType]tdmg.SimpleStructType)
				m2 = make(map[tdmg.SimpleStructType]tdmg.SimpleStructType)
			)
			for k, v := range val {
				m1[*k] = *v
			}
			for k, v := range *aval.(*tdmg.StructTypePtrStructTypePtrMapAlias) {
				m2[*k] = *v
			}
			for k, v := range m1 {
				if m2[k] != v {
					t.Errorf("Marshal - %v, result of Unmarshal - %v", m1, m2)
				}
			}
		}
	})

	t.Run("Tricky map", func(t *testing.T) {
		var (
			getVal = func(k *[]tdmg.SimpleStructType,
				m map[*[]tdmg.SimpleStructType]map[*map[int]string][2]int) (
				map[*map[int]string][2]int, bool) {
				for mk, mvl := range m {
					if reflect.DeepEqual(*k, *mk) {
						return mvl, true
					}
				}
				return nil, false
			}
			getValByMap = func(k *map[int]string, m map[*map[int]string][2]int) (
				[2]int, bool) {
				for mk, mvl := range m {
					if reflect.DeepEqual(*k, *mk) {
						return mvl, true
					}
				}
				return [2]int{}, false
			}
		)
		for _, val := range []tdmg.TrickyMapAlias{
			{
				[2]tdmg.StringAlias{"s", "k"}: {
					&[]tdmg.SimpleStructType{
						{Int: 1},
						{Int: 0},
					}: {
						&map[int]string{4: "y", 6: "o", -34: "pp"}: [2]int{
							9,
							-9,
						},
					},
				},
			},
		} {
			aval, err := testdata.ExecGeneratedCode(val)
			if err != nil {
				t.Error(err)
			}
			for k, mv := range val {
				amv := (*aval.(*tdmg.TrickyMapAlias))[k]
				for k1, vm1 := range mv {
					amv1, pst := getVal(k1, amv)
					if !pst {
						t.Errorf("Marshal - %v, result of Unmarshal - %v", val, aval)
					}
					for km2, v2 := range vm1 {
						av2, pst := getValByMap(km2, amv1)
						if !pst {
							t.Errorf("Marshal - %v, result of Unmarshal - %v", val, aval)
						}
						if v2 != av2 {
							t.Errorf("Marshal - %v, result of Unmarshal - %v", val, aval)
						}
					}
				}
			}
		}
	})

	t.Run("Buffer ends", func(t *testing.T) {
		var val tdmg.StrIntMapAlias
		err := testdata.TestBufferEnds(&val, []byte{})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Length is too big", func(t *testing.T) {
		var val tdmg.StrIntMapAlias
		buf := []byte{12, 6, 1, 1, 1, 12}
		err := testdata.TestBufferEnds(&val, buf)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Map keys validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStringIntMapAlias{
			{"one": 1, "hello": 0},
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				unmarshalErr := errors.Unwrap(err)
				if mapValueErr, ok := unmarshalErr.(*errs.MapKeyError); ok {
					if mapValueErr.Cause() != tdmg.ErrStrIsHello {
						t.Errorf("wrong error cause '%v'", mapValueErr.Cause())
					}
				} else {
					t.Errorf("wrong error '%v'", unmarshalErr)
				}
			}
		}
	})

	t.Run("Map values validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStringIntMapAlias{
			{"one": 1, "eleven": 11},
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				unmarshalErr := errors.Unwrap(err)
				if mapValueErr, ok := unmarshalErr.(*errs.MapValueError); ok {
					if mapValueErr.Cause() != tdmg.ErrBiggerThanTen {
						t.Errorf("wrong error cause '%v'", mapValueErr.Cause())
					}
				} else {
					t.Errorf("wrong error '%v'", unmarshalErr)
				}
			}
		}
	})

	t.Run("Map validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStringIntMapAlias{
			{"seven": 7, "six": 6},
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				unmarshalErr := errors.Unwrap(err)
				if unmarshalErr != tdmg.ErrMapSumBiggerThanTen {
					t.Errorf("wrong error '%v'", unmarshalErr)
				}
			}
		}
	})

	t.Run("Map length validation", func(t *testing.T) {
		for _, val := range []tdmg.ValidStringIntMapAlias{
			{"one": 1, "zero": 0, "two": 1, "three": 3},
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				unmarshalErr := errors.Unwrap(err)
				if unmarshalErr != errs.ErrMaxLengthExceeded {
					t.Errorf("wrong error '%v'", unmarshalErr)
				}
			}
		}
	})

}
