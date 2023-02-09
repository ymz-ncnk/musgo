// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/musgo/errs"

// Marshal fills buf with the MUS encoding of v.
func (v ValidInt32RawAlias) Marshal(buf []byte) int {
	i := 0
	{
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
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *ValidInt32RawAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		{
			if len(buf) < 4 {
				return i, errs.ErrSmallBuf
			}
			(*v) = ValidInt32RawAlias(buf[i])
			i++
			(*v) |= ValidInt32RawAlias(buf[i]) << 8
			i++
			(*v) |= ValidInt32RawAlias(buf[i]) << 16
			i++
			(*v) |= ValidInt32RawAlias(buf[i]) << 24
			i++
		}
		err = PositiveValidInt32AliasRaw(v)
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v ValidInt32RawAlias) Size() int {
	size := 0
	{
		{
			_ = v
			size += 4
		}
	}
	return size
}