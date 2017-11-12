//不定参数

package main

import (
	"fmt"
)

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func myfunc1(args ...int) {
	//全部传入
	myfunc(args...)
	//传切片
	myfunc(args[:]...)
}

func main() {
	//myfunc(2, 3, 4)
	//myfunc(4, 5, 6, 7)
	myfunc1(2, 3, 4)
}
