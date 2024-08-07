package testInterface

import (
	"fmt"
	"strconv"
	"time"
)

// 实现对象转字符接口
type OfString interface {
	ToString() string
}

type Book struct {
	name        string
	page        int
	publishDate time.Time
}

func (book *Book) ToString() string {
	// fmt.Println("{")
	// fmt.Println("\t name:", book.name)
	// fmt.Println("\t page:", book.page)
	// fmt.Println("\t publishDate:", book.publishDate)
	// fmt.Println("}")
	timeStr := book.publishDate.Format("2006-01-02 15:04:05")
	page := strconv.Itoa(book.page)
	return "BookToString{" + " name:" + book.name + ", page: " + page + ", publish date: " + timeStr + " }"
}

func (book Book) stringFormat() string {
	timeStr := book.publishDate.Format("2006-01-02 15:04:05")
	page := strconv.Itoa(book.page)
	return "BookFormat{" + " name:" + book.name + ", page: " + page + ", publish date: " + timeStr + " }"
}

// 空接口表示任何数据类型 使用断言判定类型
func showString(value interface{}) {
	// 空接口有 断言类型方法 val.(Type) 类似instanceOf 如下
	str, ok := value.(string)
	fmt.Println("first result: ", str)
	if ok {
		fmt.Println("value is a string type value:", value)
	} else {
		fmt.Printf("type: %T\n", value)
		fmt.Println("value is NOT a string type")

	}

	var result string
	switch typeStr := value.(type) { // 使用 value.(type) 注意 这里结合switch-case使用 判定类型 返回类型指针
	case string:
		result = value.(string)
	case Book:
		fmt.Println("Book type!")
		result = value.(Book).stringFormat()
	case OfString:
		result = value.(OfString).ToString()
	default:
		fmt.Printf("justify type: %T\n", typeStr)
		result = "can't convert string"
	}
	fmt.Println("second result:", result)
}

func TestInterfaceAssert() {

	showString("111222333")
	fmt.Println("------")
	showString(123)
	fmt.Println("------")
	showString(Book{name: "bookName", page: 101, publishDate: time.Now()})
	fmt.Println("------")
	showString(&Book{name: "bookName", page: 101, publishDate: time.Now()}) // interface 以 指针的形式传递 否则判定是 Book类型
	fmt.Println("------")

}
