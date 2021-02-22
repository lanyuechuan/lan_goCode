package main

import (
	_ "GolangStudy/init_/lib1"     // 我只想取调lib1包中的init方法，不使用其它变量或接口
	lib2_ "GolangStudy/init_/lib2" //起别名
	"fmt"
)

func main() {
	// lib1.Lib()
	a,b := 1,2
	a,b = b,a
	fmt.Println(a,b)
	lib2_.Lib()
}