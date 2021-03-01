package main

import (
	"fmt"
	"runtime"
	"time"
)

//注意看打印顺序

func main() {
	//用协程执行一个形参为空，返回值为空的函数
	go func(){
		defer fmt.Println("A.defer")


		func() {
			defer fmt.Println("B.defer")
			// return
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	//死循环，防止主线程退出
	for {
		time.Sleep(1 * time.Second)
	}
}