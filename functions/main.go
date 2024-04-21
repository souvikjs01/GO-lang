package main

import "fmt"

func simpleFnc() {
	println("simple function")
}

// func add(a, b int) int {
// 	return a + b
// }

func add(a, b int) (result int) {
	result = a + b

	return
}
func name(a, b string) string {
	return a + b
}
func main() {
	fmt.Println("learning function")
	simpleFnc()
	ans := add(3, 4)
	fmt.Println(ans)

	ans2 := name("Alex", "Hales")
	fmt.Println("Hello Mr. " + ans2)
}
