// Code generated by musgen. DO NOT EDIT.

package musgen

import (
	"math"

	"github.com/ymz-ncnk/serialization/musgo/errs"
)

// Marshal fills buf with the MUS encoding of v.
func (v Int32Float64ArrayMapAlias) Marshal(buf []byte) int {
	i := 0
	{
		length := len(v)
		{
			uv := uint64(length)
			if length < 0 {
				uv = ^(uv << 1)
			} else {
				uv = uv << 1
			}
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
		for ke, vl := range v {
			{
				uv := uint32(ke)
				if ke < 0 {
					uv = ^(uv << 1)
				} else {
					uv = uv << 1
				}
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
			{
				for _, item := range vl {
					{
						uv := math.Float64bits(float64(item))
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
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *Int32Float64ArrayMapAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		var length int
		{
			var uv uint64
			{
				if i > len(buf)-1 {
					return i, errs.ErrSmallBuf
				}
				shift := 0
				done := false
				for l, b := range buf[i:] {
					if l == 9 && b > 1 {
						return i, errs.ErrOverflow
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
					return i, errs.ErrSmallBuf
				}
			}
			if uv&1 == 1 {
				uv = ^(uv >> 1)
			} else {
				uv = uv >> 1
			}
			length = int(uv)
		}
		if length < 0 {
			return i, errs.ErrNegativeLength
		}
		(*v) = make(map[int32][2]float64)
		for ; length > 0; length-- {
			var kem int32
			var vlm [2]float64
			{
				var uv uint32
				{
					if i > len(buf)-1 {
						return i, errs.ErrSmallBuf
					}
					shift := 0
					done := false
					for l, b := range buf[i:] {
						if l == 4 && b > 15 {
							return i, errs.ErrOverflow
						}
						if b < 0x80 {
							uv = uv | uint32(b)<<shift
							done = true
							i += l + 1
							break
						}
						uv = uv | uint32(b&0x7F)<<shift
						shift += 7
					}
					if !done {
						return i, errs.ErrSmallBuf
					}
				}
				if uv&1 == 1 {
					uv = ^(uv >> 1)
				} else {
					uv = uv >> 1
				}
				kem = int32(uv)
			}
			if err != nil {
				err = errs.NewMapKeyError(kem, err)
				break
			}
			{
				for j := 0; j < 2; j++ {
					{
						var uv uint64
						{
							if i > len(buf)-1 {
								return i, errs.ErrSmallBuf
							}
							shift := 0
							done := false
							for l, b := range buf[i:] {
								if l == 9 && b > 1 {
									return i, errs.ErrOverflow
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
								return i, errs.ErrSmallBuf
							}
						}
						uv = (uv << 32) | (uv >> 32)
						uv = ((uv << 16) & 0xFFFF0000FFFF0000) | ((uv >> 16) & 0x0000FFFF0000FFFF)
						uv = ((uv << 8) & 0xFF00FF00FF00FF00) | ((uv >> 8) & 0x00FF00FF00FF00FF)
						vlm[j] = float64(math.Float64frombits(uv))
					}
					if err != nil {
						err = errs.NewArrayError(j, err)
						break
					}
				}
			}
			if err != nil {
				err = errs.NewMapValueError(kem, vlm, err)
				break
			}
			(*v)[kem] = vlm
		}
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v Int32Float64ArrayMapAlias) Size() int {
	size := 0
	{
		length := len(v)
		{
			uv := uint64(length<<1) ^ uint64(length>>63)
			{
				for uv >= 0x80 {
					uv >>= 7
					size++
				}
				size++
			}
		}
		for ke, vl := range v {
			{
				uv := uint32(ke<<1) ^ uint32(ke>>31)
				{
					for uv >= 0x80 {
						uv >>= 7
						size++
					}
					size++
				}
			}
			{
				for _, item := range vl {
					{
						uv := math.Float64bits(float64(item))
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
	}
	return size
}
