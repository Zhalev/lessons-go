package hw_4

import (
	"reflect"
	"testing"
)

func TestEnv(t *testing.T) {
	dir := "../test/2.txt"
	res := map[string]string{
		"mart": "3",
		"apr":  "4",
	}
	if l, err := ReadDir(dir); reflect.DeepEqual(l, res) == false {
		t.Fatalf("Text: %s: Result: %v; Expected: %v; Error: %v", dir, l, res, err)
	}

}
