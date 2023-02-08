package musgo

import (
	"io/fs"
	"os"

	"github.com/ymz-ncnk/musgen"
	"golang.org/x/tools/imports"
)

// FilenameExtenstion of the generated files.
const FilenameExtenstion = "mus.go"

// FilenameBuilder creates a name for the generated file.
type FilenameBuilder func(tdesc musgen.TypeDesc) string

// DefaultFilenameBuilder creates a default name for the generated file.
func DefaultFilenameBuilder(td musgen.TypeDesc) string {
	return td.Name + "." + FilenameExtenstion
}

// -----------------------------------------------------------------------------
// Persistor responsible for storing the generated data.
type Persistor interface {
	// Persist data to the specified path.
	Persist(tdesc musgen.TypeDesc, data []byte, path string) error
}

// -----------------------------------------------------------------------------
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

func (persistor HarDrivePersistor) Persist(tdesc musgen.TypeDesc, data []byte,
	path string) (err error) {
	data, err = imports.Process("", data, nil)
	if err != nil {
		return
	}
	filename := persistor.FilenameBuilder(tdesc)
	if path == "" {
		path = "."
	}
	return os.WriteFile(path+string(os.PathSeparator)+filename, data,
		persistor.Perm)
}
