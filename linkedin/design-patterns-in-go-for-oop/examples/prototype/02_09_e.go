package main 
import "fmt"


// Prototype interface
type Shape interface {
    print()
    clone() Shape
}


// Concrete prototype 
type Circle struct {
    color string
}

func (c *Circle) print() {
		fmt.Println(c.color + " circle")

}

func (c *Circle) clone() Shape {
		fmt.Println("Cloning...")
		return &Circle{color: c.color}
}



func main() {

	greenCircle := &Circle{color: "Green"}
	blueCircle := &Circle{color: "Blue"}
	redCircle := &Circle{color: "Red"}


	greenClone := greenCircle.clone()
	blueClone := blueCircle.clone()
	redClone := redCircle.clone()


	greenCircle.print()
	blueCircle.print()
	redCircle.print()
	greenClone.print()
	blueClone.print()
	redClone.print()

}