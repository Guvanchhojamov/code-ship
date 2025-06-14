package daily

import (
	"fmt"
	"strconv"
)

/*
2566. Maximum Difference by Remapping a Digit
You are given an integer num.
You know that Bob will sneakily remap one of the 10 possible digits (0 to 9) to another digit.
Return the difference between the maximum and minimum values Bob can make by remapping exactly one digit in num.
Notes:
- When Bob remaps a digit d1 to another digit d2, Bob replaces all occurrences of d1 in num with d2.
- Bob can remap a digit to itself, in which case num does not change.
- Bob can remap different digits for obtaining minimum and maximum values respectively.
- The resulting number after remapping can contain leading zeroes.

Example 1:

Input: num = 11891
Output: 99009
Explanation:
To achieve the maximum value, Bob can remap the digit 1 to the digit 9 to yield 99899.
To achieve the minimum value, Bob can remap the digit 1 to the digit 0, yielding 890.
The difference between these two numbers is 99009.

Example 2:
Input: num = 90
Output: 99
Explanation:
The maximum value that can be returned by the function is 99 (if 0 is replaced by 9) and the minimum value that can be returned by the function is 0 (if 9 is replaced by 0).
Thus, we return 99.

Constraints:
1 <= num <= 10^8
*/

/*
 What we need:
- change i val and equal to i, and make num max
- change i and equals, and make num min.
return diff(max,min)
How make num max?
How make num min?
what is max num we can replace? 9
what is the min num we can replace? 0 (need remove leading zeros)

We can use greedy approach there.
To make num max we need make front nums max.
- To do this we can go until firt oqurence where a[i]<9 then A and replace A=9 and all next nums[i]==A == 9.
After iteration we got max possible num.
- To do min nim we do same logic but for 0
go until nums[i]>0 then a[i]=B
return maxNum - minNum
tc: n+n = 2n
sc: n+n - for copy of num; 2N
*/

func minMaxDifference(num int) int {
	var s = strconv.Itoa(num)
	var maxNum = getMaxNum([]byte(s))
	var minNum = getMinNum([]byte(s))

	fmt.Println(num, string(maxNum), string(minNum))
	a, _ := strconv.Atoi(maxNum)
	b, _ := strconv.Atoi(minNum)
	return a - b
}

func getMaxNum(num []byte) string {
	i := 0
	for i < len(num)-1 && num[i] >= '9' {
		i++
	}
	el := num[i]
	for i < len(num) {
		if num[i] == el {
			num[i] = '9'
		}
		i++
	}
	return string(num)
}

func getMinNum(num []byte) string {
	i := 0
	for i < len(num)-1 && num[i] <= '0' {
		i++
	}
	el := num[i]
	for i < len(num) {
		if num[i] == el {
			num[i] = '0'
		}
		i++
	}
	i = 0
	for i < len(num)-1 && num[i] == '0' {
		i++
	}
	return string(num[i:])
}
