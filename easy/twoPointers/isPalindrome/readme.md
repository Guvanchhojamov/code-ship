
# 125. Valid Palindrome

**Difficulty:** Easy

## Topics
- Strings
- Two Pointers

## Problem Statement

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string `s`, return `true` if it is a palindrome, or `false` otherwise.

## Examples

### Example 1:
- **Input:** `s = "A man, a plan, a canal: Panama"`
- **Output:** `true`
- **Explanation:** `"amanaplanacanalpanama"` is a palindrome.

### Example 2:
- **Input:** `s = "race a car"`
- **Output:** `false`
- **Explanation:** `"raceacar"` is not a palindrome.

### Example 3:
- **Input:** `s = " "`
- **Output:** `true`
- **Explanation:** `s` is an empty string `""` after removing non-alphanumeric characters. Since an empty string reads the same forward and backward, it is a palindrome.

## Constraints
- `1 <= s.length <= 2 * 10^5`
- `s` consists only of printable ASCII characters.

## Solution 
```go
func isPalindrome(s string) bool {
	var b []rune
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			b = append(b, v)
		} else if v >= 'A' && v <= 'Z' {
			b = append(b, v+32)
		}
	}

	l, r := 0, len(b)-1
	for l <= r {
		if b[l] != b[r] {
			return false
		}
		l++
		r--
	}
	return true
}
```