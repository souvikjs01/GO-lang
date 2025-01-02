package main

import (
	"fmt"
	"sync"
	"time"
)

/*
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y

		case <-quit:
			fmt.Println("quit")
			return
		}

	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)

}

*/

/*
func main() {
	tick := time.Tick(100 * time.Millisecond)  // it will be received every 100 millisec
	boom := time.After(500 * time.Millisecond) // it will be received once at 500 millisec

	for {
		select {
		case <-tick:
			fmt.Println("tick,")
		case <-boom:
			fmt.Println("boom!")
			return
		default:
			fmt.Println("    ,")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
*/

// example 3:

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()

	c.v[key]++
	c.mu.Unlock()
}
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
