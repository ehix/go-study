package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p *Point) DistanceP(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func main() {
	p := Point{1, 1}
	q := Point{5, 4}

	// Will also use Point{1, 1}
	distanceFromP := p.Distance
	fmt.Printf("%T\n", distanceFromP)              // func(main.Point) float64
	fmt.Printf("%T\n", Point.Distance)             // func(main.Point, main.Point) float64
	fmt.Println(p.Distance(q) == distanceFromP(q)) // true

	// More like a normal closure
	distanceFromPP := p.DistanceP
	p = Point{2, 2} // p is changed
	// distanceFromPP takes the value of p right now through *
	fmt.Println(distanceFromPP(q)) // 3.6055512754639896
	p = Point{3, 3}                // p is changed again
	fmt.Println(distanceFromPP(q)) // 2.23606797749979

}
