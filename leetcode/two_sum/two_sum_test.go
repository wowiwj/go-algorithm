package two_sum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {

	// 表驱动测试
	var tests = []struct {
		inArr  []int
		target int
		out    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for i, tt := range tests {
		s := twoSum(tt.inArr, tt.target)
		if !isEqualArr(s, tt.out) {
			t.Errorf("test failed in %d", i)
		}

	}

}

func isEqualArr(arr1, arr2 []int) bool {

	if len(arr1) != len(arr2) {
		return false
	}

	for index, item := range arr1 {
		if item != arr2[index] {
			return false
		}
	}
	return true
}
