package main

import "fmt"

// 主要看顺序
func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
