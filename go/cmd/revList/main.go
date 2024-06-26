package main

import (
	"fmt"
	"revList/internal/list"
)

func main() {
	var l list.LinkedList[int]
	for i := 0; i < 10; i++ {
		l.Push(i)
	}
	fmt.Println(l)

	l.Reverse()
	fmt.Println(l)
}
