package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time study")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // month date year

	createTime := time.Date(2024, time.December, 1, 23, 20, 45, 0, time.UTC)
	fmt.Println(createTime)
	// formated date:
	fmt.Println(createTime.Format("01-02-2006 Monday"))

}
