lan yue chuan

import "fmt"

type myInt int

type Book struct {
	title string
	auth  string 
}

//结构体是值类型
func changeBook(book Book) {
	book.auth = "666"
}

func changeBook2(book *Book) {
	book.auth = "777"
}

func main() {
	var a myInt = 10
	fmt.Printf("%T",a)
	var book1 Book
	book1.title = "golang"
	book1.auth = "zhangsan"
	changeBook2(&book1)
	fmt.Println(book1)

}