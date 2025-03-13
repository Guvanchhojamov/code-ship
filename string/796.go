package string

// 796. Rotate String

/*
Given two strings s and goal, return true if and only if s can become goal after some number of shifts on s.

A shift on s consists of moving the leftmost character of s to the rightmost position.

For example, if s = "abcde", then it will be "bcdea" after one shift.


Example 1:
Input: s = "abcde", goal = "cdeab"
Output: true

Example 2:
Input: s = "abcde", goal = "abced"
Output: false


Constraints:

1 <= s.length, goal.length <= 100
s and goal consist of lowercase English letters.
*/

/*
  1. Given s and goal
     Is we can generate goal shifting s left x times.
  Brute force:
    Take s and shift elements one by one and check with goal:
        if equal return true
        else return false.
    tc: len(s)
       generation means:
        add s[i] to end of s:
        take s = s[i+1:]  O(1)
    sc: O(1)
  Can we optimize?
  Maybe.
     What is the shifting?
        Shifting to left means, shift all elements from i one position to left.
        If the char index 0 this becomes last char in this word. For example:
            axfre
            xfrea
            freax
            reaxf
            eaxfr
            axfre
        after full shift the s becomes same as started s.
         This means all shift is len(s);
         After each len(s) th shift we got same word with s.
     Take goal =  `reaxf`  is it rotated s yes.
     s = axfreaxfr
     g = reaxf
    1. if s == g retrun true.
    2. iterate s
       i=0
       for i<len(s){
           if s[i] == t[0]:
            s[i:] == goal:
             return true
        s = append(s[i])
       }
    return
*/

func rotateString(s string, goal string) bool {
	var c = []byte(s)
	var n = len(s)

	for i := 0; i < n; i++ {
		c = append(c, c[i])
		if string(c[i+1:]) == goal {
			return true
		}
	}

	return false
}
