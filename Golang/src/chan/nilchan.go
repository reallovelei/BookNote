package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	// ch = nil //
	ch <- 1

	fmt.Println(ch)
}
