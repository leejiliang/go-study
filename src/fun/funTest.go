package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	num1 := 100
	num2 := 200
	maxNum := maxInt(num1, num2)
	fmt.Println(maxNum)

	str1, str2 := swapVal("hello", "world")
	fmt.Println(str1, str2)
}

func maxInt(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
func swapVal(v1, v2 string) (string, string) {
	return v2, v1
}
