/*
给结构体赋值，一定要用指针，因为结构体是值类型
*/
package main

import (
	"fmt"
)

type aaa struct {
	a int
	b string
	c string
}

func main(){
	eee := make(map[string]*aaa, 2)
	eee["wo"] = &aaa{a:1, b:"r", c:"t"}
	eee["wo"].b = "蓝"
	fmt.Println(eee["wo"].a, eee["wo"].b, eee["wo"].c)


}