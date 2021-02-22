package main

import (
	"fmt"
)

//引用传递，传递的指针(map，channel，slice)
func printMap(myMap map[string]string){
	myMap["Japan"] = "tokyo"
	for k,v := range myMap {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func main() {
	myMap1 := make(map[string]string)
	//增加
	myMap1["China"] = "Beijing"
	myMap1["Japan"] = "Tokyo"
	myMap1["USA"] = "KONW"

	// 删除
	delete(myMap1,"USA")

	//修改
	myMap1["China"] = "beijing"

	printMap(myMap1)

	//查(遍历)
	for k,v := range myMap1 {
		fmt.Println(k)
		fmt.Println(v)
	}
}

//动态数组也叫切片