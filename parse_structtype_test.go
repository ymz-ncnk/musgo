package musgo

import (
	"reflect"
	"testing"

	"github.com/ymz-ncnk/musgo/parser"
	mgtd "github.com/ymz-ncnk/musgo/testdata/musgen"
)

func TestParseStructType(t *testing.T) {
	td, err := parser.Parse(reflect.TypeOf(mgtd.StructType{}))
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, mgtd.StructTypeDesc) {
		t.Error("StructType")
	}
}
