/* 
主要内容：
1、结构体的构造函数
2、go语言的方法，记住：方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
*/


package main

import (
	"fmt"
)

type Person struct {
    name string
    city string
    age  int8
}

// 给Person类型定义一个Dream方法
func (p Person) Dream(langu string) string{
    return p.name + "的梦想是学好" + langu + "语言！"
}


// Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个person的构造函数。
// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func newPerson(name, city string, age int8) *Person {
    return &Person{
        name: name,
        city: city,
        age:  age,
    }
}

func main() {
	// 调用构造函数
	p := newPerson("小蓝", "测试", 90)
	// 例--->简明格式用 %+v   {Name:zhangsan} 
	// 例---->golang特有模式用 %#v   main.Human{Name:"zhangsan"}
	// 例----->不需要键则不用格式化直接输出  {zhangsan}
	fmt.Printf("%+v\n", p)  

	result := p.Dream("golang")
	fmt.Println((result))
	
}