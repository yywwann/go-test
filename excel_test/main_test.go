package main

import "testing"

func TestFq(t *testing.T) {
	for i := 0; i <= 100; i++ {
		t.Log("start ", i)
		res := ff(i)
		t.Log(string(res))
	}
}
