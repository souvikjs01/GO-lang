package main

import "fmt"

func main() {
	fmt.Println("learning go switch cases")
	// day := "Sunday"
	// day := "Saturday"
	// day := "Wednessday"

	// switch day {
	// case "Sunday", "Saturday":
	// 	fmt.Println("it is the weekend")
	// case "Monday":
	// 	fmt.Println("it is the starting day")
	// case "Friday":
	// 	fmt.Println("it is the pre weekend day")
	// default:
	// 	fmt.Println("just a normal day")
	// }

	//------------------------------
	// The fallthrough keyword forces the execution of the next case, even if it doesn’t match.
	// switch num := 2; num {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// 	fallthrough
	// case 3:
	// 	fmt.Println("Three") // Executed because of fallthrough
	// default:
	// 	fmt.Println("Default")
	// }

	//-----------------------------

	describe("souvik")
	describe(true)
	describe(8)
}

func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("It's an int: %d\n", v)
	case string:
		fmt.Printf("It's a string: %s\n", v)
	case bool:
		fmt.Printf("It's a boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type\n")
	}
}
