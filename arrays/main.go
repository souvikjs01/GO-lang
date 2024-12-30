package main

import (
	"fmt"
)

func main() {
	// var arr [5]int
	// for i := 0; i < 5; i++ {
	// 	arr[i] = i + 5
	// }
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("arr[%d]= %d\n", i, arr[i])
	// }

	//------------------------------------------------------------------
	// arr := [5]int{1, 2, 3, 4, 5} // Array with 5 elements initialized
	// arr := [...]int{10, 20, 30} // Compiler infers size (3 elements)
	// for _, v := range arr {
	// 	fmt.Println("values: ", v)
	// }

	//------------------------------------------------------------------
	// arr := [...]int{3, 1, 4, 1, 5}
	// sort.Ints(arr[:]) // Convert array to slice and sort
	// fmt.Println(arr)

	//------------------------------------------------------------------
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Println(search(arr, 30)) // true

}

func search(arr [5]int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
