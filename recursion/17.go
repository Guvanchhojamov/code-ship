package recursion

import "fmt"

/*
17. Letter Combinations of a Phone Number
Given a string containing digits from 2-9 inclusive,
return all possible letter combinations that the number could represent.
Return the answer in any order.

A mapping of digits to letters
(just like on the telephone buttons) is given below.
Note that 1 does not map to any letters.

Example 1:

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:
Input: digits = ""
Output: []
Example 3:
Input: digits = "2"
Output: ["a","b","c"]

Constraints:

0 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9'].
*/
/*
   Notes:
	Generate all possible combinations in phone nums chars,
	for given (pressed) nums.
- digits [2..9]
- each digit holds 3-4 char.
- we need to map digits and chars:
2 = [a,b,c]
3 = [d,e,f]
...
9 = [w,x,y,z]
   in this problem we ned to use recursion and pick and not pick algorithm
  like similar to len(s) loop.
we need to keep track.
*/
var charsMap = map[byte][]rune{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	var res = make([]string, 0)
	var curr = make([]rune, 0)
	if len(digits) == 0 {
		return []string{}
	}
	fnLetters(&res, curr, []byte(digits), 0)
	fmt.Println(res)
	return res
}

func fnLetters(res *[]string, curr []rune, c []byte, dIdx int) {
	// base case:
	// eger comb == c uzynylklar den bolsa 1 jogap generate edildi diyip result-a goshyas.
	if len(curr) == len(c) {
		*res = append(*res, string(curr))
		return
	}

	// recursion case:
	// her i-nji digitin harpy uchin indiki digitin harplaryny chagyryas.
	chars := charsMap[c[dIdx]]
	for _, ch := range chars {
		curr = append(curr, ch)
		//fmt.Printf("path=%s, index=%d\n", string(curr), dIdx)
		fnLetters(res, curr, c, dIdx+1) // expose
		curr = curr[:len(curr)-1]       // remove chosen / backtrack
		//fmt.Printf("path=%s, index=%d\n", string(curr), dIdx)
	}
	return
}

/*
  1. create a map with chars.
  - iterate over s, take s[i] chars.
  iterate s[i].chars, for each char,
*/
