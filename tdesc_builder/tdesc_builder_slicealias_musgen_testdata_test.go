package tdesc_builder

import (
	"reflect"
	"testing"

	tdmg "github.com/ymz-ncnk/serialization/musgo/testdata/musgen"
)

func TestParseSliceAliasMusgenTestdata(t *testing.T) {
	conf := Conf{}

	td, err := Build(reflect.TypeOf(tdmg.TrickySliceAlias{}), conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, tdmg.TrickySliceAliasTypeDesc) {
		t.Error("TrickySliceAlias")
	}
}
