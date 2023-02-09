package persistor

import (
	"os"
	"reflect"
	"testing"

	"github.com/ymz-ncnk/musgen"
)

func TestHarDrivePersistor(t *testing.T) {
	dname, err := os.MkdirTemp("", "musgo")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dname)

	t.Run("Persist", func(t *testing.T) {
		var (
			wantData = []byte("package musgo\n")
			tDesc    = musgen.TypeDesc{Name: "StringAlias"}
		)
		p := NewHarDrivePersistor()
		err := p.Persist(tDesc, wantData, dname)
		if err != nil {
			t.Fatal(err)
		}
		wantFilePath := dname + string(os.PathSeparator) + tDesc.Name + "." +
			FilenameExtenstion
		data, err := os.ReadFile(wantFilePath)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(data, wantData) {
			t.Fatalf("unexpected Data, want '%v' actual '%v'", wantData, data)
		}
	})

	t.Run("Persist not source code", func(t *testing.T) {
		var (
			wantData = []byte{1, 2, 3}
			tDesc    = musgen.TypeDesc{Name: "StringAlias"}
		)
		p := NewHarDrivePersistor()
		err := p.Persist(tDesc, wantData, dname)
		if err == nil {
			t.Fatal("expects an error here")
		}
	})

}
