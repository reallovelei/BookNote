package main

import "fmt"

func SubSlice() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	n := len(data)
	fmt.Printf("data len:%d \n", n)

	l := data[n-1]
	fmt.Println(l)
	new := data[0:1]
	fmt.Println(new)
}

func for1() {
	ar := [3]int{1, 2, 3}
	for _, v := range ar {
		fmt.Println("v:", v)
	}

	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr)
}

func woke() {
	old := 6
	old &^= 2
	fmt.Printf("old:%d \n", old)
}

func main() {
	SubSlice()
	woke()
}
