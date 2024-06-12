package main

import "fmt"

func main() {
	fmt.Println(mergeWords("abcd", "pq"))
}

func mergeWords(word1, word2 string) string {
	var result []byte
	l, r := 0, 0

	for l < len(word1) || r < len(word2) {
		if l < len(word1) {
			result = append(result, word1[l])
			l++
		}

		if r < len(word2) {
			result = append(result, word2[r])
			r++
		}
	}
	return string(result)
}
