# 28.Find the Index of the First Occurrence in a String

## Problem Statement

Given two strings `needle` and `haystack`, return the index of the first occurrence of `needle` in `haystack`, or `-1` if `needle` is not part of `haystack`.

## Examples

### Example 1:
- **Input:** `haystack = "sadbutsad"`, `needle = "sad"`
- **Output:** `0`
- **Explanation:** `"sad"` occurs at index `0` and `6`. The first occurrence is at index `0`, so we return `0`.

### Example 2:
- **Input:** `haystack = "leetcode"`, `needle = "leeto"`
- **Output:** `-1`
- **Explanation:** `"leeto"` did not occur in `"leetcode"`, so we return `-1`.

## Constraints:
- `1 <= haystack.length, needle.length <= 10^4`
- `haystack` and `needle` consist of only lowercase English characters.

### Solution 
```go
func strStr(haystack string, needle string) int {
  var l,r int 
  for l<len(haystack) {
	if haystack[l]==needle[r] {
		if r == len(needle)-1 {
			return l-r   
		}
		r++
	}else if r > 0{	  
		l-=r  // left-i yene onki yerine goyyas.
		r=0
	} 
	l++
  }
  return -1  
}
```