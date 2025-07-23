package main

import (
	"fmt"
	"sync"
)

type safeCount struct {
	v  map[string]int
	mu sync.Mutex
}

func (c *safeCount) Inc(key string, wg *sync.WaitGroup) {
	c.mu.Lock()
	defer c.mu.Unlock()
	defer wg.Done() // defer works as LIFO
	c.v[key]++
}

func (c *safeCount) value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.v[key]
}

func main() {
	var wg sync.WaitGroup

	c := safeCount{
		v: make(map[string]int),
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go c.Inc("alex", &wg)
	}

	wg.Wait()

	fmt.Println("Current value of the key", c.value("alex"))
}
