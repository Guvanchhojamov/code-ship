package string

/*
1614. Maximum Nesting Depth of the Parentheses
Given a valid parentheses string s,
return the nesting depth of s.
The nesting depth is the maximum number of nested parentheses.


Example 1:

Input: s = "(1+(2*3)+((8)/4))+1"

Output: 3

Explanation:

Digit 8 is inside of 3 nested parentheses in the string.


Example 2:
Input: s = "(1)+((2))+(((3)))"
Output: 3
Explanation:
Digit 3 is inside of 3 nested parentheses in the string.

Example 3:
Input: s = "() (()) ((()()))"
Output: 3



Constraints:

1 <= s.length <= 100
s consists of digits 0-9 and characters '+', '-', '*', '/', '(', and ')'.
It is guaranteed that parentheses expression s is a VPS.
*/

/*
  What we need find max depth of parathases:
    - we need just parentheses
    - parentheseses is always valid.
    - need to find max depth.
 1. Brute force.
    iterate over s
    Since parantheses is always valid we need count '(' or ')' and this is our answer.
      For to be valid parantheses always need to be balanced, so
      opensCount-closedCount == 0;
      opensCount == closedCount
    tc: O(n)
    sc: O(1)
 2. Iterate over s:
    if '(' count++
    else:
    count--
    result = max(result, count)
   return result.
  tc:O(n)
  sc:1

*/

func maxDepth(s string) int {
	var depth int
	var result int
	for _, char := range s {
		if char == '(' {
			depth++
		} else if char == ')' {
			depth--
		}
		result = max(result, depth)
	}
	return result
}
