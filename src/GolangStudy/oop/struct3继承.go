package main

import "fmt"

type Human struct {
	Name string
	sex string
}


func (this *Human) Eat() {
	fmt.Println("human Eat被调")
}

func (this *Human) Walk() {
	fmt.Println("human Walk")
}

type SuperMan struct {
	Human // 该类继承了Human的属性和方法
	level string
}

//重定义父类的方法Eat
func (this *SuperMan) Eat(){
	fmt.Println(("superman 的Eat被调用"))
}

//子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println(("superman Fly被调用"))
}

func main() {
	h := Human{"lan","female"}
	h.Eat()
	h.Walk()

	// s := SuperMan{Human{"yue","male"},"B"}
	var s SuperMan
	s.level = "C"
	s.Walk()//调父类的Walk
	s.Eat()//调子类的方法
	s.Fly()//调子类的方法
}