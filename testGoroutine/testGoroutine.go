package testgoroutine

import (
	"fmt"
	"runtime"
	"time"
)

func newTask() {
	defer fmt.Println("new task defer") // defer 在方法执行完在执行该语句 使用栈的方式 先进后执行

	for i := 0; i < 1000; i++ {
		fmt.Println("new task count: ", i)
		time.Sleep(1 * time.Second)
	}
}

func TaskMainTask() {
	defer fmt.Println("main task defer")

	go newTask()

	go func(a int) {
		defer fmt.Println("3: anonymity func defer. start a=", a)
		a++
		defer func(v int) { fmt.Println("2: anonymity func defer. v=", v) }(a) // defer 如果有参数传入 则将该参数当前情况一起入栈
		defer func() { fmt.Println("1: in anonymity call", a) }()              // defer 匿名函数没有参数则按照现有的a的值传入
		a++
		fmt.Println("a++ ", a)
	}(0)

	go func() {
		defer fmt.Println("test goexit defer.")
		runtime.Goexit()
		fmt.Println("test goexit.") // 不会输出
	}()

	for i := 0; i < 1000; i++ {
		fmt.Println("main task count: ", i)
		time.Sleep(2 * time.Second)
	}
}
