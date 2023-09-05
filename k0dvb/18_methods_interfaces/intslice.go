package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

func (is IntSlice) String() string {
	var strs []string

	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}
	// Custom string function for IntSlice
	return "[" + strings.Join(strs, ";") + "]"
}

func main() {
	// Can be created like []int as it's underlying type.
	var v IntSlice = []int{1, 2, 3}
	// s is an interface type, meaning anything can be assigned
	var s fmt.Stringer = v

	// Can iter like the base-type
	for i, x := range v {
		fmt.Printf("%d: %d\n", i, x)
	}
	// The type is IntSlice []int
	fmt.Printf("%T %[1]v\n", v)
	// At runtime, look at the type the interface is holding onto (e.g. v's type)
	fmt.Printf("%T %[1]v\n", s)
}
