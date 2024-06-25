package main

import (
	"fmt"
	"time"
)

var (
	ee int    = 100
	ff string = "232"
)

func main() {
	fmt.Println("hello Golang!")
	time.Sleep(5 * time.Second)
	var aa int = 88
	var bb float64 = 99.8
	var cc string = "hello world!"
	dd := true
	aa1, bb1 := "Hello", "xiaoLinZi"
	fmt.Println(aa1, bb1)
	fmt.Printf(" aa type of = %T value = %d\n", aa, aa)
	fmt.Printf(" bb type of = %T value = %f\n", bb, bb)
	fmt.Printf(" cc type of = %T value = %s\n", cc, cc)
	fmt.Printf(" dd type of = %T value = %t\n", dd, dd)

	fmt.Printf(" ee type of = %T value = %d\n", ee, ee)
	fmt.Printf(" ff type of = %T value = %s\n", ff, ff)

	const gg int = 10000 // 常量是不允许修改的！

	fmt.Printf(" const gg type of = %T value = %d\n", gg, gg)
}
