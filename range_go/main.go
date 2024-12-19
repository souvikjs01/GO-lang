package main

import "fmt"

func main() {

	// nums := []int{6, 7, 9, 10}

	// sum := 0

	// for indx, n := range nums {
	// 	fmt.Println("index :", indx)
	// 	fmt.Println("num :", n)
	// 	sum += n
	// }
	// fmt.Println("sum of the slices is: ", sum)

	// iterate through map using range:
	// m2 := map[string]int{"souvik": 1, "alex": 2, "cathrina": 3, "deeps": 4}
	// for k, v := range m2 {
	// 	fmt.Println(k, v)
	// }

	for i, c := range "abcd" {
		fmt.Println(i, " ", c)
		fmt.Println(i, " ", string(c))
	}
}
