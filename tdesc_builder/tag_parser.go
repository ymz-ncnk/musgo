package tdesc_builder

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/ymz-ncnk/musgo/v2/parser"
)

// TagKey represents a key of a tag, which marks field validators and
// encodings.
const MUSTagKey = "mus"

// EncodingSep devides validator and encoding.
const EncodingSep = "#"

// RawEncoding for all int, uint and float types.
const RawEncoding = "raw"

// Tag structure:
// "mus:-" or
// "mus:Validator#Encoding,MaxLength,ElemValidator#ElemEncoding,KeyValidator#KeyEncoding"
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
		fieldProps[0], fieldProps[1], err = parseFieldValidatorAndEncoding(field,
			subTags[0])
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
		fieldProps[3], fieldProps[4], err = parseElemValidatorAndEncoding(field,
			subTags[2])
		if err != nil {
			return
		}
	}
	if len(subTags) == 4 {
		// keyValidator, keyEncoding
		fieldProps[5], fieldProps[6], err = parseKeyValidatorAndEncoding(field,
			subTags[3])
		if err != nil {
			return
		}
	}
	return
}

func parseFieldValidatorAndEncoding(field reflect.StructField, str string) (
	validator, encoding string, err error) {
	validator, encoding, err = parseValidatorAndEncoding(field, str)
	if err != nil {
		return
	}
	if encoding != "" && !supportRawEncoding(field.Type) {
		err = NewInvalidTagFormatError(field.Name)
		return
	}
	return
}

func parseElemValidatorAndEncoding(field reflect.StructField, str string) (
	validator, encoding string, err error) {
	validator, encoding, err = parseValidatorAndEncoding(field, str)
	if err != nil {
		return
	}
	if validator != "" && !supportElemValidator(field.Type) {
		err = NewInvalidTagFormatError(field.Name)
		return
	}
	if encoding != "" && !supportElemEncoding(field.Type) {
		err = NewInvalidTagFormatError(field.Name)
		return
	}
	return
}

func parseKeyValidatorAndEncoding(field reflect.StructField, str string) (
	validator, encoding string, err error) {
	validator, encoding, err = parseValidatorAndEncoding(field, str)
	if err != nil {
		return
	}
	if validator != "" && !supportKeyValidator(field.Type) {
		err = NewInvalidTagFormatError(field.Name)
		return
	}
	if encoding != "" && !supportKeyEncoding(field.Type) {
		err = NewInvalidTagFormatError(field.Name)
		return
	}
	return
}

func parseValidatorAndEncoding(field reflect.StructField, str string) (
	validator, encoding string, err error) {
	if str == "" {
		return "", "", nil
	}
	if strings.HasPrefix(str, EncodingSep) {
		return "", str[1:], nil
	}
	parts := strings.Split(str, EncodingSep)
	if len(parts) > 2 {
		return "", "", NewInvalidTagFormatError(field.Name)
	}
	if len(parts) == 2 {
		if parts[1] != RawEncoding {
			return "", "", NewInvalidTagFormatError(field.Name)
		}
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
	err = NewInvalidTagFormatError(field.Name)
	return
}

func supportRawEncoding(tp reflect.Type) bool {
	_, tp = parser.ParsePtrType(tp)
	return tp.Kind() == reflect.Int64 ||
		tp.Kind() == reflect.Int32 ||
		tp.Kind() == reflect.Int16 ||
		tp.Kind() == reflect.Int8 ||
		tp.Kind() == reflect.Int ||

		tp.Kind() == reflect.Uint64 ||
		tp.Kind() == reflect.Uint32 ||
		tp.Kind() == reflect.Uint16 ||
		tp.Kind() == reflect.Uint8 ||
		tp.Kind() == reflect.Uint ||

		tp.Kind() == reflect.Float64 ||
		tp.Kind() == reflect.Float32
}

func supportElemValidator(tp reflect.Type) bool {
	_, tp = parser.ParsePtrType(tp)
	return (tp.Kind() == reflect.Array ||
		tp.Kind() == reflect.Slice ||
		tp.Kind() == reflect.Map) && tp.PkgPath() == ""
}

func supportElemEncoding(tp reflect.Type) bool {
	if !supportElemValidator(tp) {
		return false
	}
	_, tp = parser.ParsePtrType(tp)
	return supportRawEncoding(tp.Elem())
}

func supportKeyValidator(tp reflect.Type) bool {
	_, tp = parser.ParsePtrType(tp)
	return tp.Kind() == reflect.Map && tp.PkgPath() == ""
}

func supportKeyEncoding(tp reflect.Type) bool {
	if !supportKeyValidator(tp) {
		return false
	}
	_, tp = parser.ParsePtrType(tp)
	return supportRawEncoding(tp.Elem())
}
