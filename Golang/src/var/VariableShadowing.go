package main

import "fmt"

// 变量遮蔽
func foo(n int) {
	a := 1
	a += n
}

var a = 11

func main() {
	fmt.Println("a = ", a) // 11
	foo(5)
	fmt.Println("after calling foo, a = ", a) // 11
}
