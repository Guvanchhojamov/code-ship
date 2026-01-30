package dp
##### 

Given two strings str1 and str2, return the shortest string that has both str1 and str2 as subsequences. If there are multiple valid strings, return any of them.

A string s is a subsequence of string t if deleting some number of characters from t (possibly 0) results in the string s.

/*
We need to find min.len string such like: 
	a,b both are subsequence of this string. 

Input: 
	S1,s2 2 strings 
Output:
	S supersequence string. 
Conditions:
	S1,s2 english lowercase letters. 
 	S1,s2 len; min:0 max:10^3. 
	Is this always possible? Yes. empty string or a+b string in worth case. 	
*/
/*
s1=abde  s2:bdc 
 With: abde+bdc 
 min: abdce 
  We can got abde removing, c; 
  We can got bdc removing a..e; 

How we can find this? 
How we can make this min ass possible? 
What is max common ss of both? 	
	bd so, we need additional 3 chars to make them supersequence. 
But we need build this str nether than just return count. 
 
 Can we try all possible ways and got min result from this? Maybe we can. But how we know is this are a is subsequent of this or not? 
 Approach-1: 
find lcs(a,b)
Take a add chars not lcs.
Take b add chars not in lsc. 
Add lcs to res. 
TC: N*M + N+M;
SC: N*M
 Can we use another approach?
	Yes we can. 
Find min num chars we need to add with dp or backtrack, or tabulation. 
Start from backward building string. 
If s1[i] == s[j] add any of them to res and move dp[i][j]
If !=; move to min(left,top) and add opposite char; if top add left to res, if left add top to res. 
return builded string in reversed order. 
 How we can find min add chars? 
	Base case: 
	 if empty dp[0][0]= 0 
	Otherwise dp[i][0] = i;;.. And d[0][j] = j; 
 if s1[i] == s2[j]: 
	Set prev  diagonal. 
 If s1[i] != s2[j]:
	take min(top,left)+1. 
 Dp[n][m] - is min needed chars to make supersequence from both. 
*/

func shortestCommonSupersequence(str1 string, str2 string) string {
	if str1 == str2 {
		return str1
	}
	s1, s2 := []byte(str1), []byte(str2)
	n, m := len(s1), len(s2)

	cache := make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, m+1)
		cache[i][0] = i
		for j := range cache[i] {
			cache[0][j] = j
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				cache[i][j] = 1+ cache[i-1][j-1]
			} else {
				cache[i][j] = 1 + min(cache[i-1][j], cache[i][j-1])
			}
		}
	}
    
	r := getResult(cache, s1, s2)

	return r
}

func getResult(cache [][]int, s1, s2 []byte) string {
	res := []byte{}
	n, m := len(cache), len(cache[0])
	i, j := n-1, m-1
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			res = append(res, s1[i-1])
			i--
			j--
		} else if cache[i-1][j] < cache[i][j-1] {
			res=append(res, s1[i-1])
			i--
		} else {
			res=append(res, s2[j-1])
			j--
		}
	}
    
    for i>0{
        res=append(res, s1[i-1])
        i--
    }
    for j>0{
        res=append(res, s2[j-1])
        j--
    }
	
	return reverse(res)
}

func reverse(c []byte) string {
	l, r := 0, len(c)-1
	for l < r {
		c[l], c[r] = c[r], c[l]
		l++
		r--
	}
	return string(c)
}



