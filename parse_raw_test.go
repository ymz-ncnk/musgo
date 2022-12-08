package musgo

import (
	"reflect"
	"testing"

	"github.com/ymz-ncnk/musgo/parser"
	mgtd "github.com/ymz-ncnk/musgo/testdata/musgen"
)

func TestParseRawAlias(t *testing.T) {
	{
		var v mgtd.RawStructType
		td, err := parser.ParseStruct(reflect.TypeOf(v))
		if err != nil {
			t.Error(err)
		}
		etd := mgtd.RawStructTypeDesc
		if !reflect.DeepEqual(td, etd) {
			t.Error("RawStructType")
		}
	}
}
