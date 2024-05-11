package goRange

import "testing"

func TestRange(t *testing.T) {
	var nums = [...]int{10, 20, 30}
	for _, v := range nums { // 只读取value
		t.Log(v)
	}

	for v, _ := range nums { // 只读取index
		t.Log(v)
	}

	for i, v := range nums {
		t.Logf("index: %d, value, %d", i, v)
	}

	for _, v := range nums {
		v += v
		t.Log(v)
	}
}
