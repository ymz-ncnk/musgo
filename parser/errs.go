package parser

import (
	"fmt"
)

// // ErrNilType happens if tries to parse nil type.
// // var ErrNilType = errors.New("nil type")

// // -----------------------------------------------------------------------------
// // NewInvalidTagFormatError creates new InvalidTagFormatError.
// func NewInvalidTagFormatError(fieldName string) *InvalidTagFormatError {
// 	return &InvalidTagFormatError{fieldName}
// }

// // InvalidTagFormatError happens if a tag has an invalid format.
// type InvalidTagFormatError struct {
// 	fieldName string
// }

// func (err *InvalidTagFormatError) FieldName() string {
// 	return err.fieldName
// }

// func (err *InvalidTagFormatError) Error() string {
// 	return `%v field has invalid tag, it should be: ` +
// 		`"-" or "Validator#Raw,MaxLength,ElemValidator#Raw,KeyValidator#Raw"`
// }

// // -----------------------------------------------------------------------------
// // NewInvalidTagPartFormatError creates new InvalidTagPartFormatError.
// func NewInvalidTagPartFormatError(part string) *InvalidTagPartFormatError {
// 	return &InvalidTagPartFormatError{part}
// }

// // InvalidTagPartFormatError happens if an tag part has invalid format.
// type InvalidTagPartFormatError struct {
// 	part string
// }

// func (err *InvalidTagPartFormatError) FieldName() string {
// 	return err.part
// }

// func (err *InvalidTagPartFormatError) Error() string {
// 	return fmt.Sprintf("invalid format of the %s", err.part)
// }

// // -----------------------------------------------------------------------------
// // NewInvalidMaxLengthTagError creates new InvalidMaxLengthTagError.
// func NewInvalidMaxLengthTagError(fieldName string) *InvalidMaxLengthTagError {
// 	return &InvalidMaxLengthTagError{fieldName}
// }

// // InvalidMaxLengthTagError happens if a MaxLength value is invalid, for
// // example is negative.
// type InvalidMaxLengthTagError struct {
// 	fieldName string
// }

// func (err *InvalidMaxLengthTagError) FieldName() string {
// 	return err.fieldName
// }

// func (err *InvalidMaxLengthTagError) Error() string {
// 	return fmt.Sprintf("%v field has invalid tag, because of MaxLength value",
// 		err.fieldName)
// }

// // -----------------------------------------------------------------------------
// // NewUnsupportedMaxLengthTagError creates new UnsupportedMaxLengthTagError.
// func NewUnsupportedMaxLengthTagError(
// 	fieldName string) *UnsupportedMaxLengthTagError {
// 	return &UnsupportedMaxLengthTagError{fieldName}
// }

// // UnsupportedMaxLengthTagError happens if a ЬaxLength  мфдгуis set for the not
// // supported type.
// type UnsupportedMaxLengthTagError struct {
// 	fieldName string
// }

// func (err *UnsupportedMaxLengthTagError) FieldName() string {
// 	return err.fieldName
// }

// func (err *UnsupportedMaxLengthTagError) Error() string {
// 	return fmt.Sprintf("%v field has invalid tag, only strings, slices or maps "+
// 		"could have MaxLength", err.fieldName)
// }

// // -----------------------------------------------------------------------------
// // NewUnsupportedElemValidatorTagError creates new
// // UnsupportedElemValidatorTagError.
// func NewUnsupportedElemValidatorTagError(
// 	fieldName string) *UnsupportedElemValidatorTagError {
// 	return &UnsupportedElemValidatorTagError{fieldName}
// }

// // UnsupportedElemValidatorTagError happens if an ElemValidator is set for the
// // not supported type.
// type UnsupportedElemValidatorTagError struct {
// 	fieldName string
// }

// func (err *UnsupportedElemValidatorTagError) FieldName() string {
// 	return err.fieldName
// }

// func (err *UnsupportedElemValidatorTagError) Error() string {
// 	return fmt.Sprintf("%v field has invalid tag, only arrays, slices or maps "+
// 		"could have ElemValidator or ElemEncoding", err.fieldName)
// }

// // -----------------------------------------------------------------------------
// // NewUnsupportedKeyValidatorTagError creates new
// // UnsupportedKeyValidatorTagError.
// func NewUnsupportedKeyValidatorTagError(
// 	fieldName string) *UnsupportedKeyValidatorTagError {
// 	return &UnsupportedKeyValidatorTagError{fieldName}
// }

// // UnsupportedKeyValidatorTagError happens if an KeyValidator is set for the
// // not supported type.
// type UnsupportedKeyValidatorTagError struct {
// 	fieldName string
// }

// func (err *UnsupportedKeyValidatorTagError) FieldName() string {
// 	return err.fieldName
// }

// func (err *UnsupportedKeyValidatorTagError) Error() string {
// 	return fmt.Sprintf("%v field has invalid tag, only arrays, slices or maps "+
// 		"could have KeyValidator or KeyEncoding", err.fieldName)
// }

// // -----------------------------------------------------------------------------
// // NewArrayMaxLengthTagError creates new ArrayMaxLengthTagError.
// func NewArrayMaxLengthTagError(
// 	fieldName string) *ArrayMaxLengthTagError {
// 	return &ArrayMaxLengthTagError{fieldName}
// }

// // ArrayMaxLengthTagError happens if an KeyValidator is set for the
// // not supported type.
// type ArrayMaxLengthTagError struct {
// 	fieldName string
// }

// func (err *ArrayMaxLengthTagError) FieldName() string {
// 	return err.fieldName
// }

// func (err *ArrayMaxLengthTagError) Error() string {
// 	return fmt.Sprintf("%v field has invalid tag, maxLength is specified for "+
// 		"an array", err.fieldName)
// }

// -----------------------------------------------------------------------------
// NewUnsupportedTypeError creates new UnsupportedTypeError.
func NewUnsupportedTypeError(typeName string) *UnsupportedTypeError {
	return &UnsupportedTypeError{typeName}
}

// UnsupportedTypeError happens when the parser tries to parse the not
// supported type.
type UnsupportedTypeError struct {
	typeName string
}

// Type returns the not supported type.
func (err *UnsupportedTypeError) Type() string {
	return err.typeName
}

func (err *UnsupportedTypeError) Error() string {
	return fmt.Sprintf("%v type is not supported", err.Type())
}
