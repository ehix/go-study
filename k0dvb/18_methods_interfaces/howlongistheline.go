package main

import (
	"fmt"
	"math"
)

// What's the distance between the begin and end points?

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

type Path []Point

// Takes Line by value because it's not being modified
func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

func (p Path) Distance() (sum float64) {
	// Assume between two points is a line
	for i := 1; i < len(p); i++ {
		// p[i-1], p[i] = from last point to 'this' point
		// Line literal doesn't have an address,
		// but can be created, used, and dispatched on the fly
		sum += Line{p[i-1], p[i]}.Distance()
	}
	return sum
}

// Doesn't return a line so needs a pointer receiver *Line
// Also means ScaleBy cannot be called on a literal, without an address
func (l *Line) ScaleBy(f float64) {
	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
	l.End.Y += (f - 1) * (l.End.Y - l.Begin.Y)
}

// This will work on a literal
func (l Line) ScaleByAlt(f float64) Line {
	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
	l.End.Y += (f - 1) * (l.End.Y - l.Begin.Y)
	return Line{l.Begin, Point{l.End.X, l.End.Y}}
}

// Convention in Go to take adj and turn into noun
type Distancer interface {
	Distance() float64
}

// Both Line and Point are instances of types that satisfy the Distancer interface
// bc, they both have Distance() methods
func PrintDistance(d Distancer) {
	fmt.Println(d.Distance())
}

func main() {
	// Feilds Begin, End are Point literals, in the Line literal
	side := Line{Point{1, 2}, Point{4, 6}}

	// Is a slice of points, so doesn't need Path name for literal
	perimeter := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}

	// Scaled using a * rec
	side.ScaleBy(3)
	// Printed using the Distancer interface,
	// as both side and parimeter have Distance() methods to satify the interface.
	PrintDistance(side)
	PrintDistance(perimeter)

	// Scaled using a val rec
	fmt.Println(Line{Point{1, 2}, Point{4, 6}}.ScaleByAlt(2).Distance())
}
