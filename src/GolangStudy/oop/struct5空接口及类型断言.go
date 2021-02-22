package main

import "fmt"

//空接口就是个元类，像我们常见的基本数据类型就继承了空接口，因为空接口没有方法
func myFunc(arg interface{}){
	fmt.Println("myFunc 被调用")
	fmt.Println(arg)

	//interface{}空接口 如何区分此时引用的底层数据类型到底是什么
	//golang给接口变量提供了断言机制，前者为变量值，后者为是否断言成功
	v,ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	}else {
		fmt.Println("arg is string")
		fmt.Println(v)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"Golang"}
	myFunc(book)
	myFunc(100)
	myFunc("字符串类型")
	myFunc(3.14)
}