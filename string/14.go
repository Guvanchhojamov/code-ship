package string

//  14. Longest Common Prefix
/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters if it is non-empty.
*/
/*
 1. Take first string
    check each char of array and this string,
    chars must be same in same index,
    if not same return result
    otherwhise add char, to result
*/

func longestCommonPrefix(strs []string) string {
	var res []byte

	var s = strs[0]
	for i := 0; i < len(s); i++ {
		for _, word := range strs {
			if i >= len(word) || s[i] != word[i] {
				return string(res)
			}
		}
		res = append(res, s[i])
	}
	return string(res)
}
