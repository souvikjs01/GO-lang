// variables

package main

import "fmt"

func main() {
	var name string = "souvik"
	var version = 76
	fmt.Println(name)
	fmt.Println(version)

	var money int = 8000
	fmt.Println(money)

	// for float data
	var dim float64 = 87.123
	fmt.Println(dim)

	// for boolean
	var decided bool = true
	fmt.Println(decided)

	// for constant value
	const pi = 3.14
	fmt.Println(pi)

	// important
	person := "i'm jhon from new york, brooklyn"
	fmt.Println(person)

	// variable's first name must be in capital if you wanna use it in other package or files
	var PublicVariable = "data is important"
	// small letter indicate the variable is used only in the host file
	var privateVariable = "data is private"
	fmt.Println(PublicVariable)
	fmt.Println(privateVariable)
}
