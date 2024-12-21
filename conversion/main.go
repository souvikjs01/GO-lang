package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza shop")
	fmt.Println("please rate our pizza beteen 1-5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString(('\n'))
	fmt.Println("Thanks for the rating", input)

	// problem is input value is the type of string, we can not perform any operation to it
	// numRating := input + 1

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating", numRating+1)
	}
}
