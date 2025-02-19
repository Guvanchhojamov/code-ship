package _D

import "fmt"

/*
74. Search a 2D Matrix
You are given an m x n integer matrix matrix with the following two properties:

Each row is sorted in non-decreasing order.
The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-10^4 <= matrix[i][j], target <= 10^4
*/

/*
  1. Brute force:
    Check each num in two loop.
    tc: O(n*m)
    sc: O(1)
  But we must do it in log time.
So, how can we do it?
    We can take each row and search with BS in it.
    it will
    tc: O(n*log(m))  but this is not accepted
    tc: O(1)
Let's take example:
  [1, 3, 5, 7]
  [10,11,16,20]
  [23,30,34,60]
  k =3

[1,3,5,7],
[10,11,16,20],
[23,30,34,60]
k=13
  Since our first number of row is always greater of the last num of previous,
  We be sure that
  if first number of row greater of the target we need to go less less.
  It means, Our first col is also sorted.
  So doing BS in first col we need to find with row we need to search.
  We need to find max[nums[0]] is less or equal to target
    max(nums[0]) <= target
  and this is the our needed row:
  and do BS in this row, if we dont find target return false.
  tc: lon(n)+log(m) = log(n+m)
  sc: O(1)

*/

func searchMatrix(matrix [][]int, target int) bool {
	// Do BS in first col to find row
	var low, high = 0, len(matrix) - 1
	var rowIdx = 0
	for low <= high {
		mid := low + ((high - low) / 2)
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] < target {
			rowIdx = mid
			low = mid + 1
		} else if matrix[mid][0] > target {
			high = mid - 1
		}
		fmt.Println(rowIdx)
	}

	// 1 10 23
	// Do BS in row
	low, high = 0, len(matrix[rowIdx])-1
	for low <= high {
		mid := low + ((high - low) / 2)
		if matrix[rowIdx][mid] > target {
			high = mid - 1
		} else if matrix[rowIdx][mid] < target {
			low = mid + 1
		} else {
			return true
		}
	}
	return false
}
