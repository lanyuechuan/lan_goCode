/*
主要内容：
结构体为值类型，但是如果是切片套结构体，结构体的内存会变，即会随改变而改变，具有指针类型特性
*/

package main

import "fmt"

type student struct {
    id   int
    name string
    age  int
}

func demo(ce []student) {
    //切片是引用传递，是可以改变值的
    ce[1].age = 999
    // ce = append(ce, student{3, "xiaowang", 56})
    // return ce
}
func main() {
    var ce []student  //定义一个切片类型的结构体
    ce = []student{
        student{1, "xiaoming", 22},
        student{2, "xiaozhang", 33},
    }
    fmt.Println(ce)
    demo(ce)
    fmt.Println(ce)
}