package main

import (
	"errors"
	"fmt"
)

//可以给返回值命名，被命名后，它们的值在函数开始的时候被自动初始化为空。
//在函数中执行不带任何参数的return语句时，会返回对应的返回值变量的值
func add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("必须是正数")
		return
	}

	return a + b, nil
}

func main() {
	s, err := add(5, 2)
	if err == nil {
		fmt.Println(s)
	} else {
		fmt.Println(err)
	}
}
