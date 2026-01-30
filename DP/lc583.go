package dp

/*
Given two strings word1 and word2,
return the minimum number of steps required to make word1 and word2 the same.
In one step, you can delete exactly one character in either string.
*/
/*
 Input:
	Word1, word2 string
 Output: num of steps to make strings equal.

Can delete from either 1 char.

Constraints:
Word1, word2 - only lowercase english letters.
minLen: 0; maxLen: 10^3
Always possible? Yes delete all from both, its max steps.
*/

/*
 w1=abc; w2=adf
 Res = 4; we must delete  2 from each.

w1= axbde; w2 = xbdf
Res = 3; 2 from w1; 1 from w2.
Ok, if we see to strings.
 If we can find len of common substrings, and Max long among them,
And discard this from len(w1) and len(w2) and return sum of non common substrings. Its our answer.
 We just need length not common string itself.
So.
lcss(w1,w2)
len(w1)+len(w2) - 2 * lcss(w1,w2)
This is our ans.
Tc: O(n*m) - lcss;
Sc O(n*m) - lcss;
Any edge case ?
If already equal return 0.
If empty return 0


axbdef; afrtef
Lcss=3 (aef)
6+6 - 2*lcss = 12 - 6 = 6.
We need delete 6 chars and its 6 operations.
How to find Lcs len?
With DP.
We have another way todo this with DP. instead firn  longes lcs. We find min delete ops for each step.
 Create dp arr with n+1,m+1
	for each step, if they are equal, prev min val [i-1,j-1]
	If not eaqal 1 + min(top, left)
 return n-1,m-1
TC: n*m
SC: n*m
*/

func minDistance(word1, word2 string) int {
	w1, w2 := []byte(word1), []byte(word2)
	n, m := len(w1), len(w2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		dp[i][0] = i // we need deleete i chars
		for j := range dp[i] {
			dp[0][j] = j // we need delete j chars.
		}
	}
	//dp[0][0]=0  // if both empty no need delete.

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if w1[i-1] == w2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}
