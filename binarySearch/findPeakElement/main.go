package main

// lc 162. Find Peak Element
func main() {

}

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) / 2)
		if mid > 0 && nums[mid] < nums[mid-1] {
			high = mid - 1
		} else if mid < len(nums)-1 && nums[mid] < nums[mid+1] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1 // for valid arrays this line never compiles.
}

/*
   [1,2,1,3,5,6,4]
    0 1 2 3 4 5 6
 l = 0
 h = 6
 m = 3


*/
/*
Example 2:

Input: nums = [1,2,1,3,5,6,4]

Output: 5
Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.

Constraints:

1 <= nums.length <= 1000
-231 <= nums[i] <= 231 - 1
nums[i] != nums[i + 1] for all valid i.
*/
