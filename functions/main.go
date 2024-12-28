package main

import (
	"fmt"
)

/*
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
}*/

//-----------------------------------------

// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("division by zero")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	result, err := divide(10, 0)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Result:", result)
// }

// -----------------------------------------------
// Anonymous function:
// func main() {
// 	a, b := func(x, y int) (int, int) {
// 		return x + y, x * y
// 	}(5, 7)

// 	fmt.Println("sum:", a)
// 	fmt.Println("product:", b)
// }

//--------------------------------------------------
// Chaining Multiple Return Values

func add(a, b int) (int, int) {
	return a + b, a - b
}

func multiply(sum, diff int) int {
	return sum * diff
}

func main() {
	result := multiply(add(7, 5))
	fmt.Println("Result:", result)
}
