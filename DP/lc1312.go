package dp

/*
 Input:
string s
 Output:
Number of steps int.
Conditions:
 String range: min:0 - empty - palindrome.
 	Max: 10^3
 S - only lowercase english letters? Yes; 26 chars.
 S - if no is the case matters? Yes. But its only lowercase letters.
  Is this possible always? Yes we can make palindrom any string. How?
	Add same string in backwards.
*/
/*
 s = “abdca” -?
   abdc + cdb + a. - result is 3.
   abdccdba
“abdcaacdba” - in the worth case its len(s).
But this is max we cannot do this.

“abdca”
“acdba”
To make palindrome some string we need to take max possible palindrom number of chars. Then add other chars.
In this example we can see, we have 2 common subsequence, right.
If we find len if this, we find len of others, too.
len(s) = n
X - common subsequence.
n-x - min chars what we need to make this s palindrome.
res = len(s) - num_comm_subequence(s, reversed_s)
“aaa”
“aaa”
 if s == reversed(s) return 0 immediately.
“abb”
“bba”
Lcs=2; n=3
Res = n - lcs = 3-1 =1
abb + a= abba.
Reverse s.
Compare with reversed
Find LCS(s,rev_s)
return len(s) - LCS.
TC: N-reverse; +  N*N-LCS = O(N^2)
SC: N*N - for LCS. = O(n^2).

How we reverse? Simple helper.
How we find lcs? DP helper
*/

func minInsertions(s string) int {
	str, n := []byte(s), len(s)
	rev_str := make([]byte, len(str))
	copy(rev_str, str)
	rev_str = reverse(rev_str)

	if string(str) == string(rev_str) {
		return 0
	}

	return n - lcs(str, rev_str)
}

func lcs(str, rev_str []byte) int {
	n := len(str)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if str[i-1] == rev_str[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i][j], dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][n]
}

func reverse(s []byte) []byte {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}
