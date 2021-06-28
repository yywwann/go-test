package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

func A() (int, int) {
	return 3, 1
}

func fun() (a int) {
	a = 3
	defer func() {
		log.Info().Int("a1", a).Msg("")
	}()
	b, a := A()
	log.Info().Int("a2", a).Msg("")
	log.Info().Int("b", b).Msg("")
	if a := 2; a== 2 {
		a = 2
	}
	return a
}

func main() {
	//fmt.Println(fun())
	fmt.Println((-12340.0)>>8)
	fmt.Println((-12340.0)/256.0)
}
