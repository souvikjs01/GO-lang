package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("learning JSON")

	person := Person{Name: "alex", Age: 22, IsAdult: true}
	fmt.Println("Person data: ", person)

	// convert into json: called encode or marshal:
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error marshalling ", err)
	}
	fmt.Println(string(jsonData))

	// unmarshelling:
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("error unmarshelling", err)
	}

	fmt.Println("person data after unmarshelling:", personData)

}
