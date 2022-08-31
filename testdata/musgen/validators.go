package musgen

import "errors"

var ErrBiggerThanTen = errors.New("bigger then 10")

func BiggerThanTenUint64(n uint64) error {
	if n > 10 {
		return ErrBiggerThanTen
	}
	return nil
}

func BiggerThanTenInt8(n int8) error {
	if n > 10 {
		return ErrBiggerThanTen
	}
	return nil
}

var ErrNotEmptyString = errors.New("string is not empty")

var ErrNegative = errors.New("negative")

func NotEmptyString(str string) error {
	if str != "" {
		return ErrNotEmptyString
	}
	return nil
}

func BiggerThanTenByte(n byte) error {
	if n > 10 {
		return ErrBiggerThanTen
	}
	return nil
}

func BiggerThanTenUint16(n uint16) error {
	if n > 10 {
		return ErrBiggerThanTen
	}
	return nil
}

func BiggerThanTenInt16Ptr(n *int16) error {
	if *n > 10 {
		return ErrBiggerThanTen
	}
	return nil
}

var ErrPositiveBool = errors.New("bool is positive")

func PositiveBool(b bool) error {
	if b {
		return ErrPositiveBool
	}
	return nil
}

var ErrSliceSumBiggerThanTen = errors.New("slice sum bigger than ten")

func ValidUintSliceAliasSumBiggerThanTen(s *ValidUintSliceAlias) error {
	return UintSliceSumBiggerThanTen([]uint(*s))
}

func UintSliceSumBiggerThanTen(s []uint) error {
	var sum uint
	for _, n := range s {
		sum += n
	}
	if sum > 10 {
		return ErrSliceSumBiggerThanTen
	}
	return nil
}

func BiggerThanTenUint(n uint) error {
	if n > 10 {
		return ErrBiggerThanTen
	}
	return nil
}

func Uint16SlicePtrSumBiggerThanTen(s *[]uint16) error {
	var sum uint16
	for _, n := range *s {
		sum += n
	}
	if sum > 10 {
		return ErrSliceSumBiggerThanTen
	}
	return nil
}

var ErrArraySumBiggerThanTen = errors.New("array sum is bigger than ten")

func ValidIntArrayAliasSumBiggerThanTen(s *ValidIntArrayAlias) error {
	return IntArraySumBiggerThanTen([2]int(*s))
}

func IntArraySumBiggerThanTen(s [2]int) error {
	var sum int
	for _, n := range s {
		sum += n
	}
	if sum > 10 {
		return ErrArraySumBiggerThanTen
	}
	return nil
}

func BiggerThanTenInt(n int) error {
	if n > 10 {
		return ErrBiggerThanTen
	}
	return nil
}

func Int32ArrayPtrSumBiggerThanTen(a *[2]int32) error {
	var sum int32
	for _, n := range *a {
		sum += n
	}
	if sum > 10 {
		return ErrArraySumBiggerThanTen
	}
	return nil
}

func BiggerThanTenInt32(n int32) error {
	if n > 10 {
		return ErrBiggerThanTen
	}
	return nil
}

func PositiveValidInt32AliasRaw(n *ValidInt32RawAlias) error {
	if *n < 0 {
		return ErrNegative
	}
	return nil
}

var ErrMapSumBiggerThanTen = errors.New("map sum is bigger than 10")

func ValidStringIntMapAliasSumBiggerThanTen(m *ValidStringIntMapAlias) error {
	return MapSumBiggerThanTen(map[string]int(*m))
}

func MapSumBiggerThanTen(m map[string]int) error {
	var sum int
	for _, v := range m {
		sum += v
	}
	if sum > 10 {
		return ErrMapSumBiggerThanTen
	}
	return nil
}

func MapPtrSumBiggerThanTen(m *map[string]int) error {
	var sum int
	for _, v := range *m {
		sum += v
	}
	if sum > 10 {
		return ErrMapSumBiggerThanTen
	}
	return nil
}

var ErrStrIsHello = errors.New("string is hello")

func StrIsHello(str string) error {
	if str == "hello" {
		return ErrStrIsHello
	}
	return nil
}

var ErrSimpleStructType = errors.New("invalid SimpleStructType")

func ValidSimpleStructType(s SimpleStructType) error {
	if s.Int > 10 {
		return ErrSimpleStructType
	}
	return nil
}

func ValidSimpleStructPtrType(s *SimpleStructType) error {
	if s.Int > 10 {
		return ErrSimpleStructType
	}
	return nil
}
