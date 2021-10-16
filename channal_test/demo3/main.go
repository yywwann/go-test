package main

import (
	"fmt"
	"time"
)

// MainDemo3 一个简单的通过一个非缓冲通道实现的请求/响应的例子：
func main() {
	c := make(chan int) // 一个非缓冲通道
	done := make(chan struct{})

	go func(ch <-chan int) {
		n := <-ch      // 阻塞在此，直到有值发送到c
		fmt.Println(n) // 9
		// ch <- 123   // 此操作编译不通过
		time.Sleep(time.Second)
		done <- struct{}{}
	}(c)

	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		// <-ch    // 此操作编译不通过
		ch <- x * x // 阻塞在此，直到发送的值被接收
	}(c, 3)

	<-done // 阻塞在此，直到有值发送到done
	fmt.Println("bye")
}
