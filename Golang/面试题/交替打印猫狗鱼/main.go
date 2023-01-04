package main

import (
	"fmt"
	"sync"
)

func main() {
	chCat := make(chan int)
	chDog := make(chan int)
	chFish := make(chan int)
	// chFinish := make(chan bool)

	n := 10

	var wg sync.WaitGroup
	wg.Add(3)

	go func(wg *sync.WaitGroup) {
		for i := 0; i < n; i++ {
			<-chCat
			fmt.Printf("Cat\n")
			chDog <- i
		}
		// 这里有一个需要注意的点
		// 如果这里不主动return 可能会 存在泄露风险
		wg.Done()
		return
	}(&wg)

	go func(wg *sync.WaitGroup) {
		for i := 0; i < n; i++ {
			<-chDog
			fmt.Printf("Dog\n")
			chFish <- i
		}
		wg.Done()
		return
	}(&wg)

	go func(wg *sync.WaitGroup) {
		for i := 1; i <= n; i++ {
			<-chFish
			fmt.Printf("%d Fish\n", i)
			if i < n {
				chCat <- i
			}
		}
		wg.Done()
		return
	}(&wg)

	// 点火 启动
	chCat <- 1
	wg.Wait()
	// time.Sleep(time.Second)

}
