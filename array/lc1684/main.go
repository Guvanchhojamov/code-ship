package main

import "fmt"

func main() {
	countConsistentStrings("fstqyienx", []string{"n", "eeitfns", "eqqqsfs", "i", "feniqis", "lhoa", "yqyitei", "sqtn", "kug", "z", "neqqis"})
}

/* Leetcode 1684.
You are given a string allowed consisting of distinct characters and an array of strings words.
A string is consistent if all characters in the string appear in the string allowed.
*/

//Example 1:
//
//Input: allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
//Output: 2
//Explanation: Strings "aaab" and "baa" are consistent since they only contain characters 'a' and 'b'.
//Example 2:
//
//Input: allowed = "abc", words = ["a","b","c","ab","ac","bc","abc"]
//Output: 7
//Explanation: All strings are consistent.
//
//
//Example 3:
//
//Input: allowed = "cad", words = ["cc","acd","b","ba","bac","bad","ac","d"]
//Output: 4
//Explanation: Strings "cc", "acd", "ac", and "d" are consistent.

/*
 1. create allowedMap[string]string{}
 2. compare each word with this map
 3. if char in word, not in map, continue else count when word is over.

*/

func countConsistentStrings(allowed string, words []string) int {
	var allowedChars = make(map[int32]bool, len(allowed))
	var c = 0
	for _, char := range allowed {
		allowedChars[char] = true
	}

	for _, word := range words {
		c++
		for _, char := range word {
			if _, exists := allowedChars[char]; !exists {
				c--
				break
			}
		}
		fmt.Println(c)
	}
	return c
}
