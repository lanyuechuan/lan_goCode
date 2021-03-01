//channel有无缓冲的区别就可以理解为是否有临时仓库，前者有，后者没有
package main

import (
	"fmt"
)

func main() {
	//定义一个channel
	c := make(chan int)
	go func() {
		defer fmt.Println("函数 goroutine结束。")
		fmt.Println("函数go routine正在运行。")
		c <- 666 // 将666发送给c
	}()

	num := <- c
	fmt.Println(num)
	fmt.Println("main goroutine结束。")//其实默认main函数也是个协程


}

//按理会存在go func函数未能成功执行的情况，因为主协程肯定是更快的，
// 但是为啥还是按正常程序走完了，就是因为chan管道是个阻塞操作
// 即main.go和func.go两个协程形成了一个同步的效果