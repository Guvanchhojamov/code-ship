package daily

// 1910. Remove All Occurrences of a Substring
func removeOccurrences(s string, part string) string {
	var c []byte
	m := len(part)

	for i := 0; i < len(s); i++ {
		c = append(c, s[i])
		if len(c) >= m && string(c[len(c)-m:]) == part {
			c = c[:len(c)-m]
		}
	}
	return string(c)
}

/*
s = daabcbaabcbc   part = abc
    123456789
*/

/*
Given two strings s and part,
perform the following operation on s until all occurrences of the substring part are removed:
  - Find the leftmost occurrence of the substring part and remove it from s.
Return s after removing all occurrences of part.
A substring is a contiguous sequence of characters in a string.

Example 1:
Input: s = "daabcbaabcbc", part = "abc"
Output: "dab"
Explanation: The following operations are done:
- s = "daabcbaabcbc", remove "abc" starting at index 2, so s = "dabaabcbc".
- s = "dabaabcbc", remove "abc" starting at index 4, so s = "dababc".
- s = "dababc", remove "abc" starting at index 3, so s = "dab".
Now s has no occurrences of "abc".

Example 2:
Input: s = "axxxxyyyyb", part = "xy"
Output: "ab"
Explanation: The following operations are done:
- s = "axxxxyyyyb", remove "xy" starting at index 4 so s = "axxxyyyb".
- s = "axxxyyyb", remove "xy" starting at index 3 so s = "axxyyb".
- s = "axxyyb", remove "xy" starting at index 2 so s = "axyb".
- s = "axyb", remove "xy" starting at index 1 so s = "ab".
Now s has no occurrences of "xy".

Constraints:

1 <= s.length <= 1000
1 <= part.length <= 1000
s​​​​​​ and part consists of   English letters.

dsababcc  part = abc
dsabbbcc  part = abc

dsababcc
01234567
*/

/*
   Key points:
    - remove all occurences of part.
    - remove the first occurence each time.
    - for cover all test cases each time we need to compare element with part[last_element]
    - if is are same continue from stack each char until part[0]

  1. Brute force using stack.

func removeOccurrences(s string, part string) string {
	var c []byte
	m := len(part)

	for i := 0; i < len(s); i++ {
		c = append(c, s[i])
		if len(c) >= m && string(c[len(c)-m:]) == part {
			c = c[:len(c)-m]
		}
	}
	return string(c)
}
/*

s = daabcbaabcbc   part = abc
    0123456789


*/
