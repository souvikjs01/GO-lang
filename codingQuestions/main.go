package main

/*
1. write a go routines that print 1-10 with a delay of 1 second


func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("n =", i+1)
		}
	}()

	wg.Wait()
}
*/

/*
2. write a function to reverse a string in go

func main() {
	str := "abcdef"
	fmt.Println("original string", str)
	strRunes := []rune(str)
	i := 0
	j := len(strRunes) - 1
	for i < j {
		strRunes[i], strRunes[j] = strRunes[j], strRunes[i]
		i++
		j--
	}
	revStr := string(strRunes)
	fmt.Println("reverse string", revStr)
}
*/

/**
3. write a function to find duplicates of an slice

func findDuplicates(nums []int) []int {
	hashing := make(map[int]int)
	for _, n := range nums {
		hashing[n]++
	}
	var dups []int

	for key, val := range hashing {
		if val > 1 {
			dups = append(dups, key)
		}
	}
	return dups
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 2, 4, 6, 7, 1}

	duplicates := findDuplicates(nums)

	fmt.Println("Duplicates numbers are", duplicates)
}
*/

/*
4. write a recursion function to find the factorial of a number

func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	n := 5
	ans := fact(n)
	fmt.Println("factorial", ans)
}
*/

/*
5. example for error handling in go

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("can not divide by zero")
	}
	return a / b, nil
}
func main() {
	res, err := divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", res)
}

*/
