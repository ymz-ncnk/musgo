// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/musgo/errs"

// Marshal fills buf with the MUS encoding of v.
func (v StrZeroLengthArrayAlias) Marshal(buf []byte) int {
	i := 0
	{
		for _, item := range v {
			{
				length := len(item)
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
				if len(buf[i:]) < length {
					panic(errs.ErrSmallBuf)
				}
				i += copy(buf[i:], item)
			}
		}
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *StrZeroLengthArrayAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		for j := 0; j < 0; j++ {
			{
				var length int
				{
					var uv uint64
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
								uv = uv | uint64(b)<<shift
								done = true
								i += l + 1
								break
							}
							uv = uv | uint64(b&0x7F)<<shift
							shift += 7
						}
						if !done {
							return i, errs.ErrSmallBuf
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
					return i, errs.ErrNegativeLength
				}
				if len(buf) < i+length {
					return i, errs.ErrSmallBuf
				}
				(*v)[j] = string(buf[i : i+length])
				i += length
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
func (v StrZeroLengthArrayAlias) Size() int {
	size := 0
	{
		for _, item := range v {
			{
				length := len(item)
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
				size += len(item)
			}
		}
	}
	return size
}
