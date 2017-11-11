package main

import (
	"fmt"
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func using_map1() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	personDB["12345"] = PersonInfo{"12345", "Tome", "Room 203 ..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101"}

	//查找键为“1234”的信息
	person, ok := personDB["1234"]
	//ok返回true表示找到了对应的数据
	if ok {
		fmt.Println("找到了", person.Name, " # 1234.")
	} else {
		fmt.Println("找不到 # 1234")
	}
}

func main() {
	/*
		var map1 map[string][]string
		map1 = make(map[string][]string)
		map1["nemo"] = []string{"龙南", "lly365@gmail.com"}

		str, ok := map1["nemo"]
		if ok {
			fmt.Println(str)
		} else {
			fmt.Println("not found")
		}
	*/
	//using_map1()

}
