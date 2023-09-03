package main

import "fmt"

func main() {
	a, b := 12, 345
	c, d := 1.2, 3.45
	fmt.Printf("%d %d\n", a, b)
	fmt.Printf("%x %x\n", a, b)   // hex
	fmt.Printf("%X %X\n", a, b)   // uppercase hex
	fmt.Printf("%#x %#x\n", a, b) // hex with '0x' prefix
	fmt.Printf("%f %.3f\n", c, d) // spec decimal place

	// format tables
	fmt.Printf("|%6d|%6d|\n", a, b)    // give char width
	fmt.Printf("|%06d|%06d|\n", a, b)  // give leading zero
	fmt.Printf("|%-6d|%-6d|\n", a, b)  // left justify
	fmt.Printf("|%6f|%-6.2f|\n", c, d) // won't truncate if x > 6

	// slice, similar for maps
	e := [3]rune{'a', 'b', 'c'} // has to be single quotes
	fmt.Printf("%T\n", e)       // type
	fmt.Printf("%q\n", e)       // quoted ['a' 'b' 'c']
	fmt.Printf("%v\n", e)       // ints values for chars [97 98 99]
	fmt.Printf("%#v\n", e)      // type and value [3]int32{97, 98, 99}
}
