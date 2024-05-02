package main

import "fmt"

/*
242. Valid Anagram
*/
func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sChars := map[byte]int{}
	tChars := map[byte]int{}
	for i := 0; i < len(s); i++ {
		sChars[s[i]]++
		tChars[t[i]]++
	}

	for i, v := range sChars {
		if v != tChars[i] {
			return false
		}
	}
	return true
}
