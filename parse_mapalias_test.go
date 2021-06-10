package musgo

import (
	"reflect"
	"testing"

	"github.com/ymz-ncnk/musgo/parser"
	mgtd "github.com/ymz-ncnk/musgo/testdata/musgen"
)

func TestParseMapAlias(t *testing.T) {
	td, err := parser.Parse(reflect.TypeOf(mgtd.TrickyMapAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.TrickyMapAliasTypeDesc) {
		t.Error("TrickyMapAlias")
	}
}
