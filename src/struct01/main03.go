/*
主要内容:
json字符串和结构体互转（注意，结构体内部成员属性或方法必须是公有的，即首字母大写）
*/
package main

import (
	"encoding/json"
	"fmt"
)
//结构体实例化
//只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
//结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。
type Person struct {
    Name string
    City string
    Age  int8
}

func main() {
    var p1 Person      //这一句才会真正分配内存
    p1.Name = "pprof.cn"  //结构体就是用来实现面向对象，这个小数点熟悉不，这不就是py的对象吗
    p1.City = "北京"
    p1.Age = 18
	fmt.Println(p1)  //{pprof.cn 北京 18}

	//把结构体对象转换成json字符串
	jsonByte, _ := json.Marshal(p1) // 返回的是byte类型
	jsonStr := string(jsonByte)
	fmt.Printf("%v", jsonStr)   // "{"Name":"pprof.cn","City":"北京","Age":18}"

	//把json字符串转换成结构体对象
	var str = `{"Name":"pprof.cn","City":"北京","Age":18}` //用反引号，要不然还得\转义双引号
	var p2 Person
	err := json.Unmarshal([]byte(str), &p2) //参数是byte切片类型
	if err == nil{
		fmt.Println("转换失败")
	}
	fmt.Println(p2.Name, p2.Age)
} 