package main

import "fmt"

func print_slice(s []int) {
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(s), cap(s), s, s == nil)
}

func nil_vs_empty() {
	// nil slice, takes default value.
	var s []int
	print_slice(s)
	// > 0, 0, []int,  true, []int(nil)

	t := []int{}
	print_slice(t)
	// > 0, 0, []int, false, []int{}

	u := make([]int, 5)
	print_slice(u)
	// > 5, 5, []int, false, []int{0, 0, 0, 0, 0}

	v := make([]int, 0, 5) // (len, cap)
	print_slice(v)
	// > 0, 5, []int, false, []int{}
}

func len_vs_cap() {
	a := [3]int{1, 2, 3}
	b := a[0:1]
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	c := b[0:2] //WTF?
	fmt.Println("c =", c)

	// > a = [1 2 3]
	// > b = [1]
	// > c = [1 2]

	// TIB: b and c pick up the underlying cap of a
	fmt.Println("b: len =", len(b), "cap =", cap(b))
	fmt.Println("c: len =", len(c), "cap =", cap(c))
	// > b: len = 1 cap = 3
	// > c: len = 2 cap = 3

	// different slice operator, control len and cap
	d := a[0:1:1]
	fmt.Println("d =", d)
	fmt.Println("d: len =", len(d), "cap =", cap(d))
	// > d = [1]
	// > d: len = 1 cap = 1
}

func appends_to_slice() {
	a := [...]int{1, 2, 3}
	// slices are alias' to some underlying array
	b := a[0:1]
	c := b[0:2] // WTF

	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)
	// all have same addr/point to a
	// > a[0xc00001a198] = [1 2 3]
	// > b[0xc00001a198] = [1]
	// > c[0xc00001a198] = [1 2]

	c = append(c, 5) // a gets overwritten when append to c
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)
	// > a[0xc00001a198] = [1 2 5]
	// > c[0xc00001a198] = [1 2 5]

	c = b[0:2:2]      // using 3 slice op to contr the cap
	c = append(c, 10) // a gets overwritten when append to c
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)
	// different addr, a[2] isn't overwritten again
	// > a[0xc00001a198] = [1 2 5]
	// > c[0xc00001e220] = [1 2 10]
}

func main() {
	// nil_vs_empty()
	// len_vs_cap()
	appends_to_slice()
}
