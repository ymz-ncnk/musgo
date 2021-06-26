package errs

import (
	"errors"
	"fmt"
)

// ErrSmallBuf means that an Unmarshal requires a longer buffer than was
// provided.
var ErrSmallBuf = errors.New("buf is too small")

// ErrOverflow happens on Unmarshal when bytes number limit of the type was
// exceeded.
var ErrOverflow = errors.New("overflow")

// ErrNegativeLength happens on Unmarshal when some data was encoded with
// length and value, and length is negative.
var ErrNegativeLength = errors.New("negative length")

// ErrWrongByte happens on Unmarshal when unexpected byte was caught.
var ErrWrongByte = errors.New("wrong byte")

// ErrMaxLengthExceeded is a MaxLength validator error.
var ErrMaxLengthExceeded = errors.New("max length exceeded")

// NewArrayError returns a new ArrayError.
func NewArrayError(i int, cause error) error {
	return ArrayError{i, cause}
}

// ArrayError occurs while Unmarshalling an array, if some element is not valid.
type ArrayError struct {
	i     int
	cause error
}

func (err ArrayError) Error() string {
	return fmt.Sprintf("%v element failed, cause: %v",
		err.i, err.cause)
}

// Index returns an index of the not valid element.
func (err ArrayError) Index() int {
	return err.i
}

// Cause returns a cause of the error.
func (err ArrayError) Cause() error {
	return err.cause
}

// NewSliceError returns a new SliceError.
func NewSliceError(i int, cause error) error {
	return SliceError{i, cause}
}

// SliceError occurs while Unmarshalling a slice, if some element is not valid.
type SliceError = ArrayError

// NewMapKeyError returns a new MapKeyError.
func NewMapKeyError(key interface{}, cause error) error {
	return &MapKeyError{key, cause}
}

// MapKeyError occurs while Unmarshalling one of the map's key, if it's not
// valid.
type MapKeyError struct {
	key   interface{}
	cause error
}

// Key returns a not valid key.
func (err *MapKeyError) Key() interface{} {
	return err.key
}

// Cause returns a cause of the error.
func (err *MapKeyError) Cause() error {
	return err.cause
}

func (err *MapKeyError) Error() string {
	return fmt.Sprintf("%v key failed, cause: %v", err.key, err.cause)
}

// NewMapValueError occurs while Unmarshalling one of the map's value, if it's
// not valid.
func NewMapValueError(key interface{}, value interface{}, cause error) error {
	return &MapValueError{key, value, cause}
}

// MapValueError occurs when Unmarshalling a map value, if value is not valid.
type MapValueError struct {
	key   interface{}
	value interface{}
	cause error
}

// Key returns a key of the not valid value.
func (err *MapValueError) Key() interface{} {
	return err.key
}

// Value returns a not valid value.
func (err *MapValueError) Value() interface{} {
	return err.value
}

// Cause returns a cause of the error.
func (err *MapValueError) Cause() error {
	return err.cause
}

func (err *MapValueError) Error() string {
	return fmt.Sprintf("%v value failed, cause: %v", err.value, err.cause)
}

// NewFieldError returns a new FieldError.
func NewFieldError(fieldName string, cause error) error {
	return FieldError{fieldName, cause}
}

// FieldError occurs while Unmarshalling a struct field, if field is not valid.
type FieldError struct {
	fieldName string
	cause     error
}

func (err FieldError) Error() string {
	return fmt.Sprintf("%v field failed, cause: %v", err.fieldName, err.cause)
}

// FieldName returns a field name of the not valid field.
func (err FieldError) FieldName() string {
	return err.fieldName
}

// Cause returns a cause of the error.
func (err FieldError) Cause() error {
	return err.cause
}
