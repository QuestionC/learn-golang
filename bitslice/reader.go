package bitslice

import "io"
// import "log"

type reader struct {
	in io.ByteReader
	curr_byte byte
	read_bits uint // Range from 1-8
}

func (self *reader) ReadBit() int {
	if self.read_bits == 8 {
		self.read_bits = 0
		self.curr_byte,_ = self.in.ReadByte()
	}
	mask := 1 << self.read_bits

	self.read_bits += 1

	masked_curr_byte := byte(mask) & self.curr_byte
	if masked_curr_byte > 0 {
		return 1
	} else {
		return 0
	}
}

func NewReader(buff io.ByteReader) *reader {
	result := new(reader)
	result.in = buff
	result.read_bits = 8

	return result
}
