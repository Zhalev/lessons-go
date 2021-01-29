package main

import (
	"fmt"
	"reflect"
	"time"
)

func Concat(slices [][]int) []int {
	var s []int
	for _, l := range slices {
		s = append(s, l...)
	}
	return s
}

func main1() {
	type pair struct {
		s [][]int
		r []int
	}
	test := []pair{
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{{1, 2}, {3, 4}, {6, 5}}, []int{1, 2, 3, 4, 6, 5}},
		{[][]int{{1, 2}, {}, {6, 5}}, []int{1, 2, 6, 5}},
	}
	for _, t := range test {
		s := t.s
		r := t.r
		r2 := Concat(s)
		fmt.Printf("Test: %v\n", s)
		fmt.Printf("Expected: %v\n", r)
		fmt.Printf("Result: %v\n", r2)
		if reflect.DeepEqual(r, r2) {
			fmt.Printf("OK\n\n")
		} else {
			fmt.Printf("FAIL\n\n")
		}
	}
}
func main2() {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(10 * time.Second)
		ch <- "Hello"
	}()
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case data := <-ch:
			fmt.Printf("received: %v\n", data)
		case <-ticker.C:
			fmt.Println("failed to received in 10s")
		}
	}
}
func main() {
	// A "channel"
	ch := make(chan string)

	// Start concurrent routines
	go push("Moe", ch)
	go push("Larry", ch)
	go push("Curly", ch)

	// Read 3 results
	// (Since our goroutines are concurrent,
	// the order isn't guaranteed!)
	fmt.Println(<-ch, <-ch, <-ch)
}

func push(name string, ch chan string) {
	msg := "Hey, " + name
	ch <- msg
}
