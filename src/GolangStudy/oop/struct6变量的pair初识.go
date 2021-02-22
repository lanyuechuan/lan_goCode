package main

import (
	"fmt"
	"io"
	"os"
)

//不管怎样传递指向，pair会不变，且一直跟随传递！！！

func main() {
	var a string
	//a: pair<type:string, value="asde">
	a = "asde"

	//声明一个接口类型变量allType
	//allType: pair<type:string, value="asde">
	var allType interface{}
	allType = a

	value, _ := allType.(string)
	fmt.Println(value)
	//tty: pair<type:*os.File, value:"./struct1值传递.go">
	tty, err := os.OpenFile("./struct1值传递.go", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open error",err)
		return
	}

	//r: pair<type: , value: >
	var r io.Reader
	//r: pair<type:*os.File, value:"./struct1值传递.go">
	r = tty

	//w: pair<type: , value: >
	var w io.Writer
	//w: pair<type:*os.File, value:"./struct1值传递.go">
	w = r.(io.Writer)//使用断言强制转换
	w.Write([]byte("lan yue chuan"))
}