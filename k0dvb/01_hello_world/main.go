package main

// Every program has to have a main function, to direct where the program starts.
// Go is a modular language, so it can be distributed all over the place.

import (
	"fmt"
)

func main() {
	const name, age = "Alex", 31
	fmt.Println("Hello world!")
	fmt.Println(name, "is", age, "years old.")
}
