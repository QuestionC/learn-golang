package huffman

// Just a simple huffman encoding program

import (
	"io"
	"fmt"
	"container/heap"
	"bufio"
)

func BuildHeader(r io.Reader) []int {
	// Probably faster to use buffered io reader?

	reader := bufio.NewReader(r)

	m := make(map[int]int)

	for {
		v,err := reader.ReadByte()

		m[int(v)] += 1

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
			value: []int{value},
			weight: weight,
		}
		i++
	}
	heap.Init(&myheap)

	// Now we pop the top of the heap and form the two lightest nodes into a united node
	for (myheap.Len() > 1) {
		item1 := heap.Pop(&myheap).(*heapItem)
		item2 := myheap[0]

		// fmt.Println(item1, item2)

		new_value := []int{0}
		new_value = append(new_value, item1.value...)
		new_value = append(new_value, item2.value...)

		myheap[0].value = new_value
		myheap[0].weight = item1.weight + item2.weight
		
		// fmt.Println(myheap[0])
	
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
func WriteFile(r io.Reader, dict map[int][]int, ch chan byte) {
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

		bits := dict[int(v)]

		fmt.Println(bits)

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

func ReadFile(r io.Reader, dict map[int][]int, ch chan byte) {

	close(ch)
}

// Given a bytestring describing the huffman tree, create a dictionary mapping 
//  bytes to encoding
// This is needed for writing
func Header2Dict(convert_me []int) map[int] []int {
	result := make(map[int] []int)

	// Not trivial for loops rely on pointer arithmetic.
	// Might as well use forever loops in golang.
	i := 0
	var curr_bitstring []int

	for {
		// Branch to the left
		v := convert_me[i]
		for {
			if v != 0 {
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


////////
// Priority queue implementation, based on https://golang.org/pkg/container/heap/#pkg-overview

type heapItem struct {
	value []int
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

