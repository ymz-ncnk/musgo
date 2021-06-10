package musgo

import (
	"reflect"
	"testing"

	"github.com/ymz-ncnk/musgo/parser"
	mgtd "github.com/ymz-ncnk/musgo/testdata/musgen"
)

func TestParseSliceAlias(t *testing.T) {
	td, err := parser.Parse(reflect.TypeOf(mgtd.TrickySliceAlias{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.TrickySliceAliasTypeDesc) {
		t.Error("TrickySliceAlias")
	}
}
