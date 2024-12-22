package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// its call method
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
func (u User) NewEmail() {
	u.Email = "testgo@gmail.com"
	fmt.Println("Email of the user is:", u.Email)
}
func main() {
	alex := User{"Alex", "alex@gmail.com", false, 23}

	alex.GetStatus()
	alex.NewEmail()
}
