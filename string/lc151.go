package string

import "strings"

/*
151. Reverse Words in a String
Given an input string s, reverse the order of the words.
A word is defined as a sequence of non-space characters.
The words in S will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words.
The returned string should only have a single space separating the words.
Do not include any extra spaces.



Example 1:

Input: s = "the sky is blue"
Output: "blue is sky the"
Example 2:

Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.

1 <= s.length <= 10^4
s contains English letters (upper-case and lower-case), digits, and spaces ' '.
There is at least one word in s.

Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?
*/

/*
  1. Brute force:
    res =
    Start from back if char add to c and.
    for i=n to 0:
    if s[i] == ' ' and c > 0 :
        result = append(reverseWord(c));
        c = ''
        if i>0 append(res, ' ')

    // after for loop check for c
    if c is non empty do the same logic but dont add space because it is the last word:
        result = append(reverseWord(c));
    tc: O(n*maxWordLen)
    sc: O(maxWordLen) + result string.
Can we optimize it?
"  hello world  "
"  hello world  "
 0123456789
*/
/*
   We can use split() functio and split words in words array.
   Then we iterate words array from back to front,
   using join() with " ", and after end we need trim words string for remove space in end.
   tc: O(n) + O(n) = O(n)
   sc: O(n)

*/

func reverseWords(s string) string {
	var result []string
	var words = make([]string, 0, len(s))
	words = strings.Fields(s)
	for i := len(words) - 1; i >= 0; i-- {
		result = append(result, words[i])
	}
	return strings.Join(result, " ")
}
