package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

// constructor:
func newOrder(id string, amount float32, status string) *Order {
	// initial setup goes here...
	myOrder := Order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}
func (o *Order) changeStatus(status string) {
	o.status = status
}

func main() {
	// order := Order{
	// 	id:     "1",
	// 	amount: 50.50,
	// 	status: "received",
	// }

	// order.createdAt = time.Now()
	// fmt.Println("Order struct ", order)
	// fmt.Println(order.status)

	// myOrder := Order{
	// 	id:     "1",
	// 	amount: 50.50,
	// 	status: "delivered",
	// }
	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder)

	// myOrd := newOrder("3", 30.50, "recieved")
	// fmt.Println(myOrd)

	// another way to initialize struct
	language := struct {
		name   string
		isGood bool
	}{"good", true}

	fmt.Println(language)
}
