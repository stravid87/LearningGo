package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySum(3, 10)
	if x != 13 {
		t.Error("Expected", 5, "Got", x)
	}
}