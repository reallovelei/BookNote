package main

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"
)

func main() {
	kvch, err := lru.New(3)
	fmt.Println(kvch, err)

	kvch.Add("1", "9")
	kvch.Add("2", "8")
	kvch.Add("3", "7")
	kvch.Add("4", "6")
	kvch.Add("5", "5")

	fmt.Println(kvch.Get("1"))
	fmt.Println(kvch.Get("2"))
	fmt.Println(kvch.Get("3"))
	fmt.Println("--------")
	kvch.Add("6", "6")
	fmt.Println(kvch.Keys())
	fmt.Println(kvch.Get("4"))
	fmt.Println(kvch.Get("3"))

}
