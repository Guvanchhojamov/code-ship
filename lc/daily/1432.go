package daily

import (
	"fmt"
	"strconv"
)

/*
1432. Max Difference You Can Get From Changing an Integer
You are given an integer num.
You will apply the following steps to num two separate times:

- Pick a digit x (0 <= x <= 9).
- Pick another digit y (0 <= y <= 9). Note y can be equal to x.
- Replace all the occurrences of x in the decimal representation of num by y.
- Let a and b be the two results from applying the operation to num independently.

Return the max difference between a and b.

Note that neither a nor b may have any leading zeros, and must not be 0.


Example 1:
Input: num = 555
Output: 888
Explanation: The first time pick x = 5 and y = 9 and store the new integer in a.
The second time pick x = 5 and y = 1 and store the new integer in b.
We have now a = 999 and b = 111 and max difference = 888

Example 2:
Input: num = 9
Output: 8
Explanation: The first time pick x = 9 and y = 9 and store the new integer in a.
The second time pick x = 9 and y = 1 and store the new integer in b.
We have now a = 9 and b = 1 and max difference = 8

Constraints:

1 <= num <= 10^8
*/
/*
 x == y == (0 <= x,y <= 9)

 - What  need pick any dig x from num, and replace it with y - and store a
 - Need pick any dig x from num again, and replace it with y - and store b
 To get max diff what we should do?
 1. Get max possible num
 2. Get min possible num
 3. return diff.
 1. To get max possible num, what is the max num we can replce? - 9, so y=9 for max always.
    Need to find with with dig we need replace 9? To get max we start from frond, and find first <9 dig, and replace
    it and all nex occurences with 9.
2. To get min possible num, what is the min num we can replae? - 1, so y=1, always, why not 0? because if we use 0 result myght be 0. Not acceptable.
    Do same for min, start from fron until first el >0. replace it, and all with 1. If no elements just b=1 set. b cant be 0.
4. return a-b
*/

func maxDiff(num int) int {
	var s = strconv.Itoa(num)
	var a = []byte(s)
	var b = []byte(s)
	a = makeNumMax(a)
	b = makeNumMin(b)
	fmt.Println(string(a), string(b))
	k, _ := strconv.Atoi(string(a))
	r, _ := strconv.Atoi(string(b))
	return k - r
}

func makeNumMax(a []byte) []byte {
	var x byte
	i := 0
	for i < len(a)-1 && a[i] == '9' {
		i++
	}
	x = a[i]
	for i < len(a) {
		if a[i] == x {
			a[i] = '9'
		}
		i++
	}
	return a
}

func makeNumMin(b []byte) []byte {
	var x byte
	if b[0] != '1' {
		x = b[0]
		// Replace all occurrences of x with '1'
		for j := 0; j < len(b); j++ {
			if b[j] == x {
				b[j] = '1'
			}
		}
	} else {
		for i := 1; i < len(b); i++ {
			if b[i] != '0' && b[i] != '1' {
				x = b[i]
				for j := 1; j < len(b); j++ {
					if b[j] == x {
						b[j] = '0'
					}
				}
				break
			}
		}
	}
	return b
}

/*
123456 -> 923456
123456 -> 103456
555 -> 999
555 -> 111 not 100
*/
