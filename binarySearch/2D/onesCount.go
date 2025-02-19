package _D

import "sort"

/*
2643. Row With Maximum Ones

Given a m x n binary matrix mat,
find the 0-indexed position of the row that contains the maximum count of ones,
and the number of ones in that row.
In case there are multiple rows that have the maximum count of ones,
the row with the smallest row number should be selected.
Return an array containing the index of the row, and the number of ones in it.



Example 1:

Input: mat = [[0,1],[1,0]]
Output: [0,1]
Explanation: Both rows have the same number of 1's. So we return the index of the smaller row, 0, and the maximum count of ones (1). So, the answer is [0,1].

Example 2:
Input: mat = [[0,0,0],[0,1,1]]
Output: [1,2]
Explanation: The row indexed 1 has the maximum count of ones (2). So we return its index, 1, and the count. So, the answer is [1,2].

Example 3:
Input: mat = [[0,0],[1,1],[0,0]]
Output: [1,2]
Explanation: The row indexed 1 has the maximum count of ones (2). So the answer is [1,2].

*/
/*
 Ok simple brute force solution is,
 Create result array with 2 items.
    Iterate in for loop each num in matrix
    count 1 ns count and compare with max
    if > than mas result=[index,count]
    return result.
    Tc: n*m
    sc: O(1)
 2. With binary search:
    do BS for each row m.
    Implement lowerBound(1)
    count = len(m) - lowBound(1)
    compare with max and do same thing.
    TC: n*log(m)
    sc: O(1)
 Let's implement binary search version:
 Note: rows already in sorted order.
*/

func rowAndMaximumOnes(mat [][]int) []int {
	var result = make([]int, 2)
	var maxCount = -1

	for i, row := range mat {
		sort.Ints(row)
		onesCount := lowerBoundAndOnesCount(row)
		if onesCount > maxCount {
			result[0], result[1] = i, onesCount
			maxCount = onesCount
		}
	}

	return result

}

func lowerBoundAndOnesCount(nums []int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) / 2)
		if nums[mid] >= 1 { // we can use just == here
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return len(nums) - low
}
