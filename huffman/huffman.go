package huffman

// Just a simple huffman encoding program

import (
	"io"
	"container/heap"
	"bufio"
	"log"
	"encoding/binary"
)

type Char int16

func Char2Byte(convert_me []Char) []byte {
	result := make([]byte, len(convert_me)*2)
	for i := 0; i < len(convert_me); i++ {
		binary.LittleEndian.PutUint16(result[i*2:], uint16(convert_me[i]))
	}
	return result
}

func Byte2Char(convert_me []byte) []Char {
	result := make([]Char, len(convert_me)/2)
	for i := 0; i < len(result); i++ {
		result[i] = Char(binary.LittleEndian.Uint16(convert_me[i*2:]))
	}
	return result
}

func BuildHeader(r io.Reader) []Char {
	reader := bufio.NewReader(r)

	m := make(map[Char]int)

	for {
		v,err := reader.ReadByte()

		m[Char(v)] += 1

		if err == io.EOF {
			break
		}
	}

	// -1 for EOF
	m[-1] += 1

	// Now put the items in a heap
	myheap := make(huffmanHeap, len(m))

	i := 0
	for value, weight := range m {
		myheap[i] = &heapItem {
			value: []Char{value},
			weight: weight,
		}
		i++
	}
	heap.Init(&myheap)

	// Now we pop the top of the heap and form the two lightest nodes into a united node
	for (myheap.Len() > 1) {
		item1 := heap.Pop(&myheap).(*heapItem)
		item2 := myheap[0]

		new_value := []Char{-2}

		new_value = append(new_value, item1.value...)
		new_value = append(new_value, item2.value...)

		myheap[0].value = new_value
		myheap[0].weight = item1.weight + item2.weight

		heap.Fix(&myheap, 0)
	}

	return myheap[0].value

	// I represent the tree with a prefix representation byte string like 0 0 1 4 0 3 6
	// Each 0 is a branch in the tree followed by its left and right children.
	// So that string would be
	//     / \
	//    /\ /\
	//    14 36

	// So we walk the bytestring to generate the codes
	//dict := String2Dict(myheap[0].value)
}

// Using the provided dict, encode r
func WriteFile(r io.Reader, dict map[Char][]int, ch chan byte) {
	log.Println("WriteFile")
	var next_byte byte
	var mask byte

	mask = 1

	reader := bufio.NewReader(r)

	for {
		_v,err := reader.ReadByte()
		v := int(_v)

		if err != nil {
			v = -1
		}

		bits := dict[Char(v)]

		for _,bit := range(bits) {
			if bit == 1 {
				next_byte |= mask
			} else {
				next_byte &^= mask
			}

			mask <<= 1
			if mask == 0 {
				mask = 1
				ch <- next_byte
			}
		}

		if v == -1 {
			break
		}
	}

	// If mask==1, we already sent the byte
	if mask != 1 {
		ch <- next_byte
	}

	close(ch)
}

func ReadFile(r io.Reader, tree Tree, ch chan byte) {
	mask := byte(0)
	var next_byte byte

	reader := bufio.NewReader(r)

	node := tree

	for {
		if mask == 0 {
			var err error
			next_byte,err = reader.ReadByte()

			if err != nil {
				break
			}
			mask = 1
		}

		val := mask & next_byte
		if val == 0 {
			node = *node.Left
		} else {
			node = *node.Right
		}

		if node.Value != nil {
			if *node.Value == -1 {
				break
			}
			ch <- byte(*node.Value)
			node = tree
		}
		mask <<= 1
	}

	close(ch)
}

// Given a bytestring describing the huffman tree, create a dictionary mapping 
//  bytes to encoding
// This is needed for writing
func Header2Dict(convert_me []Char) map[Char] []int {
	result := make(map[Char] []int)

	// Not trivial for loops rely on pointer arithmetic.
	// Might as well use forever loops in golang.
	i := 0
	var curr_bitstring []int

	for {
		// Branch to the left
		v := convert_me[i]
		for {
			if v != -2 {
				break
			}

			curr_bitstring = append(curr_bitstring, 0)
			i++
			v = convert_me[i]
		}

		// v is the first non-0, so it's the leaf
		result[v] = append([]int {}, curr_bitstring...)

		// Last leaf?
		if i == len(convert_me) - 1 {
			break
		}

		// Next leaf
		// Bubble up the right branches
		var last_0_index int
		for last_0_index = len(curr_bitstring) - 1; curr_bitstring[last_0_index] == 1; last_0_index-- {
		}
		curr_bitstring = curr_bitstring[:last_0_index + 1]

		// Change from left to right branch
		curr_bitstring[last_0_index] = 1

		i++
	}

	return result
}
func Header2Tree(convert_me io.Reader) *Tree { var head Tree

	buff := make([]byte, 2)
	convert_me.Read(buff)
	curr_char := Char(binary.LittleEndian.Uint16(buff))

	if curr_char == -2 {
		head.Left= Header2Tree(convert_me)
		head.Right = Header2Tree(convert_me)
	} else {
		head.Value = &curr_char
	}
	return &head
}

////////
// Priority queue implementation, based on https://golang.org/pkg/container/heap/#pkg-overview

type heapItem struct {
	value []Char
	weight int
}

type huffmanHeap []*heapItem

func (this huffmanHeap) Len() int { return len(this) }

func (this huffmanHeap) Less(i, j int) bool {
	// We want to pop the smallest node off the heap
	return this[i].weight < this[j].weight
}

// Can't figure this out, part of heap interface
func (this huffmanHeap) Swap (i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *huffmanHeap) Push (x interface{}) {
	item := x.(*heapItem) // Cast X to Item? What's the asterisk for?
	*this = append(*this, item)
}

func (this *huffmanHeap) Pop() interface{} {
	old := *this // Copy heap?
	n := len(old) // length of heap
	item := old[n - 1] // last element
	*this = old[0 : n-1]
	return item
}

///// Tree Type for reverse decoding

type Tree struct {
	Left *Tree
	Right *Tree
	Value *Char
}
