// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/muserrs"

// Marshal fills buf with the MUS encoding of v.
func (v TrickyMapAlias) Marshal(buf []byte) int {
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
				for _, item := range ke {
					{
						si := item.Marshal(buf[i:])
						i += si
					}
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
				for ke, vl := range vl {
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
							for _, el := range *ke {
								{
									si := el.Marshal(buf[i:])
									i += si
								}
							}
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
						for ke, vl := range vl {
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
									for ke, vl := range *ke {
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
												panic(muserrs.ErrSmallBuf)
											}
											i += copy(buf[i:], vl)
										}
									}
								}
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
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *TrickyMapAlias) Unmarshal(buf []byte) (int, error) {
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
		(*v) = make(map[[2]StringAlias]map[*[]SimpleStructType]map[*map[int]string][2]int)
		for ; length > 0; length-- {
			var kem [2]StringAlias
			var vlm map[*[]SimpleStructType]map[*map[int]string][2]int
			{
				for j := 0; j < 2; j++ {
					{
						var sv StringAlias
						si := 0
						si, err = sv.Unmarshal(buf[i:])
						if err == nil {
							kem[j] = sv
							i += si
						}
					}
					if err != nil {
						err = muserrs.NewArrayError(j, err)
						break
					}
				}
			}
			if err != nil {
				err = muserrs.NewMapKeyError(kem, err)
				break
			}
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
				vlm = make(map[*[]SimpleStructType]map[*map[int]string][2]int)
				for ; length > 0; length-- {
					kemm := new([]SimpleStructType)
					var vlmm map[*map[int]string][2]int
					if buf[i] == 0 {
						i++
						kemm = nil
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
							(*kemm) = make([]SimpleStructType, length)
							for j := 0; j < length; j++ {
								{
									var sv SimpleStructType
									si := 0
									si, err = sv.Unmarshal(buf[i:])
									if err == nil {
										(*kemm)[j] = sv
										i += si
									}
								}
								if err != nil {
									err = muserrs.NewSliceError(j, err)
									break
								}
							}
						}
					}
					if err != nil {
						err = muserrs.NewMapKeyError(kemm, err)
						break
					}
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
						vlmm = make(map[*map[int]string][2]int)
						for ; length > 0; length-- {
							kemmm := new(map[int]string)
							var vlmmm [2]int
							if buf[i] == 0 {
								i++
								kemmm = nil
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
									(*kemmm) = make(map[int]string)
									for ; length > 0; length-- {
										var kemmmm int
										var vlmmmm string
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
											kemmmm = int(uv)
										}
										if err != nil {
											err = muserrs.NewMapKeyError(kemmmm, err)
											break
										}
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
											vlmmmm = string(buf[i : i+length])
											i += length
										}
										if err != nil {
											err = muserrs.NewMapValueError(kemmmm, vlmmmm, err)
											break
										}
										(*kemmm)[kemmmm] = vlmmmm
									}
								}
							}
							if err != nil {
								err = muserrs.NewMapKeyError(kemmm, err)
								break
							}
							{
								for j := 0; j < 2; j++ {
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
										vlmmm[j] = int(uv)
									}
									if err != nil {
										err = muserrs.NewArrayError(j, err)
										break
									}
								}
							}
							if err != nil {
								err = muserrs.NewMapValueError(kemmm, vlmmm, err)
								break
							}
							(vlmm)[kemmm] = vlmmm
						}
					}
					if err != nil {
						err = muserrs.NewMapValueError(kemm, vlmm, err)
						break
					}
					(vlm)[kemm] = vlmm
				}
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
func (v TrickyMapAlias) Size() int {
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
				for _, item := range ke {
					{
						ss := item.Size()
						size += ss
					}
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
				for ke, vl := range vl {
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
							for _, el := range *ke {
								{
									ss := el.Size()
									size += ss
								}
							}
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
						for ke, vl := range vl {
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
									for ke, vl := range *ke {
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
	return size
}
