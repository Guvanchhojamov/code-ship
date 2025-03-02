package string

/*
The beauty of a string is the difference in frequencies between the most frequent and least frequent characters.

For example, the beauty of "abaacc" is 3 - 1 = 2.
Given a string s, return the sum of beauty of all of its substrings.



Example 1:

Input: s = "aabcb"
Output: 5
Explanation: The substrings with non-zero beauty are ["aab","aabc","aabcb","abcb","bcb"], each with beauty equal to 1.
Example 2:

Input: s = "aabcbaa"
Output: 17


Constraints:

1 <= s.length <= 500
s consists of only lowercase English letters.


*/
/*
   1. Brute force approach
     inerate string in two  loop:
     genereate all substring
     add char frequency in frequencyMAp
     take max and min from this map,
     subtract max-min and add to ans.
     tc:n*n*26
     sc:26
*/

func beautySum(s string) int {
	var ans = 0
	var freqMap = make(map[byte]int, 26)

	for i := 0; i < len(s); i++ {
		freqMap = map[byte]int{}
		for j := i; j < len(s); j++ {
			freqMap[s[j]]++
			maxFreq, minFreq := 0, len(s)
			for _, freq := range freqMap {
				if freq > maxFreq {
					maxFreq = freq
				}
				if freq < minFreq {
					minFreq = freq
				}
			}
			ans += maxFreq - minFreq
		}
	}
	return ans
}
