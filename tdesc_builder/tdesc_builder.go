package tdesc_builder

import (
	"reflect"
	"regexp"

	"github.com/ymz-ncnk/serialization/musgen"
	"github.com/ymz-ncnk/serialization/musgo/parser"
)

// Build builds musgen.TypeDesc for struct or alias type.
func Build(tp reflect.Type, conf Conf) (
	tdesc musgen.TypeDesc, err error) {
	tdesc, err = BuildForStruct(tp, conf)
	if _, ok := err.(*parser.UnsupportedTypeError); ok {
		return BuildForAlias(tp, AliasConf{Conf: conf})
	}
	return
}

// Build builds musgen.TypeDesc for struct type.
func BuildForStruct(tp reflect.Type, conf Conf) (
	tdesc musgen.TypeDesc, err error) {
	fieldsTypes, fieldsProps, err := parser.ParseStructWithTags(tp, TagParser)
	if err != nil {
		return
	}
	if fieldsTypes == nil && fieldsProps == nil {
		err = ErrNotStruct
		return
	}
	fds := make([]musgen.FieldDesc, tp.NumField())
	for i := 0; i < len(fds); i++ {
		if _, ok := fieldsProps[i][0].(bool); ok {
			continue
		}
		fds[i] = musgen.FieldDesc{
			Name: tp.Field(i).Name,
			Type: fieldsTypes[i],
		}
		if validator, ok := fieldsProps[i][0].(string); ok {
			fds[i].Validator = validator
		}
		if encoding, ok := fieldsProps[i][1].(string); ok {
			fds[i].Encoding = encoding
		}
		if maxLength, ok := fieldsProps[i][2].(int); ok {
			fds[i].MaxLength = maxLength
		}
		if elemValidator, ok := fieldsProps[i][3].(string); ok {
			fds[i].ElemValidator = elemValidator
		}
		if elemEncoding, ok := fieldsProps[i][4].(string); ok {
			fds[i].ElemEncoding = elemEncoding
		}
		if keyValidator, ok := fieldsProps[i][5].(string); ok {
			fds[i].KeyValidator = keyValidator
		}
		if keyEncoding, ok := fieldsProps[i][6].(string); ok {
			fds[i].KeyEncoding = keyEncoding
		}
	}
	tdesc = musgen.TypeDesc{
		Package: pkg(tp),
		Name:    tp.Name(),
		Unsafe:  conf.Unsafe,
		Suffix:  conf.Suffix,
		Fields:  fds,
	}
	return
}

// Build builds musgen.TypeDesc for alias type.
func BuildForAlias(tp reflect.Type, conf AliasConf) (
	tdesc musgen.TypeDesc, err error) {
	aliasOf, _, err := parser.Parse(tp)
	if err != nil {
		return
	}
	if aliasOf == "" {
		err = ErrNotAlias
		return
	}
	fds := make([]musgen.FieldDesc, 1)
	fds[0] = musgen.FieldDesc{
		Type:          aliasOf,
		Alias:         tp.Name(),
		MaxLength:     conf.MaxLength,
		Validator:     conf.Validator,
		Encoding:      conf.Encoding,
		ElemValidator: conf.ElemValidator,
		ElemEncoding:  conf.ElemEncoding,
		KeyValidator:  conf.KeyValidator,
		KeyEncoding:   conf.KeyEncoding,
	}
	tdesc = musgen.TypeDesc{
		Package: pkg(tp),
		Name:    tp.Name(),
		Unsafe:  conf.Unsafe,
		Suffix:  conf.Suffix,
		Fields:  fds,
	}
	return
}

func pkg(t reflect.Type) string {
	re := regexp.MustCompile(`^(.*)\.`)
	match := re.FindStringSubmatch(t.String())
	if len(match) != 2 {
		return ""
	}
	return match[1]
}
