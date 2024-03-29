// Code generated by musgen. DO NOT EDIT.

package musgen

import (
	"math"

	"github.com/ymz-ncnk/muserrs"
)

// Marshal fills buf with the MUS encoding of v.
func (v FloatPtrArrayAlias) Marshal(buf []byte) int {
	i := 0
	{
		for _, item := range v {
			if item == nil {
				buf[i] = 0
				i++
			} else {
				buf[i] = 1
				i++
				{
					uv := math.Float64bits(float64((*item)))
					uv = (uv << 32) | (uv >> 32)
					uv = ((uv << 16) & 0xFFFF0000FFFF0000) | ((uv >> 16) & 0x0000FFFF0000FFFF)
					uv = ((uv << 8) & 0xFF00FF00FF00FF00) | ((uv >> 8) & 0x00FF00FF00FF00FF)
					{
						for uv >= 0x80 {
							buf[i] = byte(uv) | 0x80
							uv >>= 7
							i++
						}
						buf[i] = byte(uv)
						i++
					}
				}
			}
		}
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *FloatPtrArrayAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		for j := 0; j < 3; j++ {
			(*v)[j] = new(float64)
			if buf[i] == 0 {
				i++
				(*v)[j] = nil
			} else if buf[i] != 1 {
				i++
				return i, muserrs.ErrWrongByte
			} else {
				i++
				{
					var uv uint64
					{
						if i > len(buf)-1 {
							return i, muserrs.ErrSmallBuf
						}
						shift := 0
						done := false
						for l, b := range buf[i:] {
							if l == 9 && b > 1 {
								return i, muserrs.ErrOverflow
							}
							if b < 0x80 {
								uv = uv | uint64(b)<<shift
								done = true
								i += l + 1
								break
							}
							uv = uv | uint64(b&0x7F)<<shift
							shift += 7
						}
						if !done {
							return i, muserrs.ErrSmallBuf
						}
					}
					uv = (uv << 32) | (uv >> 32)
					uv = ((uv << 16) & 0xFFFF0000FFFF0000) | ((uv >> 16) & 0x0000FFFF0000FFFF)
					uv = ((uv << 8) & 0xFF00FF00FF00FF00) | ((uv >> 8) & 0x00FF00FF00FF00FF)
					(*(*v)[j]) = float64(math.Float64frombits(uv))
				}
			}
			if err != nil {
				err = muserrs.NewArrayError(j, err)
				break
			}
		}
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v FloatPtrArrayAlias) Size() int {
	size := 0
	{
		for _, item := range v {
			size++
			if item != nil {
				{
					uv := math.Float64bits(float64((*item)))
					uv = (uv << 32) | (uv >> 32)
					uv = ((uv << 16) & 0xFFFF0000FFFF0000) | ((uv >> 16) & 0x0000FFFF0000FFFF)
					uv = ((uv << 8) & 0xFF00FF00FF00FF00) | ((uv >> 8) & 0x00FF00FF00FF00FF)
					{
						for uv >= 0x80 {
							uv >>= 7
							size++
						}
						size++
					}
				}
			}
		}
	}
	return size
}
