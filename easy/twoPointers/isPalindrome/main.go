package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {
	var b []rune
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			b = append(b, v)
		} else if v >= 'A' && v <= 'Z' {
			b = append(b, v+32)
		}
	}

	l, r := 0, len(b)-1
	for l <= r {
		if b[l] != b[r] {
			return false
		}
		l++
		r--
	}
	return true
}
