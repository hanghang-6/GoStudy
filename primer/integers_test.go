package main

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expect := 4
	if sum != expect {
		t.Errorf("Sum does not equal expect %d != %d", expect, sum)
	}
}
