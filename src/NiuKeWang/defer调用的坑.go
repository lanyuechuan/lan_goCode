package main

import "fmt"

func hello(i int) {
	fmt.Println(i) // 5
}

func main() {
	i := 5
	defer hello(i) // 先拷i放着，再往下走，最终输出15，再输出5
	defer func(k int){
		fmt.Println(k)
	}(i)
	i = i + 10
	fmt.Println(i)
}

//跟传参有关系，只认实时数据，不管后面的数据，这是defer的坑
