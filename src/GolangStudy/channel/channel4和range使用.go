//不能向一个被close发送消息或为空的管道取消息，否则会panic
package main

import (
	"fmt"
)

func main() {
	//定义一个channel
	c := make(chan int)
	
	go func() {
		for i := 0; i < 5; i ++ {
			c <- i
		}
		close(c) // 关闭一个channel
	}()

	//整一个for死循环，一直在管道中拿
	/*
	for {
		//如果ok为true表示channel未能成功关闭，如果为false表示已经关闭
		if data, ok := <- c ; ok{
			fmt.Println(data)
		}else {
			break
		}
	}
	*/

	// 可以使用range来迭代不断操作channel,可以把chan，slice，数组看成py的可迭代对象
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main主程结束")
}