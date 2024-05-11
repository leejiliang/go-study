package goSlice

import "testing"

func TestDeclareSlice(t *testing.T) {
	var nums []int // 不需要指明长度
	// 使用make函数创建切片
	names := make([]int, 0)
	// 初始化
	nums2 := []int{1, 2, 3}
	t.Log("nums len, cap", len(nums), cap(nums))
	t.Log("names len, cap", len(names), cap(names))
	t.Log("nums2 len, cap", len(nums2), cap(nums2))
}

func TestSplitSlice(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5}
	t.Log("nums len, cap", len(nums), cap(nums))
	nums2 := nums[1:3]
	t.Log("nums2 len, cap", nums2, len(nums2), cap(nums2))
}

func TestSliceAppend(t *testing.T) {

	var nums = []int{1, 2, 3, 4, 5}
	nums = append(nums, 6, 7, 8, 9, 10, 11)
	t.Logf("nums len, %d, cap, %v", len(nums), cap(nums))
	t.Log(nums)
}

func TestSliceCopy(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5}
	var nums2 = make([]int, 5, 20)
	copy(nums2, nums)
	t.Logf("nums len, %d, cap, %v", len(nums), cap(nums))
	t.Logf("nums2 len, %d, cap, %v", len(nums2), cap(nums2))
	t.Logf("nums, %v, nums2: %v", nums, nums2)
}
