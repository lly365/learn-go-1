//直接创建切片
//会有一个匿名数组被创建出来，不需要我们操心

package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Println("==============")
	for i, v := range s { //range有2个返回值，第一个是索引，第二个是元素的值
		fmt.Printf("[%d]: %d\n", i, v)
	}
	fmt.Println("==============")
}

func printSliceInfo(s []int) {
	fmt.Println("************")
	fmt.Printf("len(slice): %d, cap(slice): %d\n", len(s), cap(s))
	fmt.Println("************")
}

func main() {
	//1、创建一个初始元素个数为5的数组切片，元素初始值为0
	mySlice1 := make([]int, 5)
	printSlice(mySlice1)
	printSliceInfo(mySlice1)
	//追加3个元素
	mySlice1 = append(mySlice1, 7, 8, 9)
	printSlice(mySlice1)
	printSliceInfo(mySlice1)

	//2、创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间
	mySlice2 := make([]int, 5, 10)
	printSlice(mySlice2)
	printSliceInfo(mySlice2)

	//3、直接创建并初始化包含5个元素的数组切片
	mySlice3 := []int{1, 2, 3, 4, 5}
	printSlice(mySlice3)
	printSliceInfo(mySlice3)
	//把mySlice1里所有元素追加到mySlice3。mySlice1后面的...的意思是把切片元素打散
	mySlice3 = append(mySlice3, mySlice1...)
	printSlice(mySlice3)
	printSliceInfo(mySlice3)

}
