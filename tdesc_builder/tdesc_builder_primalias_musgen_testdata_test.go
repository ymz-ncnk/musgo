package tdesc_builder

import (
	"reflect"
	"testing"

	tdmg "github.com/ymz-ncnk/serialization/musgo/testdata/musgen"
)

func TestParsePrimAliasMusgenTestdata(t *testing.T) {
	conf := AliasConf{MaxLength: 5}

	td, err := BuildForAlias(reflect.TypeOf((*tdmg.Uint32Alias)(nil)).Elem(),
		conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, tdmg.Uint32AliasTypeDesc) {
		t.Error("Uint32Alias")
	}

	conf = AliasConf{}
	td, err = BuildForAlias(reflect.TypeOf((*tdmg.ByteAlias)(nil)).Elem(),
		conf)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(td, tdmg.ByteAliasTypeDesc) {
		t.Error("ByteAlias")
	}
}
