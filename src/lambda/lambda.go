package main

import (
	"fmt"
)

func main() {
	//赋给一个变量
	f := func(x, y int) int {
		return x + y
	}
	result := f(123, 456)
	fmt.Println(result)

	//定义时直接调用
	result = func(x, y int) int {
		return x + y
	}(789, 91011)
	fmt.Println(result)
}
