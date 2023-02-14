// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/muserrs"

// Marshal fills buf with the MUS encoding of v.
func (v BoolSliceSliceAlias) Marshal(buf []byte) int {
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
		for _, el := range v {
			{
				length := len(el)
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
				for _, el := range el {
					{
						if el {
							buf[i] = 0x01
						} else {
							buf[i] = 0x00
						}
						i++
					}
				}
			}
		}
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *BoolSliceSliceAlias) Unmarshal(buf []byte) (int, error) {
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
		(*v) = make([][]bool, length)
		for j := 0; j < length; j++ {
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
				(*v)[j] = make([]bool, length)
				for jj := 0; jj < length; jj++ {
					{
						if i > len(buf)-1 {
							return i, muserrs.ErrSmallBuf
						}
						if buf[i] == 0x01 {
							(*v)[j][jj] = true
							i++
						} else if buf[i] == 0x00 {
							(*v)[j][jj] = false
							i++
						} else {
							err = muserrs.ErrWrongByte
						}
					}
					if err != nil {
						err = muserrs.NewSliceError(jj, err)
						break
					}
				}
			}
			if err != nil {
				err = muserrs.NewSliceError(j, err)
				break
			}
		}
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v BoolSliceSliceAlias) Size() int {
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
		for _, el := range v {
			{
				length := len(el)
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
				for _, el := range el {
					{
						_ = el
						size++
					}
				}
			}
		}
	}
	return size
}
