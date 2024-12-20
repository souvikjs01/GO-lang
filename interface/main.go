package main

import "fmt"

type Shape interface {
	Area() float32
}

type Rectangle struct {
	Width  float32
	Height float32
}

func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return 3.14 * c.Radius * c.Radius
}

// func printArea(s Shape) {
// 	fmt.Println("Area: ", s.Area())
// }

// embedded interface
type Object interface {
	Name() string
	Shape
}

func (r Rectangle) Name() string {
	return "rectangle"
}
func printObject(o Object) {
	fmt.Println("Area: ", o.Area(), "Name: ", o.Name())
}
func main() {
	// rect := Rectangle{Height: 6, Width: 3}
	// cir := Circle{Radius: 7}

	// shapes := []Shape{rect, cir}

	// for _, sp := range shapes {
	// 	printArea(sp)
	// }

	// embedded interface:
	rec := Rectangle{Height: 10, Width: 5}
	printObject(rec)
}
