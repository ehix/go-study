package main

import (
	"fmt"
	"io"
	"os"
)

// opens a file, and counts number of bytes read in

type ByteCounter int

// Can go into io.Copy() as a dest, as has a Write method,
// therefore provides the capability of a io.Writer.
func (bc *ByteCounter) Write(b []byte) (int, error) {
	// figure out number of bytes passed in, and add to reciever
	l := len(b)
	// have to cast Ints from len to ByteCounter
	*bc += ByteCounter(l)
	return len(b), nil
}

func main() {
	var c ByteCounter
	f1, _ := os.Open("a.txt")
	// f2, _ := os.Create("out.txt")

	f2 := &c

	// Reads from an io.Reader, writes to an io.Writer
	n, _ := io.Copy(f2, f1)

	fmt.Println("copied", n, "bytes")
	fmt.Println(c)
}
