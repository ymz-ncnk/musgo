// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/muserrs"

// Marshal fills buf with the MUS encoding of v.
func (v TrickyArrayAlias) Marshal(buf []byte) int {
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
				for _, el := range item {
					{
						for _, item := range el {
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
								for ke, vl := range item {
									{
										si := ke.Marshal(buf[i:])
										i += si
									}
									{
										for _, item := range vl {
											{
												uv := uint64(item)
												if item < 0 {
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
					}
				}
			}
		}
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *TrickyArrayAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		for j := 0; j < 2; j++ {
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
				(*v)[j] = make([][3]map[SimpleStructType][1]int, length)
				for jj := 0; jj < length; jj++ {
					{
						for jjj := 0; jjj < 3; jjj++ {
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
								(*v)[j][jj][jjj] = make(map[SimpleStructType][1]int)
								for ; length > 0; length-- {
									var kem SimpleStructType
									var vlm [1]int
									{
										var sv SimpleStructType
										si := 0
										si, err = sv.Unmarshal(buf[i:])
										if err == nil {
											kem = sv
											i += si
										}
									}
									if err != nil {
										err = muserrs.NewMapKeyError(kem, err)
										break
									}
									{
										for j := 0; j < 1; j++ {
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
												vlm[j] = int(uv)
											}
											if err != nil {
												err = muserrs.NewArrayError(j, err)
												break
											}
										}
									}
									if err != nil {
										err = muserrs.NewMapValueError(kem, vlm, err)
										break
									}
									((*v)[j][jj][jjj])[kem] = vlm
								}
							}
							if err != nil {
								err = muserrs.NewArrayError(jjj, err)
								break
							}
						}
					}
					if err != nil {
						err = muserrs.NewSliceError(jj, err)
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
func (v TrickyArrayAlias) Size() int {
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
				for _, el := range item {
					{
						for _, item := range el {
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
								for ke, vl := range item {
									{
										ss := ke.Size()
										size += ss
									}
									{
										for _, item := range vl {
											{
												uv := uint64(item<<1) ^ uint64(item>>63)
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
					}
				}
			}
		}
	}
	return size
}
