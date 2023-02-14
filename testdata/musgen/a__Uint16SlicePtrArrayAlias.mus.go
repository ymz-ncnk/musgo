// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/muserrs"

// Marshal fills buf with the MUS encoding of v.
func (v Uint16SlicePtrArrayAlias) Marshal(buf []byte) int {
	i := 0
	{
		for _, item := range v {
			if item == nil {
				buf[i] = 0
				i++
			} else {
				buf[i] = 1
				i++
				{
					length := len((*item))
					{
						uv := uint64(length)
						if length < 0 {
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
					for _, el := range *item {
						{
							for el >= 0x80 {
								buf[i] = byte(el) | 0x80
								el >>= 7
								i++
							}
							buf[i] = byte(el)
							i++
						}
					}
				}
			}
		}
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *Uint16SlicePtrArrayAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		for j := 0; j < 3; j++ {
			(*v)[j] = new([]uint16)
			if buf[i] == 0 {
				i++
				(*v)[j] = nil
			} else if buf[i] != 1 {
				i++
				return i, muserrs.ErrWrongByte
			} else {
				i++
				{
					var length int
					{
						var uv uint64
						{
							if i > len(buf)-1 {
								return i, muserrs.ErrSmallBuf
							}
							shift := 0
							done := false
							for l, b := range buf[i:] {
								if l == 9 && b > 1 {
									return i, muserrs.ErrOverflow
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
								return i, muserrs.ErrSmallBuf
							}
						}
						if uv&1 == 1 {
							uv = ^(uv >> 1)
						} else {
							uv = uv >> 1
						}
						length = int(uv)
					}
					if length < 0 {
						return i, muserrs.ErrNegativeLength
					}
					(*(*v)[j]) = make([]uint16, length)
					for jj := 0; jj < length; jj++ {
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
									(*(*v)[j])[jj] = (*(*v)[j])[jj] | uint16(b)<<shift
									done = true
									i += l + 1
									break
								}
								(*(*v)[j])[jj] = (*(*v)[j])[jj] | uint16(b&0x7F)<<shift
								shift += 7
							}
							if !done {
								return i, muserrs.ErrSmallBuf
							}
						}
						if err != nil {
							err = muserrs.NewSliceError(jj, err)
							break
						}
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
func (v Uint16SlicePtrArrayAlias) Size() int {
	size := 0
	{
		for _, item := range v {
			size++
			if item != nil {
				{
					length := len((*item))
					{
						uv := uint64(length<<1) ^ uint64(length>>63)
						{
							for uv >= 0x80 {
								uv >>= 7
								size++
							}
							size++
						}
					}
					for _, el := range *item {
						{
							for el >= 0x80 {
								el >>= 7
								size++
							}
							size++
						}
					}
				}
			}
		}
	}
	return size
}
