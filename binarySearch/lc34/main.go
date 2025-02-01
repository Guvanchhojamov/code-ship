package main

// 34. Find First and Last Position of Element in Sorted Array
func main() {

}

func searchRange(nums []int, target int) []int {
	var low, high = 0, len(nums) - 1
	var startIndex, endIndex = -1, -1
	// find start index.
	for low <= high {
		mid := low + ((high - low) / 2)
		if nums[mid] >= target {
			//go to small lower bound
			if nums[mid] == target {
				startIndex = mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	low, high = 0, len(nums)-1
	// find end index.
	for low <= high {
		mid := low + ((high - low) / 2)
		if nums[mid] <= target {
			//go to big
			if nums[mid] == target {
				endIndex = mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return []int{startIndex, endIndex}
}

/*
 Tc must be logN, so we cant use brute force.
 Let't try to use binary search.

 Example 1:

 Input: nums = [5,7,7,8,8,10], target = 8
 Output: [3,4]
  what is the start position?
   smallest index where nums[i] == target
  what is the end position?
   biggest index where nums[i] == target ?
1. start position:
    implement like lower bound.
      if nums[mid]>=target:
        if nums[mid] == target startPosition = target
        // go to small
         high = mid-1
      else:
         if nums[mid] == target endPosition = target
         // to to big
         low=mid+1
       return start and end positions.
*/
