package main

import "fmt"

func do(b [3]int) int { // takes an array
	// pass by value, is a copy and isn't modified
	b[0] = 0
	return b[1]
}

func do_again(b []int) int { // takes a slice
	// pass by ref, is an address to original and is modified.
	b[0] = 0
	fmt.Printf("b@ %p\n", b) // >> b@ 0xc00001a1b0
	return b[1]
}

func add_to_map(m map[int]int) {
	i := len(m) + 1
	(m)[i] = i

}

func overwrite_map(m *map[int]int) {
	// deference pointer with (*m)
	*m = make(map[int]int)
	// doesn't need return
}

func main() {
	a := [3]int{1, 2, 3}
	v1 := do(a)
	fmt.Println(a, v1) // >> [1 2 3] 2

	b := []int{1, 2, 3}
	fmt.Printf("b@ %p\n", b) // >>  b@ 0xc00001a1b0
	v2 := do_again(b)
	fmt.Println(b, v2) // >> [0 2 3] 2

	m := map[int]int{1: 1, 2: 2, 3: 3}
	add_to_map(m)
	fmt.Println(m)

	overwrite_map(&m) // & passes address
	fmt.Println(m)
}
