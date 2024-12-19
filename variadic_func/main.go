package main

import "fmt"

func add(nums ...int) int { // if you want any type of data use interface{} instead of int
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {

	result := add(2, 3, 4, 7, 9, 10, 2, 24, 272, 26)
	nums := []int{5, 6, 4, 5, 8, 2, 10}
	fmt.Println(result)
	fmt.Println(add(nums...))

}
