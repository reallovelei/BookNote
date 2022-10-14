package maptest

import "testing"

func test() map[int]int {
	m := make(map[int]int)
	for i := 0; i < 1024; i++ {
		m[i] = i
	}
	return m
}

func testCap() map[int]int {
	// 预先准备空间
	m := make(map[int]int, 1680)

	for i := 0; i < 1024; i++ {
		m[i] = i
	}
	return m
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkTestCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCap()
	}
}
