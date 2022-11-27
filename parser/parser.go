package parser

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/ymz-ncnk/musgen"
)

// TagKey represents a key of a tag, which marks field validators and
// encodings.
const TagKey = "mus"

// EncodingSep devides validator and encoding.
const EncodingSep = "#"

// InvalidTagFormatErrMsg happens if a tag has an invalid format.
const InvalidTagFormatErrMsg = `%v field has invalid tag, it should be: ` +
	`"-" or "validator#raw,maxLength,elemValidator#raw,keyValidator#raw"`

// InvalidTagMaxLengthErrMsg happens if a maxLength is invalid, for example is
// negative.
const InvalidTagMaxLengthErrMsg = "%v field has invalid tag, because of " +
	"maxLength"

// InvalidTagOwnMaxLengthErrMsg happens if a maxLength is set for not the
// supported type.
const InvalidTagOwnMaxLengthErrMsg = "%v field has invalid tag, " +
	"only strings, slices or maps could have maxLength"

// InvalidTagOwnElemValidatorErrMsg happens if an elemValidator is set for the
// not supported type.
const InvalidTagOwnElemValidatorErrMsg = "%v field has invalid tag, " +
	"only arrays, slices or maps could have elemValidator or elemEncoding"

// InvalidTagOwnKeyValidatorErrMsg happens if a keyValidator is set for the not
// supported type.
const InvalidTagOwnKeyValidatorErrMsg = "%v field has invalid tag, " +
	"only maps could have keyValidator or keyEncoding"

// InvalidTagArrayMaxLengthErrMsg happens when a maxLength is speccified for an
// array.
const InvalidTagArrayMaxLengthErrMsg = "%v field has invalid tag, maxLength " +
	"is specified for an array"

const InvalidTagPartFormatErrMsg = "invalid format of the %s"

var ErrUndefinedEncoding = errors.New("undefined encoding")
var ErrUnsupportedEncoding = errors.New("unsupported encoding")
var ErrUnsupportedElemEncoding = errors.New("unsupported elem encoding")
var ErrUnsupportedKeyEncoding = errors.New("unsupported key encoding")

var ErrOnlyUintAndIntMayHaveRawEncoding = errors.New("only uint and int " +
	"types may have #raw encoding")

// NotSupportedTypeError happens when the parser tries to parse the not
// supported type.
type NotSupportedTypeError struct {
	t string
}

// Type returns the not supported type.
func (err NotSupportedTypeError) Type() string {
	return err.t
}

func (err NotSupportedTypeError) Error() string {
	return fmt.Sprintf("%v type is not supported", err.Type())
}

// Parse creates a TypeDesc from the specified type. It handles and public, and
// private fields. If the type is not an alias or struct returns an error.
// If type is an alias to a pointer type returns an error.
//
// Adds to each map type a "map number". For example, map[string]int becomes
// map-0[string]-0int. With help of map numbers we could parse map types
// correctly in situations like map[*map[string]int]int.
//
// Each field of the struct type could have a tag: `mus:-` or
// `mus:validator,maxLength,elemValidator,keyValidator`
// Only string, array, slice, map fields could have a maxLenght validator.
// The maxLength should be positive number.
// Only array, slice, map fields could have an elemValidator.
// And only map fields could have a keyValidator.
// Otherwise returns an error.
func Parse(t reflect.Type) (musgen.TypeDesc, error) {
	if t == nil {
		return musgen.TypeDesc{}, errors.New("type is nil")
	}
	if t.PkgPath() == "" {
		// if not an alias or struct
		return musgen.TypeDesc{}, NotSupportedTypeError{t.String()}
	} else if t.Kind() == reflect.Ptr {
		// if alias of the pointer type
		return musgen.TypeDesc{}, NotSupportedTypeError{t.String()}
	}
	td, ok, err := ParseAlias(t)
	if ok {
		return td, nil
	}
	if err != nil {
		return musgen.TypeDesc{}, err
	}
	td, err = ParseStruct(t)
	if err != nil {
		return musgen.TypeDesc{}, err
	}
	return td, nil
}

// ParseAlias tries to parse an alias type(not an alias to a struct).
// Returns true on success.
func ParseAlias(t reflect.Type) (musgen.TypeDesc, bool, error) {
	var ft string
	var err error
	k := t.Kind()
	if primitive(k) {
		ft = k.String()
	} else if k == reflect.Array {
		ft, _, err = ParseArrayType("", t, t.PkgPath(), 0)
		if err != nil {
			return musgen.TypeDesc{}, false, err
		}
	} else if k == reflect.Slice {
		ft, _, err = ParseSliceType("", t, t.PkgPath(), 0)
		if err != nil {
			return musgen.TypeDesc{}, false, err
		}
	} else if k == reflect.Map {
		ft, _, err = ParseMapType("", t, t.PkgPath(), 0)
		if err != nil {
			return musgen.TypeDesc{}, false, err
		}
	} else {
		return musgen.TypeDesc{}, false, nil
	}
	return musgen.TypeDesc{
		Package: pkg(t),
		Name:    t.Name(),
		Fields: []musgen.FieldDesc{{
			Alias: t.Name(),
			Type:  ft,
		}},
	}, true, nil
}

// ParseStruct tries to parse a struct type or an alias to struct type.
func ParseStruct(t reflect.Type) (musgen.TypeDesc, error) {
	var f reflect.StructField
	var ft string
	var err error
	fields := make([]musgen.FieldDesc, 0)
	k := t.Kind()
	if k == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			f = t.Field(i)
			ft, _, err = ParseType(f.Type, t.PkgPath(), 0)
			if err != nil {
				return musgen.TypeDesc{}, err
			}
			field := musgen.FieldDesc{
				Name: f.Name,
				Type: ft,
			}
			skip, err := ParseFieldTag(f, &field)
			if err != nil {
				return musgen.TypeDesc{}, err
			}
			if skip {
				continue
			}
			fields = append(fields, field)
		}
	}
	return musgen.TypeDesc{
		Package: pkg(t),
		Name:    t.Name(),
		Fields:  fields,
	}, nil
}

// ParseFieldTag tries to parse a field tag. Returns true, if the field should
// be skipped.
func ParseFieldTag(f reflect.StructField, field *musgen.FieldDesc) (skip bool,
	err error) {
	tag, pst := f.Tag.Lookup(TagKey)
	if pst {
		_, st, _ := ParseStars(f.Type)
		k := st.Kind()
		vals := strings.Split(tag, ",")
		if len(vals) > 4 {
			return false, fmt.Errorf(InvalidTagFormatErrMsg, field.Name)
		}
		if vals[0] == "-" {
			if len(vals) > 1 {
				return false, fmt.Errorf(InvalidTagFormatErrMsg, field.Name)
			}
			return true, nil
		}
		if len(vals) >= 1 {
			err = setUpValidatorAndEncoding(field, vals[0])
			if err != nil {
				return false, err
			}
		}
		if len(vals) >= 2 {
			if vals[1] != "" {
				if (k == reflect.String || k == reflect.Slice || k == reflect.Map) &&
					f.Type.PkgPath() == "" {
					err := setUpMaxLength(field, vals[1])
					if err != nil {
						return false, fmt.Errorf(InvalidTagMaxLengthErrMsg, field.Name)
					}
				} else {
					return false, fmt.Errorf(InvalidTagOwnMaxLengthErrMsg, field.Name)
				}
			}
		}
		if len(vals) >= 3 {
			if vals[2] != "" {
				if (k == reflect.Array || k == reflect.Slice || k == reflect.Map) &&
					f.Type.PkgPath() == "" {
					err = setUpElemValidatorAndEncoding(field, vals[2])
					if err != nil {
						return false, err
					}
				} else {
					return false, fmt.Errorf(InvalidTagOwnElemValidatorErrMsg, field.Name)
				}
			}
		}
		if len(vals) == 4 {
			if vals[3] != "" {
				if k == reflect.Map && f.Type.PkgPath() == "" {
					err = setUpKeyValidatorAndEncoding(field, vals[3])
					if err != nil {
						return false, err
					}
				} else {
					return false, fmt.Errorf(InvalidTagOwnKeyValidatorErrMsg, field.Name)
				}
			}
		}
	}
	return false, nil
}

// ParseType parses a type into the string representation. Complex types are
// parsed recursively.
// Translates a custom type to its name or package + name(if it's from
// another package).
// Adds a map number to the map type. In this case returns an incremented
// mapsCount.
// All other types leaves untouched.
func ParseType(t reflect.Type, pkgPath string,
	mapsCount int) (string, int, error) {
	stars, st, err := ParseStars(t)
	if err != nil {
		return "", mapsCount, err
	}
	k := st.Kind()
	if st.PkgPath() != "" {
		// alias of a simple type or struct
		if primitive(k) ||
			k == reflect.Array ||
			k == reflect.Slice ||
			k == reflect.Map ||
			k == reflect.Struct {
			if pkgPath == st.PkgPath() {
				// if type from the same package
				return stars + st.Name(), mapsCount, nil
			}
			return stars + st.String(), mapsCount, nil
		}
		// here could be an alias of an interface type
		return "", mapsCount, NotSupportedTypeError{st.String()}
	}
	if primitive(k) {
		return t.String(), mapsCount, nil
	}
	if k == reflect.Array {
		return ParseArrayType(stars, st, pkgPath, mapsCount)
	}
	if k == reflect.Slice {
		return ParseSliceType(stars, st, pkgPath, mapsCount)
	}
	if k == reflect.Map {
		return ParseMapType(stars, st, pkgPath, mapsCount)
	}
	return "", mapsCount, NotSupportedTypeError{st.String()}
}

// ParseStars returns pointer signs and real type. If real type is an alias
// to the pointer type returns an error.
func ParseStars(t reflect.Type) (stars string, st reflect.Type, err error) {
	st = t
	k := st.Kind()
	for {
		if k == reflect.Ptr {
			if st.PkgPath() != "" {
				// alias to pointer type for cases like **MyInt *int
				return stars, st, NotSupportedTypeError{st.String()}
			}
			st = st.Elem()
			k = st.Kind()
			stars += "*"
			continue
		}
		break
	}
	return
}

// ParseArrayType returns a string representation of the array type.
func ParseArrayType(stars string, t reflect.Type, pkgPath string,
	mapsCount int) (string, int, error) {
	var aelt string
	var err error
	aelt, mapsCount, err = ParseType(t.Elem(), pkgPath, mapsCount)
	if err != nil {
		return "", mapsCount, err
	}
	return stars + "[" + strconv.Itoa(t.Len()) + "]" + aelt, mapsCount, nil
}

// ParseSliceType returns a string representation of the slice type.
func ParseSliceType(stars string, t reflect.Type, pkgPath string,
	mapsCount int) (string, int, error) {
	var selt string
	var err error
	selt, mapsCount, err = ParseType(t.Elem(), pkgPath, mapsCount)
	if err != nil {
		return "", mapsCount, err
	}
	return stars + "[]" + selt, mapsCount, nil
}

// ParseMapType returns a string representation of the map type.
func ParseMapType(stars string, t reflect.Type, pkgPath string,
	mapsCount int) (string, int, error) {
	var mkt string
	var mvt string
	var err error
	mkt, mapsCount, err = ParseType(t.Key(), pkgPath, mapsCount)
	if err != nil {
		return "", mapsCount, err
	}
	mvt, mapsCount, err = ParseType(t.Elem(), pkgPath, mapsCount)
	if err != nil {
		return "", mapsCount, err
	}
	strMapsCount := strconv.Itoa(mapsCount)
	ft := stars + "map-" + strMapsCount + "[" + mkt + "]-" + strMapsCount + mvt
	mapsCount++
	return ft, mapsCount, nil
}

func setUpValidatorAndEncoding(field *musgen.FieldDesc, value string) error {
	validator, encoding, err := parseValidatorAndEncoding(field, value)
	if err != nil {
		return err
	}
	field.Validator = validator
	field.Encoding = encoding
	return nil
	// return SetEncoding(field, encoding)
}

func setUpMaxLength(field *musgen.FieldDesc, value string) error {
	if value != "" {
		maxLength, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		if maxLength < 0 {
			return errors.New("negative length")
		}
		field.MaxLength = maxLength
	}
	return nil
}

// func setUpElemValidator(field *musgen.FieldDesc, value string) {
// 	if value != "" {
// 		field.ElemValidator = value
// 	}
// }

func setUpElemValidatorAndEncoding(field *musgen.FieldDesc,
	value string) error {
	validator, encoding, err := parseValidatorAndEncoding(field, value)
	if err != nil {
		return err
	}
	field.ElemValidator = validator
	field.ElemEncoding = encoding
	return nil
}

func setUpKeyValidatorAndEncoding(field *musgen.FieldDesc, value string) error {
	validator, encoding, err := parseValidatorAndEncoding(field, value)
	if err != nil {
		return err
	}
	field.KeyValidator = validator
	field.KeyEncoding = encoding
	return nil
}

func parseValidatorAndEncoding(field *musgen.FieldDesc, value string) (
	validator, encoding string, err error) {
	if value == "" {
		return "", "", nil
	}
	if strings.HasPrefix(value, EncodingSep) {
		return "", value[1:], nil
	}
	vals := strings.Split(value, EncodingSep)
	if len(vals) > 2 {
		return "", "", fmt.Errorf(InvalidTagPartFormatErrMsg, value)
	}
	if len(vals) == 2 {
		return vals[0], vals[1], nil
	}
	return value, "", nil
}

func primitive(k reflect.Kind) bool {
	return k == reflect.Bool ||
		k == reflect.Int ||
		k == reflect.Int8 ||
		k == reflect.Int16 ||
		k == reflect.Int32 ||
		k == reflect.Int64 ||
		k == reflect.Uint ||
		k == reflect.Uint8 ||
		k == reflect.Uint16 ||
		k == reflect.Uint32 ||
		k == reflect.Uint64 ||
		k == reflect.Float32 ||
		k == reflect.Float64 ||
		// k == reflect.Complex64 ||
		// k == reflect.Complex128 ||
		k == reflect.String
}

func pkg(t reflect.Type) string {
	re := regexp.MustCompile(`^(.*)\.`)
	match := re.FindStringSubmatch(t.String())
	if len(match) != 2 {
		return ""
	}
	return match[1]
}
