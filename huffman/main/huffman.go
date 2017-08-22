package main

import "learn-golang/huffman"
import "flag"
import "fmt"
import "log"
import "os"
import "bufio"

var encode_file string
var decode_file string
var out_file string

func init() {
	flag.StringVar(&encode_file, "encode", "", "File to encode")
	flag.StringVar(&decode_file, "decode", "", "File to decode")
	flag.StringVar(&out_file, "out", "", "Output file")
}

func main() {
	flag.Parse()

	if out_file == "" {
		fmt.Println("You need so specify an out file")
		return
	}

	if encode_file == "" && decode_file == "" {
		fmt.Println("You need to specify encode or decode")
		return
	}

	if encode_file != "" && decode_file != "" {
		fmt.Println("You must either encode or decode, not both")
		return
	}

	if encode_file != "" {
		encode(encode_file, out_file)
	} else {
		decode(decode_file, out_file)
	}
}

func encode (fname_in string, fname_out string) {
	reader,err := os.Open(fname_in)
	if err != nil {
		log.Fatal("Can't open file ", fname_in, " for reading")
	}

	writer,err := os.Create(fname_out)
	if err != nil {
		log.Fatal("Can't open file ", fname_out, " for writing")
	}

	header := huffman.BuildHeader(reader)
	dict := huffman.Header2Dict(header)

	reader.Seek(0, 0)
	ch := make(chan byte, 100)
	go huffman.WriteFile(reader, dict, ch)

	header_bytes := huffman.Char2Byte(header)
	writer.Write(header_bytes)

	buff_writer := bufio.NewWriter(writer)
	for v := range(ch) {
		buff_writer.WriteByte(v)
	}
	buff_writer.Flush()
	writer.Close()
}

func decode (fname_in string, fname_out string) {
	reader,err := os.Open(fname_in)
	if err != nil {
		log.Fatal("Can't open file ", fname_in, " for reading")
	}

	writer,err := os.Create(fname_out)
	if err != nil {
		log.Fatal("Can't open file ", fname_out, " for writing")
	}

	tree := huffman.Header2Tree(reader)

	ch := make(chan byte, 100)
	go huffman.ReadFile(reader, *tree, ch)
	buff_writer := bufio.NewWriter(writer)
	for v := range(ch) {
		buff_writer.WriteByte(v)
	}
	buff_writer.Flush()
	writer.Close()
}
