package test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

func TestContains(t *testing.T) {
	t.Log(Monday, Tuesday)
	t.Log("abcd2")
}

func TestChannel(t *testing.T) {
	ch := make(chan int, 65537)
	ch <- 5
	t.Log(ch)
}
