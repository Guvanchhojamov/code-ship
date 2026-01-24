package dp

/*
You are given an integer array nums and an integer target.
You want to build an expression out of nums by adding one of the symbols '+' and '-' before each integer in nums and then concatenate all the integers.

For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and concatenate them to build the expression "+2-1".
Return the number of different expressions that you can build, which evaluates to target.

*/
/*
  input: nums []int, target int
  out: ways count int
Nums can be negative? No
Target can be negative? Yes can be.
*/
/*
nums = [1,1,1,1] target = 2
 1+1-1+1 = 2; 1
 1-1+1+1 = 2; 1
 1+1+1-1 = 2; 1
 -1+1+1+1 = 2; 1
We have only 4 ways to do this.
*/
/*
 So, How can we do this?
We can build all combinations and check when this is equal to target or not.
How we can generate? And how we can use negative, positive case?
 Ok, if we see, for each number we have 2 cases.
-negative, positive.
Can we use there 0/1 knapsack top-down approach? Maybe we can, if we define base and recursion cases.
When we use top-down our base case is 0.
So, what we need to do when we come to 0 th element.
Use helper function, fn(nums, idx, currRes, target)
[1,1,1,1]
 0 1 2 3
//
 if currRes == target: return 1; - we found some way.
 if i<0:
Return 0; - its not valid path.

// in other all other cases we try 2 states for + and -.
// recursion case
 negative:=fn(nums, i-1, currRes - nums[i], target)
 positive:=fn(nums, i-1, currRes + nums[i], target)
 return negative+positive
 TC: 2^n
 SC: 2^n time we call the recursion.
*/
/*
 We can optimize this using memoization, in that case.
TC: O(N)
Sc: O(N) + O(N)
But, we cannot use dp[n][target] because target can be negative.
Ok lets, implement recursive solution.
*/

func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	return fn(nums, n-1, 0, target)
}

func fn(nums []int, i int, currRes int, target int) int {
	//base case
	if i < 0 {
		return 0
	}
	if currRes == target {
		return 1 // we found 1 way.
	}

	// recursion case
	negative := fn(nums, i-1, currRes-nums[i], target)
	positive := fn(nums, i-1, currRes+nums[i], target)
	return negative + positive
}
