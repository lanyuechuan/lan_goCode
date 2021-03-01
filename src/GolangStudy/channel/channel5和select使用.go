//单流程下的一个go只能监控一个chan的状态，select可以完成监控多个chan的状态
package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	//虽然这儿是死循环取操作，但是不会panic，因为落点有个return退出该go程
	for {
		select {
			// 如果成功向c写入数据，则执行该case
		case c <- x:
			x = y
			y = x + y
			// 如果从quit成功读到数据，则执行该case
		case <- quit:
			return
		} 
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	
	go func() {
		for i := 0; i < 6; i ++ {
			//不断地去c中读数据并打印
			fmt.Println(<- c)
		}
		quit <- 0
	}()
	
	//先执行该主程
	fibonacii(c, quit)
}