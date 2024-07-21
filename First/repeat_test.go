package main

import "testing"

func TestRepeat(t *testing.T) {
	got := repeatChar("a", 5)
	expect := "aaaaa"
	if got != expect {
		t.Errorf("got %q, expect %q", got, expect)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeatChar("bab", 5)
	}
}
