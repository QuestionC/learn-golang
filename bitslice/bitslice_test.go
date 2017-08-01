package bitslice

import "fmt"

func ExampleShort() {
	b := New(1, 0, 1, 1)
	fmt.Println("Len", b.Len())
	fmt.Println("Cap", b.Cap())
	fmt.Println(b)
	// Output: 
	// Len 4
	// Cap 8
	// 1011
}

func ExampleSelfAppend() {
	b := New(1, 0)
	fmt.Println(b)
	b = b.Append(1)
	fmt.Println(b)
	b = b.Append(0)
	fmt.Println(b)
	// Output:
	// 10
	// 101
	// 1010
}

func ExampleLong() {
	b := New(1, 0, 1, 1, 0, 0, 0, 1, 1, 0, 1)
	fmt.Println("Len", b.Len())
	fmt.Println("Cap", b.Cap())
	fmt.Println(b)
	// Output: 
	// Len 11
	// Cap 16
	// 10110001 101
}

func ExampleAppend1() {
	b := New(0, 0, 0, 0, 0, 0)
	c := b.Append(1)
	fmt.Println(c)
	c = c.Append(1)
	fmt.Println(c)
	c = c.Append(1)
	fmt.Println(c)
	fmt.Println(b)
	// Output:
	// 0000001
	// 00000011
	// 00000011 1
	// 000000
}

func ExampleAppend0() {
	b := New(1, 1, 1, 1, 1, 1)
	c := b.Append(0)
	fmt.Println(c)
	c = c.Append(0)
	fmt.Println(c)
	c = c.Append(0)
	fmt.Println(c)
	fmt.Println(b)
	// Output:
	// 1111110
	// 11111100
	// 11111100 0
	// 111111
}

