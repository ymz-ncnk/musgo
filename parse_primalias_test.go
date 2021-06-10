package musgo

import (
	"reflect"
	"testing"

	"github.com/ymz-ncnk/musgo/parser"
	mgtd "github.com/ymz-ncnk/musgo/testdata/musgen"
)

func TestParsePrimAlias(t *testing.T) {
	{
		var v mgtd.Uint32Alias
		td, err := parser.Parse(reflect.TypeOf(v))
		if err != nil {
			t.Error(err)
		}
		etd := mgtd.Uint32AliasTypeDesc
		etd.Fields[0].MaxLength = 0
		if !reflect.DeepEqual(td, etd) {
			t.Error("Uint32Alias")
		}
	}
	{
		var v mgtd.ByteAlias
		td, err := parser.Parse(reflect.TypeOf(v))
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(td, mgtd.ByteAliasTypeDesc) {
			t.Error("ByteAlias")
		}
	}
}
