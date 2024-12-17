package main

import "fmt"

func main() {
	// fmt.Println("hello")

	var age uint8 = 78
	fmt.Println("my age is", age)

	var price float32 = 90.6
	fmt.Println("price is ", price)

	// constants
	// const age = 59
	// age = 60 , you cant do this
	// multiple const:
	const (
		port     = 5050
		host     = "localhost"
		protocol = "http"
	)
	fmt.Print(protocol, "://", host, ":", port)

}
