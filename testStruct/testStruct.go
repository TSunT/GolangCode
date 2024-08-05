package testStruct

import "fmt"

// 声明一种数据类型
type myint int // 给int 起个别名myint

// 定义一个结构体
type Book struct {
	title  string
	author string
}

// 注意这里是传递一个副本 不改变实参的值
func changeBook(book Book) {
	book.author = "nju"
}

// 注意这里是传递一个副本
func changeBookPointer(book *Book) {
	book.author = "nuaa"
}

func TestStruct() {
	var a myint = 10
	fmt.Println("a=", a)

	var book1 Book
	book1.title = "title1"
	book1.author = "js"
	fmt.Printf("%v\n", book1) // %v可以打印任意类型
	fmt.Println("-----changeBook----")
	changeBook(book1)
	fmt.Printf("%v\n", book1)
	fmt.Println("-----changeBookPointer----")
	changeBookPointer(&book1)
	fmt.Printf("%v\n", book1)
}
