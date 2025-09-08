"""
1317. Convert Integer to the Sum of Two No-Zero Integers
No-Zero integer is a positive integer that does not contain any 0 in its decimal representation.
Given an integer n, return a list of two integers [a, b] where:
 - a and b are No-Zero integers.
 - a + b = n
 The test cases are generated so that there is at least one valid solution.
 If there are many valid solutions, you can return any of them.

Example 1:

Input: n = 2
Output: [1,1]
Explanation: Let a = 1 and b = 1.
Both a and b are no-zero integers, and a + b = 2 = n.

Example 2:
Input: n = 11
Output: [2,9]
Explanation: Let a = 2 and b = 9.
Both a and b are no-zero integers, and a + b = 11 = n.
Note that there are other valid answers as [8, 3] that can be accepted.

Constraints:
2 <= n <= 10^4  
"""
"""
  given n num. 
  we need got 2 number such that. 
    - no zero in decimal representation. 
    - non zero numbers. 
    - at least 1 result are valid. 
    - can return any valid answer. 
for example: 
    n = 11. 
    we cannot use 10+1. because 10 has 0 in decimal. 
    we can use 9+2 or 8+3 both are valid. 

1. ok assume we can got all nums with 0 in decimal, how we can seperate then ? 
    start i = 1. 
    while i < n: decrement i from n. 
    if first and second numers are correct then return as result. 
    how we can check correct or not?
    if number is diveded to 10 or 5 then number has 0 value. 
    since n is big we can check dividing by 10. 
    if divide first, or left part by 10==0 return false and skip this number. 
    otherwise this number is valid. 
1010

[1009,1] 
1009 = 
"""

class Solution:
    def getNoZeroIntegers(self, n: int) -> List[int]:
        for num in range(1, n):
            first = n - num 
            if self.check_number(first) and self.check_number(num):
                return [first,num]
            
    def check_number(self, n: int) -> bool:
        while n>0:
            print(n)
            if n % 10 == 0: return False
            n = n // 10
        return True

"""


Since the given range of n is [2,10000],
which is relatively small, we can directly enumerate possible values of A in the range [1,n).
For each A, we compute B=n−A and check whether both A and B contain no zero in their decimal representation.

Implementation

Complexity Analysis
Time complexity: O(nlogn).

We enumerate all possible values of A in O(n). 
For each pair (A,B), we check whether their decimal representations contain a zero.
This check takes O(logn) time, since the number of digits in A and B is O(logn).
Space complexity: O(1).
"""
class Solution:
    def getNoZeroIntegers(self, n: int) -> List[int]:
        for A in range(1, n):
            B = n - A
            if "0" not in str(A) + str(B):
                return [A, B]
        return []