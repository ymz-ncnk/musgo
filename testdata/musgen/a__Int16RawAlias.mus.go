// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/muserrs"

// Marshal fills buf with the MUS encoding of v.
func (v Int16RawAlias) Marshal(buf []byte) int {
	i := 0
	{
		{
			buf[i] = byte(v)
			i++
			buf[i] = byte(v >> 8)
			i++
		}
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *Int16RawAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		{
			if len(buf) < 2 {
				return i, muserrs.ErrSmallBuf
			}
			(*v) = Int16RawAlias(buf[i])
			i++
			(*v) |= Int16RawAlias(buf[i]) << 8
			i++
		}
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v Int16RawAlias) Size() int {
	size := 0
	{
		{
			_ = v
			size += 2
		}
	}
	return size
}
