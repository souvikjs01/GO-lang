package main

import "fmt"

// func main() {
// 	fmt.Println("Starting the program...")
// 	result := divide(10, 0)
// 	fmt.Println("Result:", result) // This won't execute
// 	fmt.Println("End of the program.")
// }

// func divide(a, b int) int {
// 	return a / b // Causes a runtime panic if b is 0
// }

// solution:

func main() {
	fmt.Println("Starting the program...")
	result := safeDivide(10, 0)
	fmt.Println("Result:", result)
	fmt.Println("End of the program.")
}

func safeDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			result = -1 // Assign a default value
		}
	}()

	return a / b // This causes a panic if b is 0
}
