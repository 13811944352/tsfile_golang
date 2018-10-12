package encoder

import (
	"bytes"
	"tsfile/common/constant"
)

/**
 * @Package Name: encoder
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-9-28 下午5:55
 * @Description:
 */

const (
	MAX_STRING_LENGTH = "max_string_length"
	MAX_POINT_NUMBER  = "max_point_number"
)

type Encoder interface {
	Encode(value interface{}, buffer *bytes.Buffer) ()
	Flush(buffer *bytes.Buffer) ()
	GetOneItemMaxSize() (int)
	GetMaxByteSize() (int64)
}

func GetEncoder(et int16, tdt int16) (Encoder) {
	var encoder Encoder
	switch {
	case et == int16(constant.PLAIN):
		encoder, _ = NewPlainEncoder(tdt, et)
	case et == int16(constant.RLE):
		encoder, _ = NewPlainEncoder(tdt, et)
	case et == int16(constant.TS_2DIFF):
		encoder, _ = NewPlainEncoder(tdt, et)
	case et == int16(constant.GORILLA):
		encoder, _ = NewPlainEncoder(tdt, et)
	}

	return encoder
}
