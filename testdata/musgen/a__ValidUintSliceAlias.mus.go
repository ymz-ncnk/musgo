// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/musgo/errs"

// Marshal fills buf with the MUS encoding of v.
func (v ValidUintSliceAlias) Marshal(buf []byte) int {
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
				for el >= 0x80 {
					buf[i] = byte(el) | 0x80
					el >>= 7
					i++
				}
				buf[i] = byte(el)
				i++
			}
		}
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *ValidUintSliceAlias) Unmarshal(buf []byte) (int, error) {
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
		if length > 3 {
			err = errs.ErrMaxLengthExceeded
		} else {
			(*v) = make([]uint, length)
			for j := 0; j < length; j++ {
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
							(*v)[j] = (*v)[j] | uint(b)<<shift
							done = true
							i += l + 1
							break
						}
						(*v)[j] = (*v)[j] | uint(b&0x7F)<<shift
						shift += 7
					}
					if !done {
						return i, errs.ErrSmallBuf
					}
				}
				if err == nil {
					err = BiggerThanTenUint((*v)[j])
				}
				if err != nil {
					err = errs.NewSliceError(j, err)
					break
				}
			}
		}
	}
	if err == nil {
		err = ValidUintSliceAliasSumBiggerThanTen(v)
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v ValidUintSliceAlias) Size() int {
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
				for el >= 0x80 {
					el >>= 7
					size++
				}
				size++
			}
		}
	}
	return size
}
