package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// func worker(ch chan string) {
// 	ch <- "done"
// }
// func main() {
// 	ch := make(chan string)
// 	go worker(ch)
// 	fmt.Println(<-ch)
// }

// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("process number", num)
// 		time.Sleep(1 * time.Second)

// 	}
// }

// func sum(r chan int, a int, b int) {
// 	res := a + b
// 	r <- res
// }

//	func task(done chan bool) {
//		defer func() {
//			done <- true
//		}()
//		fmt.Println("processing...")
//	}

// receive only chan and send only chan
//
//	func emailSender(emailChan <-chan string, done chan<- bool) {
//		defer func() {
//			done <- true
//		}()
//		for email := range emailChan {
//			fmt.Println("sending email to", email)
//			time.Sleep(time.Second)
//		}
//	}

/*
func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
		chan1 <- 85
	}()
	go func() {
		chan2 <- "souvik"
	}()

	// receive:
	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan1", chan2Val)
		}
	}

	// -----------------------------------
	// emailChan := make(chan string, 5)
	// done := make(chan bool)

	// go emailSender(emailChan, done)
	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// fmt.Println("done sending")
	// close(emailChan)
	// <-done
	// emailChan <- "s@gmail.com"
	// emailChan <- "v@gmail.com"
	// emailChan <- "k@gmail.com"
	// emailChan <- "b@gmail.com"
	// fmt.Println("email: ", <-emailChan)
	// fmt.Println("email: ", <-emailChan)
	// <- done

	// --------------------------
	// status := make(chan bool)
	// go task(status)
	// res := <-status
	// fmt.Println("status:", res)

	//---------------------------
	// result := make(chan int)
	// go sum(result, 4, 6)
	// res := <-result  // blocking doesn't require time.sleep()
	// fmt.Println(res)

	//----------------------------
	// numChan := make(chan int)
	// go processNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }
	// numChan <- 5
	// time.Sleep(2 * time.Second)

	//-------------------------
	// messageChan := make(chan string)

	// messageChan <- "hi there"

	// msg := <- messageChan

	// fmt.Println(msg)
}

*/

/*
func fibonacci(n int, c chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
	}
	close(c)
}

func main() {
	// buffered channel:
	c := make(chan int, 10)
	go fibonacci(5, c)

	// sequencially
	for i := range c {
		fmt.Println(i)
	}
}
*/

func randValue() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(100)
}

func main() {
	dataChan := make(chan int)
	go func() {
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				someNum := randValue()
				dataChan <- someNum
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for num := range dataChan {
		fmt.Println("n =", num)
	}
}
