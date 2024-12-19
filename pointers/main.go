package main

import "fmt"

func temp(num *int) {
	*num = *num + 1
	fmt.Println("address ", num)
	fmt.Println(*num)
}
func main() {
	num := 5
	// fmt.Println("address ", &num) // address

	fmt.Println(num)
	temp(&num)
	fmt.Println(num)
}
