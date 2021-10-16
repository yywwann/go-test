package main

import "fmt"

func main() {

	trySend := func(ch chan string, v string) {
		select {
		case ch <- v:
		default: // 如果c的缓冲已满，则执行默认分支。
		}
	}
	tryReceive := func(ch chan string) string {
		select {
		case v := <-ch:
			return v
		default:
			return "-" // 如果c的缓冲为空，则执行默认分支。
		}
	}

	c := make(chan string, 2)
	trySend(c, "Hello!") // 发送成功
	trySend(c, "Hi!")    // 发送成功
	trySend(c, "Bye!")   // 发送失败，但不会阻塞。
	// 下面这两行将接收成功。
	fmt.Println(tryReceive(c)) // Hello!
	fmt.Println(tryReceive(c)) // Hi!
	// 下面这行将接收失败。
	fmt.Println(tryReceive(c)) // -
}
