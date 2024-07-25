package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	st := []string{"ant", "act", "tack"}
	tw := []string{"tack", "act", "acti"}
	wordCount(st, tw)
}

func wordCount(startWords []string, targetWords []string) int {
	stwMap := map[byte]string{}
	result := 0
	for _, sWord := range startWords {
		stwMap[getSum(sWord)] = sWord
	}
	for _, tWord := range targetWords {
		sumWithoutLast := getSum(tWord[:len(tWord)-1])
		if _, ok := stwMap[sumWithoutLast]; ok {
			fmt.Println(stwMap[sumWithoutLast], "tt", tWord)
			if isExists(tWord, tWord[len(tWord)-1]) {
				result++
			}
		}
	}
	fmt.Println(stwMap, result)
	return result
}

func getSum(s string) byte {
	var sum byte
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}

func isExists(word string, lastChar byte) bool {
	for i := 0; i < len(word)-1; i++ {
		if word[i] == lastChar {
			return false
		}
	}
	return true
}
