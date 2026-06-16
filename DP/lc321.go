package dp

import "math"

/*
You are given n balloons, indexed from 0 to n - 1.
Each balloon is painted with a number on it represented by an array nums.
You are asked to burst all the balloons.

If you burst the ith balloon, you will get
nums[i - 1] * nums[i] * nums[i + 1] coins.
If i - 1 or i + 1 goes out of bounds of the array,
then treat it as if there is a balloon with a 1 painted on it.

Return the maximum coins you can collect by bursting the balloons wisely.
*/
/*
 We are have 0..n-1 ballanods to brust.
We are given paints array, arr[i] - ballon paint.
Each time we brust ballon we got a[i-1] * a[i] * a[i+1] coins.
And we need to brust all ballnos in arr such that go get, max possible coins.
Input:
 Nums []int
Output:
	maxCoins int

Constraints:
 N range: 0...10^5
 nums range: [0]..[10^5]
 if i-1 <0 => 1
 If i+1 >len(arr) => 1;
Each ballon brust count a[i-1] * a[i] * a[i+1]

[3,1,5,8]
 0 1 2 3
 In this case to make max possible ballonds what we need ?
 If we do it one by one in this order?
 1*3*1=3 + 15+ 40 +40 = 98  coins.
  After bursting each ballon it removed from arr.
 What if we try other way? If we start with 1.
 3*1*5 +
 [3,5,8] =>

 [3,8]

 [8]
[1,3,1,5,8,1]
   l     r
     k
   1  .. 3

We do same this again again.
Each time how many options we have?
We need to try all possible options and need to return max one from them.
Since we need range [l, r]
 We try all possible ways between [l..r]
For this, we can try recursion.
How we try? We start from
l=0;r=n-1 and each time trya all possibilities between them.
So, for our function we need 2 params:
To make life easy we add 1; 1 to begin and end.
 [l, r]

solve(l, r int) int
 // base case
  if l <= 0 return {0}
  if r >= n return {0}

//recursion case
  maxVal:0
  k=l -> r;
   currCost=arr[l-1] * arr[k] * arr[r+1] +
   cost= currCost + solve(l, k-1) +
	solve(k+1,r)
   max(maxVal, cost)
*/

func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)
	n := len(nums)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	left, right := 1, len(nums)-2
	return solve(nums, memo, left, right)
}

func solve(nums []int, memo [][]int, left, right int) int {
	// base case
	if left > right {
		return 0
	}
	if memo[left][right] != -1 {
		return memo[left][right]
	}
	// recursion case
	maxCost := math.MinInt
	for k := left; k <= right; k++ {
		currCost := nums[left-1] * nums[k] * nums[right+1]
		cost := currCost + solve(nums, memo, left, k-1) + solve(nums, memo, k+1, right)
		maxCost = max(maxCost, cost)
		memo[left][right] = maxCost
	}

	return maxCost
}
