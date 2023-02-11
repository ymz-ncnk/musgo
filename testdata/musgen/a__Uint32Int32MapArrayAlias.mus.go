// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/musgo/errs"

// Marshal fills buf with the MUS encoding of v.
func (v Uint32Int32MapArrayAlias) Marshal(buf []byte) int {
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
					for ke, vl := range *item {
						{
							for ke >= 0x80 {
								buf[i] = byte(ke) | 0x80
								ke >>= 7
								i++
							}
							buf[i] = byte(ke)
							i++
						}
						{
							uv := uint32(vl)
							if vl < 0 {
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
					}
				}
			}
		}
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *Uint32Int32MapArrayAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		for j := 0; j < 4; j++ {
			(*v)[j] = new(map[uint32]int32)
			if buf[i] == 0 {
				i++
				(*v)[j] = nil
			} else if buf[i] != 1 {
				i++
				return i, errs.ErrWrongByte
			} else {
				i++
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
					(*(*v)[j]) = make(map[uint32]int32)
					for ; length > 0; length-- {
						var kem uint32
						var vlm int32
						{
							if i > len(buf)-1 {
								return i, errs.ErrSmallBuf
							}
							shift := 0
							done := false
							for l, b := range buf[i:] {
								if l == 4 && b > 15 {
									return i, errs.ErrOverflow
								}
								if b < 0x80 {
									kem = kem | uint32(b)<<shift
									done = true
									i += l + 1
									break
								}
								kem = kem | uint32(b&0x7F)<<shift
								shift += 7
							}
							if !done {
								return i, errs.ErrSmallBuf
							}
						}
						if err != nil {
							err = errs.NewMapKeyError(kem, err)
							break
						}
						{
							var uv uint32
							{
								if i > len(buf)-1 {
									return i, errs.ErrSmallBuf
								}
								shift := 0
								done := false
								for l, b := range buf[i:] {
									if l == 4 && b > 15 {
										return i, errs.ErrOverflow
									}
									if b < 0x80 {
										uv = uv | uint32(b)<<shift
										done = true
										i += l + 1
										break
									}
									uv = uv | uint32(b&0x7F)<<shift
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
							vlm = int32(uv)
						}
						if err != nil {
							err = errs.NewMapValueError(kem, vlm, err)
							break
						}
						(*(*v)[j])[kem] = vlm
					}
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
func (v Uint32Int32MapArrayAlias) Size() int {
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
					for ke, vl := range *item {
						{
							for ke >= 0x80 {
								ke >>= 7
								size++
							}
							size++
						}
						{
							uv := uint32(vl<<1) ^ uint32(vl>>31)
							{
								for uv >= 0x80 {
									uv >>= 7
									size++
								}
								size++
							}
						}
					}
				}
			}
		}
	}
	return size
}
