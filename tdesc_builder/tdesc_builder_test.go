package tdesc_builder

import (
	"reflect"
	"testing"

	"github.com/ymz-ncnk/musgen"
)

func TestTypeDescBuilder(t *testing.T) {

	t.Run("Invalid tag", func(t *testing.T) {
		type Struct struct {
			Field string `mus:"one,two,three,four,five"`
		}
		want := NewInvalidTagFormatError("Field")
		testInvalidTag(reflect.TypeOf((*Struct)(nil)).Elem(), want, t)
	})

	t.Run("Invalid tag part", func(t *testing.T) {
		type Struct struct {
			Field string `mus:"one#two#three"`
		}
		want := NewInvalidTagFormatError("Field")
		testInvalidTag(reflect.TypeOf((*Struct)(nil)).Elem(), want, t)
	})

	t.Run("Invalid skip tag", func(t *testing.T) {
		type Struct struct {
			Field string `mus:"-,one"`
		}
		want := NewInvalidTagFormatError("Field")
		testInvalidTag(reflect.TypeOf((*Struct)(nil)).Elem(), want, t)
	})

	t.Run("MaxLength for array field", func(t *testing.T) {
		type Struct struct {
			Field [5]int `mus:",10"`
		}
		want := NewInvalidTagFormatError("Field")
		testInvalidTag(reflect.TypeOf((*Struct)(nil)).Elem(), want, t)
	})

	t.Run("MaxLength for alias string field", func(t *testing.T) {
		type StringAlias string
		type Struct struct {
			Field StringAlias `mus:",10"`
		}
		want := NewInvalidTagFormatError("Field")
		testInvalidTag(reflect.TypeOf((*Struct)(nil)).Elem(), want, t)
	})

	t.Run("Negative MaxLength for map field", func(t *testing.T) {
		type Struct struct {
			Field map[string]string `mus:",-10"`
		}
		want := ErrNegativeMaxLength
		testInvalidTagPart(reflect.TypeOf((*Struct)(nil)).Elem(), want, t)
	})

	t.Run("ElemValidator for simple type field", func(t *testing.T) {
		type Struct struct {
			Field string `mus:",,elemValidator"`
		}
		want := NewInvalidTagFormatError("Field")
		testInvalidTag(reflect.TypeOf((*Struct)(nil)).Elem(), want, t)
	})

	t.Run("ElemValidator for alias field", func(t *testing.T) {
		type SliceAlias []int
		type Struct struct {
			Field SliceAlias `mus:",,elemValidator"`
		}
		want := NewInvalidTagFormatError("Field")
		testInvalidTag(reflect.TypeOf((*Struct)(nil)).Elem(), want, t)
	})

	t.Run("KeyValidator for simple type field", func(t *testing.T) {
		type Struct struct {
			Field int `mus:",,,keyValidator"`
		}
		want := NewInvalidTagFormatError("Field")
		testInvalidTag(reflect.TypeOf((*Struct)(nil)).Elem(), want, t)
	})

	t.Run("KeyValidator for array field", func(t *testing.T) {
		type Struct struct {
			Field [5]string `mus:",,,keyValidator"`
		}
		want := NewInvalidTagFormatError("Field")
		testInvalidTag(reflect.TypeOf((*Struct)(nil)).Elem(), want, t)
	})

	t.Run("KeyValidator for slice field", func(t *testing.T) {
		type Struct struct {
			Field []int `mus:",,,keyValidator"`
		}
		want := NewInvalidTagFormatError("Field")
		testInvalidTag(reflect.TypeOf((*Struct)(nil)).Elem(), want, t)
	})

	t.Run("KeyValidator for map alias field", func(t *testing.T) {
		type MapAlias map[string]string
		type Struct struct {
			Field MapAlias `mus:",,,keyValidator"`
		}
		want := NewInvalidTagFormatError("Field")
		testInvalidTag(reflect.TypeOf((*Struct)(nil)).Elem(), want, t)
	})

	t.Run("Struct fields with different tags", func(t *testing.T) {
		type Alias int
		type Struct struct {
			Field0 int               `mus:"-"`
			Field1 string            `mus:"validator1,1"`
			Field2 map[float64]uint8 `mus:"validator2,2,elemValidator2#raw,keyValidator2#raw"`
			Field3 string            `mus:"validator3,40"`
			Field4 [4]int            `mus:"validator4,,elemValidator4"`
			Field5 []uint            `mus:"validator5,5,elemValidator5"`
			Field6 Alias             `mus:"validator6"`
			Field7 int               `mus:""`
			Field8 float32
		}
		want := musgen.TypeDesc{
			Package: "tdesc_builder",
			Name:    "Struct",
			Unsafe:  false,
			Suffix:  "",
			Fields: []musgen.FieldDesc{
				{
					Name:      "Field1",
					Type:      "string",
					Validator: "validator1",
					MaxLength: 1,
				},
				{
					Name:          "Field2",
					Type:          "map-0[float64]-0uint8",
					Validator:     "validator2",
					MaxLength:     2,
					ElemValidator: "elemValidator2",
					ElemEncoding:  "raw",
					KeyValidator:  "keyValidator2",
					KeyEncoding:   "raw",
				},
				{
					Name:      "Field3",
					Type:      "string",
					Validator: "validator3",
					MaxLength: 40,
				},
				{
					Name:          "Field4",
					Type:          "[4]int",
					Validator:     "validator4",
					ElemValidator: "elemValidator4",
				},
				{
					Name:          "Field5",
					Type:          "[]uint",
					Validator:     "validator5",
					MaxLength:     5,
					ElemValidator: "elemValidator5",
				},
				{
					Name:      "Field6",
					Type:      "Alias",
					Validator: "validator6",
				},
				{
					Name: "Field7",
					Type: "int",
				},
				{
					Name: "Field8",
					Type: "float32",
				},
			},
		}
		tdesc, err := Build(reflect.TypeOf((*Struct)(nil)).Elem(), Conf{})
		if err != nil {
			t.Errorf("unexpected '%v' error", err)
			return
		}
		if !reflect.DeepEqual(tdesc, want) {
			t.Errorf("want %v, actual %v", want, tdesc)
		}
	})
}

func testInvalidTag(tp reflect.Type, want *InvalidTagFormatError,
	t *testing.T) {
	_, err := Build(tp, Conf{})
	if tagErr, ok := err.(*InvalidTagFormatError); ok {
		if tagErr.FieldName() != want.FieldName() {
			t.Errorf("unexpected '%v' field", tagErr.FieldName())
		}
		if tagErr.Error() != want.Error() {
			t.Errorf("unexpected '%v' error", err)
		}
		return
	}
	t.Errorf("unexpected '%v' error", err)
}

func testInvalidTagPart(tp reflect.Type, want error, t *testing.T) {
	_, err := Build(tp, Conf{})
	if err != want {
		t.Errorf("unexpected '%v' error", err)
	}
}
