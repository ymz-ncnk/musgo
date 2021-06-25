package testdata

import "github.com/ymz-ncnk/musgo/errs"

// MarshalMUS fills buf with the MUS encoding of v.
func (v MyMap) MarshalMUS(buf []byte) int {
	i := 0
	{
		length := len(v)
		{
			uv := uint64(length<<1) ^ uint64(length>>63)
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
					uv := uint64(length<<1) ^ uint64(length>>63)
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
				i += copy(buf[i:], ke)
			}
			{
				uv := uint64(vl<<1) ^ uint64(vl>>63)
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

// UnmarshalMUS parses the MUS-encoded buf, and sets the result to *v.
func (v *MyMap) UnmarshalMUS(buf []byte) (int, error) {
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
			uv = (uv >> 1) ^ uint64((int(uv&1)<<63)>>63)
			length = int(uv)
		}
		if length < 0 {
			return i, errs.ErrNegativeLength
		}
		if length > 3 {
			err = errs.ErrMaxLengthExceeded
		} else {
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
						uv = (uv >> 1) ^ uint64((int(uv&1)<<63)>>63)
						length = int(uv)
					}
					if length < 0 {
						return i, errs.ErrNegativeLength
					}
					if len(buf) < i+length {
						return i, errs.ErrSmallBuf
					}
					kem = string(buf[i : i+length])
					i += length
					err = NotHello(kem)
				}
				if err != nil {
					err = errs.NewMapKeyError(kem, err)
					break
				}
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
					uv = (uv >> 1) ^ uint64((int(uv&1)<<63)>>63)
					vlm = int(uv)
					err = BiggerThenTen(vlm)
				}
				if err != nil {
					err = errs.NewMapValueError(kem, vlm, err)
					break
				}
				(*v)[kem] = vlm
			}
			if err == nil {
				err = ValidateMyMap(v)
			}
		}
	}
	return i, err
}

// SizeMUS returns the size of the MUS-encoded v.
func (v MyMap) SizeMUS() int {
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
