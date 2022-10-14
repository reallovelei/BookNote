package main

// type B struct {
// A *int
// C int
// }

// // dlv 调试器
// func asmSayHello(a string) {
// 	nums := make([]int, 5)
// 	for i := 0; i < len(nums); i++ {
// 		nums[i] = i * i
// 	}
// 	fmt.Println(a)
// }

// func asmSayHello2(c string) int {
// 	a := 5
// 	fmt.Println(a)

// 	b := B{A: &a, C: a}
// 	return b.C
// }
func main() {
	// asmSayHello("aa")
	// c := asmSayHello2("bb")
	a := 0
	a++
	nums := make([]int, 5)

	for i := 0; i < len(nums); i++ {
		nums[i] = i * i
	}

	// fmt.Println(nums)
}
