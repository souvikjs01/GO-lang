package main

import (
	"fmt"
)

/*
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt(2))
	fmt.Println(f.Abs())
}
*/

type X struct {
	data []int
}

func modify(x X) {
	x.data[0] = 100
}

func main() {
	obj := X{data: []int{1, 2, 3, 4}}
	modify(obj)
	fmt.Println(obj.data[0])
}
