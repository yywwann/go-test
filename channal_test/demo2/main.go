package main

import (
	"fmt"
	"time"
)

func run(stop <-chan struct{}, done chan<- struct{}) {
	// 每一秒打印一次 hello
	for {
		select {
		case <-stop:
			fmt.Println("stop...")
			done <- struct{}{}
			return
		//case t := <-time.After(time.Second):
		//	fmt.Println("hello", t.String())
		default:
			time.Sleep(time.Second)
			fmt.Println("do nothing")
		}
	}

}

func main() {
	// 一对多
	stop := make(chan struct{}, 1)
	// 多对一
	done := make(chan struct{}, 1)
	//for i := 0; i < cap(done); i++ {
	//	go run(stop, done)
	//}
	go run(stop, done)

	// 5s 后退出
	time.Sleep(5 * time.Second)
	fmt.Println("stop")
	stop <- struct{}{}

	//for i := 0; i < cap(done); i++ {
	//	<-done
	//	fmt.Println("done stop", time.Now())
	//}
	fmt.Println("all done")
}
