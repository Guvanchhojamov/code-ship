package main

import "fmt"

func main() {
	fmt.Println(appendCharacters("z", "abcde"))
}

func appendCharacters(s string, t string) int {
	j := 0
	for i := 0; i < len(s) && j < len(t); i++ {
		if s[i] == t[j] {
			j++
		}
	}
	return len(t) - j
}
