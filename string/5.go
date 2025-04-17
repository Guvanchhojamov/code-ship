package string

/*
Given a string s, return the longest palindromic substring in s.

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"


Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.
*/

func longestPalindrome(s string) string {
	var res = make([]byte, 0, len(s))
	var resLen = 0
	var c = []byte(s)
	for i := 0; i < len(c); i++ {
		// for odd number of palindrome
		// assume center is 0 index.
		// left nd right poitners. Char with len=1 is also palindrom
		left, right := i, i
		for left >= 0 && right < len(c) && c[left] == c[right] {
			currLen := right - left + 1
			if currLen > resLen {
				res = c[left : right+1]
				resLen = currLen
			}
			left -= 1
			right += 1
		}

		// For even nummber of palindrome like //abbcc palindrome
		left, right = i, i+1
		for left >= 0 && right < len(c) && c[left] == c[right] {
			currLen := right - left + 1
			if currLen > resLen {
				res = c[left : right+1]
				resLen = currLen
			}
			left -= 1
			right += 1
		}
	}
	return string(res)
}

/*
  For palindromes we have 2 length
  odd and even.
  In odd each char, left and right to be same, just center element can be other .
  In even, each char must be same.
  1. Find for long odd len palindrome
  2. Find for long even len palindrome.

*/
