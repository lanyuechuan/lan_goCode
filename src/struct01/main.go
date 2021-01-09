package main


import (
	"fmt"
)

type student struct {
    name string
    age  int
}

func main() {
    m := make(map[string]*student)
    stus := []student{
        {name: "pprof.cn", age: 18},
        {name: "测试", age: 23},
        {name: "博客", age: 28},
    }

    for _, stu := range stus {
		fmt.Println(&stu)
        m[stu.name] = &stu
	}
	
	/*
	m = {
		"pprof.cn"： 0x0012,
		"测试": 0x0012,
		"博客": 0x0012,
	}
	
	
	*/

    for k, v := range m {
        fmt.Println(k, "=>", v.name)
    }
} 
/*
结果：
pprof.cn => 博客
测试 => 博客
博客 => 博客


*/