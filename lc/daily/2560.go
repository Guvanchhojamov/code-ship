package daily

/*
2560. House Robber IV

There are several consecutive houses along a street,
each of which has some money inside.
There is also a robber, who wants to steal money from the homes, but he refuses to steal from adjacent homes.

The capability of the robber is the maximum amount of money he steals from one house of all the houses he robbed.

You are given an integer array nums representing how much money is stashed in each house.
More formally, the i-th house from the left has nums[i] dollars.

You are also given an integer k,
representing the minimum number of houses the robber will steal from.
It is always possible to steal at least k houses.

Return the minimum capability of the robber out of all the possible ways to steal at least k houses.



Example 1:
Input: nums = [2,3,5,9], k = 2
Output: 5
Explanation:
There are three ways to rob at least 2 houses:
- Rob the houses at indices 0 and 2. Capability is max(nums[0], nums[2]) = 5.
- Rob the houses at indices 0 and 3. Capability is max(nums[0], nums[3]) = 9.
- Rob the houses at indices 1 and 3. Capability is max(nums[1], nums[3]) = 9.
Therefore, we return min(5, 9, 9) = 5.

Example 2:
Input: nums = [2,7,9,3,1], k = 2
Output: 2
Explanation: There are 7 ways to rob the houses.
The way which leads to minimum capability is to rob the house at index 0 and 4. Return max(nums[0], nums[4]) = 2.


Constraints:

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^9
1 <= k <= (nums.length + 1)/2
*/
/*
   need steal from non-adjacent homes.
   need steal at least min K homes.
   for each steal combination, the maximum is rubber capability.
   need to return minimum capability from all combinations.

    What is the range of robber capability?
		min =  min(nums) when k=1.
		max = max(nums) when k = len(nums)/2
   1. Find min and max from nums:
   2. Iterate from min possible ans to max.
   3. Helper function getStolenMaxCapability() witch returns max() of stolen K>= homes.
   4.

*/

func minCapability(nums []int, k int) int {
	var mn, mx = 1, nums[0]

	for _, val := range nums {
		mx = max(mx, val)
	}

	var low, high = mn, mx

	for low <= high {
		mid := low + ((high - low) / 2)
		if isPossibleStealFromKHomes(nums, k, mid) {
			high = mid - 1 // we need to find minimum possibility so go to the lower when its possible.
		} else {
			low = mid + 1
		}
	}
	return low
}

func isPossibleStealFromKHomes(nums []int, k, capability int) bool {
	var cnt = 0
	var j = 0

	for j < len(nums) {
		if nums[j] <= capability {
			j += 2
			cnt++
		} else {
			j++
		}
		if cnt >= k {
			break
		}
	}
	return cnt >= k
}
