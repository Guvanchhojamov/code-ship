package main

import (
	"fmt"
)

func main() {
	reverseWords("the     sky    is blue  ")
}

func reverseWords(s string) {
	c := append(reverseStr([]byte(s)), ' ')

	var result, word []byte
	for i := 0; i < len(c); i++ {
		if c[i] != ' ' {
			word = append(word, c[i])
		} else if len(word) > 0 {
			result = append(result, reverseStr(append(word, ' '))...)
			word = []byte{}
		}
	}
	result = result[1:]
	fmt.Println(string(result), len(result))
}

func reverseStr(word []byte) []byte {
	l, r := 0, len(word)-1
	for l < r {
		word[l], word[r] = word[r], word[l]
		l++
		r--
	}
	return word
}
