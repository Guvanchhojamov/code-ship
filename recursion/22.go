package recursion

/*
22. Generate Parentheses
Given n pairs of parentheses,
write a function to generate all combinations of well-formed parentheses.

Example 1:
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:
Input: n = 1
Output: ["()"]

Constraints:
1 <= n <= 8
*/
/*
  We need to generate all parentheses, for given length.
  What is the valid parentheses?
   ps witch length of open and closed arrows are same and
we keep track opening and closing parentheses so
   if open count != n we add open
   if open count > closed count we add closed
we go like this until we used all open and closed parentheses.
*/
/*
var result = []string
var curr = []rune
var open = 0
close = 0
  // base case:
	if open > n or close > n || open <  close:
		return res
	if open == n && close == n:
		res = append.curr
		return
// recursion case:
		curr = append."("
		fnP(res, curr, open+1, close)
	    curr = remove."("

		curr = append.")"
        fnP(res, curr, open, close+1)
		curr = remove.")"
*/
func generateParenthesis(n int) []string {
	var res = make([]string, 0)
	var curr = make([]rune, 0)
	var open, close = 0, 0
	fnP(&res, curr, n, open, close)
	return res
}

func fnP(res *[]string, curr []rune, n, open, close int) {
	// base case-1
	if open > n || close > n || open < close {
		return
	}
	if open == n && close == n && open == close {
		*res = append(*res, string(curr))
		return
	}
	curr = append(curr, '(')
	fnP(res, curr, n, open+1, close)
	curr = curr[:len(curr)-1]
	curr = append(curr, ')')
	fnP(res, curr, n, open, close+1)
	curr = curr[:len(curr)-1]
	return
}
