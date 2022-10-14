package test

import (
	"testing"
)

func TestSlice(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	n := len(data)
	t.Logf("data len:%d \n", n)

	l := data[n-1]
	t.Logf("Slice:%+v", l)
	return
}
