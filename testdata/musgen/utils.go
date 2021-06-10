package musgen

import (
	"errors"
	"fmt"

	"github.com/ymz-ncnk/musgo/errs"
)

var InvalidSizeErrMsg string = "%v type has invalid size\n"
var MuErrMsg string = "MarshalMUS/UnmarshalMUS of %v type failed\n"

type InitValType interface {
	SizeMUS() int
	MarshalMUS([]byte) int
}

type ZeroValType interface {
	UnmarshalMUS([]byte) (int, error)
}

type GenError struct {
	methodName string
	cause      error
}

func (err *GenError) Error() string {
	return fmt.Sprintf("%v failed, cause: %v", err.methodName, err.cause)
}

func (err *GenError) Cause() error {
	return err.cause
}

func ExecGeneratedCode(initVal InitValType, zeroVal ZeroValType,
	typeName string) error {
	var err error
	size := initVal.SizeMUS()
	bs := make([]byte, size)
	i := 0
	i = initVal.MarshalMUS(bs)
	if i != len(bs) {
		return &GenError{"Wrong length",
			fmt.Errorf("required length %v, got %v", len(bs), i)}
	}
	_, err = zeroVal.UnmarshalMUS(bs)
	if err != nil {
		return &GenError{"UnmarshalMUS", err}
	}
	return nil
}

func TestBufferEnds(val ZeroValType, buf []byte) error {
	_, err := val.UnmarshalMUS(buf)
	if err != errs.ErrSmallBuf {
		return errors.New("small buf is ok")
	}
	return nil
}
