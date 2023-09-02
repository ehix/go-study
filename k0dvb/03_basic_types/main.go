package main

import (
	"fmt"
	"os"
)

// when read from file:
// go run main.go < nums.txt

func main() {
	// default init to zero, var keyword is used for this.
	var sum float64
	var n int

	for {
		var val float64
		// reading in value from cmd line
		// passing address to float with &
		// signal end of input ctrl + D
		if _, err := fmt.Fscanln(os.Stdin, &val); err != nil {
			// e.g. EOF error, or anything else
			break
		}

		sum += val
		n++
	}

	if n == 0 {
		// print out to cmd line
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	// ints have to be cast explicitly
	fmt.Printf("the average is: %v\n", sum/float64(n))
}
