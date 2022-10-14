package main

import "fmt"

func myFunction(a int) (int, int) {
	return a + 16, a + 32
}

func main() {
	a, b := myFunction(64)
	fmt.Println(a, b)
}
