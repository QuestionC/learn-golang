// We need binary manipulation for huffman encoding
package bitslice

import "log"
import "io/ioutil"
//import "os"

type bitslice struct {
	data []byte
	spare_bits int // How many bits of the last byte we are using. Range 1-8
}

var trace *log.Logger
func init() {
	trace = log.New (ioutil.Discard, "", log.Ldate|log.Ltime|log.Lshortfile)
	//trace = log.New (os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func (self bitslice) Len() int {
	return len(self.data) * 8 - 8 + self.spare_bits
}

func (self bitslice) Cap() int {
	return len(self.data) * 8
}


func New(val ...int) bitslice {
	N := (len(val) + 7) / 8
	var result bitslice
	result.data = make([]byte, N)
	result.spare_bits = len(val) % 8
	if result.spare_bits == 0 {
		result.spare_bits = 8
	}

	var mask byte
	mask = 1
	byte_index := 0
	for i,v := range(val) {
		if v != 0 {
			result.data[byte_index] = result.data[byte_index] | mask
		}

		if i % 8 == 7 {
			mask = 1
			byte_index++
		} else {
			mask = mask * 2
		}
	}

	return result
}

func (self *bitslice) Append(b int) bitslice {
	trace.Printf("%s.Append(%d)", self, b)
	result := *self
	if result.spare_bits == 8 {
		trace.Print("spare_bits == 8, easy path")
		// Last bit is full, add a new one.
		result.spare_bits = 1
		next_bit := byte(b)
		result.data = append(result.data, next_bit)
	} else {
		trace.Print("hard path")
		result.spare_bits++

		if b != 0 {
			trace.Print("b != 0")
			mask := byte(1)
			last_byte := &result.data[len(result.data)-1]
			for i := 1; i < result.spare_bits; i++ {
				mask *= 2
			}
			*last_byte |= mask
		}
	}

	trace.Printf("Returned: %s", result)
	return result
}

func (self bitslice) String() string {
	var result string
	var mask byte

	mask = 1

	pivot := len(self.data) - 1
	front_data := self.data[:pivot]
	last_element := self.data[pivot]

	// Print the full bits
	for j,v := range(front_data) {
		mask := byte(1)
		for i := 0; i < 8; i++ {
			if mask & v != 0 {
				result += "1"
			} else {
				result += "0"
			}
			mask *= 2
		}

		if j != pivot - 1 || self.spare_bits != 0 {
			result += " "
		}
	}

	for i := 0; i < self.spare_bits; i++ {
		if mask & last_element != 0 {
			result += "1"
		} else {
			result += "0"
		}

		mask = mask * 2
	}
	return result
}
