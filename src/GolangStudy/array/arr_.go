package main

import (
	"fmt"
)

func main() {
	var myArray [10]int
	var a3 [10]int
	if a3 == nil {
		fmt.Println("a3是空")
	}
	for i := 0; i < len(myArray); i ++ {
		fmt.Println(myArray[i])
	}
	for index,value := range myArray {
		fmt.Println(index,value)
	}
	fmt.Printf("%T",myArray,a3) // 一定要注意数组的数据类型包括了长度
}

//动态数组也叫切片