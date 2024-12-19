package main

import (
	"fmt"
	"slices"
)

// slice
func main() {
	// var nums []int
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(nums == nil)

	// var nums = make([]int, 2)
	// var nums = make([]int, 0, 5) // 0 is the initial capacity:

	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// nums = append(nums, 7)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// nums := []int{}

	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// ---copy-----
	var nums = make([]int, 0, 5)
	nums = append(nums, 4)
	nums = append(nums, 3)
	var num2 = make([]int, len(nums))

	// copy function:
	copy(num2, nums)
	fmt.Println(nums, num2)

	// --------slice operator---------
	// var arr = []int{1, 2, 3, 4}
	// fmt.Println(arr[0:3])
	// fmt.Println(arr[2:4])

	// slice compare
	var arr1 = []int{1, 2, 6}
	var arr2 = []int{1, 2, 3}
	var ans = slices.Equal(arr1, arr2)
	fmt.Println(ans)

	// 2d array:
	var _arr = [][]int{{3, 4}, {8, 9}}
	fmt.Println(_arr)
}