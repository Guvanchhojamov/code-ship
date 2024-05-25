package main

import "fmt"

var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func main() {
	fmt.Println(reverseVowels("hello"))
}

func reverseVowels(s string) string {
	b := []rune(s)
	l, r := 0, len(b)-1
	for l < r {
		if vowels[b[l]] {
			if vowels[b[r]] {
				b[l], b[r] = b[r], b[l]
			}
			r--
		}
		l++
	}
	return string(b)
}
