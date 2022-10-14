package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(1 * time.Second)
	done := make(chan bool, 100)

	for {
		select {
		case <-done:
			return
		case <-timer.C:
			cnt := len(done)
			fmt.Println(cnt)
		}
	}
	timer.Stop()
}
