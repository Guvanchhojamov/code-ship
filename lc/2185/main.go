package main

func main() {

}

/*

You are given an array of strings words and a string pref.

Return the number of strings in words that contain pref as a prefix.

A prefix of a string s is any leading contiguous substring of s.

Example 1:

Input: words = ["pay","attention","practice","attend"], pref = "at"
Output: 2
Explanation: The 2 strings that contain "at" as a prefix are: "attention" and "attend".
Example 2:

Input: words = ["leetcode","win","loops","success"], pref = "code"
Output: 0
Explanation: There are no strings that contain "code" as a prefix.


Constraints:

1 <= words.length <= 100
1 <= words[i].length, pref.length <= 100
words[i] and pref consist of lowercase English letters.

*/

func prefixCount(words []string, pref string) int {
	var count = 0
	var k, l = 0, 0

	for _, word := range words {
		k, l = 0, 0
		for l < len(pref) && k < len(word) && word[k] == pref[l] {
			k++
			l++
		}
		if l == len(pref) && k > 0 && len(pref) > 0 { // additional ceck length for corner case when pref word is empty.
			count++
		}
	}
	return count
}

/*
 we can solve this problem using two pointer approach.
 1. Run loop for words
 2. Check each word like:
    Take pointers i,j and word[i] and pref.
    i=0 -> of word[i] index
    j=0 -> of pref index.

    iterate two pointers like:
     for words[wrod][i]=pref[j] && j<len(pref) {
         i++
         j++
     }
     if j==len(oref){
         continue
     }
    and after ending loop
    return count of words as answer.

Tc: O(n*m) ~ n*n
SC:O(1)

let's try some test cases..

words = ["pay","attention","practice","attend"], pref = "at"
    i=1
    k=0
    l=0

    attention
    at
    k=2
    l=2
*/
