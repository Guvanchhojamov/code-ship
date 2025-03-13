package daily

/**
1780. Check if Number is a Sum of Powers of Three

Given an integer n, return true if it is possible to represent n as the sum of distinct powers of three. Otherwise, return false.

An integer y is a power of three if there exists an integer x such that y == 3x.



Example 1:

Input: n = 12
Output: true
Explanation: 12 = 3^1 + 3^2

Example 2:
Input: n = 91
Output: true
Explanation: 91 = 3^0 + 3^2 + 3^4

Example 3:
Input: n = 21
Output: false


Constraints:
1 <= n <= 10^7
*/

/*
	Sum of distinct powers of tree.
		- powers must be distinct.
		- the main is always 3.
        - Need to check for given N
   1. Brute Force.
	Think, to get sum of three at
	n or n-1 must divided to 3 withour remainders.
	why n-1 because 3^0 = 1 - > 3^0 +... = n
    Problem is that the powers must be distinct.

     We have 2 variants to check:
		m=n
      if m/3==0:
         sum=3*1+3*3
		for pow=1;pow<m/2;pow++{
			num=1
           for j:=range pow:
				num*=3
		sum+=num
		if sum == n {
          return true
		} else if sum> n {
			return false
		 }
		}

      if n/3 != 0:


*/
