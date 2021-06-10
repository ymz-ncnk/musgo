//go:generate go run testdata/make/musable.go -musgo $ARG
package musgo

import (
	"testing"

	"github.com/ymz-ncnk/musgo/errs"
	"github.com/ymz-ncnk/musgo/testdata"
)

func TestMapAliasValidation(t *testing.T) {
	var m testdata.MyMap = testdata.MyMap{"one": 5, "two": 8}
	buf := make([]byte, m.SizeMUS())
	m.MarshalMUS(buf)
	var am testdata.MyMap
	_, err := am.UnmarshalMUS(buf)
	if err == nil {
		t.Error(err)
	}
	if err != testdata.ErrElementsSumBiggerThenTen {
		t.Error("wrong error")
	}
}

func TestMapAliasMaxLengthValidation(t *testing.T) {
	var m testdata.MyMap = testdata.MyMap{"one": 5, "two": 0, "tree": 2,
		"four": 0}
	buf := make([]byte, m.SizeMUS())
	m.MarshalMUS(buf)
	var am testdata.MyMap
	_, err := am.UnmarshalMUS(buf)
	if err == nil {
		t.Error(err)
	}
	if err != errs.ErrMaxLengthExceeded {
		t.Error("wrong error")
	}
}

func TestMapAliasElemValidation(t *testing.T) {
	var m testdata.MyMap = testdata.MyMap{"one": 5, "two": 11}
	buf := make([]byte, m.SizeMUS())
	m.MarshalMUS(buf)
	var am testdata.MyMap
	_, err := am.UnmarshalMUS(buf)
	if err == nil {
		t.Error(err)
	}
	mapValueErr, ok := err.(*errs.MapValueError)
	if !ok {
		t.Error("wrong error")
	}
	if mapValueErr.Key() != "two" {
		t.Error("wrong key")
	}
	if mapValueErr.Value() != 11 {
		t.Error("wrong value")
	}
	if mapValueErr.Cause() != testdata.ErrBiggerThenTen {
		t.Error("wrong cause")
	}
}

func TestMapAliasKeyValidation(t *testing.T) {
	var m testdata.MyMap = testdata.MyMap{"one": 5, "hello": 4}
	buf := make([]byte, m.SizeMUS())
	m.MarshalMUS(buf)
	var am testdata.MyMap
	_, err := am.UnmarshalMUS(buf)
	if err == nil {
		t.Error(err)
	}
	mapKeyErr, ok := err.(*errs.MapKeyError)
	if !ok {
		t.Error("wrong error")
	}
	if mapKeyErr.Key() != "hello" {
		t.Error("wrong key")
	}
	if mapKeyErr.Cause() != testdata.ErrHelloString {
		t.Error("wrong cause")
	}
}
