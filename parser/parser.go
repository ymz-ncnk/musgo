package parser

import (
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/ymz-ncnk/musgen"
)

const TagKey = "mus"

const InvalidTagFormatErrMsg = `%v field has invalid tag, it should be: ` +
	`"-" or "validator,maxLength,elemValidator,keyValidator"`

const InvalidTagMaxLengthErrMsg = "%v field has invalid tag, because of " +
	"maxLength"

const InvalidTagOwnMaxLengthErrMsg = "%v field has invalid tag, maxLength " +
	"could have only string, slice or map"

const InvalidTagOwnElemValidatorErrMsg = "%v field has invalid tag, " +
	"elemValidator could have only array, slice or map"

const InvalidTagOwnKeyValidatorErrMsg = "%v field has invalid tag, " +
	"keyValidator could have only map"

const InvalidTagArrayMaxLengthErrMsg = "%v field has invalid tag, maxLength " +
	"specified for array"

const InvalidTagNegativeMaxLength = "%v field has invalid tag, maxLength is " +
	"negative"

type NotSupportedTypeError struct {
	t string
}

func (err NotSupportedTypeError) Type() string {
	return err.t
}

func (err NotSupportedTypeError) Error() string {
	return fmt.Sprintf("%v type is not supported", err.Type())
}

// Parse creates TypeDesc from the specified type. It handles and public, and
// private fields. If the type is not an alias or struct returns error.
// If type is an alias to a pointer type returns error.
//
// Adds to each map type a "map number". For example, map[string]int becomes
// map-0[string]-0int. With help of map numbers we could parse map types
// correctly in situations like map[*map[string]int]int.
//
// Each field of the struct type could have a tag: `mus:-` or
// `mus:validator,maxLength,elemValidator,keyValidator`
// Only string, array, slice, map fields could have maxLenght validator.
// MaxLength should be positive number.
// Only array, slice, map fields could have elemValidator.
// And only map fields could have keyValidation.
// Otherwise returns error.
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

// ParseAlias tries to parse alias type(not an alias to a struct).
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

// ParseFieldTag tries to parse a field's tag. Returns true, if the field should
// be skipped.
func ParseFieldTag(f reflect.StructField, field *musgen.FieldDesc) (skip bool,
	err error) {
	tag, pst := f.Tag.Lookup(TagKey)
	if pst {
		_, st, _ := ParseStars(f.Type)
		k := st.Kind()
		vals := strings.Split(tag, ",")
		if len(vals) > 1 {
			if vals[0] == "-" {
				return false, fmt.Errorf(InvalidTagFormatErrMsg, field.Name)
			}
		}
		if len(vals) == 1 {
			if vals[0] == "-" {
				return true, nil
			} else {
				setUpValidator(field, vals[0])
			}
		} else if len(vals) == 2 {
			if (k == reflect.String || k == reflect.Slice || k == reflect.Map) &&
				f.Type.PkgPath() == "" {
				// if pkg != "" it's alias
				setUpValidator(field, vals[0])
				err := setUpMaxLength(field, vals[1])
				if err != nil {
					return false, fmt.Errorf(InvalidTagMaxLengthErrMsg, field.Name)
				}
			} else {
				return false, fmt.Errorf(InvalidTagOwnMaxLengthErrMsg, field.Name)
			}
		} else if len(vals) == 3 {
			if (k == reflect.Array || k == reflect.Slice || k == reflect.Map) &&
				f.Type.PkgPath() == "" {
				setUpValidator(field, vals[0])
				if k != reflect.Array {
					err := setUpMaxLength(field, vals[1])
					if err != nil {
						return false, fmt.Errorf(InvalidTagMaxLengthErrMsg, field.Name)
					}
				} else {
					if vals[1] != "" {
						return false, fmt.Errorf(InvalidTagArrayMaxLengthErrMsg, field.Name)
					}
				}
				setUpElemValidator(field, vals[2])
			} else {
				return false, fmt.Errorf(InvalidTagOwnElemValidatorErrMsg, field.Name)
			}
		} else if len(vals) == 4 {
			if k == reflect.Map && f.Type.PkgPath() == "" {
				setUpValidator(field, vals[0])
				err := setUpMaxLength(field, vals[1])
				if err != nil {
					return false, fmt.Errorf(InvalidTagMaxLengthErrMsg, field.Name)
				}
				setUpElemValidator(field, vals[2])
				setUpKeyValidator(field, vals[3])
			} else {
				return false, fmt.Errorf(InvalidTagOwnKeyValidatorErrMsg, field.Name)
			}
		} else {
			return false, fmt.Errorf(InvalidTagFormatErrMsg, field.Name)
		}
	}
	return false, nil
}

func setUpValidator(field *musgen.FieldDesc, value string) {
	if value != "" {
		field.Validator = value
	}
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

func setUpElemValidator(field *musgen.FieldDesc, value string) {
	if value != "" {
		field.ElemValidator = value
	}
}

func setUpKeyValidator(field *musgen.FieldDesc, value string) {
	if value != "" {
		field.KeyValidator = value
	}
}

// ParseType parses type into string representation. Complex types are parsed
// recursively.
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
		} else {
			// here could be an alias of an interface type
			return "", mapsCount, NotSupportedTypeError{st.String()}
		}
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
// to pointer type returns error.
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

// ParseArrayType returns string representation of the array type.
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

// ParseArrayType returns string representation of the slice type.
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

// ParseArrayType returns string representation of the map type.
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
	return filepath.Base(t.PkgPath())
}
