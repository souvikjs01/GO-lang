package main

import "fmt"

// func counter() func() int {
// 	var count int = 0

// 	return func() int {
// 		return count + 1
// 	}
// }
// func main() {
// 	increment := counter()
// 	fmt.Println(increment())
// }

func adder(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	add5 := adder(5)
	fmt.Println(add5(3))

	add10 := adder(10)
	fmt.Println(add10(5))
}
