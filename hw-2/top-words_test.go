package hw_2

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	str := "adf asdf asd adf adda? asd. asdf, asdf adf adf adf adda all, jk: ll kk lk df adda"
	res := map[string]int{
		"adf":  5,
		"asdf": 3,
		"adda": 3,
		"asd":  2,
		"all":  1,
		"jk":   1,
		"ll":   1,
		"kk":   1,
		"lk":   1,
		"df":   1,
	}

	if l := topWords(str); reflect.DeepEqual(l, res) == false {
		t.Fatalf("Text: %s: Result: %v; Expected: %v", str, l, res)
	}

	if l := topWords(str); l["adf"] == 5 {
		t.Fatalf("Text: %s: Result: %v; Expected: %v", str, l["adf"], 5)
	}

}
