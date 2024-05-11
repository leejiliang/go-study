package goRecursion

import "testing"

func TestRecursion(t *testing.T) {
	t.Logf("numsPlus result is %d", plusOne(1))
}

func plusOne(num int) int {
	if num < 5 {
		return plusOne(num + 1)
	}
	return num + 1
}
