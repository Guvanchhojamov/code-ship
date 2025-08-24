package heap

import (
	"slices"
	"strings"
)

/*
Note:
This problem was in live conding

	on my Tech Interview with Delivery Hero Backend Engineer II.

I cant solve sorting part due time, but solve with Maxiumum Heap.

692. Top K Frequent Words

Given an array of strings words and an integer k, return the k most frequent strings.
Return the answer sorted by the frequency from highest to lowest.
Sort the words with the same frequency by their lexicographical order.

Example 1:

Input: words = ["i","love","leetcode","i","love","coding"], k = 2
Output: ["i","love"]
Explanation: "i" and "love" are the two most frequent words.
Note that "i" comes before "love" due to a lower alphabetical order.
Example 2:

Input: words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
Output: ["the","is","sunny","day"]
Explanation: "the", "is", "sunny" and "day" are the four most frequent words, with the number of occurrence being 4, 3, 2 and 1 respectively.

Constraints:

1 <= words.length <= 500
1 <= words[i].length <= 10
words[i] consists of lowercase English letters.
k is in the range [1, The number of unique words[i]]

Follow-up: Could you solve it in O(n log(k)) time and O(n) extra space?
*/
type wordFreq struct {
	Word string
	Freq int
}

func topKFrequentWords(words []string, k int) []string {
	var mp = make(map[string]int, len(words))
	var w = []wordFreq{}
	for _, v := range words {
		mp[v]++
	}

	for key, v := range mp {
		w = append(w, wordFreq{Word: key, Freq: v})
	}

	slices.SortStableFunc(w, func(a, b wordFreq) int {
		if a.Freq != b.Freq {
			return b.Freq - a.Freq
		}
		return strings.Compare(a.Word, b.Word)
	})
	// fmt.Println(w)
	var res = []string{}
	for i := 0; i < k; i++ {
		res = append(res, w[i].Word)
	}
	return res
}

/*
 1. split words.
 2. add to map.
 3. extract map to struct.
 4. sort struct by woords freq.
	- if freq equal return lex greater element.
 5.
*/
