package testdata

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/ymz-ncnk/musgo/v2/errs"
)

type MUSMarshalType interface {
	Size() int
	Marshal([]byte) int
}

type MUSUnmarshalType interface {
	Unmarshal([]byte) (int, error)
}

func TestGeneratedCode(val MUSMarshalType) (err error) {
	aval, err := ExecGeneratedCode(val)
	if err != nil {
		return err
	}
	result := reflect.ValueOf(aval).Elem().Interface().(MUSMarshalType)
	if !reflect.DeepEqual(val, result) {
		return fmt.Errorf("want - %v, actual - %v", val, result)
	}
	return nil
}

func ExecGeneratedCode(val MUSMarshalType) (aval MUSUnmarshalType, err error) {
	var (
		buf = make([]byte, val.Size())
		n   = val.Marshal(buf)
	)
	if n != len(buf) {
		return nil,
			fmt.Errorf("returned by Marshal num %v is not equal to buf size %v", n,
				len(buf))
	}
	aval = zeroValOf(val)
	_, err = aval.Unmarshal(buf)
	if err != nil {
		return nil, fmt.Errorf("unexpected Unmarshal error: %w", err)
	}
	return aval, nil
}

func TestBufferEnds(val MUSUnmarshalType, buf []byte) error {
	_, err := val.Unmarshal(buf)
	if err != errs.ErrSmallBuf {
		return errors.New("small buf is ok")
	}
	return nil
}

func TestFieldErr(err error, wantField string, t *testing.T) (cause error) {
	unmarshalErr := errors.Unwrap(err)
	if fieldErr, ok := unmarshalErr.(*errs.FieldError); ok {
		if fieldErr.FieldName() != wantField {
			t.Errorf("wrong error field name '%v'", fieldErr.FieldName())
		}
		cause = fieldErr.Cause()
	} else {
		t.Errorf("wrong error '%v'", unmarshalErr)
	}
	return
}

func TestFieldErrAndCause(err error, wantField string, wantCause error,
	t *testing.T) {
	cause := TestFieldErr(err, wantField, t)
	if cause != nil {
		if cause != wantCause {
			t.Errorf("wrong error cause '%v'", cause)
		}
	}
}

func TestSliceElemErr(err error, wantField string, wantElemIndex int,
	wantCause error,
	t *testing.T,
) {
	cause := TestFieldErr(err, wantField, t)
	if cause != nil {
		sliceErr, ok := cause.(*errs.SliceError)
		if !ok {
			t.Errorf("wrong error cause '%v'", cause)
		}
		if sliceErr.Index() != wantElemIndex {
			t.Errorf("wrong slice error index %v", sliceErr.Index())
		}
		if sliceErr.Cause() != wantCause {
			t.Errorf("wrong slice error cause '%v'", sliceErr.Cause())
		}
	}
}

func TestMapKeyErr(err error, wantField string, wantKey string,
	wantCause error,
	t *testing.T,
) {
	cause := TestFieldErr(err, wantField, t)
	if cause != nil {
		mapKeyErr, ok := cause.(*errs.MapKeyError)
		if !ok {
			t.Errorf("wrong error cause '%v'", cause)
		}
		if mapKeyErr.Key() != wantKey {
			t.Errorf("wrong slice error index %v", mapKeyErr.Key())
		}
		if mapKeyErr.Cause() != wantCause {
			t.Errorf("wrong slice error cause '%v'", mapKeyErr.Cause())
		}
	}
}

func TestMapValueErr(err error, wantField string, wantKey string, wantValue int,
	wantCause error,
	t *testing.T) {
	cause := TestFieldErr(err, wantField, t)
	if cause != nil {
		mapValueErr, ok := cause.(*errs.MapValueError)
		if !ok {
			t.Errorf("wrong error cause '%v'", cause)
		}
		if mapValueErr.Key() != wantKey {
			t.Errorf("wrong slice error index %v", mapValueErr.Key())
		}
		if mapValueErr.Value() != wantValue {
			t.Errorf("wrong slice error index %v", mapValueErr.Value())
		}
		if mapValueErr.Cause() != wantCause {
			t.Errorf("wrong slice error cause '%v'", mapValueErr.Cause())
		}
	}
}

func zeroValOf(val MUSMarshalType) MUSUnmarshalType {
	if reflect.TypeOf(val).Kind() == reflect.Ptr {
		return reflect.New(
			reflect.ValueOf(val).Elem().Type(),
		).Interface().(MUSUnmarshalType)
	} else {
		return reflect.New(reflect.TypeOf(val)).Interface().(MUSUnmarshalType)
	}
}
