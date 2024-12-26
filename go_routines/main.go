package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("say hello to everyone")
	time.Sleep(1 * time.Second)
	fmt.Println("say hello to everyone after 1 sec")
}
func sayHi() {
	fmt.Println("say hi to everyone")
	time.Sleep(1 * time.Second)
	fmt.Println("say hi after 1 sec")
}
func main() {
	fmt.Println("Learning most important topic of GoLang:")
	go sayHello()
	go sayHi()

	// wait a moment to allow goroutine to finish:
	time.Sleep(1 * time.Second)
}
