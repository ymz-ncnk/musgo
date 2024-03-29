//go:generate go run gen/main.go $ARG
package musgo

import (
	"reflect"

	"github.com/ymz-ncnk/musgen/v2"
	musgen_textmpl "github.com/ymz-ncnk/musgen/v2/text_template"
	"github.com/ymz-ncnk/musgo/v2/parser"
	"github.com/ymz-ncnk/musgo/v2/tdesc_builder"
	persistor_mod "github.com/ymz-ncnk/persistor"
	"golang.org/x/tools/imports"
)

// FilenameExtenstion of the generated files.
const FilenameExtenstion = ".mus.go"

// DefConf is the default configuration for a struct type.
var DefConf = Conf{Suffix: "MUS", Path: "."}

// DefAliasConf is the default configuration for an alias type.
var DefAliasConf = AliasConf{Conf: DefConf}

// New creates a new MusGo.
func New() (musGo MusGo, err error) {
	musGen, err := musgen_textmpl.New()
	if err != nil {
		return
	}
	return NewWith(musGen, persistor_mod.NewHarDrivePersistor())
}

// NewWith creates a configurable MusGo.
func NewWith(musGen musgen.MusGen, persistor persistor_mod.Persistor) (
	musGo MusGo, err error) {
	return MusGo{musGen, persistor}, nil
}

// MusGo is a Go code generator for the MUS format.
type MusGo struct {
	musGen    musgen.MusGen
	persistor persistor_mod.Persistor
}

// Generate accepts a struct or alias type. Returns an error if receives an
// alias to the pointer type. It generates Marshal, Unmarshal, and Size methods
// for the specified type. Generated file with 'name of the type'.mus.go name
// is placed to the current directory.
// Each of the struct field can have a tag:
// mus:"Validator#raw,MaxLength,ElemValidator#raw,KeyValidator#raw"
func (musGo MusGo) Generate(tp reflect.Type, unsafe bool) error {
	conf := DefConf
	conf.Unsafe = unsafe
	return musGo.GenerateAs(tp, conf)
}

// GenerateAs performs like Generate. With help of this method you can configure
// the generation process.
func (musGo MusGo) GenerateAs(tp reflect.Type, conf Conf) (err error) {
	tDesc, err := tdesc_builder.Build(tp, tdesc_builder.Conf{
		Unsafe: conf.Unsafe,
		Suffix: conf.Suffix,
	})
	if err != nil {
		return
	}
	return musGo.generate(tDesc, conf.Path)
}

// GenerateAliasAs performs like Generate. With help of this method Validators
// and Encodings can be set for an alias type.
func (musGo MusGo) GenerateAliasAs(tp reflect.Type, conf AliasConf) (
	err error) {
	aliasOf, _, _, err := parser.Parse(tp, nil)
	if err != nil {
		return
	}
	if aliasOf == "" {
		return ErrNotAlias
	}
	tDesc, err := tdesc_builder.BuildForAlias(tp, aliasOf,
		tdesc_builder.AliasConf{
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
	return musGo.generate(tDesc, conf.Path)
}

func (musGo MusGo) generate(tDesc musgen.TypeDesc, path string) (err error) {
	data, err := musGo.musGen.Generate(tDesc, musgen.GoLang)
	if err != nil {
		return
	}
	data, err = imports.Process("", data, nil)
	if err != nil {
		return
	}
	name := tDesc.Name + FilenameExtenstion
	return musGo.persistor.Persist(name, data, path)
}
