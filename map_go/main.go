package main

import (
	"fmt"

	"maps"
)

func main() {
	// create map

	m := make(map[string]string)
	m["name"] = "souvik"
	m["address"] = "remote"
	m["role"] = "fullstack devs"

	fmt.Println(m["role"])

	m2 := map[string]int{"souvik": 1, "alex": 2, "cathrina": 3, "deeps": 4}
	fmt.Println(m2)
	fmt.Println(len(m2))

	// to check whether the key exist or not
	v, ok := m2["alex"]

	if ok {
		fmt.Println("all ok")
		fmt.Println("value ", v)
	} else {
		fmt.Println("not ok")
	}

	// to check whether two map are equal or not
	m3 := map[string]int{"souvik": 1, "alex": 2, "cathrina": 3, "deeps": 4}

	fmt.Println(maps.Equal(m2, m3))
}
