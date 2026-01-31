package dp

/*
Given two strings s and t, return the number of distinct subsequences of s which equals t.
The test cases are generated so that the answer fits on a 32-bit signed integer.
*/

/*
 Input:
   S,t string
 Output:
   count int

Conditions:
S,t: range ? 0<=len(s,t)<= 10^3.
S,t: are only lowercase english? Yes.
s >= t always. Right? Yes.
What if no any subsec? Return 0. Yes.
Subsequences must be taken from unique chars? Yes,
For example 0,3 is “a” then we take each just once.
It must be unique seubcequence.
*/
/*
 s=baceef; t=bef; res=2
 b+e+f;
   3
 b+e+f
   4
 We have only 2 unique ssq in s.
“a”; “”; return 0. If one empty also return 0.
*/
/*
 s=baceef; t=bef;
Generate all subsequences of S. and store them in map. Key-sqq; val=count of sub. We can do this with backtracking; 2^N.
At the end return from this map map[t].val. It return count.
TC: 2^n
SC: 2^n -we store in map.

#2
 s=baceef; t=bef;
       i      j
We can generate ssq with top-down approach, and when ssq = target just increment totalCount.
Because we do this with recursion, we ned 2 cases:
 // base case
   If i<0 and j>=0: return 0.
   if i<0 and j<0: return 1
   if i>=0 and j<0: return 1
 // recursion case.
   if s[i] == s[j]:
 f: fn(i-1,j-1) + fn(i-1,j)
   Else:
 s: fn(i-1,j)
  return sum(f,s)

TC: 2^n
SC: 2^n - for rec. Stack.

Since we do some overlapping problems we can memoization this.
Ok, lets implement first recursive way, then optimze it.
*/

func numDistinct(s string, t string) int {
	n, m := len(s), len(t)
	if m > n || (n == 0 || m == 0) {
		return 0
	}
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return fn(s, t, memo, n-1, m-1)
}

func fn(s, t string, memo [][]int, i, j int) int {
	//base case
	if (i < 0 && j < 0) || (i >= 0 && j < 0) {
		return 1
	}
	if i < 0 && j >= 0 {
		return 0
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	//recursive case
	equal, nonEqual := 0, 0
	if s[i] == t[j] {
		equal = fn(s, t, memo, i-1, j-1)
	}
	nonEqual = fn(s, t, memo, i-1, j)

	memo[i][j] = equal + nonEqual
	return memo[i][j]
}

/*
 Lets try optimize with tabulation.
*/

func numDistinctTb(s string, t string) int {
	n, m := len(s), len(t)
	tb := make([][]int, n+1)
	for i := range tb {
		tb[i] = make([]int, m+1)
	}
	for i := range tb {
		tb[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				tb[i][j] = tb[i-1][j-1] + tb[i-1][j]
			} else {
				tb[i][j] = tb[i-1][j]
			}
		}
	}
	return tb[n][m]
}

/*
s=baceef; t=bef;
      b e f
    0 1 2 3
  0 1 0 0 0
b 1 1 1 0 0
a 2 1 1 0 0
c 3 1 1 0 0
e 4 1 1 1 0
e 5 1 1 2 0
f 6 1 1 2 2

if s[i] == t[j]:
	tb[i][j]=tb[i-1] + tb[i-1][j-1]
Else:
	tb[i][j]=tb[i-1]
return tb[n][m]

*/
