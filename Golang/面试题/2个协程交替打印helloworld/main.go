package main

import (
	"fmt"
	"time"
)

func printStr(str string) {
	ch := make(chan byte)

	go func() {
		for i := 0; i < len(str); i++ {
			ch <- str[i] // 协程1先输出  让channel 转起来。  由于是无缓存chan 写入后就会阻塞休眠
			if i%2 == 0 {
				fmt.Printf("go1 %c\n", str[i])
			}
		}
		return
	}()

	go func() {
		for i := 0; i < len(str); i++ {
			<-ch
			if i%2 == 1 {
				fmt.Printf("go2 %c\n", str[i])
			}
		}
		return
	}()
}

func main() {
	str := "Hello, world!"

	printStr(str)

	time.Sleep(time.Second)

	return

}
