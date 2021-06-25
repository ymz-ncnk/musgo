package testdata

import "errors"

// ErrElementsSumBiggerThenTen happens when sum of the elements is bigger then
// ten.
var ErrElementsSumBiggerThenTen = errors.New("elements sum bigger then ten")

// ErrBiggerThenTen happens when number is bigger then ten.
var ErrBiggerThenTen = errors.New("bigger then ten")

// ErrHelloString happens when string is a "hello" string.
var ErrHelloString = errors.New("hello string")

// MyMap is a simple map alias.
type MyMap map[string]int

// ValidateMyMap validates MyMap type.
func ValidateMyMap(sl *MyMap) error {
	sum := 0
	for _, el := range *sl {
		sum += el
	}
	if sum > 10 {
		return ErrElementsSumBiggerThenTen
	}
	return nil
}

// BiggerThenTen checks if n is bigger then ten.
func BiggerThenTen(n int) error {
	if n > 10 {
		return ErrBiggerThenTen
	}
	return nil
}

// NotHello checks if str is not "hello".
func NotHello(str string) error {
	if str == "hello" {
		return ErrHelloString
	}
	return nil
}
