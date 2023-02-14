// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/muserrs"

// Marshal fills buf with the MUS encoding of v.
func (v UintRawAlias) Marshal(buf []byte) int {
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
		buf[i] = byte(v >> 32)
		i++
		buf[i] = byte(v >> 40)
		i++
		buf[i] = byte(v >> 48)
		i++
		buf[i] = byte(v >> 56)
		i++
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *UintRawAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		if len(buf) < 8 {
			return i, muserrs.ErrSmallBuf
		}
		(*v) = UintRawAlias(buf[i])
		i++
		(*v) |= UintRawAlias(buf[i]) << 8
		i++
		(*v) |= UintRawAlias(buf[i]) << 16
		i++
		(*v) |= UintRawAlias(buf[i]) << 24
		i++
		(*v) |= UintRawAlias(buf[i]) << 32
		i++
		(*v) |= UintRawAlias(buf[i]) << 40
		i++
		(*v) |= UintRawAlias(buf[i]) << 48
		i++
		(*v) |= UintRawAlias(buf[i]) << 56
		i++
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v UintRawAlias) Size() int {
	size := 0
	{
		_ = v
		size += 8
	}
	return size
}
