package musgo

import (
	"os"
	"reflect"

	"github.com/ymz-ncnk/serialization/musgen"
	"github.com/ymz-ncnk/serialization/musgo/tdesc_builder"
	"golang.org/x/tools/imports"
)

// FilenameExtenstion of the generated files.
const FilenameExtenstion = "mus.go"

// Conf configures the generation process.
type Conf struct {
	Unsafe bool   // generate unsafe code or not
	Suffix string // suffix for Marshal, Unmarshal, Size methods
	Path   string // folder of the generated file
}

// AliasConf configures the generation process for an alias type.
type AliasConf struct {
	Conf
	Validator string // validates value
	Encoding  string // sets encoding
	MaxLength int    // if alias to string, array, slice, or map, restricts
	// length, should be positive number
	ElemValidator string // if alias to array, slice, or map, validates elements
	ElemEncoding  string // if alias to array, slic, or map, sets encoding
	KeyValidator  string // if alias to map, validates keys
	KeyEncoding   string // if alias to map, sets encoding
}

type FilenameBuilder func(tdesc musgen.TypeDesc) string

func DefaultFilenameBuilder(td musgen.TypeDesc) string {
	return td.Name + "." + FilenameExtenstion
}

// New returns a new MusGo.
func New() (musGo MusGo, err error) {
	musGen, err := musgen.New()
	if err != nil {
		return
	}
	return MusGo{musGen, DefaultFilenameBuilder}, nil
}

// MusGo is a Go code generator for the MUS format.
type MusGo struct {
	musGen          musgen.MusGen
	FilenameBuilder FilenameBuilder
}

// Generate accepts a struct or alias type. Returns an error if receives an
// alias to the pointer type. It generates Marshal, Unmarshal, and Size methods
// for the specified type. Generated file with 'name of the type'.mus.go name
// is placed to the current directory.
// Each of the struct field can have a tag:
// mus:"Validator#raw,MaxLength,ElemValidator#raw,KeyValidator#raw"
func (musGo MusGo) Generate(tp reflect.Type, unsafe bool) error {
	conf := Conf{}
	conf.Unsafe = unsafe
	return musGo.GenerateAs(tp, conf)
}

// GenerateAs performs like Generate. With help of this method you can configure
// generation process.
func (musGo MusGo) GenerateAs(tp reflect.Type, conf Conf) (err error) {
	tdesc, err := tdesc_builder.Build(tp, tdesc_builder.Conf{
		Unsafe: conf.Unsafe,
		Suffix: conf.Suffix,
	})
	if err != nil {
		return
	}
	return musGo.generate(tdesc, conf.Path)
}

// GenerateAliasAs performs like Generate. With help of this method Validators
// and Encodings can be set for an alias type.
func (musGo MusGo) GenerateAliasAs(tp reflect.Type, conf AliasConf) (
	err error) {
	tdesc, err := tdesc_builder.BuildForAlias(tp, tdesc_builder.AliasConf{
		Conf: tdesc_builder.Conf{
			Unsafe: conf.Unsafe,
			Suffix: conf.Suffix,
		},
		Validator:     conf.Validator,
		Encoding:      conf.Encoding,
		MaxLength:     conf.MaxLength,
		ElemValidator: conf.ElemValidator,
		ElemEncoding:  conf.ElemEncoding,
		KeyValidator:  conf.KeyValidator,
		KeyEncoding:   conf.KeyEncoding,
	})
	if err != nil {
		return
	}
	return musGo.generate(tdesc, conf.Path)
}

func (musGo MusGo) generate(tdesc musgen.TypeDesc, path string) (err error) {
	b, err := musGo.musGen.Generate(tdesc, musgen.GoLang)
	if err != nil {
		return
	}
	if path == "" {
		path = "."
	}
	return musGo.persist(tdesc, b, path)
}

func (musGo MusGo) persist(tdesc musgen.TypeDesc, b []byte, path string) (
	err error) {
	b, err = imports.Process("", b, nil)
	if err != nil {
		return
	}
	filename := musGo.FilenameBuilder(tdesc)
	if path == "" {
		path = "."
	}
	return os.WriteFile(path+string(os.PathSeparator)+filename, b, os.ModePerm)
}
