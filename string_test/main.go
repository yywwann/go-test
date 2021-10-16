package main

import (
	"fmt"
	"strings"
	"time"
)

var (
	endpoints       = []string{"127.0.0.1:2379", "127.0.0.1:2378"}
	EndpointSepChar = ','
	EndpointSep     = fmt.Sprintf("%c", EndpointSepChar)
)

func main() {

	authority := strings.Join(endpoints, EndpointSep)
	fmt.Println(authority)

	hosts := strings.FieldsFunc(authority, func(r rune) bool {
		return r == EndpointSepChar
	})
	fmt.Println(hosts)

	fmt.Println(time.Now().Unix())
}
