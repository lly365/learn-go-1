//使用 interface{} 传递任意类型

package main

import (
	"fmt"
)

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "是个整数")
		case string:
			fmt.Println(arg, "是个字符串")
		case int64:
			fmt.Println(arg, "是个int64")
		default:
			fmt.Println(arg, "是个未知类型")
		}
	}
}

func main() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234

	MyPrintf(v1, v2, v3, v4)
}
