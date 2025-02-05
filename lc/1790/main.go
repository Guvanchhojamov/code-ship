package main

func main() {

}

/*
1790. Check if One String Swap Can Make Strings Equal
You are given two strings s1 and s2 of equal length. A string swap is an operation where you choose two indices in a string
 (not necessarily different) and swap the characters at these indices.
Return true if it is possible to make both strings equal by performing at most one string swap on exactly one of the strings. Otherwise, return false.


Example 1:
Input: s1 = "bank", s2 = "kanb"
Output: true
Explanation: For example, swap the first character with the last character of s2 to make "bank".

Example 2:
Input: s1 = "attack", s2 = "defend"
Output: false
Explanation: It is impossible to make them equal with one string swap.

Example 3:
Input: s1 = "kelb", s2 = "kelb"
Output: true
Explanation: The two strings are already equal, so no string swap operation is required.


Constraints:

1 <= s1.length, s2.length <= 100
s1.length == s2.length
s1 and s2 consist of only lowercase English letters.
*/
/*
  We need try to swap chars to get s1 and s2 equal.
    Key poins:
        - s1, s2 length are equal.
        - We need to do at most 1 operation, not more.
        - We need swap chars in exacty one string.
  Since range is small we can use brute force.
  1. Brute force.
     We need to genereate, s2 to get s1 if we cant return flase.
     If swaps more than 1 return false.
     n = len(s1) // == s2
     i,j = 0,0
       for i<n, j<n:
			if s1[i] != s2[j]:
             for s1[i]== s2[j] and j<n:
                 j++
			 if j >= n return
               false
             else
                   swap(s2[i], s2[j])
                    count++
            i++
        	j=i
		if s1 != s2 or count > 1 return false.
  return true
  ex:1
	s1 = "bank", s2 = "kanb"
   i=0
   j=0
"kelb"  , "kelb"

*/
/*
class Solution:
    def swapChars(self, s: str, i: int, j: int) -> str:
        if i >= len(s) or j >= len(s) or i < 0 or j < 0:
            return s

        lst = list(s)
        lst[i], lst[j] = lst[j], lst[i]
        return ''.join(lst)
    def areAlmostEqual(self, s1: str, s2: str) -> bool:
        n = len(s1)
        i,j,count=0,0,0
        while i<n and j<n:
            if s1[i] != s2[j]:
                while j<n and s2[i]!=s2[j]:
                    j+=1
                if j>=n:
                    return False
                else:
                    self.swapChars(s2,i,j)
                    count+=1
            if s1 != s2 or count > 1:
                return False
            i+=1
            j=i
        return True
*/
func areAlmostEqual(s1 string, s2 string) bool {
	var i, j, count = 0, 0, 0
	var n = len(s1)
	var c1, c2 = []byte(s1), []byte(s2)
	for i < n && j < n {
		if c1[i] != c2[j] {
			swapIndex, i, j := i, i+1, j+1
			for i < n && j < n && c1[i] == c2[j] {
				i++
				j++
			}
			if j >= n || i >= n {
				return false
			} else if c1[i] != c2[j] {
				c2[swapIndex], c2[j] = c2[j], c2[swapIndex]
				count++
			}
			if string(c1) != string(c2) || count > 1 {
				return false
			}
		}
		i++
		j++
	}
	return true
}

/*
 1. Brute foce:
    Key points:
        - setirngs are equal
        - swap must be 1
        - we can swap exactly one string

*/
