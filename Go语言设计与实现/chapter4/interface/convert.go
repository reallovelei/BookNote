package main

import "fmt"

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

//go:noinline
func (c *Cat) Quack() {
	fmt.Println(c.Name + " meow")
}

func main() {
	var c Duck = &Cat{Name: "draven"}
	c.Quack()
}