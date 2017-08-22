package huffman

import "strings"
import "os"
import "bufio"
import "fmt"

func ExampleA() {
	input := "Lbh penpxrq gur pbqr!"
	s := strings.NewReader(input)

	header := BuildHeader(s)
	d := Header2Dict(header)

	fmt.Println(d)

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

	file,_ = os.Open("test.qz")

	ch = make (chan byte, 100)

	go ReadFile(file, d, ch)

	out := bufio.NewWriter(os.Stdout)
	for v := range(ch) {
		out.WriteByte(v)
	}
	out.Flush()

	// Output:
}
