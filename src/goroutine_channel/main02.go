package main

import (
	"fmt"
	_ "sync"
)



func main(){
	ch := make(chan int, 4)
	ch <- 10
	fmt.Println(ch)
	a := <- ch
	fmt.Println(ch, a)
}