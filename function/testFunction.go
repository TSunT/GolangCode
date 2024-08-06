package function

import "fmt"

func TestFunc() {
	a := foo1("abc", 123)
	fmt.Println("a = ", a)
	b, c := foo2("bcd", 234)
	fmt.Printf("b = %d, c = %d\n", b, c)
	e, f := foo3("cde", 345)
	fmt.Printf("e = %d, f = %d\n", e, f)
}

// 返回单个值
func foo1(v1 string, v2 int) int {
	fmt.Printf("v1 = %s, v2= %d\n", v1, v2)
	return 100
}

// 返回两个值
func foo2(v1 string, v2 int) (int, int) {
	fmt.Printf("v1 = %s, v2= %d\n", v1, v2)
	return 100, 200 // 按顺序返回两个值
}

// 返回两个值
func foo3(v1 string, v2 int) (r1 int, r2 int) { // r1 r2 属于foo3的形参 初始化为0
	fmt.Printf("v1 = %s, v2= %d\n", v1, v2)
	fmt.Printf("r1 = %d, r2 = %d \n", r1, r2)
	r1 = 101
	r2 = 201
	return
}
