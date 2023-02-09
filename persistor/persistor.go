package persistor

import (
	"github.com/ymz-ncnk/musgen"
)

// FilenameExtenstion of the generated files.
const FilenameExtenstion = "mus.go"

// FilenameBuilder creates a name for the generated file.
type FilenameBuilder func(tDesc musgen.TypeDesc) string

// DefaultFilenameBuilder creates a default name for the generated file.
func DefaultFilenameBuilder(tDesc musgen.TypeDesc) string {
	return tDesc.Name + "." + FilenameExtenstion
}

// -----------------------------------------------------------------------------
// Persistor responsible for storing the generated data.
type Persistor interface {
	// Persist data to the specified path.
	Persist(tDesc musgen.TypeDesc, data []byte, path string) error
}
