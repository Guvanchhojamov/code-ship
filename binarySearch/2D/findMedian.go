package _D

import "math"

/*
Given a row-wise sorted matrix where the number of rows and columns is always odd, find the median of the matrix.

Examples:

Input: mat = [[1, 3, 5], [2, 6, 9], [3, 6, 9]]
Output: 5
Explanation: Sorting matrix elements gives us {1,2,3,3,5,6,6,9,9}. Hence, 5 is median.
Input: mat = [[1], [2], [3]]
Output: 2
Explanation: Sorting matrix elements gives us {1,2,3}. Hence, 2 is median
Input: mat = [[3], [5], [8]]
Output: 5
Explanation: Sorting matrix elements gives us {3,5,8}. Hence, 5 is median.


Constraints:
1 <= mat.size(), mat[0].size() <= 400
1 <= mat[i][j] <= 2000

*/

/*
  Key points:
    - rows are sorted
    - rows and cols always odd then n*m is odd too always.
    - need to find median, like:
        if index x = median
        x = m*n/2
        left x right
        left == right
     need to find median in sorted manner
  [1, 3, 5],
  [2, 6, 9],
  [3, 6, 9]     res = 5
   What we can do for simple brute foce?:
    - flatten 2D array to 1D array.
    - Sort 1d array
    - return arr[len(arr)-1/2]
    tc: n*m +nlogn
    sc: n*m

  Can we optimize? Maybe we can.
    Can we use binary search in there?
    1. Define range
    2. Find elimiation point.
  We need iterate from min to max
  The min is min element of the first column, because on each row the 1st elements is min we need to find min of minimums.
    low = min(m[0][i])
  What is the max?
  The max is max of last column elmeents, since all max elemnts in each row in last position:
    high = max(m[i][len(m[0])-1])
  Oru answer lies between low and high,so range is [low,high]
  What is the elimination point?
  1 3 3 2 6 9 3 6 9
  1 2 3 3 3 6 6 9 9 - sorted.

  1 2 3 4 5 6 7 8 9
  if after sorting  elementsCountOnTheLeft == m*n/2 then this is our ans.
   for example:
    low=1
    how many elements on the left of 1 after sorting?
    0!= 9/2=4
     5/3 =1
     5%3 =2

    1. From -inf to inf
    2. From low to high, min to max check for eaxh number countsEqual or small nums of given.
    3. count equal or small elemensts for each row using upper bound, and sum all rows counts.  Upper bound is first greater element after target.
       for exampe:
        1 3 4 5 6 7 8 9   target=5 first greater index = 4 it means 4 nums are greater or equal to target.
    4. if smallsCunt == n*m/2 then return i
       if < then got to bigger low=mid+1
       else got to smaller high = mid-1


*/

func medianMatrix(mat [][]int) int {
	var low, high = 1, math.MaxInt
	var n = len(mat)
	var m = len(mat[0])
	// find correct range.
	for i := range mat {
		low = min(low, mat[i][0])
		high = max(high, mat[i][m-1])
	}

	reqMedian := (n * m) / 2
	for low <= high {
		mid := (low + high) / 2
		smallOrEqualsCount := smallOrEqCount(mat, mid)
		if smallOrEqualsCount == reqMedian {
			return mid
		}
		if smallOrEqualsCount < reqMedian {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func smallOrEqCount(mat [][]int, target int) int {
	var count = 0
	for _, col := range mat {
		count += upperBoundOfRow(col, target)
	}
	return count
}

func upperBoundOfRow(nums []int, target int) int {
	// upperBound == smallOrEqualsCount
	var low, high = 0, len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
