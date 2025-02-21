package daily

func main() {

}

/*

You are given two 0-indexed integer permutations A and B of length n.

A prefix common array of A and B is an array C such that
C[i] is equal to the count of numbers that are present at or before the
index i in both A and B.

Return the prefix common array of A and B.

A sequence of n integers is called a permutation if it contains all integers from 1 to n exactly once.

Example 2:
Input: A = [2,3,1], B = [3,1,2]
Output: [0,1,3]
Explanation: At i = 0: no number is common, so C[0] = 0.
At i = 1: only 3 is common in A and B, so C[1] = 1.
At i = 2: 1, 2, and 3 are common in A and B, so C[2] = 3.
*/

/*
2:0
3:1
1:2

3:0
1:1
2:2

for i->B:
  for bMap[b[i]] >= i{
      count++
   }
*/
