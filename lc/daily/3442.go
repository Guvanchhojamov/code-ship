package daily

import "math"

/*3442. Maximum Difference Between Even and Odd Frequency I

You are given a string s consisting of lowercase English letters.

Your task is to find the maximum difference diff = freq(a1) - freq(a2) between the frequency of characters a1 and a2 in the string such that:

- a1 has an odd frequency in the string.
- a2 has an even frequency in the string.
 Return this maximum difference.

Example 1:

Input: s = "aaaaabbc"

Output: 3

Explanation:

The character 'a' has an odd frequency of 5, and 'b' has an even frequency of 2.
The maximum difference is 5 - 2 = 3.

Example 2:
Input: s = "abcabcab"

Output: 1

Explanation:

The character 'a' has an odd frequency of 3, and 'c' has an even frequency of 2.
The maximum difference is 3 - 2 = 1.

Constraints:

3 <= s.length <= 100
s consists only of lowercase English letters.
s contains at least one character with an odd frequency and one with an even frequency.
*/

/*
	We need odd frequency and even frequency from string.
	a1-a2 must be maximum.

Brute force:
- create hash map char:frequency.
- sort in increasing order.
- take last 2 non equal nums.
- return abs(last1-last2)
tc: n+nlogn+n = n+nlogn.
sc: N
2. Without sorting.
To get max difference we need subtract, possible max value from possible minimum value. So
for odds we need to find max, and for evens we need to find min possible value to make result maximum as possible.
  - create hashmap char:freq
  - create oddMax, evenMax = minInt, minInt
  - iterate over map:
    if freq mod 2==0:
    min(evenMin,freq)
    if freq mod 2!=0:
    oddMax = min(oddMax, freq)

return oddMax - evenMax.
tc: N+N
sc: N
*/
func maxDifference(s string) int {
	var freqMap = make(map[rune]int, 0)
	var evenMin, oddMax = math.MaxInt, math.MinInt
	for _, v := range s {
		freqMap[v]++
	}
	for _, freq := range freqMap {
		if freq%2 != 0 {
			oddMax = max(oddMax, freq)
		} else {
			evenMin = min(evenMin, freq)
		}
	}
	return oddMax - evenMin
}
