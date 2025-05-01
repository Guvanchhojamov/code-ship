package recursion

/*
A digit string is good if the digits
(0-indexed) at even indices are even and
the digits at odd indices are prime (2, 3, 5, or 7).

For example, "2582" is good because the digits (2 and 8) at even positions are even
and the digits (5 and 2) at odd positions are prime.
However, "3245" is not good because 3 is at an even index but is not even.
Given an integer n,
return the total number of good digit strings of length n.
Since the answer may be large, return it modulo 109 + 7.

A digit string is a string consisting of digits 0 through 9 that may contain leading zeros.
Example 1:
Input: n = 1
Output: 5
Explanation: The good numbers of length 1 are "0", "2", "4", "6", "8".
Example 2:
Input: n = 4
Output: 400
Example 3:

Input: n = 50
Output: 564908303

Constraints:
1 <= n <= 10^15
*/
/*
   simple brute force approach is that,
   take all nums from 1+n*(0) to 9+(n*0).
But since 1...,3...,5..,7..,9.. are not good nums we can ignore them.
	take 2.., 4.., 6...
  Turns out there is math way how we can calculate the good numbers count.
even pos = 0,2,4,6,8
odd pos = 2,3,5,7
  For even position there might be just 5 nums.
for odd position there might be just 4 nums
 if n = even num:
  n/2 odd
  n/2 even

 Half of evens half of n are primes.
 if n== even so:
  evensCount = 5^(n/2)
  oddsCount = 4^(n/2)
if n == odd then we have one extra even number, with is last, so
   evensCount = 5^(n+1/2)
   oddsCount = 4^(n/2)  stay same.
so, calculate the pow of x^y  with modulo 10^9
and sum evens and odds count.
*/

func countGoodNumbers(n int64) int {
	const mod = 1_000_000_007
	evenCount := (n + 1) / 2
	oddCount := n / 2

	evens := powerMod(5, evenCount, mod)
	odds := powerMod(4, oddCount, mod)

	return int(evens * odds % mod)
}

func powerMod(base, exp, mod int64) int64 {
	result := int64(1)
	for exp > 0 {
		if exp%2 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		exp /= 2
	}
	return result
}
