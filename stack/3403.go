package stack

/*
3403. Find the Lexicographically Largest String From the Box I

You are given a string word, and an integer numFriends.
Alice is organizing a game for her numFriends friends.
There are multiple rounds in the game, where in each round:
  - word is split into numFriends non-empty strings, such that no previous round has had the exact same split.

All the split words are put into a box.
Find the lexicographically largest string from the box after all the rounds are finished.

Example 1:
Input: word = "dbca", numFriends = 2
Output: "dbc"

Explanation:
All possible splits are:

"d" and "bca".
"db" and "ca".
"dbc" and "a".

Example 2:
Input: word = "gggg", numFriends = 4
Output: "g"

Explanation:
The only possible split is: "g", "g", "g", and "g".

Constraints:

1 <= word.length <= 5 * 10^3

	word consists only of lowercase English letters.

1 <= numFriends <= word.length

Hint: Find lexicographically largest substring of size n - numFriends + 1 or less starting at every index.
Brute force:
- for each i take n-numFriends+1 chars and check is it lexicographically larger.
keep track larger subarray with len(n-friends+1)
return this subarray as ans.
- helper function isLexGreater(a,b)
- isLexGreater(a,b string):
tc: n*n
sc: 1
i- starting index of substring
i+n - numFriends + 1 - ending index of each substring.
min(i+n - numFriends + 1, n) - or n if out of bounds.
*/
func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	var n = len(word)
	var res = ""
	for i := 0; i < n; i++ {
		subStr := word[i:min(i+n-numFriends+1, n)]
		//  greater := lexGreater(res,subStr)
		res = max(res, subStr)
	}
	return res
}

func lexGreater(a, b string) string {
	var minlen = min(len(a), len(b))
	for i := 0; i < minlen; i++ {
		if a[i] != b[i] {
			if a[i] > b[i] {
				return a
			} else {
				return b
			}
		}
	}
	// if the string is prefix or same until minlen. then return larger one.
	// if equal return any of them
	if len(a) > len(b) {
		return a
	} else if len(b) > len(a) {
		return b
	}
	return a // or b
}
