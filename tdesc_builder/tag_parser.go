package tdesc_builder

import (
	"reflect"
	"strconv"
	"strings"
)

// TagKey represents a key of a tag, which marks field validators and
// encodings.
const MUSTagKey = "mus"

// EncodingSep devides validator and encoding.
const EncodingSep = "#"

// Tag structure:
// "mus:-" or
// "mus:validator,encoding,maxLength,elemValidator,elemEncoding,keyValidator,keyEncoding"
func TagParser(tp reflect.Type, field reflect.StructField,
	tag reflect.StructTag) (fieldProps []any, err error) {
	musTag, pst := tag.Lookup(MUSTagKey)
	if !pst || tag == "" {
		return
	}
	subTags := strings.Split(musTag, ",")
	if len(subTags) > 4 {
		err = NewInvalidTagFormatError(field.Name)
		return
	}
	if subTags[0] == "-" {
		fieldProps = make([]any, 1)
		if len(subTags) > 1 {
			err = NewInvalidTagFormatError(field.Name)
			return
		}
		// skip
		fieldProps[0] = true
		return
	}
	fieldProps = make([]any, 7)
	if len(subTags) >= 1 {
		// validator, encoding
		fieldProps[0], fieldProps[1], err = parseValidatorAndEncoding(subTags[0])
		if err != nil {
			return
		}
	}
	if len(subTags) >= 2 {
		// maxLength
		fieldProps[2], err = parseMaxLength(field, subTags[1])
		if err != nil {
			return
		}
	}
	if len(subTags) >= 3 {
		// elemValidator, elemEncoding
		fieldProps[3], fieldProps[4], err = parseValidatorAndEncoding(subTags[2])
		if err != nil {
			return
		}
	}
	if len(subTags) == 4 {
		// keyValidator, keyEncoding
		fieldProps[5], fieldProps[6], err = parseValidatorAndEncoding(subTags[3])
		if err != nil {
			return
		}
	}
	return
}

func parseValidatorAndEncoding(str string) (validator, encoding string,
	err error) {
	if str == "" {
		return "", "", nil
	}
	if strings.HasPrefix(str, EncodingSep) {
		return "", str[1:], nil
	}
	parts := strings.Split(str, EncodingSep)
	if len(parts) > 2 {
		return "", "", NewInvalidTagPartFormatError(str)
	}
	if len(parts) == 2 {
		return parts[0], parts[1], nil
	}
	return str, "", nil
}

func parseMaxLength(field reflect.StructField, str string) (maxLength int,
	err error) {
	if str == "" {
		maxLength = 0
		return
	}
	kind := field.Type.Kind()
	if (kind == reflect.String || kind == reflect.Slice || kind == reflect.Map) &&
		field.Type.PkgPath() == "" {
		maxLength, err = strconv.Atoi(str)
		if err != nil {
			return
		}
		if maxLength < 0 {
			err = ErrNegativeMaxLength
			return
		}
		return
	}
	err = NewInvalidTagPartFormatError(str)
	return
}
