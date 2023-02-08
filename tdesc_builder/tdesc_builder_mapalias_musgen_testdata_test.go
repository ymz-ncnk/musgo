package tdesc_builder

import (
	"reflect"
	"testing"

	tdmg "github.com/ymz-ncnk/musgo/testdata/musgen"
)

func TestParseMapAliasMusgenTestdata(t *testing.T) {
	conf := Conf{}

	td, err := Build(reflect.TypeOf(tdmg.TrickyMapAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, tdmg.TrickyMapAliasTypeDesc) {
		t.Error("TrickyMapAlias")
	}
}
