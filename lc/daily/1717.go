package daily

/*

1717. Maximum Score From Removing Substrings
You are given a string s and two integers x and y.
You can perform two types of operations any number of times.
- Remove substring "ab" and gain x points.
For example, when removing "ab" from "cabxbae" it becomes "cxbae".
- Remove substring "ba" and gain y points.
For example, when removing "ba" from "cabxbae" it becomes "cabxe".
- Return the maximum points you can gain after applying the above operations on s.



Example 1:

Input: s = "cdbcbbaaabab", x = 4, y = 5
Output: 19
Explanation:
- Remove the "ba" underlined in "cdbcbbaaabab". Now, s = "cdbcbbaaab" and 5 points are added to the score.
- Remove the "ab" underlined in "cdbcbbaaab". Now, s = "cdbcbbaa" and 4 points are added to the score.
- Remove the "ba" underlined in "cdbcbbaa". Now, s = "cdbcba" and 5 points are added to the score.
- Remove the "ba" underlined in "cdbcba". Now, s = "cdbc" and 5 points are added to the score.
Total score = 5 + 4 + 5 + 5 = 19.

Example 2:
Input: s = "aabbaaxybbaabb", x = 5, y = 4
Output: 20

Constraints:

1 <= s.length <= 10^5
1 <= x, y <= 10^4
s consists of lowercase English letters.
*/
/*
 Given string and x,y we need to remove "ab" or "ba"
  - need gain max possible points.
  - for each remove we can gain x or y points:
	"ab" - x ; "ba" - y;
- Since we need max possible result we need to do with grind approach, well
we try first remove "ab" or "ba" using x,y points:
	if x>y:
	 then priority of "ab" is 1, "ba" is 0
	if x<y:
	then priority of "ba" is 1, "ab" is 0.
	if x==y: order is doesnt matter.
		anyone can 1 or 0.
Simple approach:
	1. Define priority.
	2. create helper func witch interates over s:
		removes given str
		returns reamining string.
	3. for "ab" and "ba" call helper funcs.
tc: 2 times n; n+n =2n =n
sc: n+n in helper func we use stack ds. = n.
 Can we optimze the solution?
Input: s = "cdbcbbaaabab", x = 4, y = 5
Input: s = "aabbaaxybbaabb", x = 5, y = 4
Output: 20
1. define priority map:"ab":1;"ba":0 or reverse using x,y valuses.
2. iterate over s and create stack:
	if st.last_two is in map["ab"] and priority==1:
		add to res this pointMap[x]
	else if last_two==map and priority==0:
		then check with next char:
		if last+next in map and priority=1:
			add to res points[x]
			remove 1 from stack and skip index.
		else:
		  default behavior.
	continue adding to stack.
tc: n
sc: n
 ok, lets try implement second approach maybe we can do.
 No needed the optima lis two pass and gready approach
*/

// /.
func removeSubstringWithResult(s string, target string, points int) (int, string) {
	stack := make([]byte, 0, len(s))
	score := 0

	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])

		if len(stack) >= 2 &&
			stack[len(stack)-2] == target[0] &&
			stack[len(stack)-1] == target[1] {
			stack = stack[:len(stack)-2]
			score += points
		}
	}

	return score, string(stack)
}

func maximumGain(s string, x int, y int) int {
	var score1, score2 int
	var remaining string

	if x >= y {
		score1, remaining = removeSubstringWithResult(s, "ab", x)
		score2, _ = removeSubstringWithResult(remaining, "ba", y)
	} else {
		score1, remaining = removeSubstringWithResult(s, "ba", y)
		score2, _ = removeSubstringWithResult(remaining, "ab", x)
	}

	return score1 + score2
}
