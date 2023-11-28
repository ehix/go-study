package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"slices"
	"sort"

	"github.com/google/uuid"
)

type Person struct {
	UUID string
	Age  int
}

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []Person

// These methods are required to implement sort.Interface.
// The Interface interface requires three methods: Len, Less, and Swap.
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func generateRandomFamily() []Person {
	var family []Person
	for i := 0; i < 5; i++ {
		uuid := uuid.NewString()[:8]
		age := rand.Intn(100)
		family = append(family, Person{uuid, age})
	}
	return family
}

func sortWithCustomComparator() {
	family := generateRandomFamily()
	var a, b []Person = family, family

	// Sort by Age.
	// sort.Slice takes a slice and a function that compares two elements.
	sort.Slice(a, func(i, j int) bool {
		return a[i].Age < a[j].Age
	})
	fmt.Println("sort.Slice:      ", a)

	// Sort by Age, keeping original order of equal elements (if there is any).
	// sort.SliceStable is like sort.Slice, but it keeps the original order of equal elements.
	sort.SliceStable(b, func(i, j int) bool {
		return b[i].Age < b[j].Age
	})
	fmt.Println("sort.SliceStable:", b)
}

// Use the generic sort.Sort and sort.Stable functions to sort slices of custom types.
// Any sort any collection that implements the sort.Interface interface.
// See above for ByAge type and methods.
func sortCustomTypes() {
	family := generateRandomFamily()
	sort.Sort(ByAge(family))
	fmt.Println("sort.Sort:", family)
}

// Use sort.Strings to sort a map by the key.
func sortMapsByKeys() {
	m := map[string]int{"Cecil": 1, "Alice": 2, "Bob": 3}
	fmt.Println("init map:", m) // <!> this seems already ordered?

	keys := make([]string, 0, len(m))
	for k := range m { // Loop over map keys.
		keys = append(keys, k) // Add key to slice.
	}

	sort.Strings(keys)
	fmt.Println("map keys:", keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}

// See https://yourbasic.org/golang/how-to-sort-in-go/
func yourBasicSort() {
	sortWithCustomComparator()
	sortCustomTypes()
	sortMapsByKeys()
}

func somethingAboutRadix() {
}

// Compare two slices of the same type.
func compareSlices[T cmp.Ordered](a []T, b []T) {
	if len(a) != len(b) {
		panic("slices are not the same length")
	}
	for i := range a {
		if a[i] != b[i] {
			panic(fmt.Sprintf("slices differ at index %d", i))
		}
	}
}

// Helper function to print the type and value of a slice before and after sorting.
// Uses type parameterization to sort a slice of any ordered type.
// Type parameters of a function appear in square brackets before the function name.
// cpm.Ordered is an interface that requires a type to implement the Less method.
// The Less method is used to determine the order of two values.
// See https://pkg.go.dev/cmp#Ordered
func sortWithSlices[T cmp.Ordered](s []T) []T {
	// The slices package can be used to check if a slice is already in order.
	fmt.Printf("Type: %10[1]T - Pre:  %[1]v - Sorted: %v\n", s, slices.IsSorted(s))
	// Use slices package to implement sorting for built-in and user defined types.
	slices.Sort(s)
	fmt.Printf("Type: %10[1]T - Post: %[1]v - Sorted: %v\n", s, slices.IsSorted(s))
	return s
}

// Example of using the slices package to sort slices of different types.
// See https://gobyexample.com/sorting
func goByExample() {
	// 1. Strings:
	// Create a slice of unordered strings.
	strs := []string{"c", "a", "b"}
	// Use helper function to sort the slice and print results.
	a := sortWithSlices(strs)
	// Compare the above method to the built-in sort function.
	b := []string{"c", "a", "b"} // same slice as above
	sort.Strings(b)              // different sort function, also see sort.Ints and sort.Floats.
	compareSlices(a, b)          // will panic if the slices are not the same

	// 2. Ints:
	// Sorting functions are generic, and work for any ordered built-in type.
	ints := []int{7, 2, 4}
	sortWithSlices(ints)

	// 3. Floats:
	// So that's largely, ints, floats, and strings.
	floats := []float32{7.1, 2.2, 4.3}
	sortWithSlices(floats)
}

func main() {
	goByExample()
	yourBasicSort()
}
