// Code generated by musgen. DO NOT EDIT.

package musgo

import "github.com/ymz-ncnk/musgo/errs"

// Marshal fills buf with the MUS encoding of v.
func (v IntAlias) Marshal(buf []byte) int {
	i := 0
	{
		uv := uint64(v)
		if v < 0 {
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
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *IntAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
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
		(*v) = IntAlias(uv)
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v IntAlias) Size() int {
	size := 0
	{
		uv := uint64(v<<1) ^ uint64(v>>63)
		{
			for uv >= 0x80 {
				uv >>= 7
				size++
			}
			size++
		}
	}
	return size
}
