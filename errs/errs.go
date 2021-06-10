package errs

import (
	"errors"
	"fmt"
)

// ErrSmallBuf means that an Unmarshal required a longer buffer than was
// provided.
var ErrSmallBuf error = errors.New("buf is too small")

// ErrOverflow happens on Unmarshal when bytes number limit of the type was
// exceeded.
var ErrOverflow error = errors.New("overflow")

// ErrNegativeLength happens on Unmarshal when some data was encoded with
// length and value, and length is negative.
var ErrNegativeLength error = errors.New("negative length")

// ErrWrongByte happends on Unmarshal when unexpected byte was caught.
var ErrWrongByte error = errors.New("wrong byte")

// ErrMaxLengthExceeded is a MaxLength validator error.
var ErrMaxLengthExceeded error = errors.New("max length exceeded")

func NewSliceError(i int, cause error) error {
	return SliceError{i, cause}
}

func NewArrayError(i int, cause error) error {
	return ArrayError{i, cause}
}

// ArrayError occurs when Unmarshalling an array, if some element is not valid.
type ArrayError struct {
	i     int
	cause error
}

func (err ArrayError) Error() string {
	return fmt.Sprintf("%v element failed, cause: %v",
		err.i, err.cause)
}

func (err ArrayError) Index() int {
	return err.i
}

func (err ArrayError) Cause() error {
	return err.cause
}

// SliceError occurs when Unmarshalling a slice, if some element is not valid.
type SliceError = ArrayError

func NewMapKeyError(key interface{}, cause error) error {
	return &MapKeyError{key, cause}
}

// MapKeyError occurs when Unmarshalling a map key, if key is not valid.
type MapKeyError struct {
	key   interface{}
	cause error
}

func (err *MapKeyError) Key() interface{} {
	return err.key
}

func (err *MapKeyError) Cause() error {
	return err.cause
}

func (err *MapKeyError) Error() string {
	return fmt.Sprintf("%v key failed, cause: %v", err.key, err.cause)
}

func NewMapValueError(key interface{}, value interface{}, cause error) error {
	return &MapValueError{key, value, cause}
}

// MapValueError occurs when Unmarshalling a map value, if value is not valid.
type MapValueError struct {
	key   interface{}
	value interface{}
	cause error
}

func (err *MapValueError) Key() interface{} {
	return err.key
}

func (err *MapValueError) Value() interface{} {
	return err.value
}

func (err *MapValueError) Cause() error {
	return err.cause
}

func (err *MapValueError) Error() string {
	return fmt.Sprintf("%v value failed, cause: %v", err.value, err.cause)
}

func NewFieldError(fieldName string, cause error) error {
	return FieldError{fieldName, cause}
}

// FieldError occurs when Unmarshalling a struct field, if field is not valid.
type FieldError struct {
	fieldName string
	cause     error
}

func (err FieldError) Error() string {
	return fmt.Sprintf("%v field failed, cause: %v", err.fieldName, err.cause)
}

func (err FieldError) FieldName() string {
	return err.fieldName
}

func (err FieldError) Cause() error {
	return err.cause
}
