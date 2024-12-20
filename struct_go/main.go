package main

import (
	"fmt"
	"time"
)

// type Order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time
// }

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

// func (o *Order) changeStatus(status string) {
// 	o.status = status
// }

// for embedded struct
type Customer struct {
	id    string
	name  string
	phone string
}

func newCustomer(id string, name string, phone string) *Customer {
	customer := Customer{
		id:    id,
		name:  name,
		phone: phone,
	}
	return &customer
}

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	Customer
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
	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{"good", true}

	// fmt.Println(language)

	// embedded struct
	newCustomer := Customer{
		id:    "1",
		name:  "souvik",
		phone: "1234567890",
	}

	myOrder := Order{
		id:       "1",
		amount:   50.50,
		status:   "delivered",
		Customer: newCustomer,
	}

	fmt.Println(myOrder.Customer)

	myOrder.Customer.name = "alex"
	fmt.Println(myOrder.Customer)

}
