package main

import (
	"fmt"
	"time"
)

func main() {
	//定义一个channel
	c := make(chan int, 3)
	fmt.Println(len(c), cap(c))

	//channel只能容纳三个元素，当满了之后，会阻塞，直到main程取出一个之后
	go func() {
		defer fmt.Println("子go程结束")
		for i := 0; i < 4; i ++ {
			c <- i
			fmt.Println("子go程正在运行", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i ++ {
		num := <- c //从管道中取出值给到num
		fmt.Println(num)
	}
	fmt.Println("main程结束")
}