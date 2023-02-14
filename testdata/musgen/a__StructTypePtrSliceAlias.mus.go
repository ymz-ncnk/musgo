// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/musgo/v2/errs"

// Marshal fills buf with the MUS encoding of v.
func (v StructTypePtrSliceAlias) Marshal(buf []byte) int {
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
			if el == nil {
				buf[i] = 0
				i++
			} else {
				buf[i] = 1
				i++
				{
					si := (*el).Marshal(buf[i:])
					i += si
				}
			}
		}
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *StructTypePtrSliceAlias) Unmarshal(buf []byte) (int, error) {
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
		(*v) = make([]*SimpleStructType, length)
		for j := 0; j < length; j++ {
			(*v)[j] = new(SimpleStructType)
			if buf[i] == 0 {
				i++
				(*v)[j] = nil
			} else if buf[i] != 1 {
				i++
				return i, errs.ErrWrongByte
			} else {
				i++
				{
					var sv SimpleStructType
					si := 0
					si, err = sv.Unmarshal(buf[i:])
					if err == nil {
						(*(*v)[j]) = sv
						i += si
					}
				}
			}
			if err != nil {
				err = errs.NewSliceError(j, err)
				break
			}
		}
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v StructTypePtrSliceAlias) Size() int {
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
			size++
			if el != nil {
				{
					ss := (*el).Size()
					size += ss
				}
			}
		}
	}
	return size
}
