package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// many return type
func temp() (string, string, string) {
	return "go", "js", "rust"
}

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}
func main() {
	result := add(5, 3)
	fmt.Println(result)

	val1, val2, val3 := temp()
	fmt.Println(val1 + " " + val2 + " " + val3)

	fn := processIt()
	val := fn(6)
	fmt.Println(val)

}
