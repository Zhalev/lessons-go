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

		if unicode.IsLetter(l) && esc == false {
			r = l
			b.WriteRune(l)
		}
		if l == '\\' && esc == true {
			r = l
			esc = false
			b.WriteRune(l)
		}
		if l == '\\' && r != l {
			esc = true
		}

		if unicode.IsNumber(l) {
			if r != 0 && esc == false {
				c := int(l - '0')
				for i := 1; i < c; i++ {
					b.WriteRune(r)
				}
			}
			if esc == true {
				r = l
				b.WriteRune(r)
				esc = false
			}

		}
	}
	return b.String()
}
