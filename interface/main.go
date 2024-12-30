package main

import "fmt"

/*
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

*/

//------------------------------------------------------
// type Speaker interface {
// 	Speak() string
// }

// type Person struct {
// 	Name string
// }

// func (p Person) Speak() string {
// 	return "Hello, I am " + p.Name
// }

// func main() {
// 	var s Speaker
// 	s = Person{Name: "Alice"}
// 	fmt.Println(s.Speak())
// }

//----------------------------------
//empty interface:
// func describe(i interface{}) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// func main() {
// 	describe(42)
// 	describe("hello")
// 	describe([]int{1, 2, 3})
// }

//----------------------------------

// type Reader interface {
// 	Read() string
// }

// type Writer interface {
// 	Write(s string)
// }

// type ReadWrite interface {
// 	Reader
// 	Writer
// }

// type Device struct{}

// // Implement Read method
// func (d Device) Read() string {
// 	return "reading"
// }

// // Implement Write method
// func (d Device) Write(s string) {
// 	fmt.Println("writing:", s)
// }
// func main() {
// 	var rw ReadWrite = Device{} // Now works
// 	fmt.Println(rw.Read())
// 	rw.Write("data")
// }

//-------------------------------------------------
type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type AllInOne struct{}

func (AllInOne) Print() {
	fmt.Println("Printing...")
}

func (AllInOne) Scan() {
	fmt.Println("Scanning...")
}

func main() {
	var p Printer = AllInOne{}
	var s Scanner = AllInOne{}

	p.Print()
	s.Scan()
}
