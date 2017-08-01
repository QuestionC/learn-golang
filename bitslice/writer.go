package bitslice

import "io"
import "log"

type writer struct {
	out io.ByteWriter
	next_byte byte
	used_bits uint // Range from 0-7
}

/// Write a 0 or 1 to the stream
func (self *writer) WriteBit(write_me int) {
	if write_me != 0 && write_me != 1 {
		log.Fatalf("WriteBit(%d)\n", write_me)
	}

	var mask byte
	mask = byte(write_me) << self.used_bits
	self.next_byte |= mask
	self.used_bits++

	if self.used_bits == 8 {
		self.flush()
	}
}

func (self *writer) flush() {
	if self.used_bits != 8 {
		// log.Printf("writer.flush() expected writer.used_bits 8, got %d\n", self.used_bits)
	}
	self.out.WriteByte(self.next_byte)

	self.next_byte = 0
	self.used_bits = 0
}

func (self *writer) close() {
	if self.used_bits != 0 {
		self.flush()
	}
}

func NewWriter(buff io.ByteWriter) *writer {
	result := new(writer)
	result.out = buff
	return result
}
