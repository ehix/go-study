package main

import (
	"fmt"
	"sort"
)

// Example of composition

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

// Methods only applicable to Organs (plural, not singular)
func (s Organs) Len() int      { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Can sort by either criteria, Name or Weight.
// Len and Swap are promoted in.
type ByName struct{ Organs }
type ByWeight struct{ Organs }

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main() {
	s := []Organ{{"brain", 1340}, {"liver", 1494}, {"spleen", 162}, {"pancreas", 132}, {"heart", 209}}
	fmt.Println(s)

	// Sorts inline, and in ascending order.
	sort.Sort(ByWeight{s})
	fmt.Println(s)

	sort.Sort(ByName{s})
	fmt.Println(s)

}
