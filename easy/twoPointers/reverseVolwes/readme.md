# 345. Reverse Vowels of a String

## Problem Statement
Given a string `s`, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

## Examples

### Example 1:
- **Input:** `s = "hello"`
- **Output:** `"holle"`

### Example 2:
- **Input:** `s = "leetcode"`
- **Output:** `"leotcede"`

## Constraints:
- `1 <= s.length <= 3 * 10^5`
- `s` consists of printable ASCII characters.

## Solution

```go
var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func reverseVowels(s string) string {
	b := []rune(s)
	l, r := 0, len(b)-1
	for l < r {
		if vowels[b[l]] {
			if vowels[b[r]] {
				b[l], b[r] = b[r], b[l]
                l++
                r--
			}else{
                r--
            }
		}else{
            l++
        }
	}
	return string(b)
}

```