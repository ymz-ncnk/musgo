// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/muserrs"

// Marshal fills buf with the MUS encoding of v.
func (v Uint16Alias) Marshal(buf []byte) int {
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
func (v *Uint16Alias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		if i > len(buf)-1 {
			return i, muserrs.ErrSmallBuf
		}
		shift := 0
		done := false
		for l, b := range buf[i:] {
			if l == 2 && b > 3 {
				return i, muserrs.ErrOverflow
			}
			if b < 0x80 {
				(*v) = (*v) | Uint16Alias(b)<<shift
				done = true
				i += l + 1
				break
			}
			(*v) = (*v) | Uint16Alias(b&0x7F)<<shift
			shift += 7
		}
		if !done {
			return i, muserrs.ErrSmallBuf
		}
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v Uint16Alias) Size() int {
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
