// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/serialization/musgo/errs"

// Marshal fills buf with the MUS encoding of v.
func (v UintIntStringMapPtrMapAlias) Marshal(buf []byte) int {
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
				length := len((*vl))
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
				for ke, vl := range *vl {
					{
						uv := uint64(ke)
						if ke < 0 {
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
					{
						length := len(vl)
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
						i += copy(buf[i:], vl)
					}
				}
			}
		}
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *UintIntStringMapPtrMapAlias) Unmarshal(buf []byte) (int, error) {
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
		(*v) = make(map[uint16]*map[int]string)
		for ; length > 0; length-- {
			var kem uint16
			vlm := new(map[int]string)
			{
				if i > len(buf)-1 {
					return i, errs.ErrSmallBuf
				}
				shift := 0
				done := false
				for l, b := range buf[i:] {
					if l == 2 && b > 3 {
						return i, errs.ErrOverflow
					}
					if b < 0x80 {
						kem = kem | uint16(b)<<shift
						done = true
						i += l + 1
						break
					}
					kem = kem | uint16(b&0x7F)<<shift
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
				(*vlm) = make(map[int]string)
				for ; length > 0; length-- {
					var kemm int
					var vlmm string
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
						kemm = int(uv)
					}
					if err != nil {
						err = errs.NewMapKeyError(kemm, err)
						break
					}
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
						vlmm = string(buf[i : i+length])
						i += length
					}
					if err != nil {
						err = errs.NewMapValueError(kemm, vlmm, err)
						break
					}
					(*vlm)[kemm] = vlmm
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
func (v UintIntStringMapPtrMapAlias) Size() int {
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
			{
				for ke >= 0x80 {
					ke >>= 7
					size++
				}
				size++
			}
			{
				length := len((*vl))
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
				for ke, vl := range *vl {
					{
						uv := uint64(ke<<1) ^ uint64(ke>>63)
						{
							for uv >= 0x80 {
								uv >>= 7
								size++
							}
							size++
						}
					}
					{
						length := len(vl)
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
						size += len(vl)
					}
				}
			}
		}
	}
	return size
}
