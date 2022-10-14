package main

import (
	"fmt"
	"time"
)

func TestList() {
	// 预先准备空间
	m := make(map[int]int, 1000)
	for i := 0; i < 10; i++ {
		m[i] = i
	}

	go func(m map[int]int) {
		for k, v := range m {
			fmt.Printf("K:%d, v:%d \n", k, v)
		}
	}(m)

	go func(m map[int]int) {
		for k, v := range m {
			fmt.Printf("K:%d, v:%d \n", k, v)
		}
	}(m)
	time.Sleep(time.Second)
	return
}

func main() {
	TestList()
}
