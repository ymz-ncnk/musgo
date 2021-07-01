package musgen

import (
	"errors"
	"fmt"

	"github.com/ymz-ncnk/musgo/errs"
)

var InvalidSizeErrMsg string = "%v type has invalid size\n"
var MuErrMsg string = "MarshalMUS/UnmarshalMUS of %v type failed\n"

type InitValType interface {
	Size() int
	Marshal([]byte) int
}

type ZeroValType interface {
	Unmarshal([]byte) (int, error)
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
	size := initVal.Size()
	bs := make([]byte, size)
	i := 0
	i = initVal.Marshal(bs)
	if i != len(bs) {
		return &GenError{"Wrong length",
			fmt.Errorf("required length %v, got %v", len(bs), i)}
	}
	_, err = zeroVal.Unmarshal(bs)
	if err != nil {
		return &GenError{"UnmarshalMUS", err}
	}
	return nil
}

func TestBufferEnds(val ZeroValType, buf []byte) error {
	_, err := val.Unmarshal(buf)
	if err != errs.ErrSmallBuf {
		return errors.New("small buf is ok")
	}
	return nil
}
