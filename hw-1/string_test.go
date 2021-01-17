package main

import "testing"

func TestCount(t *testing.T) {
	str := "as2d4fa3"
	res := "assddddfaaa"
	if l := count(str); l != res {
		t.Fatalf("Parsing string %s: Result: %s; Expected: %s", str, l, res)
	}
	str = "asdf"
	res = "asdf"
	if l := count(str); l != res {
		t.Fatalf("Parsing string %s: Result: %s; Expected: %s", str, l, res)
	}
	str = "442"
	res = ""
	if l := count(str); l != res {
		t.Fatalf("Parsing string %s: Result: %s; Expected: %s", str, l, res)
	}
	str = `qwe\4\5`
	res = `qwe45`
	if l := count(str); l != res {
		t.Fatalf("Parsing string %s: Result: %s; Expected: %s", str, l, res)
	}
	str = `qwe\45`
	res = `qwe44444`
	if l := count(str); l != res {
		t.Fatalf("Parsing string %s: Result: %s; Expected: %s", str, l, res)
	}
	str = `qwe\\5`
	res = `qwe\\\\\`
	if l := count(str); l != res {
		t.Fatalf("Parsing string %s: Result: %s; Expected: %s", str, l, res)
	}
}
