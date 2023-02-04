// Code generated by musgen. DO NOT EDIT.

package musgen

import "github.com/ymz-ncnk/serialization/musgo/errs"

// Marshal fills buf with the MUS encoding of v.
func (v StructTypePtrArrayAlias) Marshal(buf []byte) int {
	i := 0
	{
		for _, item := range v {
			{
				si := (*item).Marshal(buf[i:])
				i += si
			}
		}
	}
	return i
}

// Unmarshal parses the MUS-encoded buf, and sets the result to *v.
func (v *StructTypePtrArrayAlias) Unmarshal(buf []byte) (int, error) {
	i := 0
	var err error
	{
		for j := 0; j < 3; j++ {
			(*v)[j] = new(SimpleStructType)
			{
				var sv SimpleStructType
				si := 0
				si, err = sv.Unmarshal(buf[i:])
				if err == nil {
					(*(*v)[j]) = sv
					i += si
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
func (v StructTypePtrArrayAlias) Size() int {
	size := 0
	{
		for _, item := range v {
			{
				ss := (*item).Size()
				size += ss
			}
		}
	}
	return size
}
