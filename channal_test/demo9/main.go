package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
)

//向一个通道发送一个值来实现单对单通知
//我们已知道，如果一个通道中无值可接收，则此通道上的下一个接收操作将阻塞到另一个协程发送一个值到此通道为止。
//所以一个协程可以向此通道发送一个值来通知另一个等待着从此通道接收数据的协程。
//
//在下面这个例子中，通道done被用来做为一个信号通道来实现单对单通知。
func main() {
	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}) // 也可以是缓冲的

	// 排序协程
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{} // 通知排序已完成
	}()

	// 并发地做一些其它事情...

	<-done // 等待通知
	fmt.Println(values[0], values[len(values)-1])
}
