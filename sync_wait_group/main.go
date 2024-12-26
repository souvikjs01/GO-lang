package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started working\n", i)
	// some task is going on

	fmt.Printf("worker %d ended\n", i)
}
func main() {
	// fmt.Println("exploring sync wait")

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1) // increment wait group counter
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("worker has completed the task")
}
