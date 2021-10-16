package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var ball = make(chan string)
	kickBall := func(playerName string) {
		for {
			fmt.Print(<-ball, "传球", "\n")
			time.Sleep(time.Millisecond * time.Duration(rand.Int()%1000))
			ball <- playerName
		}
	}
	go kickBall("张三")
	go kickBall("李四")
	go kickBall("王二麻子")
	go kickBall("刘大")
	ball <- "裁判"    // 开球
	var c chan bool // 一个零值nil通道
	<-c             // 永久阻塞在此
}
