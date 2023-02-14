// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/muserrs"

// Marshal fills buf with the MUS encoding of v.
func (v BoolArrayArrayAlias) Marshal(buf []byte) int {
	i := 0
	{
		for _, item := range v {
			{
				for _, item := range item {
					{
						if item {
							buf[i] = 0x01
						} else {
							buf[i] = 0x00
						}
						i++
					}
				}
			}
		}
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *BoolArrayArrayAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		for j := 0; j < 3; j++ {
			{
				for jj := 0; jj < 3; jj++ {
					{
						if i > len(buf)-1 {
							return i, muserrs.ErrSmallBuf
						}
						if buf[i] == 0x01 {
							(*v)[j][jj] = true
							i++
						} else if buf[i] == 0x00 {
							(*v)[j][jj] = false
							i++
						} else {
							err = muserrs.ErrWrongByte
						}
					}
					if err != nil {
						err = muserrs.NewArrayError(jj, err)
						break
					}
				}
			}
			if err != nil {
				err = muserrs.NewArrayError(j, err)
				break
			}
		}
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v BoolArrayArrayAlias) Size() int {
	size := 0
	{
		for _, item := range v {
			{
				for _, item := range item {
					{
						_ = item
						size++
					}
				}
			}
		}
	}
	return size
}
