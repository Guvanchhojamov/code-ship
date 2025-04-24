package daily

import "fmt"

/*
You are given an array nums consisting of positive integers.

We call a subarray of an array complete if the following condition is satisfied:
The number of distinct elements in the subarray is equal
to the number of distinct elements in the whole array.
Return the number of complete subarrays.

A subarray is a contiguous non-empty part of an array.

Example 1:

Input: nums = [1,3,1,2,2]
Output: 4
Explanation: The complete subarrays are the following: [1,3,1,2], [1,3,1,2,2], [3,1,2] and [3,1,2,2].

Example 2:
Input: nums = [5,5,5,5]
Output: 10
Explanation: The array consists only of the integer 5, so any subarray is complete.
The number of subarrays that we can choose is 10.

Constraints:
1 <= nums.length <= 1000
1 <= nums[i] <= 2000
1 <= nums[i] <= 2000
*/
/*
  1 simple brute force solution is generate all subarrays and check is it satisfied or not.
	How we can generate subarrays?
	subarray is cont elements so we can generate them in 2 loops.
 for i=0 to n:
	for j=0 to i:
		if targetUniqLen == countUniq(curr):
			res++
return res
tc: n*n*k
sc: n
is it ok?
	How can we optimize? or can we count them without generate all subarrays?
let try wth brute force since the range is small.
*/

// our solution.

func countCompleteSubarrays(nums []int) int {
	var res = 0
	target := countUniq(nums)

	if target == 1 {
		return len(nums) * 2 // all ements are same in nums, we can take all subarrays, the max subarray is n*2
	}
	k := 0
	for k < len(nums) {
		curr := []int{}
		for i := k; i < len(nums); i++ {
			curr = []int{}
			for j := k; j <= i; j++ {
				curr = append(curr, nums[j])
			}
			fmt.Println(target, curr, countUniq(curr))
			if target == countUniq1(curr) {
				res++
			}
		}
		k++
	}

	return res
}

func countUniq1(nums []int) int {
	var freqMap = make(map[int]int, 0)
	var count = 0

	for _, v := range nums {
		if _, ex := freqMap[v]; ex {
			continue
		}
		freqMap[v]++
		count++
	}
	return count
}

// showed solution.
func countCompleteSubarrays(nums []int) int {
	totalUnique := countUniq(nums)
	res := 0
	freq := map[int]int{}
	left := 0

	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++

		for len(freq) == totalUnique {
			// All subarrays from left to the end with this right are valid
			res += len(nums) - right
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				delete(freq, nums[left])
			}
			left++
		}
	}

	return res
}

func countUniq(nums []int) int {
	seen := map[int]bool{}
	for _, v := range nums {
		seen[v] = true
	}
	return len(seen)
}
