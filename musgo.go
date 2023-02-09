package musgo

import (
	"reflect"

	"github.com/ymz-ncnk/musgen"
	musgen_textmpl "github.com/ymz-ncnk/musgen/text_template"
	"github.com/ymz-ncnk/musgo/parser"
	persist "github.com/ymz-ncnk/musgo/persistor"
	"github.com/ymz-ncnk/musgo/tdesc_builder"
)

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

var DefConf = Conf {Path: "."}

// New creates a new MusGo.
func New() (musGo MusGo, err error) {
	musGen, err := musgen_textmpl.New()
	if err != nil {
		return
	}
	return NewWith(musGen, persist.NewHarDrivePersistor())
}

// NewWith creates a configurable MusGo.
func NewWith(musGen musgen.MusGen, persistor persist.Persistor) (musGo MusGo,
	err error) {
	return MusGo{musGen, persistor}, nil
}

// MusGo is a Go code generator for the MUS format.
type MusGo struct {
	musGen    musgen.MusGen
	persistor persist.Persistor
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
	aliasOf, _, _, err := parser.Parse(tp, nil)
	if err != nil {
		return
	}
	if aliasOf == "" {
		return ErrNotAlias
	}
	tdesc, err := tdesc_builder.BuildForAlias(tp, aliasOf, 
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
	return musGo.generate(tdesc, conf.Path)
}

func (musGo MusGo) generate(tdesc musgen.TypeDesc, path string) (err error) {
	b, err := musGo.musGen.Generate(tdesc, musgen.GoLang)
	if err != nil {
		return
	}
	return musGo.persistor.Persist(tdesc, b, path)
}
