package statistics

import (
	"tsfile/common/constant"
	"tsfile/common/utils"
)

type Boolean struct {
	max     bool
	min     bool
	first   bool
	last    bool
	sum     float64
	isEmpty bool
}

func (s *Boolean) Deserialize(reader *utils.FileReader) {
	s.min = reader.ReadBool()
	s.max = reader.ReadBool()
	s.first = reader.ReadBool()
	s.last = reader.ReadBool()
	s.sum = reader.ReadDouble()
}

func (b *Boolean) SizeOfDaum() int {
	return 1
}

func (b *Boolean) GetMaxByte(tdt int16) []byte {
	return utils.BoolToByte(b.max, 0)
}

func (b *Boolean) GetMinByte(tdt int16) []byte {
	return utils.BoolToByte(b.min, 0)
}

func (b *Boolean) GetFirstByte(tdt int16) []byte {
	return utils.BoolToByte(b.first, 0)
}

func (b *Boolean) GetLastByte(tdt int16) []byte {
	return utils.BoolToByte(b.last, 0)
}

func (b *Boolean) GetSumByte(tdt int16) []byte {
	return utils.Float64ToByte(b.sum, 0)
}

func (b *Boolean) UpdateStats(iValue interface{}) {
	value := iValue.(bool)
	if !b.isEmpty {
		b.InitializeStats(value, value, value, value, 0)
		b.isEmpty = true
	} else {
		b.UpdateValue(value, value, value, value, 0)
		b.isEmpty = false
	}
}

func (b *Boolean) UpdateValue(max bool, min bool, first bool, last bool, sum float64) {
	if max && !b.max {
		b.max = max
	}
	if !min && b.min {
		b.min = min
	}
	b.last = last
}

func (b *Boolean) InitializeStats(max bool, min bool, first bool, last bool, sum float64) {
	b.max = max
	b.min = min
	b.first = first
	b.last = last
}

func (s *Boolean) GetSerializedSize() int {
	return 4*constant.BOOLEAN_LEN + constant.DOUBLE_LEN
}

//func NewBool() (*Statistics, error) {
//
//	return &Statistics{
//		isEmpty:true,
//	},nil
//}
