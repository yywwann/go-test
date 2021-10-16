package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source3() <-chan int32 {
	c := make(chan int32, 1) // 必须为一个缓冲通道
	go func() {
		ra, rb := rand.Int31(), rand.Intn(3)+1
		time.Sleep(time.Duration(rb) * time.Second)
		c <- ra
	}()
	return c
}

func Demo3() {
	rand.Seed(time.Now().UnixNano())

	var rnd int32
	// 阻塞在此直到某个数据源率先回应。
	select {
	case rnd = <-source3():
	case rnd = <-source3():
	case rnd = <-source3():
	}
	fmt.Println(rnd)
}
