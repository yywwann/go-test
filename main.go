package main

import "fmt"

type A struct {
	name string
	age  int32
	sex  int32
}

func (a A) String() string {
	return a.name
}

type B struct {
	A
	Role int32
}

type C struct {
	A
}

func main() {
	b := B{
		A:    A{},
		Role: 0,
	}
	b.name = "1"
	fmt.Println(b, b.String(), b.A.String())
	b.A.name = "2"
	fmt.Println(b, b.String(), b.A.String())
	c := C{
		A: A{},
	}
	c.name = "3"
	fmt.Println(c, c.String(), c.A.String())
	c.A.name = "4"
	fmt.Println(c, c.String(), c.A.String())

}
