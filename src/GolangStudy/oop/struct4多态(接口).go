package main

import "fmt"

//本质是一个指针,必须实现以下三个方法，才证明struct实现量这个接口
type AnimalF interface {
	Sleep()
	GetColor() string //代表返回值是个string
	GetType() string
} 

//具体的类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

//具体的类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}


//比main方法中更显然，即传什么对象就调那个对象的Sleep方法
func showAnimal(animal AnimalF) {
	animal.Sleep() //
}

func main() {

	//证明多态现象，含义就是不同struct调同样方法最终执行的是struct各自的方法
	var animal AnimalF //接口类型
	animal = &Cat{"Green"}
	animal.Sleep() //调用的就是Cat的Sleep方法
	animal = &Dog{"Green"}
	animal.Sleep() //调用的就是Dog的Sleep方法
}

//多态要素：
// 1、有一个父类接口
// 2、有子类并且实现了父类接口中的全部方法
// 3、父类类型的变量(指针)，指向了之类指向了子类的具体变量