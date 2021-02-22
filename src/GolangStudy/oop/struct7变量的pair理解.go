package main

import "fmt"

//l来个接口
type Reader interface{
	ReadBook()
}

//再来个接口
type Writer interface{
	WriteBook()
}

//来个结构体实现上面的两个接口
type Book struct {

}

func (this *Book) ReadBook(){
	fmt.Println("read book func.")
}

func (this *Book) WriteBook(){
	fmt.Println("write book func.")
}

func main(){
	//b: pair<type:Book, value:Book{}地址>
	b := &Book{}

	//来一个接口变量r: pair<type:, value:>
	var r Reader
	//r: pair<type:Book, value:Book{}地址>
	r = b
	r.ReadBook()

	var w Writer
	w,_ = r.(Writer) //断言成功拿到w
	w.WriteBook()
}