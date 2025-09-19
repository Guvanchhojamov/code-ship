package sliding_window

/*
424. Longest Repeating Character Replacement

You are given a string s and an integer k.
You can choose any character of the string and change it to any other uppercase English character.
You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

Example 1:
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.

Example 2:
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achieve this answer too.

Constraints:

1 <= s.length <= 10^5
s consists of only uppercase English letters.
0 <= k <= s.length
*/
/*
- Need replace k chars with any.
- return max length we can generate.
brute force:
1. Gerate all substrings.
2. if a[i] != a[i-1] count r;
3. if r>k break.
4. each time check length with j-i+1.
return max len.
tc: n*n
ec: 1

Can we optimze?
	maybe use sliding window wich containes less<k diff elements and keep window size,
	comparing each time with max?
Input: s = "AABABBA", k = 1
Output: 4
i=j=0
1. take i,j pointers from 0,0 index.
2. increase window size each time j++
2. if a[i]!=a[j] then count r.
3. if r>k then:
	until a[i]!=a[j] : i++   // shrink the window.
   r-- decrease r;
4. each time check with max(mx,j-i+1)
 return mx.
 tc: O(n)
 sc: O(1)
*/
func characterReplacement(s string, k int) int {
	var c = []byte(s)
	var freq = make(map[byte]int, 26)
	var mxFreq = 0
	var mxLen = 0
	var i, j = 0, 0

	for j < len(c) {
		freq[c[j]]++
		mxFreq = max(mxFreq, freq[c[j]])
		if (j-i+1)-mxFreq > k {
			freq[c[i]]--
			i++
		}
		mxLen = max(mxLen, j-i+1)
		j++
	}
	return mxLen
}
