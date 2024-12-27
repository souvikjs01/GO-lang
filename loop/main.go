package main

import "fmt"

func main() {
	// s := []int{5, 6, 7, 8, 9, 10}
	// fmt.Println("original slice:", s)

	// reverseSlice(s)
	// fmt.Println("reverse slice:", s)

	// ---------------------------------
	// str := "learning go"
	// printStr(str)

	//----------------------------------
	matrix := [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{7, 8, 9},
	}
	printMatrix(matrix)

}
func printMatrix(m [][]int) {
	for i, row := range m {
		for j, val := range row {
			fmt.Printf("matrix[%d][%d] = %d", i, j, val)
		}
		fmt.Println()
	}
}

func printStr(s string) {
	for idx, val := range s {
		fmt.Printf("%dth index char is %c\n", idx, val)
	}
}

// func reverseSlice(s []int) {
// 	n := len(s)
// 	// for i := 0; i < n/2; i++ {
// 	// 	s[i], s[n-1-i] = s[n-1-i], s[i]
// 	// }

// 	i := 0
// 	j := n - 1
// 	for i < j {
// 		s[i], s[j] = s[j], s[i]
// 		i++
// 		j--
// 	}
// }
