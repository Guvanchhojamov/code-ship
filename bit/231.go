package bit

/*
  231. Power of Two
Given an integer n, return true if it is a power of two. Otherwise, return false.
An integer n is a power of two, if there exists an integer x such that n == 2^x.

Example 1:

Input: n = 1
Output: true
Explanation: 2^0 = 1
Example 2:
Input: n = 16
Output: true
Explanation: 24 = 16
Example 3:

Input: n = 3
Output: false


Constraints:

-2^31 <= n <= 2^31 - 1


Follow up: Could you solve it without loops/recursion?
*/
/*
 1. use loop until n==0, n/2==1 return false, otherwise true.
 2. convert to bin, count set if setsCount ==1 return true other false.
 3. using bitwise if x&x-1 == 1 then its power of two, if 0 then not.
 4. count set bits until n=0 shift to right and check if count ==1 then return true. else false.
*/

func isPowerOfTwo(n int) bool {
	if n > 0 && n&(n-1) == 0 {
		return true
	}
	return false
}

/*
 011
 010 -> 010
 1000 & 0111

 10000
  1111
*/
