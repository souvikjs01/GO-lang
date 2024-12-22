package main

import "fmt"

func main() {
	// defer works as LIFO (Last In First Out)

	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

	fmt.Println("hello")
	myDefer(5)
}

func myDefer(n int) {
	for i := 0; i < n; i++ {
		defer fmt.Println(i)
	}
}
