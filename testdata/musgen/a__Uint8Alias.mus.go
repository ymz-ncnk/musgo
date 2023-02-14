// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/musgo/v2/errs"

// Marshal fills buf with the MUS encoding of v.
func (v Uint8Alias) Marshal(buf []byte) int {
	i := 0
	{
		buf[i] = byte(v)
		i++
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *Uint8Alias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		if i > len(buf)-1 {
			return i, errs.ErrSmallBuf
		}
		(*v) = Uint8Alias(buf[i])
		i++
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v Uint8Alias) Size() int {
	size := 0
	{
		_ = v
		size++
	}
	return size
}
