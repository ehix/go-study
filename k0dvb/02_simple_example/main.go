package main

import (
	"example/simple/hello"
	"fmt"
	"os"
)

func main() {
	// os.Args[0] will be the name of the program.
	if len(os.Args) > 1 {
		fmt.Println(hello.Say(os.Args[1]))
	} else {
		fmt.Println(hello.Say("world!"))
	}
}
