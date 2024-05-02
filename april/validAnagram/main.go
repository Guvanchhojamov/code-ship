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

	var sSum, tSum uint64
	var sM = uint64(1)
	var tM = uint64(1)

	for i := 0; i < len(s); i++ {
		sM *= uint64(s[i])
		sSum += uint64(s[i])
		tM *= uint64(t[i])
		tSum += uint64(t[i])

	}
	return sM+sSum == tM+tSum
}

//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//
//	sChars := map[byte]int{}
//	tChars := map[byte]int{}
//	for i := 0; i < len(s); i++ {
//		sChars[s[i]]++
//		tChars[t[i]]++
//	}
//
//	for i, v := range sChars {
//		if v != tChars[i] {
//			return false
//		}
//	}
//	return true
//}
