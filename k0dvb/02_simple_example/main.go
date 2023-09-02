package main

import (
	"example/simple/hello"
	"fmt"
	"os"
)

// func main() {
// 	// os.Args[0] will be the name of the program.
// 	if len(os.Args) > 1 {
// 		fmt.Println(hello.Say(os.Args[1]))
// 	} else {
// 		fmt.Println(hello.Say("world!"))
// 	}
// }

func main() {
	// if nothing passed, will be a slice of len = 0
	fmt.Println(hello.Say(os.Args[1:]))
}
