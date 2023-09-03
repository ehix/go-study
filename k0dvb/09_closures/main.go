package main

import "fmt"

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	f := fib() // is calling the inner function,
	// and using a, b each time it's called.
	// As long as f exists, a and b exist.
	for x := f(); x < 100; x = f() {
		fmt.Println(x)
	}

	g, h := fib(), fib()
	fmt.Println("g:", g(), g(), g(), g(), g())
	fmt.Println("h:", h(), h(), h(), h(), h())
	// > g: 1 2 3 5 8
	// > h: 1 2 3 5 8
	// g and h get different a and b variables, and are closed over
	// by the anon function. tf, in the heap, there's distinct a and b vars.

	s := make([]func(), 4)
	for i := 0; i < 4; i++ {
		i2 := i // closure capture
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i2, &i2)
		}
	}
	for i := 0; i < 4; i++ {
		s[i]()
	}
	// > 0 @ 0xc0000120b8
	// > 1 @ 0xc0000120d0
	// > 2 @ 0xc0000120d8
	// > 3 @ 0xc0000120e0
}
