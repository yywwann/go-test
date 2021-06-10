package main

import (
	"fmt"
	"os"
)

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


func main() {
	//erre := os.Mkdir("./data/program/goapp/golang", 0777)
	err := os.MkdirAll("./log", 0777)
	if err != nil {
		fmt.Println(err)
	}
}
