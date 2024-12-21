package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// taking user input:

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter ratting of our pizza")
	input, _ := reader.ReadString('\n')

	fmt.Println("pizza ratting given by the user", input)
}
