package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data []int
		answer int
	}

	tests := []test{
		test{[]int{3, 4}, 8},
		test{[]int{4, 5, 6, 7}, 22},
		test{[]int{1, 0}, 1},
		test{[]int{3, 5, -2}, 6},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "But Got", x)
		}
	}
}