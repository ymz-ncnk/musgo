// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/muserrs"

// Marshal fills buf with the MUS encoding of v.
func (v StrIntMapAlias) Marshal(buf []byte) int {
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
				length := len(ke)
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
					panic(muserrs.ErrSmallBuf)
				}
				i += copy(buf[i:], ke)
			}
			{
				uv := uint64(vl)
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
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *StrIntMapAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
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
		(*v) = make(map[string]int)
		for ; length > 0; length-- {
			var kem string
			var vlm int
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
				if len(buf) < i+length {
					return i, muserrs.ErrSmallBuf
				}
				kem = string(buf[i : i+length])
				i += length
			}
			if err != nil {
				err = muserrs.NewMapKeyError(kem, err)
				break
			}
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
				vlm = int(uv)
			}
			if err != nil {
				err = muserrs.NewMapValueError(kem, vlm, err)
				break
			}
			(*v)[kem] = vlm
		}
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v StrIntMapAlias) Size() int {
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
				length := len(ke)
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
				size += len(ke)
			}
			{
				uv := uint64(vl<<1) ^ uint64(vl>>63)
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
	return size
}
