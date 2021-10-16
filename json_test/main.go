package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	A string `json:"A,omitempty"`
	B string `json:"B,omitempty"`
	C string `json:"C,omitempty"`
}

func main() {
	jsonStr := `
{
	"A": "123",
	"B": "",
	"D": ""
}
`

	t := &T{}
	err := json.Unmarshal([]byte(jsonStr), t)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", t)

	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(b, "\n", string(b))
}
