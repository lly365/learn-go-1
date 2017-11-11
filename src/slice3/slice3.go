package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{7, 8, 9}

	//copy(s1, s2) //s1: [7 8 9 4 5], s2: [7 8 9]

	copy(s2, s1) //s1: [1 2 3 4 5], s2: [1 2 3]
	fmt.Println(s1, s2)
}
