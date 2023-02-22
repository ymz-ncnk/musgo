package musgo

import (
	"os"
	"reflect"
	"testing"
)

func TestMusGoIntegration(t *testing.T) {
	type IntAlias int
	dname, err := os.MkdirTemp("", "musgo")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dname)

	musGo, err := New()
	if err != nil {
		t.Fatal(err)
	}
	conf := Conf{}
	conf.Path = dname
	err = musGo.GenerateAs(reflect.TypeOf((*IntAlias)(nil)).Elem(), conf)
	if err != nil {
		t.Fatal(err)
	}
	wantPath := conf.Path + string(os.PathSeparator) + "IntAlias" +
		FilenameExtenstion
	if _, err := os.Stat(wantPath); err != nil {
		t.Error("file was not generated")
	}
}
