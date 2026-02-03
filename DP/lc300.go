package dp

/*
Given an integer array nums, return the length of the longest strictly increasing subsequence.
*/
/*
 Input:
	nums []int
 Output:
	Len int

Conditions:
 Range for nums? Min:0; max:10^5.
 Range for nums[i]? min: -10^5; max: 10^5.
 What is min possible answer? [1] - one number itself.
To be a subsequence, we need to keep order and take continuous or non-continuous, sequence of nums.
There can be multiple answers? We just return size, its not the problem.
Examples:
 Nums = [4,1,5,2,6,3];
 1,2,3;
 4,5,6;  result is 3.
*/
/*
What we can do?
[4,1,5,2,6,3]
 0 1 2 3 4 5
 i
 p
maxLen:
Approach:
Ok, since we need to chek some value again-agaon,can we use recursion?
Maybe we can? How we can use recursion?
For recursion, we need base case and recursion case.
What arguments do we need for this?
 fn(i int, prev int)
 //base case
  If i == n {return 0}
 //recursion case
 if nums[i] > prev:
	return 1 + fn(i+1, nums[i])

 return fn(i+1, prev)

Ok, can we do better with taboulation maybe?
Yes, we can do this with tabulation but again it will take n^2 time.
For tabulation we need base case and loop case. What we are doing here?
[4,1,5,2,6,3]
 0 1 2 3 4 5
We declare another dp array to store previous length and take from this arr.
dp:=make([][]int,n+1)
Dp = [i][previ] nxn
And check each time,
for each index initially len is 1 so we have at least 1 length ssq for each one.

[4,1,5,2,6,3]
 0 1 2 3 4 5
[1,1,2,1,3,2]

max:3
Count:3
*/

func lengthOfLISdp(nums []int) int {
	if len(nums) < 0 {
		return 0
	}
	n := len(nums)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 1
	}
	maxLen := 1
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])
				maxLen = max(maxLen, dp[i])
			}
		}
	}
	// fmt.Println(dp)
	return maxLen
}

/*

Ok, how we can solve with binary search approach?

[4,1,5,2,6,3]
 0 1 2 3 4 5
[1,1,1,1,1,1]

Ok, to solve this we need to create array for each starting point.
What is starting point. First its 0. After this each min point is starting one.
Example:
[4,1,5,-2,6,3]

[4,5,6]
[1,3]
[-2]
 In this case, the max len is 3.
How do we know where we need to paste? And which array? The longest array is first priority so, we need set of arrays for this. But we can simplify this solution.
Using only 1 array and replacing elements in positions. Since we need only lengths at the end we got correct length from this arr.
[1,3,6] = 3
[-2,3,6] = 3 . . .
 In any case, replacing nums we got same length max array.
Ok, Now how we can find where we need to paste element?
Since elements are in increasing order, we can use some efficient algo to search position to replace or add back.
[4,5] - assume we have this and next is -2.
4,5
0 1 to find position we need upper bound +1 after the num < -2 in this case upper bound is -1 idx. So we add right after this. -1 + 1 = 0. We replace with 0th num. Or assume we need add 6.
4,5; 6
0 1   in this case we need to add 1 + 1; so i > len. So we add to the end.
So if up_b < len(arr) -> replace else append to arr.

So since, for each number we do up search.
TC: n * logN
SC: N - the additional array.
*/

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	arr := []int{} // we can do len =n but we dont need this
	arr = append(arr, nums[0])
	for i := 1; i < n; i++ {
		idx := lowerBound(arr, nums[i])
		if idx == len(arr) {
			arr = append(arr, nums[i])
		} else {
			arr[idx] = nums[i]
		}
	}
	return len(arr)
}

/*
[4,5]; t=6

	0 1

l:2
r:1
m:1
*/
func lowerBound(arr []int, target int) int {
	n := len(arr)
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
