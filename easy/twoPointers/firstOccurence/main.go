package main

import "fmt"

func main() {
	fmt.Println(strStr("mississippi", "issip"))
}

func strStr(haystack string, needle string) int {
	var l, r int
	for l < len(haystack) {
		if haystack[l] == needle[r] {
			if r == len(needle)-1 {
				return l - r
			}
			fmt.Println(l, r)
			r++
		} else if r > 0 {
			l -= r // left-i yene onki yerine goyyas.
			fmt.Println(l)
			r = 0
		}
		l++
	}
	return -1
}
