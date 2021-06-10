package musgo

import (
	"reflect"
	"testing"

	"github.com/ymz-ncnk/musgo/parser"
	mgtd "github.com/ymz-ncnk/musgo/testdata/musgen"
)

func TestParseArrayAlias(t *testing.T) {
	td, err := parser.Parse(reflect.TypeOf(mgtd.StrArrayAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.StrArrayAliasTypeDesc) {
		t.Error("StrArrayAlias")
	}

	td, err = parser.Parse(reflect.TypeOf(mgtd.FloatPtrArrayAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.FloatPtrArrayAliasTypeDesc) {
		t.Error("FloatPtrArrayAlias")
	}

	td, err = parser.Parse(reflect.TypeOf(mgtd.IntAliasArrayAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.IntAliasArrayAliasTypeDesc) {
		t.Error("IntAliasArrayAlias")
	}

	td, err = parser.Parse(reflect.TypeOf(mgtd.Uint64PtrAliasArrayAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.Uint64PtrAliasArrayAliasTypeDesc) {
		t.Error("Uint64PtrAliasArrayAlias")
	}

	td, err = parser.Parse(reflect.TypeOf(mgtd.Int64SliceArrayAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.Int64SliceArrayAliasTypeDesc) {
		t.Error("Int64SliceArrayAlias")
	}

	td, err = parser.Parse(reflect.TypeOf(mgtd.Uint16SlicePtrArrayAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.Uint16SlicePtrArrayAliasTypeDesc) {
		t.Error("Uint16SlicePtrArrayAlias")
	}

	td, err = parser.Parse(reflect.TypeOf(mgtd.BoolArrayArrayAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.BoolArrayArrayAliasTypeDesc) {
		t.Error("BoolArrayArrayAlias")
	}

	td, err = parser.Parse(reflect.TypeOf(mgtd.BytePtrArrayPtrArrayAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.BytePtrArrayPtrArrayAliasTypeDesc) {
		t.Error("BytePtrArrayPtrArrayAlias")
	}

	td, err = parser.Parse(reflect.TypeOf(mgtd.IntStrMapArrayAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.IntStrMapArrayAliasTypeDesc) {
		t.Error("IntStrMapArrayAlias")
	}

	td, err = parser.Parse(reflect.TypeOf(mgtd.Uint32Int32MapArrayAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.Uint32Int32MapArrayAliasTypeDesc) {
		t.Error("Uint32Int32MapArrayAlias")
	}

	td, err = parser.Parse(reflect.TypeOf(mgtd.StructTypeArrayAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.StructTypeArrayAliasTypeDesc) {
		t.Error("StructTypeArrayAlias")
	}

	td, err = parser.Parse(reflect.TypeOf(mgtd.StructTypePtrArrayAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.StructTypePtrArrayAliasTypeDesc) {
		t.Error("StructTypePtrArrayAlias")
	}

	td, err = parser.Parse(reflect.TypeOf(mgtd.TrickyArrayAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.TrickyArrayAliasTypeDesc) {
		t.Error("TrickyArrayAlias")
	}
}
