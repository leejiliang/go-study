package common

import "fmt"

func init() {
	fmt.Printf("init %d", 1)
}

func init() {
	fmt.Printf("init %d", 2)
}
func Square(v int) int {
	return v * v
}
