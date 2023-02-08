package tdesc_builder

import (
	"errors"
	"fmt"
)

// ErrNegativeMaxLength happens on parsing a tag with negative MaxLength.
var ErrNegativeMaxLength = errors.New("negative MaxLength")

// -----------------------------------------------------------------------------
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
	return fmt.Sprintf(`'%v' field has invalid tag, it should be: `+
		`"-" or "Validator#Raw,MaxLength,ElemValidator#Raw,KeyValidator#Raw"`,
		err.fieldName)
}
