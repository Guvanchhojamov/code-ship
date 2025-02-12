package daily

func main() {

}

/*
You are given two string arrays words1 and words2.

A string b is a subset of string a if every letter in b occurs in a including multiplicity.

For example, "wrr" is a subset of "warrior" but is not a subset of "world".
A string a from words1 is universal if for every string b in words2, b is a subset of a.

Return an array of all the universal strings in words1. You may return the answer in any order.

Example 1:
Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["e","o"]
Output: ["facebook","google","leetcode"]

Example 2:
Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["l","e"]
Output: ["apple","google","leetcode"]
*/

/*
Input: words1 = ["amazon","appleb","facebook","google","leeacebode"], words2 = ["ac","eb"]
  out: [facebook, leeacebode]
*/

/*
  1. First Solution coming to mind is brute force.
  	2 loops
      for i -> to len(words1):
         if isSubset(words1[i], words2):
            res = append(res, words1[i])
  return res.
*/

/*
   func isSubset(word1, words2):
       for i -> to words2:
          if !contains(word1, words2[i]):
              return false
    return true
*/
