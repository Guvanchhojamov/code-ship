package bit

/*
784. Letter Case Permutation
Given a string s, you can transform every letter individually to be lowercase or uppercase to create another string.
Return a list of all possible strings we could create. Return the output in any order.

Example 1:
Input: s = "a1b2"
Output: ["a1b2","a1B2","A1b2","A1B2"]

Example 2:
Input: s = "3z4"
Output: ["3z4","3Z4"]

Constraints:
1 <= s.length <= 12
s consists of lowercase English letters, uppercase English letters, and digits.
*/
/*
brute force or recursion:
we go with pick and not pick method here:
Since we need ony chars, if the s[i] is num we just skip.
base case:
	if i == len(s):
		ans = ans.append(s)
		return
	// not pick
	fn(ans, s, i+1):
	// pick
	if s[i] in [a..z]{
		s[i] = s[i] ^ 32  // make it uppercase or lower case.
	}
	fn(ans,s,i+1)
*/

func letterCasePermutation(s string) []string {
	var ans = make([]string, 0, 1<<len(s)) // 2^n
	var c = []byte(s)
	fn(&ans, c, 0)
	return ans
}

/*
s = "c"
*/
func fn(ans *[]string, c []byte, i int) {
	// base case
	if i == len(c) {
		*ans = append(*ans, string(c))
		return
	}
	// not pick - not change to uppercase.
	fn(ans, c, i+1)

	//pick - change to uppercase if ith element is char.
	if isLetter(c[i]) {
		c[i] ^= 32 // toggle the case upper | letter
		fn(ans, c, i+1)
		c[i] ^= 32 // toggle back to original
	}
}
func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// how can we do this using bit manipulation?
