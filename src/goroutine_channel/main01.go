/*
多次执行下面的代码，会发现每次打印的数字的顺序都不一致。
这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println(i)
}

func main(){
	for a := 0; a < 10 ;a ++{
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(a)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}