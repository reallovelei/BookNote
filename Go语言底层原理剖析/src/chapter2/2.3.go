package main

import (
	"fmt"
	"math"
)

func main() {
	var number float32 = 0.85
	fmt.Printf("Starting Number :%f \n\n", number)

	bits := math.Float32bits(number)
	binary := fmt.Sprintf("%.32b", bits)
	fmt.Printf("Bit Pattern: %s | %s %s | %s %s %s %s %s %s \n",
		binary[0:1],
		binary[1:5], binary[5:9],
		binary[9:12], binary[12:16],
		binary[16:20], binary[20:24],
		binary[24:28], binary[28:32])

	bias := 127
	sign := bits & (1 << 31)
	exponentRaw := int(bits >> 23)
	exponent := exponentRaw - bias
	var mantissa float64

	for index, bit := range binary[9:32] {
		fmt.Printf("-- index:%d bit:%d  \n", index, bit)

		// 1 的ascll码是49, 其实这里就是判断 是不是1
		if bit == 49 {
			position := index + 1
			bitValue := math.Pow(2, float64(position))
			fractional := 1 / bitValue
			mantissa = mantissa + fractional
			fmt.Printf("---- index:%d bit:%f %f %f \n", index, bitValue, fractional, mantissa)
		}
	}

	value := (1 + mantissa) * math.Pow(2, float64(exponent))

	fmt.Printf("sign:%d Exponent:%d (%d) Mantissa:%f Value:%f \n\n", sign, exponentRaw, exponent, mantissa, value)

}

// 2.4 判断浮点数是否为整数
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
