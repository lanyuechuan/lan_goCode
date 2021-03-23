package main

import (
	"reflect"
	"fmt"
)

type Flower struct {

}

func (this *Flower) Move(){
	fmt.Println("花儿飘动")
}

func main() {
	f := Flower{}
	value := reflect.ValueOf(&f)
	v := value.MethodByName("Move")
	// v.Call(nil)
	v.Call([]reflect.Value{})
}