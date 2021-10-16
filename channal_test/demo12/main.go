package main

import (
	"fmt"
	"time"
)

func AfterDuration(d time.Duration) <-chan struct{} {
	//c := make(chan struct{}, 1)
	//go func() {
	//	time.Sleep(d)
	//	c <- struct{}{}
	//}()
	//return c

	c := make(chan struct{}, 1)
	<-time.After(d)
	c <- struct{}{}
	return c
}

func main() {
	fmt.Println("Hi!")
	<-AfterDuration(time.Second)
	fmt.Println("Hello!")
	<-AfterDuration(time.Second)
	fmt.Println("Bye!")

	fmt.Println("Hi!")
	<-time.After(time.Second)
	fmt.Println("Hello!")
	<-time.After(time.Second)
	fmt.Println("Bye!")
}
