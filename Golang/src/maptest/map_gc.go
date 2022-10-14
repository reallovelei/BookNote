package main

import (
	"fmt"
	"runtime"
	"time"
)

type CountConf struct {
	id     string
	status map[int]int
}

func main() {

	stats := new(runtime.MemStats)
	fmt.Printf("Before: %5.2fk\n", float64(stats.Alloc)/1024)

	cc := make(map[int]*CountConf)
	for i := 0; i < 3; i++ {
		cc[i] = &CountConf{
			id: "cdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvdcdsvdfvd",
			status: map[int]int{
				200: 343,
				302: 454,
			},
		}
		// fmt.Println(i)
	}
	// debug.FreeOSMemory()s
	runtime.ReadMemStats(stats)
	fmt.Printf("After add: %5.2fk\n", float64(stats.Alloc)/1024)

	for i := 0; i < 2; i++ {
		delete(cc, i)
		// cc[i] = nil
	}
	runtime.ReadMemStats(stats)
	fmt.Printf("After delete: %5.2fk\n", float64(stats.Alloc)/1024)

	// cc = make(map[int]*CountConf)
	// fmt.Println(cc)
	fmt.Println("addddd")
	time.Sleep(time.Second * 120)
	fmt.Println("addddd-120")
	time.Sleep(time.Second * 30)
	runtime.ReadMemStats(stats)
	fmt.Printf("After set map empty: %5.2fk\n", float64(stats.Alloc)/1024)
}
