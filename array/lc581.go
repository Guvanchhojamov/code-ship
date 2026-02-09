package array

import "math"

/*
Given an array, Return the minimum possible subarray which when sorted, the entire array will become sorted.


We given array, we need to find min possible subarray which when sorted, the entire array will be sorted.
 Input:
	Nums []int
 Output:
	Subarray []int ? or just length?  Ok we need subarray itself.  Ok is order of response matters? No it doesnt matter.

Constraints:
	Nums length? 0..10^5.
	nums[i]  -1000...1000 - can be negative? Assume not. Only positive nums.
	We need to find min subarray.
	By sorted we means the non-decreasing order? Right? Yes.
Ex:
 [1,4,3,6,7]
 res: [4,3] but not our result is 1 [4] or [3] lets assume we need return length of subarray
 The longest subarray is n.
 We need to find min.
What is brute force.
We try to sort all possible subarrays and check is array is sorted or not?
Keeping track min length of this subarray.
TC: N * NlogN
SC: O(N)
Ok how we can optimze?
 We can think from other vision, what min elements we can remove to become other elements are sorted? If we find this we find our anser.
 But, how we can find this? Longest sorted array.
We need 2 functionalities to make this happen,
Got correctly, and efficiently possible subarrays.
Can efficiently check is array sorted or not ?
[1,4,3,6,7]
   i
   i+1
Can we go greedy here? Maybe we can maybe not.
Check i and j=i+1 is sorted a[i] <= a[j]
If not sorted go until j++ a[i-1] > a[j]; keeping counter..
And compare length with min count
At the end return minLen
 [1,4,3,6,7]
        i j
 maxLen=1
 [6,1,4,3,5]
          ij
  min:1
 Res:1 we only need remove 6.

[2,6,4,8,10,9,15]
     L    R

*/

func findUnsortedSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)

	left := 0
	for left < n-1 && nums[left] <= nums[left+1] {
		left++
	}

	if left == n-1 {
		return 0
	}

	right := n - 1

	for right > 0 && nums[right] >= nums[right-1] {
		right--
	}

	minVal, maxVal := math.MaxInt, math.MinInt
	for i := left; i <= right; i++ {
		minVal = min(minVal, nums[i])
		maxVal = max(maxVal, nums[i])
	}
	for left > 0 && nums[left-1] > minVal {
		left--
	}
	for right < n-1 && nums[right+1] < maxVal {
		right++
	}
	return right - left + 1
}
