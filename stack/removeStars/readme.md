# 2390. Problem: Removing Stars From a String

## Description

You are given a string `s`, which contains stars `*`.

In one operation, you can:
- Choose a star in `s`.
- Remove the closest non-star character to its left, as well as remove the star itself.

Return the string after all stars have been removed.

**Note:**
- The input will be generated such that the operation is always possible.
- It can be shown that the resulting string will always be unique.

## Examples

### Example 1
**Input:**
s = "leet**cod*e"

**Explanation:**
- Performing the removals from left to right:
    - The closest character to the 1st star is 't' in "leet**cod*e". `s` becomes "lee*cod*e".
    - The closest character to the 2nd star is 'e' in "lee*cod*e". `s` becomes "lecod*e".
    - The closest character to the 3rd star is 'd' in "lecod*e". `s` becomes "lecoe".
    - There are no more stars, so we return "lecoe".
## Solution: 
```go
func removeStars(s string) string {
	var stack []rune
	for _, v := range s {
		if v == '*' && len(stack) > 0 {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, v)
	}
	return string(stack)
}
```