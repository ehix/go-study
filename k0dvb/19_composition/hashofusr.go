package main

import (
	"fmt"
	"path/filepath"
)

type Pair struct {
	Path string
	Hash string
}

// Methods are also promoted.
func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

// func Filename(p Pair) string {
// 	return filepath.Base(p.Path)
// }

func (p Pair) Filename() string {
	return filepath.Base(p.Path)
}

type Filenamer interface {
	Filename() string
}

// Pair is embedded in PairWithLength.
// Path and Hash are "promoted to PairWithLength".
type PairWithLength struct {
	Pair
	Length int
}

// If the types have the same method, it's own will be prefered.
func (p PairWithLength) String() string {
	return fmt.Sprintf("Hash of %s is %s; length %v", p.Path, p.Hash, p.Length)
}

type Fizgig struct {
	*PairWithLength
	Broken bool
}

func main() {
	p := Pair{"/usr", "0xfdfe"}
	// Literals have to explicity name the emb type.
	pwl := PairWithLength{Pair{"/usr", "0xfdfe"}, 121}
	// Not.pwl.x.Path, because Path and Length are at the same
	// level in PairWithLength as they are in Pair.
	// fmt.Println(pwl.Path, pwl.Length)

	fmt.Println(p)
	fmt.Println(pwl)

	// This won't run, bc a PairWithLength cannot be passed as a param to Pair.
	// Even though PairWithLength has a Pair within it.
	// The Pair has to be passed with dot notation.
	// fmt.Println(Filename(p)) 		// allowed
	// fmt.Println(Filename(pwl.Pair))	// not allowed

	// This can be fixed by implementing an interface
	fmt.Println(p.Filename())
	fmt.Println(pwl.Filename())
	// This will also work, bc the Filename method is on Pair,
	// which is promoted into PairWithLength.
	var fn Filenamer = PairWithLength{Pair{"/usr/lib", "0xdead"}, 999}
	fmt.Println(fn.Filename())

	// Can embed a pointer to another type; promotion works the same way.
	fg := Fizgig{&PairWithLength{Pair{"/usr", "Oxfdfe"}, 121}, false}
	fmt.Println(fg.Filename())

}
