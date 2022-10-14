package main

import (
	"fmt"
	"unsafe"
)

func generate8191() {
	nums := make([]int, 8191) // < 64KB
	nums[0] = 1
	// for i := 0; i < 8191; i++ {
	// 	nums[i] = i
	// }
}

func generate8192() {
	nums := make([]int, 8193) // > 64KB
	for i := 0; i < 8193; i++ {
		nums[i] = i
	}
}

func generate(n int) {
	nums := make([]int, n) // 不确定大小
	for i := 0; i < n; i++ {
		nums[i] = i
	}
}

// go build -gcflags=-m escape_stack.go
func main() {
	i := 1
	fmt.Printf("%d  size: %d\n", i, unsafe.Sizeof(i))
	generate8191()
	generate8192()
	generate(1)
}
