// Code generated by musgen. DO NOT EDIT.

package musgen

import (
	"math"

	"github.com/ymz-ncnk/serialization/musgo/errs"
)

// Marshal fills buf with the MUS encoding of v.
func (v RawStructType) Marshal(buf []byte) int {
	i := 0
	{
		buf[i] = byte(v.UintRaw)
		i++
		buf[i] = byte(v.UintRaw >> 8)
		i++
		buf[i] = byte(v.UintRaw >> 16)
		i++
		buf[i] = byte(v.UintRaw >> 24)
		i++
		buf[i] = byte(v.UintRaw >> 32)
		i++
		buf[i] = byte(v.UintRaw >> 40)
		i++
		buf[i] = byte(v.UintRaw >> 48)
		i++
		buf[i] = byte(v.UintRaw >> 56)
		i++
	}
	{
		uv := math.Float32bits(float32(v.Float32Raw))
		{
			buf[i] = byte(uv)
			i++
			buf[i] = byte(uv >> 8)
			i++
			buf[i] = byte(uv >> 16)
			i++
			buf[i] = byte(uv >> 24)
			i++
		}
	}
	{
		{
			buf[i] = byte((**v.IntRawPtrPtr))
			i++
			buf[i] = byte((**v.IntRawPtrPtr) >> 8)
			i++
			buf[i] = byte((**v.IntRawPtrPtr) >> 16)
			i++
			buf[i] = byte((**v.IntRawPtrPtr) >> 24)
			i++
			buf[i] = byte((**v.IntRawPtrPtr) >> 32)
			i++
			buf[i] = byte((**v.IntRawPtrPtr) >> 40)
			i++
			buf[i] = byte((**v.IntRawPtrPtr) >> 48)
			i++
			buf[i] = byte((**v.IntRawPtrPtr) >> 56)
			i++
		}
	}
	{
		length := len((*v.MapRawPtr))
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
		for ke, vl := range *v.MapRawPtr {
			{
				uv := math.Float64bits(float64((*ke)))
				{
					buf[i] = byte(uv)
					i++
					buf[i] = byte(uv >> 8)
					i++
					buf[i] = byte(uv >> 16)
					i++
					buf[i] = byte(uv >> 24)
					i++
					buf[i] = byte(uv >> 32)
					i++
					buf[i] = byte(uv >> 40)
					i++
					buf[i] = byte(uv >> 48)
					i++
					buf[i] = byte(uv >> 56)
					i++
				}
			}
			{
				{
					buf[i] = byte((*vl))
					i++
					buf[i] = byte((*vl) >> 8)
					i++
				}
			}
		}
	}
	{
		length := len(v.SliceRaw)
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
		for _, el := range v.SliceRaw {
			{
				buf[i] = byte(el)
				i++
				buf[i] = byte(el >> 8)
				i++
				buf[i] = byte(el >> 16)
				i++
				buf[i] = byte(el >> 24)
				i++
				buf[i] = byte(el >> 32)
				i++
				buf[i] = byte(el >> 40)
				i++
				buf[i] = byte(el >> 48)
				i++
				buf[i] = byte(el >> 56)
				i++
			}
		}
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *RawStructType) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		if len(buf) < 8 {
			return i, errs.ErrSmallBuf
		}
		v.UintRaw = uint(buf[i])
		i++
		v.UintRaw |= uint(buf[i]) << 8
		i++
		v.UintRaw |= uint(buf[i]) << 16
		i++
		v.UintRaw |= uint(buf[i]) << 24
		i++
		v.UintRaw |= uint(buf[i]) << 32
		i++
		v.UintRaw |= uint(buf[i]) << 40
		i++
		v.UintRaw |= uint(buf[i]) << 48
		i++
		v.UintRaw |= uint(buf[i]) << 56
		i++
		err = BiggerThanTenUint(v.UintRaw)
	}
	if err != nil {
		return i, errs.NewFieldError("UintRaw", err)
	}
	{
		var uv uint32
		{
			if len(buf) < 4 {
				return i, errs.ErrSmallBuf
			}
			uv = uint32(buf[i])
			i++
			uv |= uint32(buf[i]) << 8
			i++
			uv |= uint32(buf[i]) << 16
			i++
			uv |= uint32(buf[i]) << 24
			i++
		}
		v.Float32Raw = float32(math.Float32frombits(uv))
	}
	if err != nil {
		return i, errs.NewFieldError("Float32Raw", err)
	}
	{
		tmp0 := new(int)
		v.IntRawPtrPtr = &tmp0
	}
	{
		{
			if len(buf) < 8 {
				return i, errs.ErrSmallBuf
			}
			(**v.IntRawPtrPtr) = int(buf[i])
			i++
			(**v.IntRawPtrPtr) |= int(buf[i]) << 8
			i++
			(**v.IntRawPtrPtr) |= int(buf[i]) << 16
			i++
			(**v.IntRawPtrPtr) |= int(buf[i]) << 24
			i++
			(**v.IntRawPtrPtr) |= int(buf[i]) << 32
			i++
			(**v.IntRawPtrPtr) |= int(buf[i]) << 40
			i++
			(**v.IntRawPtrPtr) |= int(buf[i]) << 48
			i++
			(**v.IntRawPtrPtr) |= int(buf[i]) << 56
			i++
		}
	}
	if err != nil {
		return i, errs.NewFieldError("IntRawPtrPtr", err)
	}
	v.MapRawPtr = new(map[*float64]*int16)
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
		(*v.MapRawPtr) = make(map[*float64]*int16)
		for ; length > 0; length-- {
			kem := new(float64)
			vlm := new(int16)
			{
				var uv uint64
				{
					if len(buf) < 8 {
						return i, errs.ErrSmallBuf
					}
					uv = uint64(buf[i])
					i++
					uv |= uint64(buf[i]) << 8
					i++
					uv |= uint64(buf[i]) << 16
					i++
					uv |= uint64(buf[i]) << 24
					i++
					uv |= uint64(buf[i]) << 32
					i++
					uv |= uint64(buf[i]) << 40
					i++
					uv |= uint64(buf[i]) << 48
					i++
					uv |= uint64(buf[i]) << 56
					i++
				}
				(*kem) = float64(math.Float64frombits(uv))
			}
			if err != nil {
				err = errs.NewMapKeyError(kem, err)
				break
			}
			{
				{
					if len(buf) < 2 {
						return i, errs.ErrSmallBuf
					}
					(*vlm) = int16(buf[i])
					i++
					(*vlm) |= int16(buf[i]) << 8
					i++
				}
				err = BiggerThanTenInt16Ptr(vlm)
			}
			if err != nil {
				err = errs.NewMapValueError(kem, vlm, err)
				break
			}
			(*v.MapRawPtr)[kem] = vlm
		}
	}
	if err != nil {
		return i, errs.NewFieldError("MapRawPtr", err)
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
		v.SliceRaw = make([]uint, length)
		for j := 0; j < length; j++ {
			{
				if len(buf) < 8 {
					return i, errs.ErrSmallBuf
				}
				v.SliceRaw[j] = uint(buf[i])
				i++
				v.SliceRaw[j] |= uint(buf[i]) << 8
				i++
				v.SliceRaw[j] |= uint(buf[i]) << 16
				i++
				v.SliceRaw[j] |= uint(buf[i]) << 24
				i++
				v.SliceRaw[j] |= uint(buf[i]) << 32
				i++
				v.SliceRaw[j] |= uint(buf[i]) << 40
				i++
				v.SliceRaw[j] |= uint(buf[i]) << 48
				i++
				v.SliceRaw[j] |= uint(buf[i]) << 56
				i++
				err = BiggerThanTenUint(v.SliceRaw[j])
			}
			if err != nil {
				err = errs.NewSliceError(j, err)
				break
			}
		}
	}
	if err != nil {
		return i, errs.NewFieldError("SliceRaw", err)
	}
	return i, err
}

// Size returns the size of the MUS-encoded v.
func (v RawStructType) Size() int {
	size := 0
	{
		_ = v.UintRaw
		size += 8
	}
	{
		{
			_ = v.Float32Raw
			size += 4
		}

	}
	{
		{
			_ = (**v.IntRawPtrPtr)
			size += 8
		}
	}
	{
		length := len((*v.MapRawPtr))
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
		for ke, vl := range *v.MapRawPtr {
			{
				{
					_ = (*ke)
					size += 8
				}

			}
			{
				{
					_ = (*vl)
					size += 2
				}
			}
		}
	}
	{
		length := len(v.SliceRaw)
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
		for _, el := range v.SliceRaw {
			{
				_ = el
				size += 8
			}
		}
	}
	return size
}
