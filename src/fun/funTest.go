package main

import (
	"fmt"
)

const WEEK = "MONDAY" // 全局变量
var testStr string
var globalStr string = "globalStr out of method"

func main() {
	fmt.Println("Hello World")
	num1 := 100 // 局部变量
	num2 := 200
	maxNum := maxInt(num1, num2)
	fmt.Println(maxNum)

	str1, str2 := swapVal("hello", "world")
	fmt.Println(str1, str2)
	fmt.Println(WEEK)
	testStr = "Global val"
	var globalStr string = "hello global str in method"
	fmt.Println(testStr)
	fmt.Println(globalStr)
}

func maxInt(num1 int, num2 int) int { //num1 num2 形参
	if num1 > num2 {
		return num1
	}
	return num2
}
func swapVal(v1, v2 string) (string, string) {
	return v2, v1
}
