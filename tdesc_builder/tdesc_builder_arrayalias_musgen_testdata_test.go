package tdesc_builder

import (
	"reflect"
	"testing"

	tdmg "github.com/ymz-ncnk/serialization/musgo/testdata/musgen"
)

func TestParseArrayAliasMusgenTestdata(t *testing.T) {
	conf := Conf{}

	tdesc, err := Build(reflect.TypeOf(tdmg.StrArrayAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tdesc, tdmg.StrArrayAliasTypeDesc) {
		t.Error("StrArrayAlias")
	}

	tdesc, err = Build(reflect.TypeOf(tdmg.FloatPtrArrayAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tdesc, tdmg.FloatPtrArrayAliasTypeDesc) {
		t.Error("FloatPtrArrayAlias")
	}

	tdesc, err = Build(reflect.TypeOf(tdmg.IntAliasArrayAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tdesc, tdmg.IntAliasArrayAliasTypeDesc) {
		t.Error("IntAliasArrayAlias")
	}

	tdesc, err = Build(reflect.TypeOf(tdmg.Uint64PtrAliasArrayAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tdesc, tdmg.Uint64PtrAliasArrayAliasTypeDesc) {
		t.Error("Uint64PtrAliasArrayAlias")
	}

	tdesc, err = Build(reflect.TypeOf(tdmg.Int64SliceArrayAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tdesc, tdmg.Int64SliceArrayAliasTypeDesc) {
		t.Error("Int64SliceArrayAlias")
	}

	tdesc, err = Build(reflect.TypeOf(tdmg.Uint16SlicePtrArrayAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tdesc, tdmg.Uint16SlicePtrArrayAliasTypeDesc) {
		t.Error("Uint16SlicePtrArrayAlias")
	}

	tdesc, err = Build(reflect.TypeOf(tdmg.BoolArrayArrayAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tdesc, tdmg.BoolArrayArrayAliasTypeDesc) {
		t.Error("BoolArrayArrayAlias")
	}

	tdesc, err = Build(reflect.TypeOf(tdmg.BytePtrArrayPtrArrayAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tdesc, tdmg.BytePtrArrayPtrArrayAliasTypeDesc) {
		t.Error("BytePtrArrayPtrArrayAlias")
	}

	tdesc, err = Build(reflect.TypeOf(tdmg.IntStrMapArrayAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tdesc, tdmg.IntStrMapArrayAliasTypeDesc) {
		t.Error("IntStrMapArrayAlias")
	}

	tdesc, err = Build(reflect.TypeOf(tdmg.Uint32Int32MapArrayAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tdesc, tdmg.Uint32Int32MapArrayAliasTypeDesc) {
		t.Error("Uint32Int32MapArrayAlias")
	}

	tdesc, err = Build(reflect.TypeOf(tdmg.StructTypeArrayAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tdesc, tdmg.StructTypeArrayAliasTypeDesc) {
		t.Error("StructTypeArrayAlias")
	}

	tdesc, err = Build(reflect.TypeOf(tdmg.StructTypePtrArrayAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tdesc, tdmg.StructTypePtrArrayAliasTypeDesc) {
		t.Error("StructTypePtrArrayAlias")
	}

	tdesc, err = Build(reflect.TypeOf(tdmg.TrickyArrayAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tdesc, tdmg.TrickyArrayAliasTypeDesc) {
		t.Error("TrickyArrayAlias")
	}
}
