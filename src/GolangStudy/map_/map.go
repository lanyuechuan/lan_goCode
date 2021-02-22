package main

import (
	"fmt"
)

func main() {
	// 三种声明方式
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("空map")
	}

	
	myMap1 := make(map[string]string)
	myMap1["one"] = "1"


	myMap2 := map[string] string {
		"two":"2",
	}
	fmt.Println(myMap1,"\n",myMap2)
}

//动态数组也叫切片