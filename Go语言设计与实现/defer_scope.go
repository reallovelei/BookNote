package main

import "fmt"

// 主要看顺序
func main() {
	{
		defer fmt.Println("defer runs")
		fmt.Println("block ends")
	}
	fmt.Println("main ends")
}
