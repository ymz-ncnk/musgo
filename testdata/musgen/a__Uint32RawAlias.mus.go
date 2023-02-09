// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/musgo/errs"

// Marshal fills buf with the MUS encoding of v.
func (v Uint32RawAlias) Marshal(buf []byte) int {
	i := 0
	{
		buf[i] = byte(v)
		i++
		buf[i] = byte(v >> 8)
		i++
		buf[i] = byte(v >> 16)
		i++
		buf[i] = byte(v >> 24)
		i++
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *Uint32RawAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		if len(buf) < 4 {
			return i, errs.ErrSmallBuf
		}
		(*v) = Uint32RawAlias(buf[i])
		i++
		(*v) |= Uint32RawAlias(buf[i]) << 8
		i++
		(*v) |= Uint32RawAlias(buf[i]) << 16
		i++
		(*v) |= Uint32RawAlias(buf[i]) << 24
		i++
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v Uint32RawAlias) Size() int {
	size := 0
	{
		_ = v
		size += 4
	}
	return size
}