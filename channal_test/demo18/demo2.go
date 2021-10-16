package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source2(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1
	// 休眠1秒/2秒/3秒
	time.Sleep(time.Duration(rb) * time.Second)
	select {
	case c <- ra:
	default:
	}
}

func Demo2() {
	rand.Seed(time.Now().UnixNano())

	c := make(chan int32, 1) // 此通道容量必须至少为1
	for i := 0; i < 5; i++ {
		go source2(c)
	}
	rnd := <-c // 只采用第一个成功发送的回应数据
	fmt.Println(rnd)
}
