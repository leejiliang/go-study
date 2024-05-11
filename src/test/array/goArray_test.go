package array

import "testing"

func TestArrayInit(t *testing.T) {
	var bal = [5]int{1, 2, 3, 4, 5}
	var bal2 = [5]int{1: 10, 2, 3: 30, 4}
	t.Log(bal)
	t.Log(bal2)
}
