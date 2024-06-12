# 151. Reverse Words in a String

## Problem Statement

Given an input string `s`, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in `s` will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that `s` may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

## Examples

### Example 1:
- **Input:** `s = "the sky is blue"`
- **Output:** `"blue is sky the"`

### Example 2:
- **Input:** `s = "  hello world  "`
- **Output:** `"world hello"`
- **Explanation:** Your reversed string should not contain leading or trailing spaces.

### Example 3:
- **Input:** `s = "a good   example"`
- **Output:** `"example good a"`
- **Explanation:** You need to reduce multiple spaces between two words to a single space in the reversed string.

## Constraints
- \( 1 \leq s.length \leq 10^4 \)
- `s` contains English letters (upper-case and lower-case), digits, and spaces `' '`.
- There is at least one word in `s`.

## Follow-up

If the string data type is mutable in your language, can you solve it in-place with \( O(1) \) extra space?

## Solution 
```go

func reverseWords(s string) string {
    c := reverseStr([]byte(s))

	var result, word []byte
	for i := 0; i < len(c); i++ {
		if c[i] != ' ' {
			word = append(word, c[i])
		} else if len(word) > 0 {
			result = append(result, reverseStr(word)...)
			word = []byte{}
		}
	}
	return string(result[:len(result)-1])
}

func reverseStr(word []byte) []byte {
	l, r := 0, len(word)-1
	for l < r {
		word[l], word[r] = word[r], word[l]
		l++
		r--
	}
	return append(word, ' ')
}
```