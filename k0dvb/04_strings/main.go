package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func main() {
// 	s := "København"
// 	fmt.Printf("type: %8T value: %[1]v len: %d\n", s, len(s))
// 	a := []rune(s)
// 	fmt.Printf("type: %8T value: %[1]v len: %d\n", a, len(a))
// 	b := []byte(s)
// 	fmt.Printf("type: %8T value: %[1]v len: %d\n", b, len(b))
// }

func main() {
	// index an arg that's not there, so fail.
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	// buffered io tool around the input, if input large won't read it all.
	scan := bufio.NewScanner(os.Stdin)
	// scan.Scan will return true, acts as while loop.
	for scan.Scan() {
		s := strings.Split(scan.Text(), old) // returns a slice
		t := strings.Join(s, new)            // takes a slice

		fmt.Println(t)
	}
	// go run main.go alex neil < test.txt
}
