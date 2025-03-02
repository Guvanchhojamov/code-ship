package string

/*
451. Sort Characters By Frequency
Given a string s,
sort it in decreasing order based on the frequency of the characters.
The frequency of a character is the number of times it appears in the string.
Return the sorted string. If there are multiple answers, return any of them.



Example 1:

Input: s = "tree"
Output: "eert"
Explanation: 'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
Example 2:

Input: s = "cccaaa"
Output: "aaaccc"
Explanation: Both 'c' and 'a' appear three times, so both "cccaaa" and "aaaccc" are valid answers.
Note that "cacaca" is incorrect, as the same characters must be together.

Example 3:
Input: s = "Aabb"
Output: "bbAa"
Explanation: "bbaA" is also a valid answer, but "Aabb" is incorrect.
Note that 'A' and 'a' are treated as two different characters.


Constraints:

1 <= s.length <= 5 * 10^5
s consists of uppercase and lowercase English letters and digits.
*/

/*
  Need to sort chars by frequency. From max frequency to min.
     What we can do ?
       Key points:
        - string has lower and upper cases.
        - answer could me more than one.
        - Return any correct answer
        - len(s) > 0  always.
   1. Brute force:
        a. Use any build in function if we have in used language.
        b. Create slice with same chars in same index, []string{'aa', 'bb' 'c'}
         - sort array
         - join sorted array and return.
        tc: n+nlogn+n = n+nlogn
        sc: n+n = N
   Since we need frequncy of chars.
   Can we use hashmap?
   key=char
   val = count
   'Aabb'
   b:2
   a:1
   A:1
  We need to use another hash map witch groups of same chars in one group:
    key = count
    val = chars
*/
/*
  tree
  t:1
  r:1
  e:2
Then how we can create string with sorted order of  chars.
We need to find range:
  min = 1
  what is the max.
 The max is when all chars are same,
 so max= len(s)


*/

func frequencySort(s string) string {
	var charsMap = make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		charsMap[s[i]]++
	}

	var freqMap = make(map[int][]byte, len(charsMap))
	for char, count := range charsMap {
		freqMap[count] = append(freqMap[count], char)
	}

	// from max to min
	// take frequency and generate string.
	var res = []byte{}
	for i := len(s); i >= 0; i-- {
		chars := freqMap[i]
		for _, c := range chars {
			count := charsMap[c]
			for i := 0; i < count; i++ {
				res = append(res, c)
			}
		}
	}
	return string(res)
}

/*
  tc: n+n+ (n+len(s[i])) n+n = n
  sc: n+n = N
*/
