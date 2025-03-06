package daily

/*
2965. Find Missing and Repeated Values

You are given a 0-indexed 2D integer matrix grid of size n * n with values in the range [1, n^2].
Each integer appears exactly once except A which appears twice and B which is missing.

The task is to find the repeating and missing numbers a and b.

Return a 0-indexed integer array ans of size 2 where ans[0] equals to A and ans[1] equals to B.



Example 1:
Input: grid = [[1,3],[2,2]]
Output: [2,4]
Explanation: Number 2 is repeated and number 4 is missing so the answer is [2,4].

Example 2:
Input: grid = [[9,1,7],[8,9,2],[3,4,6]]
Output: [9,5]
Explanation: Number 9 is repeated and number 5 is missing so the answer is [9,5].


Constraints:
2 <= n == grid.length == grid[i].length <= 50
1 <= grid[i][j] <= n * n
For all x that 1 <= x <= n * n there is exactly one x that is not equal to any of the grid members.
For all x that 1 <= x <= n * n there is exactly one x that is equal to exactly two of the grid members.
For all x that 1 <= x <= n * n except two of them there is exatly one pair of i, j that 0 <= i, j <= n - 1 and grid[i][j] == x.

*/
/*
   9 1 7
   8 9 2
   3 4 6
*/

func findMissingAndRepeatedValues(grid [][]int) []int {
	var d, m = -1, -1
	var n = len(grid)
	var freqMap = make(map[int]int, n*n)

	for _, row := range grid {
		for _, num := range row {
			freqMap[num]++
		}
	}

	for i := 0; i <= n*n; i++ {
		if freqMap[i] == 0 {
			m = i
		}
		if freqMap[i] > 1 {
			d = i
		}
	}
	return []int{d, m}
}

/*
1. Brute force:
	1. Create map with n*n size and fill with zer=0;  key=number; val=frequency.
	2. Iterate over 2D array and count number.
	3. If count > 1 add to response. This is the duplicate value.
	4. Iterate 2D once and if val = 0 add to response. This is missing number.
 tc: n*n+n*n = n*n
 sc: n*n
2. Brute force.
   1. Convert ot 1D array.
   2. Sort the array.
   3. Iterate in sorted array.
   4. Find missing number:
   5 Find a duplicate number.
3. Can we find missing and duplicate with sum of nums?
 	1. Sum nums 1 to n*n
    2. Sum nums of 2D Arr.

   1.. n*n  = 45
   arr  sum = 49
    arr sum / n*n = 5 missing num
    (arrSum / n*n )+(arrSum % n*n) = 5.4.. = 5+4= 9 = duplicate num.



*/
