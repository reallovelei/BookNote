package main

import (
	"fmt"
	"time"
)

func Check() bool {
	nowHM := time.Now().Format("15:04")
	// start := "7:00"
	//end := "23:00"
	currentTime := time.Now()

	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 7, 0, 0, 0, currentTime.Location())
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 0, 0, 0, currentTime.Location())

	if currentTime.Unix() < startTime.Unix() {
		fmt.Printf("现在是%s,还未到上班时间 \n", nowHM)
		return false
	}

	if currentTime.Unix() > endTime.Unix() {
		fmt.Printf("现在是%s,已下班\n", nowHM)
		return false
	}

	fmt.Printf("现在是%s, working\n", nowHM)
	return true
}

func main() {
	fmt.Println(time.Now().UnixNano() / 1000000)
	fmt.Println(time.Now().Unix())
	Check()

}
