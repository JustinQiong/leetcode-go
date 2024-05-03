package leetcode

import "testing"

func TestAverage1(t *testing.T) {
	arr := []int{4000, 3000, 1000, 2000}
	ans := average(arr)
	if ans != 2500.00000 {
		t.Error()
	}
}

func TestAverage2(t *testing.T) {
	arr := []int{1000, 2000, 3000}
	ans := average(arr)
	if ans != 2000.00000 {
		t.Error()
	}
}
