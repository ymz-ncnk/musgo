// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/musgo/errs"

// Marshal fills buf with the MUS encoding of v.
func (v StrPtrIntPtrMapAlias) Marshal(buf []byte) int {
	i := 0
	{
		length := len(v)
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
		for ke, vl := range v {
			if ke == nil {
				buf[i] = 0
				i++
			} else {
				buf[i] = 1
				i++
				{
					length := len((*ke))
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
					i += copy(buf[i:], (*ke))
				}
			}
			if vl == nil {
				buf[i] = 0
				i++
			} else {
				buf[i] = 1
				i++
				{
					uv := uint64((*vl))
					if (*vl) < 0 {
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
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *StrPtrIntPtrMapAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
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
		(*v) = make(map[*string]*int)
		for ; length > 0; length-- {
			kem := new(string)
			vlm := new(int)
			if buf[i] == 0 {
				i++
				kem = nil
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
					if len(buf) < i+length {
						return i, errs.ErrSmallBuf
					}
					(*kem) = string(buf[i : i+length])
					i += length
				}
			}
			if err != nil {
				err = errs.NewMapKeyError(kem, err)
				break
			}
			if buf[i] == 0 {
				i++
				vlm = nil
			} else if buf[i] != 1 {
				i++
				return i, errs.ErrWrongByte
			} else {
				i++
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
					(*vlm) = int(uv)
				}
			}
			if err != nil {
				err = errs.NewMapValueError(kem, vlm, err)
				break
			}
			(*v)[kem] = vlm
		}
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v StrPtrIntPtrMapAlias) Size() int {
	size := 0
	{
		length := len(v)
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
		for ke, vl := range v {
			size++
			if ke != nil {
				{
					length := len((*ke))
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
					size += len((*ke))
				}
			}
			size++
			if vl != nil {
				{
					uv := uint64((*vl)<<1) ^ uint64((*vl)>>63)
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
	return size
}
