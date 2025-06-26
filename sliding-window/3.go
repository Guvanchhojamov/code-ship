package sliding_window

/*
3. Longest Substring Without Repeating Characters
Given a string s, find the length of the longest substring without duplicate characters.

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/
/*
 1. brute force solution is:
  generate all subarrays and check.
  check repeating using hashmap.
  i=0->n; j=i->n
  if a[j] !in map:
	map.add(a[j])
	mx = max(mx, len(map))
*/
func lengthOfLongestSubstring(s string) int {
	var c = []byte(s)
	var mp = make(map[byte]int, len(s))
	var mx = 0
	for i := 0; i < len(c); i++ {
		mp = map[byte]int{}
		for j := i; j < len(c); j++ {
			if _, ok := mp[c[j]]; ok {
				break
			}
			mp[c[j]]++
			mx = max(mx, len(mp))
		}
	}
	return mx
}

// the answer is accepted but, can we optimize it?
/*
Input: s = "abcabcbb" Output: 3
 abc
*/

func lengthOfLongestSubstringSlWindow(s string) int {
	var mx = 0
	var l, r = 0, 0
	var c = []byte(s)
	var mp = make(map[byte]int, len(c))
	for r < len(c) {
		if _, ok := mp[c[r]]; ok {
			delete(mp, c[l])
			l++
		} else {
			mp[c[r]]++
			r++
		}
		mx = max(mx, len(mp))
	}
	return mx
}
