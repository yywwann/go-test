package main

import (
	"bytes"
	"fmt"
)

func demo1() {
	path := []byte("AAAA/BBBBBBBBB")
	fmt.Println("path =>", string(path), "path.len =>", len(path), "path.cap =>", cap(path))
	// path => AAAA/BBBBBBBBB path.len => 14 path.cap => 14

	sepIndex := bytes.IndexByte(path, '/')
	dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1), "dir1.len =>", len(dir1), "dir1.cap =>", cap(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2), "dir2.len =>", len(dir2), "path.cap =>", cap(dir2)) //prints: dir2 => BBBBBBBBB
	dir1 = append(dir1, "suffix"...)
	fmt.Println("path =>", string(path), "path.len =>", len(path), "path.cap =>", cap(path))
	fmt.Println("dir1 =>", string(dir1), "dir1.len =>", len(dir1), "dir1.cap =>", cap(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2), "dir2.len =>", len(dir2), "path.cap =>", cap(dir2)) //prints: dir2 => BBBBBBBBB
}
