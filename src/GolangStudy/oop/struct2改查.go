package main

import "fmt"

type Hero struct {
	Name string
	Ad int
}

//这个this就是面向对象的提醒(没有硬性要求，但是我习惯用this)，他表明当前方法只能给Hero类来调用
//如果要修改结构体类的值，必须传指针(值传递)，并且他是唯一一个不用显式用*来赋值的(内部自动加*)
//如果类的属性和方法首字母大写，代表外部包可以访问。如果只在本包，那就大小写无所谓

//改
func (this *Hero) SetName(newName string) {
	this.Name = newName
}

//查询
func (this Hero) GetName() {
	fmt.Println(this.Name)
}

func main() {
	hero := Hero{Name : "lan",Ad : 100}
	hero.GetName()
	hero.SetName("li4")
	hero.GetName()
}