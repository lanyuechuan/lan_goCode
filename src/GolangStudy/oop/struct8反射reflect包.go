package main

import (
	"fmt"
	"reflect"

)

type User struct {
	Id int
	name string
	age int
}

func (this *User) Call() {
	fmt.Println(this)
}

//最简单的一种数据类型反射获取type和value
func reflectNum(arg interface{}) {
	argType := reflect.TypeOf(arg)
	fmt.Println(argType)
	argValue := reflect.ValueOf(arg)
	fmt.Println(argValue)
}

func DoFileAndMethod(arg interface{}) {
	//获取input的type和value
	argType := reflect.TypeOf(arg)
	fmt.Println(argType)
	argValue := reflect.ValueOf(arg)
	fmt.Println(argValue)
	//获取field进行变量
	for i:= 0;i < argType.NumField(); i ++ {
		field := argType.Field(i)
		value := argValue.Field(i).Interface()
		fmt.Println(field.Name,field.Type,value)
	}
	//通过type获取
	for i:= 0; i < argType.NumMethod(); i ++ {
		m := argType.Method(i)
		fmt.Println(m.Name, m.Type)
	}

}

func main() {
	num := 1.0
	reflectNum(num)

	user := User{1,"lan",1}
	DoFileAndMethod(user)
}