package main

import (
	"flag"
	"os"
	"reflect"
	"runtime"

	"github.com/ymz-ncnk/amock"
	"github.com/ymz-ncnk/musgen/v2"
	musgen_textmpl "github.com/ymz-ncnk/musgen/v2/text_template"
	"github.com/ymz-ncnk/musgo/v2"
	testdata_musgen "github.com/ymz-ncnk/musgo/v2/testdata/musgen"
	persistor_mod "github.com/ymz-ncnk/persistor"
	"golang.org/x/tools/imports"
)

const folderName = "testdata/musgen"

func main() {
	err := generateMocks()
	if err != nil {
		panic(err)
	}
	err = generateTestdata()
	if err != nil {
		panic(err)
	}
}

func generateMocks() (err error) {
	aMock, err := amock.New()
	if err != nil {
		return
	}
	err = aMock.Generate(reflect.TypeOf((*persistor_mod.Persistor)(nil)).Elem())
	if err != nil {
		return
	}
	err = aMock.Generate(reflect.TypeOf((*musgen.MusGen)(nil)).Elem())
	if err != nil {
		return
	}
	return
}

func generateTestdata() (err error) {
	flag.Parse()
	unsafe := false
	args := flag.Args()
	if args[0] == "unsafe" {
		unsafe = true
	}
	err = generatePrimAlias(unsafe)
	if err != nil {
		return
	}
	err = generateArrayAlias(unsafe)
	if err != nil {
		return
	}
	err = generateSliceAlias(unsafe)
	if err != nil {
		return
	}
	err = generateMapAlias(unsafe)
	if err != nil {
		return
	}
	err = generateStructType(unsafe)
	if err != nil {
		return
	}
	return generateRaw(unsafe)
}

func generatePrimAlias(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, testdata_musgen.Uint64AliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint64RawAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint32AliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint16AliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint8AliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.UintAliasTypeDesc)

	allTypes = append(allTypes, testdata_musgen.Int64AliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Int32AliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Int16AliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Int8AliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.IntAliasTypeDesc)

	allTypes = append(allTypes, testdata_musgen.Float64AliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Float32AliasTypeDesc)

	allTypes = append(allTypes, testdata_musgen.BoolAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.ByteAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StringAliasTypeDesc)

	return generate(allTypes, folderName, unsafe)
}

func generateArrayAlias(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, testdata_musgen.StrArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.IntAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.IntAliasArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.FloatPtrArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint64PtrAliasArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Int64SliceArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint16SlicePtrArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint64AliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.BoolArrayArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.BytePtrArrayPtrArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.IntStrMapArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint32Int32MapArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.SimpleStructTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StructTypeArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StructTypePtrArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.TrickyArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StrZeroLengthArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.IntPtrPtrPtrAliasArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.ValidIntArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.ValidPtrIntArrayAliasTypeDesc)

	return generate(allTypes, folderName, unsafe)
}

func generateSliceAlias(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, testdata_musgen.StrSliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.IntAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.IntAliasSliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.FloatPtrSliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint64PtrAliasSliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint64AliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.BoolSliceSliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.ByteArraySliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.FloatArrayPtrSliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.IntStrMapSliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint32Int32MapPtrSliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint32Int32MapPtrSliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.SimpleStructTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StructTypeSliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StructTypePtrSliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StringAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.TrickySliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.ValidUintSliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.ValidPtrStringSliceAliasTypeDesc)

	return generate(allTypes, folderName, unsafe)
}

func generateMapAlias(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, testdata_musgen.StrIntMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StrPtrIntPtrMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StringAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.IntAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StrAliasIntAliasMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StrAliasPtrIntAliasPtrMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.BoolInt16SliceMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.ByteUint16SlicePtrMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Int32Float64ArrayMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Float32Uint32ArrayPtrMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.FloatByteBoolMapMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.UintIntStringMapPtrMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.SimpleStructTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StructTypeStructTypeMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StructTypePtrStructTypePtrMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.TrickyMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.ValidStringIntMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.ValidPtrIntPtrIntMapAliasTypeDesc)

	return generate(allTypes, folderName, unsafe)
}

func generateStructType(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, testdata_musgen.IntAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StringAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.SimpleStructTypeDesc)
	allTypes = append(allTypes, testdata_musgen.StructTypeDesc)
	allTypes = append(allTypes, testdata_musgen.ValidStructTypeDesc)
	allTypes = append(allTypes, testdata_musgen.FieldlessStructTypeDesc)

	return generate(allTypes, folderName, unsafe)
}

func generateRaw(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, testdata_musgen.Uint64RawAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint32RawAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint16RawAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint8RawAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.UintRawAliasTypeDesc)

	allTypes = append(allTypes, testdata_musgen.Int64RawAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Int32RawAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Int16RawAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Int8RawAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.IntRawAliasTypeDesc)

	allTypes = append(allTypes, testdata_musgen.Float64RawAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Float32RawAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.IntRawArrayAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Uint16Int32RawMapAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.Float64RawPtrPtrPtrAliasSliceAliasTypeDesc)
	allTypes = append(allTypes, testdata_musgen.RawStructTypeDesc)
	allTypes = append(allTypes, testdata_musgen.ValidInt32RawTypeDesc)

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
			err = os.WriteFile(filename, b, 0644)
		} else {
			err = os.WriteFile(folder+"/"+filename, b, 0644)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func makeFileName(td musgen.TypeDesc) string {
	return "a__" + td.Name + musgo.FilenameExtenstion
}
