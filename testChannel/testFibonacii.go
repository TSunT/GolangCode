package testchannel

import (
	"fmt"
)

func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, y+x
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Text_Fibonacii() {
	ch := make(chan int)
	quit := make(chan int)
	count := 10

	go func(c1 chan int) {
		for i := 0; i < count; i++ {
			fmt.Println(<-c1)
		}
		quit <- 0
	}(ch)
	fibonacii(ch, quit)
}
