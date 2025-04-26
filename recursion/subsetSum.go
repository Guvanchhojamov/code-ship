package recursion

import "fmt"

/*
Given an array print all the sum of the subset generated from it, in the increasing order.
Examples:
Example 1:
Input: N = 3, arr[] = {5,2,1}
Output: 0,1,2,3,5,6,7,8
Explanation: We have to find all the subset’s sum and print them.
in this case the generated subsets are [ [], [1], [2], [2,1], [5], [5,1], [5,2]. [5,2,1],so the sums we get will be  0,1,2,3,5,6,7,8

Input: N=3,arr[]= {3,1,2}
Output: 0,1,2,3,3,4,5,6
Explanation: We have to find all the subset’s sum and print them.in
this case the generated subsets are [ [], [1], [2], [2,1], [3], [3,1], [3,2]. [3,2,1],so the sums we get will be  0,1,2,3,3,4,5,6
*/

/*
 Let's first repeat the subset sum generation problem, after this the sum is should nbe follow-up question.
  How to generate subset sum using recursion?
	to generate subset we already used the pick and not pick method, let try to implement this
*/

func subsetSum(nums []int) {
	var res = make([][]int, 0, 0)
	var arr = make([]int, 0, 0)
	r := fnn(res, nums, arr, 0)
	fmt.Println(r)

	var sumRes = make([]int, 0)
	var s = 0
	ss := fnSum(sumRes, nums, arr, 0, s)
	fmt.Println("sums: ", ss)
}

func fnSum(res, nums, arr []int, i, s int) []int {
	// base case
	if i == len(nums) {
		res = append(res, s)
		return res
	}

	// not pick
	res = fnSum(res, nums, arr, i+1, s)

	//pick
	arr = append(arr, nums[i])
	res = fnSum(res, nums, arr, i+1, s+nums[i])
	return res
}

func fnn(res [][]int, nums, arr []int, i int) [][]int {
	// base case
	if i == len(nums) {
		res = append(res, append([]int{}, arr...)) // avoid mutation
		return res
	}

	// not pick
	res = fnn(res, nums, arr, i+1)
	// pick
	arr = append(arr, nums[i])
	res = fnn(res, nums, arr, i+1)
	return res
}

/*
  base case:
    i == len(nums)

  not pick
    fn(res, nums, arr, i+1)
  pick
    arr = appen(arr, nums[i])
    fn(res, nums, arr, i+1)

  return res
*/
