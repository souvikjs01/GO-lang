// GET Request:

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todos struct {
	Id        int    `json:userId`
	Todo      string `json:title`
	Completed bool   `json:completed`
	UserId    int    `json:userId`
}

func main() {
	fmt.Println("learning Get request")

	res, err := http.Get("https://dummyjson.com/todos/1")

	if err != nil {
		fmt.Println("error while fetching", err)
		return
	}

	// close the network resourses:
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("error in getting response", res.Status)
		return
	}

	// data, err1 := ioutil.ReadAll(res.Body)

	// if err1 != nil {
	// 	fmt.Println("error while reading", err1)
	// 	return
	// }
	// fmt.Println(string(data))

	var todo Todos
	err3 := json.NewDecoder(res.Body).Decode(&todo)

	if err3 != nil {
		fmt.Println("error while decode", err3)
		return
	}
	fmt.Println("todo:", todo)

}
