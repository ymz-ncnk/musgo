// +build ignore

package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/ymz-ncnk/musgen"
	"github.com/ymz-ncnk/musgo"
	"github.com/ymz-ncnk/musgo/testdata"
	mgtd "github.com/ymz-ncnk/musgo/testdata/musgen"
	"golang.org/x/tools/imports"
)

// main could generate unsafe code if unsafe 'ARG=unsafe' is present
func main() {
	primAlias := flag.Bool("prim", false, "generate prim alias")
	arrayAlias := flag.Bool("array", false, "generate array alias")
	sliceAlias := flag.Bool("slice", false, "generate slice alias")
	mapAlias := flag.Bool("map", false, "generate map alias")
	structType := flag.Bool("struct", false, "generate struct")
	intRaw := flag.Bool("intraw", false, "generate intraw")
	musgo := flag.Bool("musgo", false, "generate musgo")
	flag.Parse()
	unsafe := false
	args := flag.Args()
	if args[0] == "unsafe" {
		unsafe = true
	}
	var err error
	if *primAlias && !*arrayAlias && !*sliceAlias && !*mapAlias && !*structType &&
		!*musgo && !*intRaw {
		err = generatePrimAlias(unsafe)
	} else if !*primAlias && *arrayAlias && !*sliceAlias && !*mapAlias &&
		!*structType && !*musgo && !*intRaw {
		err = generateArrayAlias(unsafe)
	} else if !*primAlias && !*arrayAlias && *sliceAlias && !*mapAlias &&
		!*structType && !*musgo && !*intRaw {
		err = generateSliceAlias(unsafe)
	} else if !*primAlias && !*arrayAlias && !*sliceAlias && *mapAlias &&
		!*structType && !*musgo && !*intRaw {
		err = generateMapAlias(unsafe)
	} else if !*primAlias && !*arrayAlias && !*sliceAlias && !*mapAlias &&
		*structType && !*musgo && !*intRaw {
		err = generateStructType(unsafe)
	} else if !*primAlias && !*arrayAlias && !*sliceAlias && !*mapAlias &&
		!*structType && *musgo && !*intRaw {
		err = generateMusGo(unsafe)
	} else if !*primAlias && !*arrayAlias && !*sliceAlias && !*mapAlias &&
		!*structType && !*musgo && *intRaw {
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
	allTypes = append(allTypes, mgtd.Uint64AliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint64RawAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint32AliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint16AliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint8AliasTypeDesc)
	allTypes = append(allTypes, mgtd.UintAliasTypeDesc)

	allTypes = append(allTypes, mgtd.Int64AliasTypeDesc)
	allTypes = append(allTypes, mgtd.Int32AliasTypeDesc)
	allTypes = append(allTypes, mgtd.Int16AliasTypeDesc)
	allTypes = append(allTypes, mgtd.Int8AliasTypeDesc)
	allTypes = append(allTypes, mgtd.IntAliasTypeDesc)

	allTypes = append(allTypes, mgtd.Float64AliasTypeDesc)
	allTypes = append(allTypes, mgtd.Float32AliasTypeDesc)

	allTypes = append(allTypes, mgtd.BoolAliasTypeDesc)
	allTypes = append(allTypes, mgtd.ByteAliasTypeDesc)
	allTypes = append(allTypes, mgtd.StringAliasTypeDesc)

	return generate(allTypes, "musgen", unsafe)
}

func generateArrayAlias(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, mgtd.StrArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.IntAliasTypeDesc)
	allTypes = append(allTypes, mgtd.IntAliasArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.FloatPtrArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint64PtrAliasArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Int64SliceArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint16SlicePtrArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint64AliasTypeDesc)
	allTypes = append(allTypes, mgtd.BoolArrayArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.BytePtrArrayPtrArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.IntStrMapArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint32Int32MapArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.SimpleStructTypeDesc)
	allTypes = append(allTypes, mgtd.StructTypeArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.StructTypePtrArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.TrickyArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.StrZeroLengthArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.IntPtrPtrPtrAliasArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.ValidIntArrayAliasTypeDesc)

	return generate(allTypes, "musgen", unsafe)
}

func generateSliceAlias(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, mgtd.StrSliceAliasTypeDesc)
	allTypes = append(allTypes, mgtd.IntAliasTypeDesc)
	allTypes = append(allTypes, mgtd.IntAliasSliceAliasTypeDesc)
	allTypes = append(allTypes, mgtd.FloatPtrSliceAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint64PtrAliasSliceAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint64AliasTypeDesc)
	allTypes = append(allTypes, mgtd.BoolSliceSliceAliasTypeDesc)
	allTypes = append(allTypes, mgtd.ByteArraySliceAliasTypeDesc)
	allTypes = append(allTypes, mgtd.FloatArrayPtrSliceAliasTypeDesc)
	allTypes = append(allTypes, mgtd.IntStrMapSliceAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint32Int32MapPtrSliceAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint32Int32MapPtrSliceAliasTypeDesc)
	allTypes = append(allTypes, mgtd.SimpleStructTypeDesc)
	allTypes = append(allTypes, mgtd.StructTypeSliceAliasTypeDesc)
	allTypes = append(allTypes, mgtd.StructTypePtrSliceAliasTypeDesc)
	allTypes = append(allTypes, mgtd.StringAliasTypeDesc)
	allTypes = append(allTypes, mgtd.TrickySliceAliasTypeDesc)
	allTypes = append(allTypes, mgtd.ValidUintSliceAliasTypeDesc)

	return generate(allTypes, "musgen", unsafe)
}

func generateMapAlias(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, mgtd.StrIntMapAliasTypeDesc)
	allTypes = append(allTypes, mgtd.StrPtrIntPtrMapAliasTypeDesc)
	allTypes = append(allTypes, mgtd.StringAliasTypeDesc)
	allTypes = append(allTypes, mgtd.IntAliasTypeDesc)
	allTypes = append(allTypes, mgtd.StrAliasIntAliasMapAliasTypeDesc)
	allTypes = append(allTypes, mgtd.StrAliasPtrIntAliasPtrMapAliasTypeDesc)
	allTypes = append(allTypes, mgtd.BoolInt16SliceMapAliasTypeDesc)
	allTypes = append(allTypes, mgtd.ByteUint16SlicePtrMapAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Int32Float64ArrayMapAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Float32Uint32ArrayPtrMapAliasTypeDesc)
	allTypes = append(allTypes, mgtd.FloatByteBoolMapMapAliasTypeDesc)
	allTypes = append(allTypes, mgtd.UintIntStringMapPtrMapAliasTypeDesc)
	allTypes = append(allTypes, mgtd.SimpleStructTypeDesc)
	allTypes = append(allTypes, mgtd.StructTypeStructTypeMapAliasTypeDesc)
	allTypes = append(allTypes, mgtd.StructTypePtrStructTypePtrMapAliasTypeDesc)
	allTypes = append(allTypes, mgtd.TrickyMapAliasTypeDesc)
	allTypes = append(allTypes, mgtd.ValidStringIntMapAliasTypeDesc)

	return generate(allTypes, "musgen", unsafe)
}

func generateStructType(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, mgtd.IntAliasTypeDesc)
	allTypes = append(allTypes, mgtd.StringAliasTypeDesc)
	allTypes = append(allTypes, mgtd.SimpleStructTypeDesc)
	allTypes = append(allTypes, mgtd.StructTypeDesc)
	allTypes = append(allTypes, mgtd.ValidStructTypeDesc)
	allTypes = append(allTypes, mgtd.FieldlessStructTypeDesc)

	return generate(allTypes, "musgen", unsafe)
}

func generateMusGo(unsafe bool) error {
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filepath.Dir(filename))
	musGo, err := musgo.New()
	if err != nil {
		return err
	}
	var v testdata.MyMap
	conf := musgo.NewAliasConf()
	conf.T = reflect.TypeOf(v)
	conf.Unsafe = unsafe
	conf.Validator = "ValidateMyMap"
	conf.MaxLength = 3
	conf.ElemValidator = "BiggerThenTen"
	conf.KeyValidator = "NotHello"
	conf.Path = dir
	conf.Filename = "xxxMyMap.musgen.go"
	err = musGo.GenerateAliasAs(conf)
	if err != nil {
		return err
	}
	return nil
}

func generateRaw(unsafe bool) error {
	var allTypes []musgen.TypeDesc
	allTypes = append(allTypes, mgtd.Uint64RawAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint32RawAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint16RawAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint8RawAliasTypeDesc)
	allTypes = append(allTypes, mgtd.UintRawAliasTypeDesc)

	allTypes = append(allTypes, mgtd.Int64RawAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Int32RawAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Int16RawAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Int8RawAliasTypeDesc)
	allTypes = append(allTypes, mgtd.IntRawAliasTypeDesc)

	allTypes = append(allTypes, mgtd.Float64RawAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Float32RawAliasTypeDesc)
	allTypes = append(allTypes, mgtd.IntRawArrayAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Uint16Int32RawMapAliasTypeDesc)
	allTypes = append(allTypes, mgtd.Float64RawPtrPtrPtrAliasSliceAliasTypeDesc)
	allTypes = append(allTypes, mgtd.RawStructTypeDesc)
	allTypes = append(allTypes, mgtd.ValidInt32RawTypeDesc)

	return generate(allTypes, "musgen", unsafe)
}

func generate(tds []musgen.TypeDesc, folder string, unsafe bool) error {
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filepath.Dir(filename))
	var bs []byte
	musGen, err := musgen.New()
	if err != nil {
		return err
	}
	for _, td := range tds {
		td.Unsafe = unsafe
		filename = makeFileName(td)
		bs, err = musGen.Generate(td, musgen.GoLang)
		if err != nil {
			return err
		}
		bs, err = imports.Process("", bs, nil)
		if err != nil {
			return err
		}
		if folder == "" {
			err = ioutil.WriteFile(dir+"/"+filename, bs, 0644)
		} else {
			err = ioutil.WriteFile(dir+"/"+folder+"/"+filename, bs, 0644)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func makeFileName(td musgen.TypeDesc) string {
	return "xxx" + td.Name + ".musgen.go"
}
