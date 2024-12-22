package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://mocki.io/v1/d4867d8b-b5d5-4a48-a4ab-79131b5809b8"

func main() {

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of thpe: %T\n", res)

	defer res.Body.Close() // caller's responsibility to close the connection

	databyte, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		panic(err2)
	}

	content := string(databyte)
	fmt.Println(content)

}
