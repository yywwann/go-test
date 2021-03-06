package main

import (
	"fmt"
	"time"
)

// 从一个通道接收一个值来实现单对单通知
//如果一个通道的数据缓冲队列已满（非缓冲的通道的数据缓冲队列总是满的）但它的发送协程队列为空，则向此通道发送一个值将阻塞，直到另外一个协程从此通道接收一个值为止。 所以我们可以通过从一个通道接收数据来实现单对单通知。一般我们使用非缓冲通道来实现这样的通知。
//
//这种通知方式不如上例中介绍的方式使用得广泛。
func main() {
	done := make(chan struct{})
	// 此信号通道也可以缓冲为1。如果这样，则在下面
	// 这个协程创建之前，我们必须向其中写入一个值。

	go func() {
		fmt.Print("Hello")
		// 模拟一个工作负载。
		time.Sleep(time.Second * 2)

		// 使用一个接收操作来通知主协程。
		<-done
	}()

	done <- struct{}{} // 阻塞在此，等待通知
	fmt.Println(" world!")
}
