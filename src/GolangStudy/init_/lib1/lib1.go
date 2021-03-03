package lib1

import "fmt"

func Lib() {
	fmt.Println("我是lib1的Lib")
}

func init() {
	fmt.Println("我是lib1 的 init")
}