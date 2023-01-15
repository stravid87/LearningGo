package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySumchanged(3, 10)
	if x != 13 {
		t.Error("Expected", 13, "Got", x)
	}
}