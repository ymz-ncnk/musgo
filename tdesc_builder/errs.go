package tdesc_builder

import (
	"errors"
	"fmt"
)

var ErrNotAlias = errors.New("not an alias")
var ErrNotStruct = errors.New("not a struct")
var ErrNegativeMaxLength = errors.New("negative MaxLength")

// NewInvalidTagFormatError creates new InvalidTagFormatError.
func NewInvalidTagFormatError(fieldName string) *InvalidTagFormatError {
	return &InvalidTagFormatError{fieldName}
}

// InvalidTagFormatError happens if a tag has an invalid format.
type InvalidTagFormatError struct {
	fieldName string
}

func (err *InvalidTagFormatError) FieldName() string {
	return err.fieldName
}

func (err *InvalidTagFormatError) Error() string {
	return `%v field has invalid tag, it should be: ` +
		`"-" or "Validator#Raw,MaxLength,ElemValidator#Raw,KeyValidator#Raw"`
}

// -----------------------------------------------------------------------------
// NewInvalidTagPartFormatError creates new InvalidTagPartFormatError.
func NewInvalidTagPartFormatError(part string) *InvalidTagPartFormatError {
	return &InvalidTagPartFormatError{part}
}

// InvalidTagPartFormatError happens if an tag part has invalid format.
type InvalidTagPartFormatError struct {
	part string
}

func (err *InvalidTagPartFormatError) FieldName() string {
	return err.part
}

func (err *InvalidTagPartFormatError) Error() string {
	return fmt.Sprintf("invalid format of the %s", err.part)
}
