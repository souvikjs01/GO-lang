// GET Request:

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todos struct {
	Id        int    `json:userId`
	Todo      string `json:title`
	Completed bool   `json:completed`
	UserId    int    `json:userId`
}

func getReq() {
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
	// it already a json data
	// data, err1 := ioutil.ReadAll(res.Body)

	// if err1 != nil {
	// 	fmt.Println("error while reading", err1)
	// 	return
	// }
	// fmt.Println(string(data))

	// it is preffered
	var todo Todos
	err3 := json.NewDecoder(res.Body).Decode(&todo)

	if err3 != nil {
		fmt.Println("error while decode", err3)
		return
	}
	fmt.Println("todo:", todo)
}

type Todo struct {
	UserId    int    `json:userId`
	Id        int    `json:id`
	Title     string `json:title`
	Completed bool   `json:completed`
}

func postReq() {
	todo := Todo{30, 2, "go to gym", false}
	// convert it to json data:
	fmt.Println(todo.UserId)
	jsonData, _ := json.Marshal(todo)

	jsonStr := string(jsonData)
	// convert string data to reader:
	jsonReader := strings.NewReader(jsonStr)

	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)

	if err != nil {
		fmt.Println("error while posting", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response: ", string(data))

}

func deleteMethod() {
	const url = "https://jsonplaceholder.typicode.com/todos/1"
	// delete request:
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println("error while deleteing")
		return
	}
	// send the request:
	client := http.Client{}
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println("error sending request", err2)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response: ", string(data))
	fmt.Println("status: ", res.Status)

}

func main() {
	// getReq()

	// postReq()

	// updateMethod()

	deleteMethod()
}
