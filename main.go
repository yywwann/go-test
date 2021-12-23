package main

import (
	"fmt"
	"time"
)

type testInt struct {
	Num int
}

func (t *testInt) print() {
	fmt.Println(t.Num)
}

func pts(t *testInt) {
	fmt.Println(t.Num)
}

func main() {
	a := make([]int, 0)
	for i := 1; i < 1000; i++ {
		a = append(a, i)
	}
	t := &testInt{
		Num: 0,
	}

	t.print()
	go func() {
		time.Sleep(1 * time.Second)
		t.print()
	}()

	Test(t)
	t.print()
	select {}
}

func Test(t *testInt) {
	t.Num = 2
}
