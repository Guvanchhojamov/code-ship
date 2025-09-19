package main

import (
	"fmt"
)

func main() {
	str := makeFancyString("leeetcode")
	fmt.Println(str)
}

func makeFancyString(s string) string {
	var result = []byte{}
	var count = 1
	result = append(result, s[0])
	for i := 1; i < len(s); i++ {
		if len(result) > 0 {
			switch s[i] == s[i-1] {
			case true:
				count++
				if count == 3 {
					count--
					continue
				} else {
					result = append(result, s[i])
				}
			default:
				result = append(result, s[i])
				count = 1
			}
		}
	}
	return string(result)
}

/*
 Input: s = "aaabaaaa"
 Output: "aabaa"

 res = aaa
 count = 3
  charStack = []byte
  for 1 to len(s):
    if s[i]==s[i-1]:
        c++
        if c == 3:
          continue
        res = append(res,s[i])
    else:
    c=1
     res = append(res,s[i])

   return string(stack)
*/

/*
 Repeat from daily lc.

1957. Delete Characters to Make Fancy String
A fancy string is a string where no three consecutive characters are equal.
Given a string s, delete the minimum possible number of characters from s to make it fancy.
Return the final string after the deletion.
It can be shown that the answer will always be unique.



Example 1:

Input: s = "leeetcode"
Output: "leetcode"
Explanation:
Remove an 'e' from the first group of 'e's to create "leetcode".
No three consecutive characters are equal, so return "leetcode".

Example 2:
Input: s = "aaabaaaa"
Output: "aabaa"
Explanation:
Remove an 'a' from the first group of 'a's to create "aabaaaa".
Remove two 'a's from the second group of 'a's to create "aabaa".
No three consecutive characters are equal, so return "aabaa".

Example 3:
Input: s = "aab"
Output: "aab"
Explanation: No three consecutive characters are equal, so return "aab".

Constraints:

1 <= s.length <= 10^5
s consists only of lowercase English letters.

-We need delete 1 char if char freq is >=3.
-return remain string after deletion.
"aaabaaaa"
1. Brute force.
   iterate from start to end.
   what delete means - is skip element not add to result.
   for i=0 until n-1:
	if a[i] == a[i+1]: count++
	for count >=2 && i<n-1:
		i++
	if a[i]!=a[i+1]: count--
	else:
		res.add(a[i])
	return res
tc: O(n)
sc: O(1) - if not take result string.
2. With stack ds:
	while i=0 to n-1:
	if st>0 and st.last == a[i]:
		count++
	while count >=2 and st.size>0 :
		st = st.pop()
		if st.Last != a[i] count--

	st.add(a[i])
	return st.
	tc: O(n)
	sc: O(n) for stack ds.
*/

func makeFancyString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result = []byte{}
	result = append(result, s[0])
	var count = 1

	for i := 1; i < len(s); i++ {
		if s[i] == result[len(result)-1] {
			count++
			if count < 3 {
				result = append(result, s[i])
			}
		} else {
			result = append(result, s[i])
			count = 1
		}
	}
	return string(result)
}
