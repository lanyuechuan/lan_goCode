package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func hello(i int) {
	defer wg.Done()
	fmt.Println(i)
}

func main(){
	for a := 0; a < 10 ;a ++{
		wg.Add(1)
		go hello(a)
	}
	wg.Wait()
}