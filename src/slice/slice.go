// 基于数组创建切片

package main

import (
	"fmt"
)

func main() {
	//定义一个数组先
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//基于数组创建一个数组切片
	var mySlice []int = myArray[:5]

	fmt.Println("myArrary元素：")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nmySlice元素：")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
}
