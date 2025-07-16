package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
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

*/

// example :2
/*
type Message struct {
	chats   []string
	friends []string
}

func main() {
	now := time.Now()
	id := getUserIdByName("john")
	fmt.Println(id)

	wg := &sync.WaitGroup{}
	ch := make(chan *Message, 2)

	wg.Add(2)

	go getUserChats(id, ch, wg)
	go getUserFriends(id, ch, wg)

	wg.Wait()
	close(ch)

	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println(time.Since(now))
}

func getUserIdByName(name string) string {
	time.Sleep(2 * time.Second)

	return fmt.Sprintf("%s-2", name)
}

func getUserChats(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(4 * time.Second)

	ch <- &Message{
		chats: []string{
			"hello",
			"there",
			"jane",
			"you are",
			"great",
		},
	}
	wg.Done()
}

func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)

	ch <- &Message{
		friends: []string{
			"emma",
			"alex",
			"jane",
			"armin",
			"sasha",
		},
	}
	wg.Done()
}
*/

func randHun() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(100)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n := randHun()
			fmt.Println("i =", i+1, "n =", n)
		}()
	}
	wg.Wait()
	fmt.Println("Process complete...")
}
