package daily

import "strings"

/*

1935. Maximum Number of Words You Can Type
There is a malfunctioning keyboard where some letter keys do not work.
All other keys on the keyboard work properly.

Given a string text of words separated by a single space (no leading or trailing spaces)
and a string brokenLetters of all distinct letter keys that are broken,
return the number of words in text you can fully type using this keyboard.




Example 1:
Input: text = "hello world", brokenLetters = "ad"
Output: 1
Explanation: We cannot type "world" because the 'd' key is broken.

Example 2:
Input: text = "leet code", brokenLetters = "lt"
Output: 1
Explanation: We cannot type "leet" because the 'l' and 't' keys are broken.

Example 3:
Input: text = "leet code", brokenLetters = "e"
Output: 0
Explanation: We cannot type either word because the 'e' key is broken.


Constraints:

1 <= text.length <= 10^4
0 <= brokenLetters.length <= 26
- text consists of words separated by a single space without any leading or trailing spaces.
- Each word only consists of lowercase English letters.
- brokenLetters consists of distinct lowercase English letters.
*/
/*
 We need to find count where we can write full world.
	- worlds seperated with 1 space, always valid.
	- brokenLetters onlu inuque chars.
	- contains only lowercase English letters.

brute force:
	1. create words array seperated with space.
	2. Take each word, iterate ower word and brokenLetters if mach then skip.
	3. If not mach increment count.
	4. check each word and return count at the end.
tc: n+26*n*wordlen.
sc: n - for store in words array.

 Can we optimze?  How?
	- since our brokenWords at max always 26, we can create bw map; key-char, val- bool.
	- seperate array to words.
	 iterate over array:
	 	iterate over word:
			if word[i] in map: continue
		conuter++
	return counter at end.
	tc: n+ n*wordLen.
	sc: n
*/
func canBeTypedWords(text string, brokenLetters string) int {
	var result = 0
	var bl = make(map[rune]bool, len(brokenLetters))

	for _, ch := range brokenLetters {
		bl[ch] = true
	}

	words := strings.Split(text, " ")

	for _, word := range words {
		if !isBroken(word, bl) {
			result++
		}
	}

	return result
}

func isBroken(word string, bl map[rune]bool) bool {
	for _, char := range word {
		if bl[char] {
			return true
		}
	}
	return false
}
