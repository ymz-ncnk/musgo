package musgo

import (
	"errors"
	"os"
	"reflect"

	"github.com/ymz-ncnk/musgo/parser"

	"github.com/ymz-ncnk/musgen"
	"golang.org/x/tools/imports"
)

// ErrNotAliasType happens on GenerateAlias if type is not an alias.
var ErrNotAliasType = errors.New("not alias type")

// DefaultSuffix for Marshal, Unmarsha, Size methods.
const DefaultSuffix = "MUS"

// FilenameExtenstion of the generated files.
const FilenameExtenstion = "musgen.go"

// NewConf creates a Conf for the MusGo.
func NewConf() Conf {
	return Conf{
		Suffix: DefaultSuffix,
	}
}

// Conf configures the generation process.
type Conf struct {
	T        reflect.Type // generate code for this type
	Unsafe   bool         // generate unsafe code or not
	Path     string       // folder of the generated file
	Filename string       // name of the generated file
	Suffix   string       // suffix for Marshal, Unmarshal, Size methods
}

// New returns a new MusGo.
func New() (MusGo, error) {
	musGen, err := musgen.New()
	if err != nil {
		return MusGo{}, err
	}
	return MusGo{musGen: musGen, FilenameBuilder: DefaultFilenameBuilder}, nil
}

// MusGo is a Go code generator for the MUS format.
type MusGo struct {
	musGen          musgen.MusGen
	FilenameBuilder func(td musgen.TypeDesc) string
}

// Generate generates Marshal, Unmarshal, and Size methods for the specified
// type. Generated file with 'name of the type'.musgen.go name is placed to the
// current directory.
// If type is an alias to a pointer type returns error.
func (musGo MusGo) Generate(t reflect.Type, unsafe bool) error {
	conf := NewConf()
	conf.T = t
	conf.Unsafe = unsafe
	return musGo.GenerateAs(conf)
}

// GenerateAs performs like Generate.
func (musGo MusGo) GenerateAs(conf Conf) error {
	td, err := parser.Parse(conf.T)
	td.Unsafe = conf.Unsafe
	td.Suffix = conf.Suffix
	if err != nil {
		return err
	}
	return musGo.generate(td, conf.Path, conf.Filename)
}

// NewAliasConf creates an AliasConf for the MusGo.
func NewAliasConf() AliasConf {
	return AliasConf{
		Conf: NewConf(),
	}
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

// GenerateAlias performs like the Generate method. Use it if you want to
// provide validation for an alias type.
func (musGo MusGo) GenerateAlias(t reflect.Type, unsafe bool, validator string,
	encoding string, maxLength int, elemValidator, elemEncoding, keyValidator,
	keyEncoding string) error {
	conf := NewAliasConf()
	conf.Conf.T = t
	conf.Conf.Unsafe = unsafe
	conf.Validator = validator
	conf.Encoding = encoding
	conf.MaxLength = maxLength
	conf.ElemValidator = elemValidator
	conf.ElemEncoding = elemEncoding
	conf.KeyValidator = keyValidator
	conf.KeyEncoding = keyEncoding
	return musGo.GenerateAliasAs(conf)
}

// GenerateAliasAs performs like the GenerateAlias method.
func (musGo MusGo) GenerateAliasAs(conf AliasConf) error {
	td, err := parser.Parse(conf.T)
	if err != nil {
		return err
	}
	if !musgen.Alias(td) {
		return ErrNotAliasType
	}
	td.Unsafe = conf.Unsafe
	td.Suffix = conf.Suffix
	td.Fields[0].Validator = conf.Validator
	td.Fields[0].Encoding = conf.Encoding
	td.Fields[0].MaxLength = conf.MaxLength
	td.Fields[0].ElemValidator = conf.ElemValidator
	td.Fields[0].ElemEncoding = conf.ElemEncoding
	td.Fields[0].KeyValidator = conf.KeyValidator
	td.Fields[0].KeyEncoding = conf.KeyEncoding
	return musGo.generate(td, conf.Path, conf.Filename)
}

func (musGo MusGo) generate(td musgen.TypeDesc, path, filename string) error {
	musGen, err := musgen.New()
	if err != nil {
		return err
	}
	var bs []byte
	bs, err = musGen.Generate(td, musgen.GoLang)
	if err != nil {
		return err
	}
	bs, err = imports.Process("", bs, nil)
	if err != nil {
		return err
	}
	if filename == "" {
		filename = musGo.FilenameBuilder(td)
	}
	if path == "" {
		path = "."
	}
	return os.WriteFile(path+"/"+filename, bs, os.ModePerm)
}

func DefaultFilenameBuilder(td musgen.TypeDesc) string {
	return td.Name + "." + FilenameExtenstion
}
