package stack

/*
402. Remove K Digits
Given string num representing a non-negative integer num,
and an integer k, return the smallest possible integer after removing k digits from num.

Example 1:
Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.

Example 2:
Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.

Example 3:
Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.


Constraints:
1 <= k <= num.length <= 10^5
num consists of only digits.
num does not have any leading zeros except for the zero itself.
To got smallest number we need to go with greedy approach.
to get smaller num we need remove first k big numbers.
How can we know bigger nums and remove them?
We will use stack ds to store minimum numbers from front to end.

Approach:
- start from 0->n
- keep monotonic stack witch is keeps nums on increasing order, from beginning.
- if removed from stack elements == k then add remaining elements to stack too.
and return the number in stack. Also handle edge cases too.

Input: num = "1432219", k = 3
Output: "1219"
for this, add to stack.
i=0
s=1
i=1
s=1<4 = 14
i=2
s=4>3  st.pop; st.add(3)
*/

func removeKdigits(num string, k int) string {
	var st = make([]byte, 0, len(num))

	for i := range num {
		for len(st) > 0 && k > 0 && num[i] < st[len(st)-1] {
			st = st[:len(st)-1]
			k--
		}
		st = append(st, num[i])
	}
	for len(st) > 0 && k > 0 {
		st = st[:len(st)-1]
		k--
	}
	for len(st) > 0 && st[0]-'0' == 0 {
		st = st[1:]
	}
	if len(st) == 0 && k == 0 {
		return "0"
	}

	return string(st)
}
