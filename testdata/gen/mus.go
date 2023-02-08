//go:build ignore

package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"runtime"

	"github.com/ymz-ncnk/musgen"
	musgen_textmpl "github.com/ymz-ncnk/musgen/text_template"
	tdmg "github.com/ymz-ncnk/musgo/testdata/musgen"
	"golang.org/x/tools/imports"
)

const folderName = "testdata/musgen"

// main could generate unsafe code if 'ARG=unsafe' is present:
// $ ARG=unsafe go generate
func main() {
	primAlias := flag.Bool("prim", false, "generate prim alias")
	arrayAlias := flag.Bool("array", false, "generate array alias")
	sliceAlias := flag.Bool("slice", false, "generate slice alias")
	mapAlias := flag.Bool("map", false, "generate map alias")
	structType := flag.Bool("struct", false, "generate struct")
	intRaw := flag.Bool("intraw", false, "generate intraw")
	flag.Parse()
	unsafe := false
	args := flag.Args()
	if args[0] == "unsafe" {
		unsafe = true
	}
	var err error
	if *primAlias && !*arrayAlias && !*sliceAlias && !*mapAlias && !*structType &&
		!*intRaw {
		err = generatePrimAlias(unsafe)
	} else if !*primAlias && *arrayAlias && !*sliceAlias && !*mapAlias &&
		!*structType && !*intRaw {
		err = generateArrayAlias(unsafe)
	} else if !*primAlias && !*arrayAlias && *sliceAlias && !*mapAlias &&
		!*structType && !*intRaw {
		err = generateSliceAlias(unsafe)
	} else if !*primAlias && !*arrayAlias && !*sliceAlias && *mapAlias &&
		!*structType && !*intRaw {
		err = generateMapAlias(unsafe)
	} else if !*primAlias && !*arrayAlias && !*sliceAlias && !*mapAlias &&
		*structType && !*intRaw {
		err = generateStructType(unsafe)
	} else if !*primAlias && !*arrayAlias && !*sliceAlias && !*mapAlias &&
		!*structType && *intRaw {
		err = generateRaw(unsafe)
	} else {
		err = errors.New("invalid flag")
	}
	if err != nil {
		panic(err)
	}
}

func generatePrimAlias(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, tdmg.Uint64AliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint64RawAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint32AliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint16AliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint8AliasTypeDesc)
	allTypes = append(allTypes, tdmg.UintAliasTypeDesc)

	allTypes = append(allTypes, tdmg.Int64AliasTypeDesc)
	allTypes = append(allTypes, tdmg.Int32AliasTypeDesc)
	allTypes = append(allTypes, tdmg.Int16AliasTypeDesc)
	allTypes = append(allTypes, tdmg.Int8AliasTypeDesc)
	allTypes = append(allTypes, tdmg.IntAliasTypeDesc)

	allTypes = append(allTypes, tdmg.Float64AliasTypeDesc)
	allTypes = append(allTypes, tdmg.Float32AliasTypeDesc)

	allTypes = append(allTypes, tdmg.BoolAliasTypeDesc)
	allTypes = append(allTypes, tdmg.ByteAliasTypeDesc)
	allTypes = append(allTypes, tdmg.StringAliasTypeDesc)

	return generate(allTypes, folderName, unsafe)
}

func generateArrayAlias(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, tdmg.StrArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.IntAliasTypeDesc)
	allTypes = append(allTypes, tdmg.IntAliasArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.FloatPtrArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint64PtrAliasArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Int64SliceArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint16SlicePtrArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint64AliasTypeDesc)
	allTypes = append(allTypes, tdmg.BoolArrayArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.BytePtrArrayPtrArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.IntStrMapArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint32Int32MapArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.SimpleStructTypeDesc)
	allTypes = append(allTypes, tdmg.StructTypeArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.StructTypePtrArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.TrickyArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.StrZeroLengthArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.IntPtrPtrPtrAliasArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.ValidIntArrayAliasTypeDesc)

	return generate(allTypes, folderName, unsafe)
}

func generateSliceAlias(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, tdmg.StrSliceAliasTypeDesc)
	allTypes = append(allTypes, tdmg.IntAliasTypeDesc)
	allTypes = append(allTypes, tdmg.IntAliasSliceAliasTypeDesc)
	allTypes = append(allTypes, tdmg.FloatPtrSliceAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint64PtrAliasSliceAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint64AliasTypeDesc)
	allTypes = append(allTypes, tdmg.BoolSliceSliceAliasTypeDesc)
	allTypes = append(allTypes, tdmg.ByteArraySliceAliasTypeDesc)
	allTypes = append(allTypes, tdmg.FloatArrayPtrSliceAliasTypeDesc)
	allTypes = append(allTypes, tdmg.IntStrMapSliceAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint32Int32MapPtrSliceAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint32Int32MapPtrSliceAliasTypeDesc)
	allTypes = append(allTypes, tdmg.SimpleStructTypeDesc)
	allTypes = append(allTypes, tdmg.StructTypeSliceAliasTypeDesc)
	allTypes = append(allTypes, tdmg.StructTypePtrSliceAliasTypeDesc)
	allTypes = append(allTypes, tdmg.StringAliasTypeDesc)
	allTypes = append(allTypes, tdmg.TrickySliceAliasTypeDesc)
	allTypes = append(allTypes, tdmg.ValidUintSliceAliasTypeDesc)

	return generate(allTypes, folderName, unsafe)
}

func generateMapAlias(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, tdmg.StrIntMapAliasTypeDesc)
	allTypes = append(allTypes, tdmg.StrPtrIntPtrMapAliasTypeDesc)
	allTypes = append(allTypes, tdmg.StringAliasTypeDesc)
	allTypes = append(allTypes, tdmg.IntAliasTypeDesc)
	allTypes = append(allTypes, tdmg.StrAliasIntAliasMapAliasTypeDesc)
	allTypes = append(allTypes, tdmg.StrAliasPtrIntAliasPtrMapAliasTypeDesc)
	allTypes = append(allTypes, tdmg.BoolInt16SliceMapAliasTypeDesc)
	allTypes = append(allTypes, tdmg.ByteUint16SlicePtrMapAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Int32Float64ArrayMapAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Float32Uint32ArrayPtrMapAliasTypeDesc)
	allTypes = append(allTypes, tdmg.FloatByteBoolMapMapAliasTypeDesc)
	allTypes = append(allTypes, tdmg.UintIntStringMapPtrMapAliasTypeDesc)
	allTypes = append(allTypes, tdmg.SimpleStructTypeDesc)
	allTypes = append(allTypes, tdmg.StructTypeStructTypeMapAliasTypeDesc)
	allTypes = append(allTypes, tdmg.StructTypePtrStructTypePtrMapAliasTypeDesc)
	allTypes = append(allTypes, tdmg.TrickyMapAliasTypeDesc)
	allTypes = append(allTypes, tdmg.ValidStringIntMapAliasTypeDesc)

	return generate(allTypes, folderName, unsafe)
}

func generateStructType(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, tdmg.IntAliasTypeDesc)
	allTypes = append(allTypes, tdmg.StringAliasTypeDesc)
	allTypes = append(allTypes, tdmg.SimpleStructTypeDesc)
	allTypes = append(allTypes, tdmg.StructTypeDesc)
	allTypes = append(allTypes, tdmg.ValidStructTypeDesc)
	allTypes = append(allTypes, tdmg.FieldlessStructTypeDesc)

	return generate(allTypes, folderName, unsafe)
}

func generateRaw(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, tdmg.Uint64RawAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint32RawAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint16RawAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint8RawAliasTypeDesc)
	allTypes = append(allTypes, tdmg.UintRawAliasTypeDesc)

	allTypes = append(allTypes, tdmg.Int64RawAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Int32RawAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Int16RawAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Int8RawAliasTypeDesc)
	allTypes = append(allTypes, tdmg.IntRawAliasTypeDesc)

	allTypes = append(allTypes, tdmg.Float64RawAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Float32RawAliasTypeDesc)
	allTypes = append(allTypes, tdmg.IntRawArrayAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Uint16Int32RawMapAliasTypeDesc)
	allTypes = append(allTypes, tdmg.Float64RawPtrPtrPtrAliasSliceAliasTypeDesc)
	allTypes = append(allTypes, tdmg.RawStructTypeDesc)
	allTypes = append(allTypes, tdmg.ValidInt32RawTypeDesc)

	return generate(allTypes, folderName, unsafe)
}

func generate(tds []musgen.TypeDesc, folder string, unsafe bool) error {
	var (
		_, filename, _, _ = runtime.Caller(1)
		b                 []byte
	)
	musGen, err := musgen_textmpl.New()
	if err != nil {
		return err
	}
	for _, td := range tds {
		td.Unsafe = unsafe
		filename = makeFileName(td)
		b, err = musGen.Generate(td, musgen.GoLang)
		if err != nil {
			return err
		}
		b, err = imports.Process("", b, nil)
		if err != nil {
			return err
		}
		if folder == "" {
			err = ioutil.WriteFile(filename, b, 0644)
		} else {
			err = ioutil.WriteFile(folder+"/"+filename, b, 0644)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func makeFileName(td musgen.TypeDesc) string {
	return "a__" + td.Name + ".mus.go"
}
