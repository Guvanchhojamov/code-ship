package daily

/*
Given a string s and an integer k, return true if you can use all the characters in s to construct k palindrome strings or false otherwise.

Example 1:

Input: s = "annabelle", k = 2
Output: true
Explanation: You can construct two palindromes using all characters in s.
Some possible constructions "anna" + "elble", "anbna" + "elle", "anellena" + "b"
Example 2:

Input: s = "leetcode", k = 3
Output: false
Explanation: It is impossible to construct 3 palindromes using all the characters of s.
Example 3:

Input: s = "true", k = 4
Output: true
Explanation: The only possible solution is to put each character in a separate string.
*/

func main() {

}

/*
Given a string s and an integer k, return true if you can use all the characters in s to construct k palindrome strings or false otherwise.

Example 1:
Input: s = "annabelle", k = 2
Output: true
Explanation: You can construct two palindromes using all characters in s.
Some possible constructions "anna" + "elble", "anbna" + "elle", "anellena" + "b"

Example 2:
Input: s = "leetcode", k = 3
Output: false
Explanation: It is impossible to construct 3 palindromes using all the characters of s.
Example 3:

Input: s = "true", k = 4
Output: true
Explanation: The only possible solution is to put each character in a separate string.

Constraints:
1 <= s.length <= 105
s consists of lowercase English letters.
1 <= k <= 105

*/

/*
   1. 1 char is palindrome so if k==len(s) then true.

*/

/*
  Input: s = "annabelle", k = 2
  a = 2
  n = 2
  b = 1
  e = 2
  l = 2

  if map[i]==1 cnt++ continue
     s+=map[i]
  if s/2==0 cnt++
      if cnt == k || (k=1 && len(s)== s+cnt) return true
  return false.

  Input: s = "leeotcoode", k = 3
   l = 1
   e = 3
   t = 1
   c = 1
   o = 3
   d = 1

   s = 6 cnt= 4

*/

/*
s = "trueurt", k = 2   out= true
trrt + ueu
trueurt
rueur + tt
trueurt

 t=2
 r=2
 u=2
 e=1
 s=6  cnt=1

We need keep and track
    count of chars sequence
 and count odd sequence number chars
 and even sequence number chars.

if odd seq > k then the min numbers of polindromes we can construct > k so, it will we false.
Othervise it will be true.

*/

func canConstruct(s string, k int) bool {
	var oddCount, evenCount int
	var charSeq = make(map[rune]int, 0)
	if len(s) < k {
		return false
	}
	for _, char := range s {
		charSeq[char]++
	}

	for _, seq := range charSeq {
		if seq > 0 && seq%2 != 0 {
			oddCount++
		} else {
			evenCount++
		}
		if oddCount > k {
			return false
		}
	}
	return true
}
