package bitslice

import "fmt"
import "bytes"

func ExampleBitWriter() {
	var b bytes.Buffer

	bitwrite := NewWriter(&b)

	// 5
	bitwrite.WriteBit(1)
	bitwrite.WriteBit(0)
	bitwrite.WriteBit(1)
	bitwrite.WriteBit(0)
	bitwrite.WriteBit(0)
	bitwrite.WriteBit(0)
	bitwrite.WriteBit(0)
	bitwrite.WriteBit(0)

	// 2
	bitwrite.WriteBit(0)
	bitwrite.WriteBit(1)

	bitwrite.flush()

	fmt.Println(b.Bytes())

	// Output:
	// [5 2]
}

func ExampleBitReader() {
	b := []byte{5, 2}

	buff := bytes.NewBuffer(b)

	r := NewReader(buff)

	for i := 0; i < 16; i++ {
		fmt.Print(r.ReadBit())
	}

	// Output:
	// 1010000001000000
}

func ExampleWriterReader() {
	var b []byte

	writeBuff := bytes.NewBuffer(b)

	w := NewWriter(writeBuff)

	w.WriteBit(1)
	w.WriteBit(0)
	w.WriteBit(1)
	w.WriteBit(0)
	w.WriteBit(1)
	w.WriteBit(1)
	w.WriteBit(0)
	w.WriteBit(0)

  w.WriteBit(0)
  w.WriteBit(0)
  w.WriteBit(0)
  w.WriteBit(0)
  w.WriteBit(1)
  w.WriteBit(1)
  w.WriteBit(1)
  w.WriteBit(1)

	readBuff := bytes.NewBuffer(writeBuff.Bytes())
	r := NewReader(readBuff)

	for i := 0; i < 8; i++ {
		fmt.Printf("%d", r.ReadBit())
	}
	fmt.Printf(" ")
	for i := 0; i < 8; i++ {
		fmt.Printf("%d", r.ReadBit())
	}

	// Output:
	// 10101100 00001111
}
