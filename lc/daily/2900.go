package daily

/*
2900. Longest Unequal Adjacent Groups Subsequence I
You are given a string array words and a binary array groups both of length n,
where words[i] is associated with groups[i].

Your task is to select the longest alternating subsequence from words.
A subsequence of words is alternating
if for any two consecutive strings in the sequence,
their corresponding elements in the binary array groups differ.
Essentially, you are to choose strings such that adjacent elements
have non-matching corresponding bits in the groups array.

Formally, you need to find the longest subsequence of an array of indices
[0, 1, ..., n - 1] denoted as [i0, i1, ..., ik-1],
such that groups[ij] != groups[ij+1]
for each 0 <= j < k - 1 and then find the words corresponding to these indices.

Return the selected subsequence. If there are multiple answers, return any of them.

Note: The elements in words are distinct.

Example 1:
Input: words = ["e","a","b"], groups = [0,0,1]
Output: ["e","b"]
Explanation: A subsequence that can be selected is ["e","b"]
because groups[0] != groups[2].
Another subsequence that can be selected is ["a","b"] because groups[1] != groups[2].
It can be demonstrated that the length of the longest subsequence of indices that satisfies the condition is 2.

Example 2:
Input: words = ["a","b","c","d"], groups = [1,0,1,1]
Output: ["a","b","c"]

Explanation: A subsequence that can be selected is ["a","b","c"]
because groups[0] != groups[1] and groups[1] != groups[2].
Another subsequence that can be selected is ["a","b","d"] because
groups[0] != groups[1] and groups[1] != groups[3].
It can be shown that the length of the longest subsequence of indices that satisfies the condition is 3.

Constraints:
1 <= n == words.length == groups.length <= 100
1 <= words[i].length <= 10
groups[i] is either 0 or 1.
words consists of distinct strings.
words[i] consists of lowercase English letters.
*/
/*
  given words, and  groups
   we need to find subsequence where:
	groups[i] != groups[i+1]
	in all these subsequences we need to find longest.
	len(words) == len(groups)
	words elements are - distinct.
  brute force:
	res = []string{w[0]}
	count = 1
	mx = math.MinInt
	si, ei = 0,1
	for i=1 to len(g)-1
	if 	g[i] != g[i-1]:
		ei=i
		count++
	else:
		if count > mx:
			mx = count
			res = w[si:ei+1]
			si=ei
			count=1
	// check again.
	if count > mx:
			mx = count
			res = w[si:ei+1]
			si=ei
	return res

words = ["a","b","c","d"], groups = [1,0,1,1]
res = a
si=0 ei=1
res = a,b
si=1 ei=2 c=2
res = bc
 ok sime works:
	tc: n
	sc: n
but, i think sliding window would be best strategy in there.
	start window with size=1
	start=0 end=1
	if i!=i+1:
		end++
	else:
	  start++
*/
/*
 After showing the editorial we can use second approach
	We can implement the second greedy approach
*/
// 1011
func getLongestSubsequence(words []string, groups []int) []string {
	var res = make([]string, 0, len(words))
	res = append(res, words[0])
	for i := 1; i < len(groups); i++ {
		if groups[i] != groups[i-1] {
			res = append(res, words[i])
			continue
		}
	}
	return res
}

/*
 iterate over groups and select the firs index of segment. To get longest we need to take from all segments.
*/
