package testdata

import "errors"

var ErrElementsSumBiggerThenTen = errors.New("elements sum bigger then ten")
var ErrBiggerThenTen = errors.New("bigger then ten")
var ErrHelloString = errors.New("hello string")

type MyMap map[string]int

func ValidateMyMap(sl MyMap) error {
	sum := 0
	for _, el := range sl {
		sum += el
	}
	if sum > 10 {
		return ErrElementsSumBiggerThenTen
	}
	return nil
}

func BiggerThenTen(n int) error {
	if n > 10 {
		return ErrBiggerThenTen
	}
	return nil
}

func NotHello(str string) error {
	if str == "hello" {
		return ErrHelloString
	}
	return nil
}
