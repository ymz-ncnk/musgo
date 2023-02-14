package persistor

import (
	"io/fs"
	"os"

	"github.com/ymz-ncnk/musgen/v2"
	"golang.org/x/tools/imports"
)

// NewHarDrivePersistor creates new HarDrivePersistor.
func NewHarDrivePersistor() HarDrivePersistor {
	return HarDrivePersistor{
		FilenameBuilder: DefaultFilenameBuilder,
		Perm:            os.ModePerm,
	}
}

// HarDrivePersistor saves data to hard drive.
type HarDrivePersistor struct {
	FilenameBuilder FilenameBuilder
	Perm            fs.FileMode
}

func (persistor HarDrivePersistor) Persist(tDesc musgen.TypeDesc, data []byte,
	path string) (err error) {
	data, err = imports.Process("", data, nil)
	if err != nil {
		return
	}
	filename := persistor.FilenameBuilder(tDesc)
	if path == "" {
		path = "."
	}
	return os.WriteFile(path+string(os.PathSeparator)+filename, data,
		persistor.Perm)
}
