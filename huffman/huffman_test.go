package huffman

import "strings"
import "os"
import "bufio"
import "bytes"


func ExampleA() {
	input := "Lbh penpxrq gur pbqr!"
	s := strings.NewReader(input)

	header := BuildHeader(s)
	d := Header2Dict(header)

	s.Reset(input)

	ch := make(chan byte, 100)

	go WriteFile(s, d, ch)

	file,_ := os.Create("test.qz")

	write := bufio.NewWriter(file)

	for v := range(ch) {
		write.WriteByte(v)
	}
	write.Flush()
	file.Close()

	header_buff := Char2Byte(header)
	header_reader := bytes.NewReader(header_buff)
	tree := Header2Tree(header_reader)

	file,_ = os.Open("test.qz")

	ch = make (chan byte, 100)

	go ReadFile(file, *tree, ch)

	out := bufio.NewWriter(os.Stdout)
	for v := range(ch) {
		out.WriteByte(v)
	}
	out.Flush()

	// Output:
	// Lbh penpxrq gur pbqr!
}

func ExampleB() {
	// Encode a file containing a string including the header
	test_string := "How now brown cow?"
	string_reader := strings.NewReader(test_string)

	header := BuildHeader(string_reader)
	dict := Header2Dict(header)

	string_reader.Reset(test_string)

	ch := make(chan byte, 100)

	go WriteFile(string_reader, dict, ch)

	out_file,_ := os.Create("ExampleB.qz")

	header_bytes := Char2Byte(header)

	// First write the header
	out_file.Write(header_bytes)

	// Write the encoded data
	writer := bufio.NewWriter(out_file)
	for v := range(ch) {
		writer.WriteByte(v)
	}
	writer.Flush()
	out_file.Close()

	// Now decode the file
	in_file,_ := os.Open("ExampleB.qz")
	ch = make(chan byte, 100)

	tree := Header2Tree(in_file)

	ReadFile(in_file, *tree, ch)
	out := bufio.NewWriter(os.Stdout)
	for v:= range(ch) {
		out.WriteByte(v)
	}
	out.Flush()

	// Output:
	// How now brown cow?
}
