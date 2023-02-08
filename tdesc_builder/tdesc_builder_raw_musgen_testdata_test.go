package tdesc_builder

import (
	"reflect"
	"testing"

	tdmg "github.com/ymz-ncnk/musgo/testdata/musgen"
)

func TestParseRawAliasMusgenTestdata(t *testing.T) {
	conf := Conf{}

	td, err := Build(reflect.TypeOf(tdmg.RawStructType{}), conf)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(td, tdmg.RawStructTypeDesc) {
		t.Error("RawStructType")
	}
}
