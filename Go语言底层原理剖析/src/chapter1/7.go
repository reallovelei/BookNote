// 变量捕获
package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1
	b := 2

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(a, b)
	}()

	a = 99
	time.Sleep(3 * time.Second)
}
