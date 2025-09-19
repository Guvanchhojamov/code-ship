package greedy

/*
678. Valid Parenthesis String

Given a string s containing only three types of characters: '(', ')' and '*', return true if s is valid.

The following rules define a valid string:
Any left parenthesis '(' must have a corresponding right parenthesis ')'.
Any right parenthesis ')' must have a corresponding left parenthesis '('.
Left parenthesis '(' must go before the corresponding right parenthesis ')'.
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string "".

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "(*)"
Output: true
Example 3:

Input: s = "(*))"
Output: true


Constraints:

1 <= s.length <= 100
s[i] is '(', ')' or '*'.

*/

/*
  Brute force solution here, try generate all possible ways with keep track
  pairs count.
  - if '(' increment count and continue.
  - if ')' decrement count and continue.
  - if '*' we have 3 ways:
	icemenet count ')'
	decrement count '('
	or do nothing ''
  how we can do it and check all combinations?
	When we come to some combinations we need recursion.
	We can do it with recursion.
 to use recursion we need to define 2 points.
 - break point.
 - call recussion point.
   break point is if:
	idx reachen len(s)
	if s starts with ')' --- ?
   conitue points:
	if '(' then idx+1, pairs+1.
	if ')' then idx+1, pairs-1
	if '*' then:
		'(' idx+1, pairs+1.
		')' idx+1, pairs-1
		''  idx+1 - do nothing.
	if at the end of each call if pairs != 0 or pairs < 0 then return false.
*/

func checkValidString(s string) bool {
	var c = []rune(s)
	var minopen, maxopen = 0, 0

	for _, v := range c {
		if v == '(' {
			minopen++
			maxopen++
		} else if v == ')' {
			minopen--
			maxopen--
		} else {
			minopen-- // best case we got ) - for *
			maxopen++ // worst case we got another ( - for *
		}
		if minopen < 0 {
			minopen = 0
		}
		if maxopen < 0 {
			return false
		}
	}
	return minopen == 0 // make sure its balanced after all
}

/*
   Since our brute force solution gives us TLE, we need use greedy with range method.
   we need keep track 2 vars.
   minOpened, maxOpened.
   We need to keep range of possibilities.
   and at the end minOpened should be 0.
   Why?
   minOpened - is minimum possible opened parh.
   maxOpened - is maximum possible opened paranth.
   if '(' we increase minOpened, maxOpened.
   if ')' we decrease minOpened, maxOpened - because we have found 1 pair ')' ing paranthases.
   if '*' - we need:
       decrease minOpen - for ')' because we assume if ')' and it decreases openCount so, minopen--
       increase maxOpen - for '(' because we assume if '(' this increses opencount so, maxopen++
    if minopen < 0 we need reset it to 0.
    Why?
    Because most optimistic case if where minopen ==0. it is impossible if minopen < 0.
    * ) in there minopen == -1 but in relistic range is [0,1] always bigger than 0.
*/

func checkValidString2(s string) bool {
	var c = []rune(s)

	if len(c) > 0 && c[0] == ')' {
		return false
	}
	return fn(c, 0, 0)

}

func fn(c []rune, idx int, pairs int) bool {
	// break
	if pairs < 0 {
		return false
	}
	if idx >= len(c) {
		return pairs == 0
	}

	// call point
	if c[idx] == '(' {
		return fn(c, idx+1, pairs+1)
	}
	if c[idx] == ')' {
		return fn(c, idx+1, pairs-1)
	}

	if c[idx] == '*' {
		return fn(c, idx+1, pairs+1) || fn(c, idx+1, pairs-1) || fn(c, idx+1, pairs)
	}

	return false
}
