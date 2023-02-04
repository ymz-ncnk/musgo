package parser

import (
	"reflect"
	"strconv"
)

type TagParser func(tp reflect.Type, field reflect.StructField,
	tag reflect.StructTag) (arr []any, err error)

// TODO
// Parse creates a gen4type.TypeDesc from the specified type. It handles and
// public, and private fields. If the type is not an alias or struct returns an
// error.
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

// Parse accepts alias or sturct types. For other types returns
// UnsupportedTypeError.
// For alias type creates string representation of the underlying type and
// returns it as aliasOf value.
// For struct type (alias to struct is also a sturct) - for each struct field
// creates string representation of its type, puts all this in fieldsTypes
// value.
// Adds to each map type a "map number". For example, map[string]int becomes
// map-0[string]-0int. With help of map numbers we could parse map types
// correctly in situations like map[*map[string]int]int.
func Parse(tp reflect.Type) (aliasOf string, fieldsTypes []string, err error) {
	if tp == nil {
		err = NewUnsupportedTypeError("nil")
		return
	}
	if aliasType(tp) {
		aliasOf, fieldsTypes, err = parseAlias(tp)
		return
	}
	// TODO Should check hasPkg()?
	if structType(tp) {
		aliasOf, fieldsTypes, _, err = parseStruct(tp, nil)
		return
	}
	err = NewUnsupportedTypeError(tp.String())
	return
}

// ParseStructWithTags accepts struct type (alias to struct is also a sturct).
// For each struct field creates string representation of its type (fieldsTypes)
// + with help of the TagParser cretes its properties (fieldsProps value).
func ParseStructWithTags(tp reflect.Type, tagParser TagParser) (
	fieldsTypes []string,
	fieldsProps [][]any, err error,
) {
	if tp == nil {
		err = NewUnsupportedTypeError("nil")
		return
	}
	if !structType(tp) {
		err = NewUnsupportedTypeError(tp.String())
		return
	}
	_, fieldsTypes, fieldsProps, err = parseStruct(tp, tagParser)
	return
}

// parseAlias tries to parse an alias type(not an alias to a struct).
// Returns true on success.
func parseAlias(tp reflect.Type) (aliasOf string, fieldsTypes []string,
	err error) {
	if primitiveType(tp) {
		aliasOf = tp.Kind().String()
	} else {
		switch tp.Kind() {
		case reflect.Array:
			aliasOf, _, err = parseArrayType("", tp, tp.PkgPath(), 0)
		case reflect.Slice:
			aliasOf, _, err = parseSliceType("", tp, tp.PkgPath(), 0)
		case reflect.Map:
			aliasOf, _, err = parseMapType("", tp, tp.PkgPath(), 0)
		default:
			err = NewUnsupportedTypeError(tp.String())
		}
	}
	return
}

// parseStruct tries to parse a struct type or an alias to struct type.
func parseStruct(tp reflect.Type, tagParser TagParser) (aliasOf string,
	fieldsTypes []string,
	fieldsProps [][]any,
	err error,
) {
	var (
		field     reflect.StructField
		fieldType string
	)
	fieldsTypes = []string{}
	if tagParser != nil {
		fieldsProps = make([][]any, tp.NumField())
	}
	for i := 0; i < tp.NumField(); i++ {
		field = tp.Field(i)
		fieldType, _, err = parseType(field.Type, tp.PkgPath(), 0)
		if err != nil {
			return
		}
		if tagParser != nil {
			fieldsProps[i], err = tagParser(tp, field, field.Tag)
		}
		fieldsTypes = append(fieldsTypes, fieldType)
	}
	return
}

// parseType parses a type into the string representation. Complex types are
// parsed recursively.
// Translates a custom type to its name or package + name(if it's from
// another package).
// Adds a map number to the map type. In this case returns an incremented
// mapsCount.
// All other types leaves untouched.
func parseType(tp reflect.Type, currPkg string, mapsCount int) (tpStr string,
	mapsCountOut int, err error) {
	stars, tp, err := parsePtrType(tp)
	if err != nil {
		return
	}
	mapsCountOut = mapsCount
	// TODO Do not types like struct{}{}?
	if aliasType(tp) || structType(tp) {
		if currPkg == tp.PkgPath() {
			tpStr = stars + tp.Name()
			return
		}
		tpStr = stars + tp.String()
		return
	}
	if primitiveType(tp) {
		tpStr = stars + tp.String()
		return
	}
	if tp.Kind() == reflect.Array {
		return parseArrayType(stars, tp, currPkg, mapsCount)
	}
	if tp.Kind() == reflect.Slice {
		return parseSliceType(stars, tp, currPkg, mapsCount)
	}
	if tp.Kind() == reflect.Map {
		return parseMapType(stars, tp, currPkg, mapsCount)
	}
	err = NewUnsupportedTypeError(tp.String())
	return
}

// parsePtrType returns pointer signs and an underlying type. If the underlying
// type is an alias to the pointer type returns an error.
func parsePtrType(tp reflect.Type) (stars string, atp reflect.Type, err error) {
	atp = tp
	for {
		if !ptrType(atp) {
			return
		}
		atp = atp.Elem()
		stars += "*"
		continue
	}
}

// parseArrayType returns a string representation of the array type.
func parseArrayType(stars string, tp reflect.Type, pkgPath string,
	mapsCount int) (tpStr string, mapsCountOut int, err error) {
	elemTpStr, mapsCount, err := parseType(tp.Elem(), pkgPath, mapsCount)
	if err != nil {
		return "", mapsCount, err
	}
	tpStr = stars + "[" + strconv.Itoa(tp.Len()) + "]" + elemTpStr
	mapsCountOut = mapsCount
	return
}

// parseSliceType returns a string representation of the slice type.
func parseSliceType(stars string, tp reflect.Type, pkgPath string,
	mapsCount int) (tpStr string, mapsCountOut int, err error) {
	elemTpStr, mapsCount, err := parseType(tp.Elem(), pkgPath, mapsCount)
	if err != nil {
		return "", mapsCount, err
	}
	tpStr = stars + "[]" + elemTpStr
	mapsCountOut = mapsCount
	return
}

// parseMapType returns a string representation of the map type.
func parseMapType(stars string, tp reflect.Type, pkgPath string,
	mapsCount int) (tpStr string, mapsCountOut int, err error) {
	keyTpStr, mapsCount, err := parseType(tp.Key(), pkgPath, mapsCount)
	if err != nil {
		return "", mapsCount, err
	}
	elemTpStr, mapsCount, err := parseType(tp.Elem(), pkgPath, mapsCount)
	if err != nil {
		return "", mapsCount, err
	}
	strMapsCount := strconv.Itoa(mapsCount)
	tpStr = stars + "map-" + strMapsCount + "[" + keyTpStr + "]-" + strMapsCount +
		elemTpStr
	mapsCountOut = mapsCount + 1
	return
}

func hasPkg(tp reflect.Type) bool {
	return tp.PkgPath() != ""
}

// struct or alias to struct, also returns true for struct{}{}
func structType(tp reflect.Type) bool {
	return tp.Kind() == reflect.Struct
}

func interfaceType(tp reflect.Type) bool {
	return tp.Kind() == reflect.Interface
}

// alias to primitive type, array, slice or map
func aliasType(tp reflect.Type) bool {
	return hasPkg(tp) && !structType(tp) && !interfaceType(tp)
}

func primitiveType(tp reflect.Type) bool {
	kind := tp.Kind()
	return kind == reflect.Bool ||
		kind == reflect.Int ||
		kind == reflect.Int8 ||
		kind == reflect.Int16 ||
		kind == reflect.Int32 ||
		kind == reflect.Int64 ||
		kind == reflect.Uint ||
		kind == reflect.Uint8 ||
		kind == reflect.Uint16 ||
		kind == reflect.Uint32 ||
		kind == reflect.Uint64 ||
		kind == reflect.Float32 ||
		kind == reflect.Float64 ||
		// kind == reflect.Complex64 ||
		// kind == reflect.Complex128 ||
		kind == reflect.String
}

func ptrType(tp reflect.Type) bool {
	return tp.Kind() == reflect.Pointer
}
