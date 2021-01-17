package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(count("Хочу вам рассказать про всю Одессу. Одесса очень велика"))
}
func count(s string) string {

	var r rune
	var b strings.Builder
	var esc bool = false

	for _, l := range s {
		if unicode.IsLetter(l) {
			r = l
		}
		if l == '\\' {
			esc = true
		}
		if unicode.IsNumber(l) && !esc {
			if r != 0 {
				c := int(l - '0')
				for i := 1; i < c; i++ {
					b.WriteRune(r)
				}
			}

		} else {
			b.WriteRune(l)
		}
	}
	return b.String()
}
