package main

import (
	"fmt"
	"time"
)

//func main() {
//	path := []byte("AAAA/BBBBBBBBB")
//	fmt.Println("path =>", string(path), "path.len =>", len(path), "path.cap =>", cap(path))
//	// path => AAAA/BBBBBBBBB path.len => 14 path.cap => 14
//
//	sepIndex := bytes.IndexByte(path,'/')
//	dir1 := path[:sepIndex: sepIndex]
//	dir2 := path[sepIndex+1:]
//	fmt.Println("dir1 =>",string(dir1), "dir1.len =>", len(dir1), "dir1.cap =>", cap(dir1)) //prints: dir1 => AAAA
//	fmt.Println("dir2 =>",string(dir2), "dir2.len =>", len(dir2), "path.cap =>", cap(dir2)) //prints: dir2 => BBBBBBBBB
//	dir1 = append(dir1,"suffix"...)
//	fmt.Println("path =>", string(path), "path.len =>", len(path), "path.cap =>", cap(path))
//	fmt.Println("dir1 =>",string(dir1), "dir1.len =>", len(dir1), "dir1.cap =>", cap(dir1)) //prints: dir1 => AAAA
//	fmt.Println("dir2 =>",string(dir2), "dir2.len =>", len(dir2), "path.cap =>", cap(dir2)) //prints: dir2 => BBBBBBBBB
//}

type user struct {
	name string
	age  int8
}

var u = user{name: "Ankur", age: 25}
var g = &u

func modifyUser(pu *user) {
	fmt.Println("modifyUser Received Vaule", pu)
	pu.name = "Anand"
}
func printUser(u <-chan *user) {
	time.Sleep(2 * time.Second)
	fmt.Println("printUser goRoutine called", <-u)
}

func main() {
	c := make(chan *user, 5)
	c <- g
	fmt.Println(g)
	// modify g
	g = &user{name: "Ankur Anand", age: 100}
	u.age += 1
	go printUser(c)
	go modifyUser(g)
	time.Sleep(5 * time.Second)
	fmt.Println(g)
}
