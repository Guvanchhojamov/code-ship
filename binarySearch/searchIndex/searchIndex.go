package searchIndex

// 35. Search Insert Position
func main() {

}

func searchInsert(nums []int, target int) int {
	var low, high = 0, len(nums) - 1
	var res = len(nums) // if we can't find the target from array it wold be last+1==n

	for low <= high {
		mid := low + ((high - low) / 2)
		if nums[mid] >= target {
			res = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return res
}

/*
Example 2:
Input: nums = [1,3,5,6], target = 2
Output: 1

simple solution:
    for i->to nums:
        return nums[i] >= target

This will returns where
 if equal to target or
    reuturns first nums[i]> target.
    this will right becaues our array is sorted in ascending order.
     Tc: O(n) - > in worst case.
But we need to solve this in logN:
So, what we can do.
   somthing simialar to lower bound in binary search
     the smallest index where equal or greater than target, right?
      where smallest index where  nums[i]>=target.
       if nums[i] == target it will return target
    otherwhise it will return first nums[i]> target, so we can insert our target
    in there because the array is sorted.
    Lets implement lower BOund then.
     Tc: is O(logN)

*/
