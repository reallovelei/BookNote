package maptest

import "testing"

func TestCap(t *testing.T) {
	// 预先准备空间
	m := make(map[int]int, 1000)
	for i := 0; i < 10; i++ {
		m[i] = i
	}
}

func TestList(t *testing.T) {
	// 预先准备空间
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}

	go func(m map[int]int) {
		for k, v := range m {
			t.Logf("K:%d, v:%d \n", k, v)
		}
	}(m)

	go func(m map[int]int) {
		for k, v := range m {
			t.Logf("K:%d, v:%d \n", k, v)
		}
	}(m)
	return
}
