package main

import (
	"fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

func main() {
	var a Integer = 3
	if a.Less(5) {
		fmt.Println("小于")
	} else {
		fmt.Println("不小于")
	}
	a.Add(4)
	fmt.Println(a)
}
