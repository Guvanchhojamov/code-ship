package recursion

/*
131. Palindrome Partitioning
Given a string s,
partition s such that every substring of the partition is a palindrome.
Return all possible palindrome partitioning of s.

Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

Input: s = "a"
Output: [["a"]]


Constraints:

1 <= s.length <= 16
s contains only lowercase English letters.
*/

/*
 How can we solve this problem?
	1. Brute force
 		Generate all substrings and check if all parts are palindrome
	add to result?
 1. in one loop, we need partition each element and ad to and, this is by default 1 char is palindrome,
so we can partition them 1 char on each partition.
  we do same thing each time, take 1 part, and take 2 part if both palindrome then,
add it to ans, starting from each element in string.
  Can we use recursion in there, exactly this is the main recursion problem.
But, how?
*/
/*
"aab"
  Output: [["a","a","b"],["aa","b"]]
  1. string to []byte.
   res = [][]string the main result.
   partitions = []string.
  base case:
	if index == len(s):
       res = append.partitions;
    	return
  // recursion case.
	// each time start from index and check for each sustr is it palindrome
    //  if palindrome we can partition it.
	for i:=index; i<len(c); i++:
		partition = c[index:i+1]
		if isPalindrome(partition):
			partitions = append(p, string(partition))
			fn(res, partitions, i)
			partitions = partitions[:len(p)-1]
  return res

func isPalindrome(c []byte) bool:
*/

func partition(s string) [][]string {
	var res = make([][]string, 0)
	var ps = make([]string, 0)
	var c = []byte(s)
	fnPart(&res, ps, c, 0)
	return res
}

func fnPart(res *[][]string, ps []string, c []byte, index int) {
	if index == len(c) {
		*res = append(*res, append([]string{}, ps...))
		return
	}

	// recursion case
	for i := index; i < len(c); i++ {
		partition := c[index : i+1]
		if isPalindrome(partition) {
			ps = append(ps, string(partition))
			fnPart(res, ps, c, i+1)
			ps = ps[:len(ps)-1]
		}
	}
	return
}

func isPalindrome(c []byte) bool {
	i, j := 0, len(c)-1
	for i <= j {
		if c[i] != c[j] {
			return false
		}
		i++
		j--
	}
	return true
}
