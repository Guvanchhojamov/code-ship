package sliding_window

import "math"

/*

Given two strings s and t of lengths m and n respectively,
return the minimum window substring of s such that every character in t
(including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.


Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.


Constraints:

m == s.length
n == t.length
1 <= m, n <= 105
s and t consist of uppercase and lowercase English letters.


Follow up: Could you find an algorithm that runs in O(m + n) time?

*/
/*
1. brute force with hash map:
	tc: n*n
	sc: m
2. Can we optimize with sliding window approach?
	maybe we can.
	Input: s = "ADOBECODEBANC", t = "ABC"
1. create preinserted map with t string chars.
2. start from s[0] decrease char freq if:
	tch[s[r]] --
	if tm[s[r]] < 0: count ++
	while count == k:
		tm[s[l]]++
		if tm[s[l]] > 0: count--
		if min_len < r-l+1:
			min_len = r-l+1
			starti=left.
		l++
	r++
	return [starti:starti+min_len-1]

	tc: O(n+n)
	sc: (n+m)
*/
func minWindow(s string, t string) string {
	if len(s) == 0 {
		return ""
	}
	var mp = make(map[byte]int, len(s)+len(t))
	var min_len = math.MaxInt
	var starti = -1
	for i := range t {
		mp[t[i]]++
	}

	var i, j = 0, 0
	var count = 0
	for j < len(s) {
		if mp[s[j]] > 0 {
			count++
		}
		mp[s[j]]--
		for i < len(s) && count == len(t) {
			if mp[s[i]] == 0 {
				count--
			}
			if j-i+1 < min_len {
				min_len = j - i + 1
				starti = i
			}
			mp[s[i]]++

			i++
		}
		j++
	}
	//	fmt.Println(min_len, count, starti)
	if min_len > len(s) || starti < 0 {
		return ""
	}
	return s[starti : starti+min_len]
}
