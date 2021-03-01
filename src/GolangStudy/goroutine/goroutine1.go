//go的协程创建很简单，只需要用go关键字加一个函数，或接一个自执行函数


package main

import (
	"fmt"	
	"time"
)


func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("goroutine i:%v\n",i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//创建一个协程，去执行这个方法
	go newTask()

	//再来个死循环保证主线程不退出
	i := 0
	for {
		i ++
		fmt.Printf("main goroutine %v\n",i)
		time.Sleep(1 * time.Second)
	}

	//如果不加死循环，主线程会马上退出，goroutine也会很快退出
}