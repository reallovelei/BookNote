package main

// 逃逸分析
// go tool compile -m 9_escape.go
/**
结果
9_escape.go:7:6: can inline test
9_escape.go:11:6: can inline main
9_escape.go:12:6: inlining call to test
9_escape.go:8:2: moved to heap: a
9_escape.go:12:6: moved to heap: a
*/

var e *int

func test() {
	a := 1
	e = &a
}
func main() {
	test()
}
