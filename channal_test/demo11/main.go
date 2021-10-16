package main

import "log"
import "time"

type T = struct{}

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready // 阻塞在此，等待通知
	log.Print("Worker#", id, "开始工作")
	// 模拟一个工作负载。
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("Worker#", id, "工作完成")
	done <- T{} // 通知主协程（N-to-1）
}

//多对单和单对多通知

//事实上，上例中展示的多对单和单对多通知实现方式在实践中用的并不多。
//在实践中，我们多使用sync.WaitGroup来实现多对单通知，使用关闭一个通道的方式来实现单对多通知

func main() {
	log.SetFlags(0)

	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	// 模拟一个初始化过程
	time.Sleep(time.Second * 3 / 2)
	// 单对多通知
	ready <- T{}
	ready <- T{}
	ready <- T{}
	// 从一个已关闭的通道可以接收到无穷个值这一特性也将被用在很多其它在后面将要介绍的用例中。
	// 实际上，这一特性被广泛地使用于标准库包中。比如，context标准库包使用了此特性来传达操作取消消息。
	// close(ready) // 群发通知
	// 等待被多对单通知
	<-done
	<-done
	<-done
}
