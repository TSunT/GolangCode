package testchannel

import (
	"fmt"
	"time"
)

// 无缓冲channel
func TestChannelPrimitive() {
	// 定义一个channel
	c := make(chan int) // 无缓冲

	go func() {
		defer fmt.Println("receive goroutine end...")
		fmt.Println("before receive int.")
		num, ok := <-c // 从channel里面接收一个666 并返回给num 注意：这里如果发送方未发送这里会阻塞 等待发送方发送
		fmt.Println("after receive int.")
		if !ok {
			fmt.Println("receive one int failed!")
		} else {
			fmt.Println("receive one int: ", num)
		}
	}()
	time.Sleep(2 * time.Second)
	go func() {
		defer fmt.Println("send goroutine end...")
		fmt.Println("before send.")
		c <- 666 // 往channel里发送一个666 注意 如果此时接收方未取到该数据发送方阻塞等待接收方获取数据
		fmt.Println("after send.")
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
	}
}

// 有缓冲channel
func TestChannelWithCache() {
	// 定义一个有缓冲的channel
	c := make(chan int, 3) // 3 表示缓存数量
	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("receive goroutine end...")
		for i := 0; i < 10; i++ {
			num, ok := <-c // 从channel里面接收一个666 并返回给num 注意：这里如果发送方未发送这里会阻塞 等待发送方发送
			if !ok {
				fmt.Println("receive one int failed!")
			} else {
				fmt.Println("receive one int: ", num)
			}
		}

	}()
	time.Sleep(2 * time.Second)
	go func() {
		defer fmt.Println("send goroutine end...")
		fmt.Println("before send.")
		for i := 0; i < 10; i++ {
			c <- i // 往channel里发送一个666 注意 如果此时接收方未取到 超过缓存数量的时候 该数据发送方阻塞等待接收方获取数据
		}
		fmt.Println("after send.")
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
	}

}
