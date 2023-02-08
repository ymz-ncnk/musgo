package musgo

import (
	"os"
	"reflect"
	"testing"

	tdmg "github.com/ymz-ncnk/musgo/testdata/musgo"
)

func TestMusGoIntegration(t *testing.T) {
	musGo, err := New()
	if err != nil {
		t.Fatal(err)
	}
	conf := Conf{}
	conf.Path = "testdata/musgo"
	err = musGo.GenerateAs(reflect.TypeOf((*tdmg.IntAlias)(nil)).Elem(), conf)
	if err != nil {
		t.Fatal(err)
	}
	wantPath := conf.Path
	if _, err := os.Stat(wantPath); err != nil {
		t.Error("file was not generated")
	}
}
