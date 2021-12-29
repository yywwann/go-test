package main

import "fmt"

type IA interface {
	A() string
	B() string
}

type TAT struct {
	name string
}

func NewTAT(name string) TAT {
	return TAT{
		name: name,
	}
}

func (a TAT) A() string {
	return "A()" + a.name
}

func (a TAT) B() string {
	return "B()" + a.name
}

func (a TAT) C() string {
	return "C()" + a.name
}

func Print(t IA) string {
	return t.B()
}

func main() {
	var a IA
	a = NewTAT("a")
	fmt.Println(a.A())
	fmt.Println(Print(a))
	c := NewTAT("c")
	fmt.Println(c.C())
	fmt.Println(Print(c))
}
