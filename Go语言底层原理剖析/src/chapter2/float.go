package main

import "fmt"

// 2.1 浮点数的陷阱
func main() {
	var f1 float64 = 0.3
	var f2 float64 = 0.6
	fmt.Println(f1 + f2)
	var f3 float64 = 0.84
	var f4 float64 = 0.01
	fmt.Println(f3 + f4)
}
