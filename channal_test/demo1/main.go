package main

import (
	"fmt"
	"sync"
	"time"
)

// 这里只能读
func read(c <-chan int) {
	fmt.Println("read:", <-c)
}

// 这里只能写
func write(c chan<- int) {
	c <- 0
}

func run(done chan int, i int, wg *sync.WaitGroup) {
	for {
		fmt.Println(i, "start select")
		select {
		case <-done:
			time.Sleep(2 * time.Second)
			fmt.Println("do something ", i)
			fmt.Println("stop goroutine ", i)
			wg.Done()
			return
			//default:

			//fmt.Println("do nothing")
		}
		fmt.Println(i, "stop select")
	}
}

func main() {
	done := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go run(done, i, &wg)
	}

	time.Sleep(5 * time.Second)
	//done<-1
	//done<-2
	close(done)
	time.Sleep(time.Second)
	wg.Wait()
	fmt.Println("exit main")
}
