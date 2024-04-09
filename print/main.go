package main

import "fmt"

func main() {
	age := 25
	name := "Alex"
	height := 5.821749

	fmt.Println("age: ", age, "name: ", name, "height: ", height)
	fmt.Println("hi alex")
	fmt.Printf("age is: %d\n", age)
	fmt.Printf("height is: %.2f\n", height)
	// to check the type of variable:-
	fmt.Printf("type of name is: %T\n", name)
	fmt.Printf("type of age is: %T\n", height)
	fmt.Printf("Name: %s,  Age: %d,  Height: %.3f\n", name, age, height)
}
