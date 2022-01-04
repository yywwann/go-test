package main

import "fmt"

func demo2() {
	a := []int{}
	n := 9
	for i := 0; i < n; i++ {
		a = append(a, i)
	}

	printSlice(a)

	fmt.Println("b remove 5")
	b := remove2(a, 5)
	printSlice(a)
	printSlice(b)

	fmt.Println("b append 1")
	b = append(b, 1)
	printSlice(a)
	printSlice(b)

	fmt.Println("b append 2")
	b = append(b, 2)
	printSlice(a)
	printSlice(b)

	fmt.Println("a append 3")
	a = append(a, 3)
	printSlice(a)
	printSlice(b)
}

// 同一个底层数组

func remove1(a []int, i int) []int {
	return append(a[:i], a[i+1:]...)
}

func remove2(a []int, index int) []int {
	b := []int{}
	for i, val := range a {
		if i == index {
			continue
		}

		b = append(b, val)
	}
	return b
}

func printSlice(a []int) {
	for _, val := range a {
		fmt.Print(val, " ")
	}
	fmt.Println("")
	fmt.Println("cap = ", cap(a), ", len = ", len(a)) // 输出: 8, 8
}
