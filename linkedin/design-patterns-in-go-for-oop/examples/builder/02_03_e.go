package main

import (
	"fmt"
)

// Product
type Burger struct {
	breadType string
	meatType  string
	toppings  string
}

// Builder Interface
type BurgerBuilder interface {
	setBreadType()
	setMeatType()
	setToppings()
	getBurger() Burger
}

// Concrete Builder 1
type RegularBurgerBuilder struct {
	breadType string
	meatType  string
	toppings  string
}

func newRegularBurgerBuilder() *RegularBurgerBuilder {
	return &RegularBurgerBuilder{}
}

func (b *RegularBurgerBuilder) setBreadType() {
	b.breadType = "Sesame"
}

// Set the Meat Type
func (b *RegularBurgerBuilder) setMeatType() {
	b.meatType = "Beef"
}

// Set the Toppings
func (b *RegularBurgerBuilder) setToppings() {
	b.toppings = "Lettuce, Tomato, Bacon, and Cheese"
}

func (b *RegularBurgerBuilder) getBurger() Burger {
	return Burger{
		breadType: b.breadType,
		meatType:  b.meatType,
		toppings:  b.toppings,
	}
}

// Concrete Builder 2
type VeganBurgerBuilder struct {
	breadType string
	meatType  string
	toppings  string
}

func newVeganBurgerBuilder() *VeganBurgerBuilder {
	return &VeganBurgerBuilder{}
}

func (b *VeganBurgerBuilder) setBreadType() {
	b.breadType = "Gluten Free"
}

// Set the Meat Type
func (b *VeganBurgerBuilder) setMeatType() {
	b.meatType = "Black Bean"
}

// Set the Toppings
func (b *VeganBurgerBuilder) setToppings() {
	b.toppings = "Lettuce and Tomato"
}

func (b *VeganBurgerBuilder) getBurger() Burger {
	return Burger{
		breadType: b.breadType,
		meatType:  b.meatType,
		toppings:  b.toppings,
	}
}

// Director
type Director struct {
	builder BurgerBuilder
}

func newDirector(b BurgerBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b BurgerBuilder) {
	d.builder = b
}

func getBuilder(builderType string) BurgerBuilder {
	if builderType == "regular" {
		return newRegularBurgerBuilder()
	}

	if builderType == "vegan" {
		return newVeganBurgerBuilder()
	}

	return nil
}

func (d *Director) buildBurger() Burger {
	d.builder.setBreadType()
	d.builder.setMeatType()
	d.builder.setToppings()
	return d.builder.getBurger()
}

func main() {
	regularBurgerBuilder := getBuilder("regular")
	veganBurgerBuilder := getBuilder("vegan")

	// Initialize the director
	director := newDirector(regularBurgerBuilder)

	// Create the regular burger
	regularBurger := director.buildBurger()

	fmt.Printf("Regular Burger Bread Type: %s\n", regularBurger.breadType)
	fmt.Printf("Regular Burger Meat Type: %s\n", regularBurger.meatType)
	fmt.Printf("Regular Burger Toppings: %s\n", regularBurger.toppings)

	// Build the vegan burger
	director.setBuilder(veganBurgerBuilder)
	veganBurger := director.buildBurger()

	fmt.Printf("Vegan Burger Bread Type: %s\n", veganBurger.breadType)
	fmt.Printf("Vegan Burger Meat Type: %s\n", veganBurger.meatType)
	fmt.Printf("Vegan Burger Toppings: %s\n", veganBurger.toppings)

}
