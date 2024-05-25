# 344. Reverse String

## Problem Description

**Difficulty:** Easy

Write a function that reverses a string. The input string is given as an array of characters `s`.

You must do this by modifying the input array in-place with O(1) extra memory.

## Examples

**Example 1:**

```
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
```

### Solution
```go
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[r], s[l] = s[l], s[r]
		l++
		r--
	}
}
```