// Code generated by "stringer -type=AttrType"; DO NOT EDIT.

package yoti

import "strconv"

const _AttrType_name = "AttrTypeTimeAttrTypeStringAttrTypeJPEGAttrTypePNGAttrTypeJSONAttrTypeBoolAttrTypeInterface"

var _AttrType_index = [...]uint8{0, 12, 26, 38, 49, 61, 73, 90}

func (i AttrType) String() string {
	i -= 1
	if i >= AttrType(len(_AttrType_index)-1) {
		return "AttrType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _AttrType_name[_AttrType_index[i]:_AttrType_index[i+1]]
}
