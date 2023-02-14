// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/musgo/v2/errs"

// Marshal fills buf with the MUS encoding of v.
func (v UintAlias) Marshal(buf []byte) int {
	i := 0
	{
		for v >= 0x80 {
			buf[i] = byte(v) | 0x80
			v >>= 7
			i++
		}
		buf[i] = byte(v)
		i++
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *UintAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
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
				(*v) = (*v) | UintAlias(b)<<shift
				done = true
				i += l + 1
				break
			}
			(*v) = (*v) | UintAlias(b&0x7F)<<shift
			shift += 7
		}
		if !done {
			return i, errs.ErrSmallBuf
		}
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v UintAlias) Size() int {
	size := 0
	{
		for v >= 0x80 {
			v >>= 7
			size++
		}
		size++
	}
	return size
}
