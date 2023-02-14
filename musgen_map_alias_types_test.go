//go:generate go run testdata/gen/mus.go -map $ARG
package musgo

import (
	"errors"
	"math"
	"reflect"
	"testing"

	"github.com/ymz-ncnk/muserrs"
	"github.com/ymz-ncnk/musgo/v2/testdata"
	tdmg "github.com/ymz-ncnk/musgo/v2/testdata/musgen"
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
			{"sky": math.MaxInt},
			{"town": math.MinInt},
			{"yetee": -1028474},
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
			str4 tdmg.StringAlias = "st"
			num1 tdmg.IntAlias    = math.MaxInt
			num2 tdmg.IntAlias    = math.MinInt
			num3 tdmg.IntAlias    = 0
		)
		for _, val := range []tdmg.StrAliasPtrIntAliasPtrMapAlias{
			{&str1: &num1, nil: nil, &str4: nil},
			{&str2: &num2},
			{&str3: &num3},
			{nil: nil},
			{&str4: nil},
		} {
			aval, err := testdata.ExecGeneratedCode(val)
			if err != nil {
				t.Error(err)
			}
			if val[nil] != (*aval.(*tdmg.StrAliasPtrIntAliasPtrMapAlias))[nil] {
				t.Error("not equal nil values")
			}
			delete(val, nil)
			delete((*aval.(*tdmg.StrAliasPtrIntAliasPtrMapAlias)), nil)
			var (
				m1 = make(map[tdmg.StringAlias]tdmg.IntAlias)
				m2 = make(map[tdmg.StringAlias]tdmg.IntAlias)
			)
			for k, v := range val {
				if v != nil {
					m1[*k] = *v
				}
			}
			for k, v := range *aval.(*tdmg.StrAliasPtrIntAliasPtrMapAlias) {
				if v == nil {
					if *k != str4 {
						t.Error("wrong key have nil value")
					}
				} else {
					m2[*k] = *v
				}
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
			{true: {123, math.MaxInt16, 0}},
			{false: {-12, math.MinInt16, 123}},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map of pointer slices", func(t *testing.T) {
		for _, val := range []tdmg.ByteUint16SlicePtrMapAlias{
			{0x50: &[]uint16{98, 333, 0}, 0x3: nil, 0x4: nil},
			{0x8: &[]uint16{12, 1123, math.MaxUint16}},
			{0x9: &[]uint16{}},
			{0x51: nil},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map of arrays", func(t *testing.T) {
		for _, val := range []tdmg.Int32Float64ArrayMapAlias{
			{math.MaxInt32: [2]float64{0.19, 434.937}},
			{0: [2]float64{}},
			{-12847: {-12, -math.MaxFloat64}},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map of pointer arrays", func(t *testing.T) {
		for _, val := range []tdmg.Float32Uint32ArrayPtrMapAlias{
			{0.8: nil, 19283.12: &[2]uint32{234, 434}, 10.1: nil},
			{0: &[2]uint32{}},
			{-12847: &[2]uint32{3, 9374}},
			{5.5: nil, 4.4: nil, 0: nil},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map of maps", func(t *testing.T) {
		for _, val := range []tdmg.FloatByteBoolMapMapAlias{
			{-math.MaxFloat32: {0x03: true, 0x20: true}},
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
			{math.MaxUint16: nil, 56: {-123554: "true", math.MaxInt: "city"}},
			{0: {0: ""}},
			{1918: {12918923131211111: "bird", -1: "&&&92-2=!", math.MinInt: "iiii"}},
			{2: nil, 0: nil},
		} {
			if err := testdata.TestGeneratedCode(val); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Map of CustomType", func(t *testing.T) {
		for _, val := range []tdmg.StructTypeStructTypeMapAlias{
			{
				tdmg.SimpleStructType{Int: math.MinInt}: tdmg.SimpleStructType{Int: math.MaxInt},
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
				&tdmg.SimpleStructType{Int: 12}:          &tdmg.SimpleStructType{Int: math.MinInt},
				nil:                                      nil,
				&tdmg.SimpleStructType{Int: math.MaxInt}: &tdmg.SimpleStructType{Int: -2},
			},
			{
				&tdmg.SimpleStructType{Int: -12}: &tdmg.SimpleStructType{Int: 111},
				nil:                              nil,
			},
		} {
			aval, err := testdata.ExecGeneratedCode(val)
			if err != nil {
				t.Error(err)
			}
			if val[nil] != (*aval.(*tdmg.StructTypePtrStructTypePtrMapAlias))[nil] {
				t.Error("not equal nil values")
			}
			delete(val, nil)
			delete((*aval.(*tdmg.StructTypePtrStructTypePtrMapAlias)), nil)

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
				if mapKeyErr, ok := unmarshalErr.(*muserrs.MapKeyError); ok {
					if mapKeyErr.Cause() != tdmg.ErrStrIsHello {
						t.Errorf("wrong error cause '%v'", mapKeyErr.Cause())
					}
				} else {
					t.Errorf("wrong error '%v'", unmarshalErr)
				}
			}
		}
	})

	t.Run("Map pointer keys validation", func(t *testing.T) {
		var (
			n  = 1
			n1 = &n
			m  = 2
			m1 = &m
		)
		for _, val := range []tdmg.ValidPtrIntPtrIntMapAlias{
			{m1: m1, nil: n1},
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				unmarshalErr := errors.Unwrap(err)
				if mapKeyErr, ok := unmarshalErr.(*muserrs.MapKeyError); ok {
					if !reflect.ValueOf(mapKeyErr.Key()).IsNil() {
						t.Errorf("wrong error key '%v'", mapKeyErr.Key())
					}
					if mapKeyErr.Cause() != tdmg.ErrNil {
						t.Errorf("wrong error cause '%v'", mapKeyErr.Cause())
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
				if mapValueErr, ok := unmarshalErr.(*muserrs.MapValueError); ok {
					if mapValueErr.Cause() != tdmg.ErrBiggerThanTen {
						t.Errorf("wrong error cause '%v'", mapValueErr.Cause())
					}
				} else {
					t.Errorf("wrong error '%v'", unmarshalErr)
				}
			}
		}
	})

	t.Run("Map pointer values validation", func(t *testing.T) {
		var (
			n  = 1
			n1 = &n
			m  = 2
			m1 = &m
		)
		for _, val := range []tdmg.ValidPtrIntPtrIntMapAlias{
			{m1: nil, n1: n1},
		} {
			if _, err := testdata.ExecGeneratedCode(val); err != nil {
				unmarshalErr := errors.Unwrap(err)
				if mapValueErr, ok := unmarshalErr.(*muserrs.MapValueError); ok {
					if *mapValueErr.Key().(*int) != m {
						t.Errorf("wrong error key '%v'", mapValueErr.Cause())
					}
					if mapValueErr.Cause() != tdmg.ErrNil {
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
				if unmarshalErr != muserrs.ErrMaxLengthExceeded {
					t.Errorf("wrong error '%v'", unmarshalErr)
				}
			}
		}
	})

}
