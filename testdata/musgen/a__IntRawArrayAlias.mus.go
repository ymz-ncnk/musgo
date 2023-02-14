// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/musgo/v2/errs"

// Marshal fills buf with the MUS encoding of v.
func (v IntRawArrayAlias) Marshal(buf []byte) int {
	i := 0
	{
		for _, item := range v {
			{
				{
					buf[i] = byte(item)
					i++
					buf[i] = byte(item >> 8)
					i++
					buf[i] = byte(item >> 16)
					i++
					buf[i] = byte(item >> 24)
					i++
					buf[i] = byte(item >> 32)
					i++
					buf[i] = byte(item >> 40)
					i++
					buf[i] = byte(item >> 48)
					i++
					buf[i] = byte(item >> 56)
					i++
				}
			}
		}
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *IntRawArrayAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		for j := 0; j < 3; j++ {
			{
				{
					if len(buf) < 8 {
						return i, errs.ErrSmallBuf
					}
					(*v)[j] = int(buf[i])
					i++
					(*v)[j] |= int(buf[i]) << 8
					i++
					(*v)[j] |= int(buf[i]) << 16
					i++
					(*v)[j] |= int(buf[i]) << 24
					i++
					(*v)[j] |= int(buf[i]) << 32
					i++
					(*v)[j] |= int(buf[i]) << 40
					i++
					(*v)[j] |= int(buf[i]) << 48
					i++
					(*v)[j] |= int(buf[i]) << 56
					i++
				}
			}
			if err != nil {
				err = errs.NewArrayError(j, err)
				break
			}
		}
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v IntRawArrayAlias) Size() int {
	size := 0
	{
		for _, item := range v {
			{
				{
					_ = item
					size += 8
				}
			}
		}
	}
	return size
}
