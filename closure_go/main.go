package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		return count + 1
	}
}
func main() {
	increment := counter()
	fmt.Println(increment())
}
