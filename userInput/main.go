package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey, what's your name")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello, Mr.", name)

	// right way to take onputs:-
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, Mr.", name)

}
