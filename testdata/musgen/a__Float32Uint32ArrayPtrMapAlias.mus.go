// Code generated by musgen. DO NOT EDIT.

package musgen

import (
	"math"

	"github.com/ymz-ncnk/muserrs"
)

// Marshal fills buf with the MUS encoding of v.
func (v Float32Uint32ArrayPtrMapAlias) Marshal(buf []byte) int {
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
				uv := math.Float32bits(float32(ke))
				uv = (uv << 16) | (uv >> 16)
				uv = ((uv << 8) & 0xFF00FF00) | ((uv >> 8) & 0x00FF00FF)
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
			if vl == nil {
				buf[i] = 0
				i++
			} else {
				buf[i] = 1
				i++
				{
					for _, item := range *vl {
						{
							for item >= 0x80 {
								buf[i] = byte(item) | 0x80
								item >>= 7
								i++
							}
							buf[i] = byte(item)
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
func (v *Float32Uint32ArrayPtrMapAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		var length int
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
			if uv&1 == 1 {
				uv = ^(uv >> 1)
			} else {
				uv = uv >> 1
			}
			length = int(uv)
		}
		if length < 0 {
			return i, muserrs.ErrNegativeLength
		}
		(*v) = make(map[float32]*[2]uint32)
		for ; length > 0; length-- {
			var kem float32
			vlm := new([2]uint32)
			{
				var uv uint32
				{
					if i > len(buf)-1 {
						return i, muserrs.ErrSmallBuf
					}
					shift := 0
					done := false
					for l, b := range buf[i:] {
						if l == 4 && b > 15 {
							return i, muserrs.ErrOverflow
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
						return i, muserrs.ErrSmallBuf
					}
				}
				uv = (uv << 16) | (uv >> 16)
				uv = ((uv << 8) & 0xFF00FF00) | ((uv >> 8) & 0x00FF00FF)
				kem = float32(math.Float32frombits(uv))
			}
			if err != nil {
				err = muserrs.NewMapKeyError(kem, err)
				break
			}
			if buf[i] == 0 {
				i++
				vlm = nil
			} else if buf[i] != 1 {
				i++
				return i, muserrs.ErrWrongByte
			} else {
				i++
				{
					for j := 0; j < 2; j++ {
						{
							if i > len(buf)-1 {
								return i, muserrs.ErrSmallBuf
							}
							shift := 0
							done := false
							for l, b := range buf[i:] {
								if l == 4 && b > 15 {
									return i, muserrs.ErrOverflow
								}
								if b < 0x80 {
									(*vlm)[j] = (*vlm)[j] | uint32(b)<<shift
									done = true
									i += l + 1
									break
								}
								(*vlm)[j] = (*vlm)[j] | uint32(b&0x7F)<<shift
								shift += 7
							}
							if !done {
								return i, muserrs.ErrSmallBuf
							}
						}
						if err != nil {
							err = muserrs.NewArrayError(j, err)
							break
						}
					}
				}
			}
			if err != nil {
				err = muserrs.NewMapValueError(kem, vlm, err)
				break
			}
			(*v)[kem] = vlm
		}
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v Float32Uint32ArrayPtrMapAlias) Size() int {
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
				uv := math.Float32bits(float32(ke))
				uv = (uv << 16) | (uv >> 16)
				uv = ((uv << 8) & 0xFF00FF00) | ((uv >> 8) & 0x00FF00FF)
				{
					for uv >= 0x80 {
						uv >>= 7
						size++
					}
					size++
				}
			}
			size++
			if vl != nil {
				{
					for _, item := range *vl {
						{
							for item >= 0x80 {
								item >>= 7
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
