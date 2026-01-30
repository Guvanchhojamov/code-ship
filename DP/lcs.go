package dp

/*
Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.

Example 1:

Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length 3.

/*
What is range for len text1, text2 ? 10^5.
Whats is text chars? Only lowercase english chars.
If there can be no common sbsequence? Return -1.
Input 2 text, output len of lognest common subsequence.
Lengs of 2 strings can be difference. Not equal always.
*/

/*
 Ok, we need to return common subsequense of 2 strings. What is common subsequence? Common non-continious characters with keeped order.
 So, whats is brute force?
text1 = "abcde", text2 = "ace"; out=3.
 How we can do it with simple way?
 Generate all possible subsequences of both texts. Store them in somewhere. And take longest from them.
It will take 2^n * 2^m TC and similar space complexity.
 Maybe we can dompare them simpler.
Where there come to generate all possibilities we use recursion.
Recursion with top-down approach.
How we can do this with top-down approach?
Start from top (end of both texts).
  n-1, m-1.
"abcde", text2 = "ace"
 What we need? We need common chars. From end->begin.
Ok, there can be some cases, and on each case we need to do something.
For recursion we need 2 case. Base case. Recursion case.
 // base case.
 We have 2 pointers so they can go to <0.
What if i < 0? We dont have we return 0.
What if j<0? We dont have other ops return 0. 0 -there is min possible number. We dont need return maxInt or minint.
If i<0 ||  j<0:
return 0 - we tried all and didnt find anything.

  // we come back after.
What cases ther can be with chase? And what we need to do in each case?
Assume i,j each char.
S1[i] == s2[j]:
In this case we found some equality, some subsequence so increase count 1+ and go further to find anothers. Incrementing both pointers
return 1+ fn(i-1,j-1)
If s1[i] != s2[i]:
In this case we have 2 choices we can search s1.car in s2 incrementing j pointer.
Or we can search in s1; s2.char incementing i to next char.
What we need is the longest, so we do both. And take MAX from them.

At the end we got max possible count.
Ok, lets implement this one.
"abcde", text2 = "ace"
i
What if one of them empty?
If longest is 0 return -1	       j
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	dp := make([][]int, n)
	for i := range dp[i] {
		dp[i] = make([]int, m)
		for i := range dp[i] {
			dp[i][j] = -1
		}
	}
	return fnlcs(n-1, m-1, dp, text1, text2)
}
func fnlcs(i, j int, dp [][]int, s1, s2 string) int {
	// base case
	if i < 0 || j < 0 {
		return 0
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	//recursion case
	if s1[i] == s2[j] {
		return 1 + fnlcs(i-1, j-1, dp, s1, s2)
	}
	first := fnlcs(i-1, j, dp, s1, s2)
	second := fnlcs(i, j-1, dp, s1, s2)
	dp[i][j] = max(first, second)
	return dp[i][j]
}

/*
 How we can do this using tabulation?
What is base case?
How many sub. we can have with 0 s1, s2 ?
 its.0
If any of texts are 0 we got 0 always.

start from 1,1.
if they are equal add +1 to
max of text1 until now, max of text2 until now.
and store curr equal cell.
if they are not equal just store max of
 text1 until know text2 until know.
At the end return dp[n][m] value.
This is our answer.

*/

func longestCommonSubsequenceOpt(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1] // because we move both pointers.
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[n][m]
}

/* wrong answer…
    j m j k b k j k v
 [0 0 0 0 0 0 0 0 0 0]
b[0 0 0 0 0 1 1 1 1 1]
s[0 0 0 0 0 1 1 1 1 1]
b[0 0 0 0 0 2 2 2 2 2]
i[0 0 0 0 0 2 2 2 2 2]
n[0 0 0 0 0 2 2 2 2 2]
i[0 0 0 0 0 2 2 2 2 2]
n[0 0 0 0 0 2 2 2 2 2]
m[0 0 1 1 1 2 2 2 2 2]
*/
