package parser

import (
	"reflect"
	"strconv"
)

// TagParser parses tags, used by Parse() function.
type TagParser func(tp reflect.Type, field reflect.StructField,
	tag reflect.StructTag) (arr []any, err error)

// Parse can accepts alias or sturct type, for other types returns
// ErrUnsupportedType.
// For alias type creates string representation of the underlying type and
// returns it as the aliasOf value.
// For struct type (alias to struct is also a sturct) - for each struct 
// field creates a string representation of its type, puts all this in 
// the fieldsTypes value.
// Adds to each map type a "map number". For example, map[string]int becomes
// map-0[string]-0int. With help of map numbers we could parse map types
// correctly in situations like map[*map[string]int]int.
func Parse(tp reflect.Type, tagParser TagParser) (aliasOf string, 
	fieldsTypes []string, 
	fieldsProps [][]any, 
	err error,
	) {
		if tp == nil {
			err = ErrUnsupportedType
			return
		}
		if aliasType(tp) {
			aliasOf, err = parseAlias(tp)
			return
		}
		if definedStrucType(tp) {
			fieldsTypes, fieldsProps, err = parseStruct(tp, tagParser)
			return
		}
		err = ErrUnsupportedType
		return
	}

	// parseAlias tries to parse an alias type(not an alias to a struct).
func parseAlias(tp reflect.Type) (aliasOf string, err error) {
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
			err = ErrUnsupportedType
		}
	}
	return
}

// parseStruct tries to parse a struct type or an alias to struct type.
func parseStruct(tp reflect.Type, tagParser TagParser) (fieldsTypes []string,
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
			if err != nil {
				return
			}
		}
		fieldsTypes = append(fieldsTypes, fieldType)
	}
	return
}

// parseType parses a type into the string representation. Complex types are
// parsed recursively.
// Translates a custom type to its name or package + name(if it's from
// another package).
// Adds a "map number" to the map type. In this case returns an incremented
// mapsCount.
// All other types leaves untouched.
func parseType(tp reflect.Type, currPkg string, mapsCount int) (tpStr string,
	mapsCountOut int, err error) {
	stars, tp := ParsePtrType(tp)
	mapsCountOut = mapsCount
	if aliasType(tp) || definedStrucType(tp) {
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
	err = ErrUnsupportedType
	return
}

// ParsePtrType returns pointer signs and an underlying type. If the underlying
// type is an alias to the pointer type returns an error.
func ParsePtrType(tp reflect.Type) (stars string, atp reflect.Type) {
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
	mapsCountStr := strconv.Itoa(mapsCount)
	tpStr = stars + "map-" + mapsCountStr + "[" + keyTpStr + "]-" + mapsCountStr +
		elemTpStr
	mapsCountOut = mapsCount + 1
	return
}

func hasPkg(tp reflect.Type) bool {
	return tp.PkgPath() != ""
}

func definedStrucType(tp reflect.Type) bool {
	return hasPkg(tp) && strucType(tp)
}

// struct or alias to struct, also returns true for struct{}{}
func strucType(tp reflect.Type) bool {
	return tp.Kind() == reflect.Struct
}

func interfaceType(tp reflect.Type) bool {
	return tp.Kind() == reflect.Interface
}

// alias to primitive type, array, slice or map
func aliasType(tp reflect.Type) bool {
	return hasPkg(tp) && !strucType(tp) && !interfaceType(tp)
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
