package main

import (
	"fmt"
)

/*
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
*/

func main() {
	// m := make(map[string]int)

	// m["answer"] = 23
	// fmt.Println("the value", m["answer"])

	// m["answer"] = 46
	// fmt.Println("the value", m["answer"])

	// delete(m, "answer")
	// fmt.Println("the value", m["answer"])
	// // but key doesnt exists

	// v, ok := m["answer"]
	// fmt.Println("the value", v, "present", ok)

	//---------------------------------------------
	// m := map[string]int{
	// 	"ten":      10,
	// 	"eleven":   11,
	// 	"twelve":   12,
	// 	"thirteen": 13,
	// }

	// for k, v := range m {
	// 	// random ordered:
	// 	fmt.Println(k, " ", v)
	// }

	//-----------------------------------------------
	type Product struct {
		Name  string
		Price float64
	}
	productMap := map[string]Product{
		"p1": {"Laptop", 899.99},
		"p2": {"SmartPhone", 102.99},
		"p3": {"watch", 99.09},
	}

	for _, v := range productMap {
		fmt.Println(v.Name, " ", v.Price)
	}
}
