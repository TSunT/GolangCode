package testPair

import (
	"fmt"
	"io"
	"os"
)

// simple 1
func TestPair() {
	//a pair<staticType:string, value:123>
	var a string
	a = "123"
	var allType interface{}

	//allType pair<staticType:string, value:123> 与a一直
	allType = a

	str, _ := allType.(string)

	fmt.Println(str)

}

// simple 2
func TestOsIo() {
	//tty: pair<type:*os.File, value:"/dev/tty"文件表示符> 初始pair
	userHome, _ := os.UserHomeDir()
	fmt.Println("user home: ", userHome)
	os.Create(userHome + "\\testGo.txt")
	tty, err := os.OpenFile(userHome+"\\testGo.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open File error!", err)
		return
	}

	var r io.Reader
	//r: pair<type:*os.File, value:"/dev/tty"文件表示符> 与tty一直
	r = tty
	var w io.Writer
	//w: pair<type:*os.File, value:"/dev/tty"文件表示符> 断言转换后 pair与tty保持一直
	w = r.(io.Writer)
	w.Write([]byte("hello goLang!!!"))
}

// simple 3
type ReadBookIf interface {
	DoReadBook() string
}

type WriteBookIf interface {
	DoWriteBook(strData string)
}

type Book struct {
	bookData string
}

func (book *Book) DoWriteBook(strData string) {
	book.bookData = strData
}

func (book *Book) DoReadBook() string {
	return book.bookData
}

func TestReadWriteBook() {

	//book: pair<type:Book, value:"hello world"> 初始pair
	book := Book{bookData: "my tears ricochet!!!"}
	var bookReader ReadBookIf
	var bookWriter WriteBookIf
	//bookReader: pair<type:Book, value:"hello world"> 与book初始pair一致
	bookReader = &book
	//bookWriter: pair<type:Book, value:"hello world"> 与book初始pair一致
	bookWriter = &book
	fmt.Println("before write data: ", bookReader.DoReadBook())
	bookWriter.DoWriteBook("bejeweled!")
	fmt.Println("after write data: ", bookReader.DoReadBook())
}
