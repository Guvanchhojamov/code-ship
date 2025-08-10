package daily

import (
	"sort"
	"strconv"
)

/*

869. Reordered Power of 2
Solved
Medium
Topics
premium lock icon
Companies
You are given an integer n. We reorder the digits in any order (including the original order) such that the leading digit is not zero.

Return true if and only if we can do this so that the resulting number is a power of two.



Example 1:

Input: n = 1
Output: true
Example 2:

Input: n = 10
Output: false

*/

/*
  we know how to check power of two, n&n-1==0   if 1 bit set then the number is power of two.
- the leading digit must not be zero. we can move zero any place, but not 0-s position.
  brute force:
   - generate all combinations.
   - check bits.
	return true or false.
	tc: n!.
	sc: 1.  got time limit.
Can we optimize:
	- first check this number itself.
	- what max val we can get changin order in num?  sort in desc.
	- what min value we can get changing order in num? sort in asc. num chars.
	- from min to max check each number for n & n-1 ==0.
	- if any num in min->max ok, then return true. we can got this num replacing places in num.
	- if no one ok for this condition then any change chars of num, not got power of 2. return false.
 how we can not take leading zeros optimally?
	ex:
	 4720014 - >
	 max: 744210 -> this is the max possible val on change.
	 min: 00012447 -> for min check, go to first non zero num, an replace with 0-th index with this num, and this num with 0.

	 while [min->max]:
	 	if n & n-1 == 0:
			return true.

	 return false.
	 1 = true.
	 10 = 01 => 10 => false.

*/

func reorderedPowerOf2(n int) bool {
	target := sortDigits(n)

	for i := 1; i <= 1e9; i <<= 1 {
		if sortDigits(i) == target {
			return true
		}
	}
	return false
}

func sortDigits(num int) string {
	str := strconv.Itoa(num)
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

/*
 Try to solve with count digits frquency.
*/
