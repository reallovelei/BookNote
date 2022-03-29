package main

import "fmt"

// 判断浮点数是否为整数
func IsInt(bits uint32, bias int) {
	exponent := int(bits>>23) - bias - 23
	coefficient := (bits & ((1 << 23) - 1)) | (1 << 23)
	intTest := (coefficient & (1<<uint32(exponent) - 1))

	fmt.Printf("\nExponent: %d Coefficient: %d IntTest: %d\n", exponent, coefficient, intTest)

	if exponent < -23 {
		fmt.Printf("Not Integer\n")
		return
	}
	if exponent < 0 && intTest != 0 {
		fmt.Printf("Not Integer\n")
		return
	}
	fmt.Printf("INTEGER \n")
}
