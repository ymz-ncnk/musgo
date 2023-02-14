package tdesc_builder

import (
	"reflect"
	"regexp"

	"github.com/ymz-ncnk/musgen/v2"
	"github.com/ymz-ncnk/musgo/parser"
)

// Build builds musgen.TypeDesc for struct or alias type.
func Build(tp reflect.Type, conf Conf) (tdesc musgen.TypeDesc, err error) {
	aliasOf, fieldsTypes, fieldsProps, err := parser.Parse(tp, TagParser)
	if err != nil {
		return
	}
	if aliasOf != "" {
		return BuildForAlias(tp, aliasOf, AliasConf{Conf: conf})
	}
	return BuildForStruct(tp, fieldsTypes, fieldsProps, conf)
}

// Build builds musgen.TypeDesc for alias type.
func BuildForAlias(tp reflect.Type, aliasOf string, conf AliasConf) (
	tdesc musgen.TypeDesc, err error) {
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

// Build builds musgen.TypeDesc for struct type.
func BuildForStruct(tp reflect.Type, fieldsTypes []string, fieldsProps [][]any,
	conf Conf) (tdesc musgen.TypeDesc, err error) {
	var (
		fd  musgen.FieldDesc
		fds = []musgen.FieldDesc{}
	)
	for i := 0; i < len(fieldsProps); i++ {
		if len(fieldsProps[i]) == 0 {
			fd = musgen.FieldDesc{
				Name: tp.Field(i).Name,
				Type: fieldsTypes[i],
			}
		} else {
			if _, ok := fieldsProps[i][0].(bool); ok {
				continue
			}
			fd = musgen.FieldDesc{
				Name: tp.Field(i).Name,
				Type: fieldsTypes[i],
			}
			if validator, ok := fieldsProps[i][0].(string); ok {
				fd.Validator = validator
			}
			if encoding, ok := fieldsProps[i][1].(string); ok {
				fd.Encoding = encoding
			}
			if maxLength, ok := fieldsProps[i][2].(int); ok {
				fd.MaxLength = maxLength
			}
			if elemValidator, ok := fieldsProps[i][3].(string); ok {
				fd.ElemValidator = elemValidator
			}
			if elemEncoding, ok := fieldsProps[i][4].(string); ok {
				fd.ElemEncoding = elemEncoding
			}
			if keyValidator, ok := fieldsProps[i][5].(string); ok {
				fd.KeyValidator = keyValidator
			}
			if keyEncoding, ok := fieldsProps[i][6].(string); ok {
				fd.KeyEncoding = keyEncoding
			}
		}
		fds = append(fds, fd)
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
