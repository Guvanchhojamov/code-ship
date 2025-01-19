package main

// 33. Search in Rotated Sorted Array
func main() {
	search([]int{1, 3, 5, 6}, 0)
}

func search(nums []int, target int) int {
	var low, high = 0, len(nums) - 1
	for low <= high {
		mid := low + ((high - low) / 2)
		if nums[mid] == target {
			return mid
		}
		if nums[low] <= nums[mid] {
			// Left part is sorted
			if nums[low] <= target && target <= nums[mid] {
				high = mid - 1 // illuminate right part
			} else {
				// target is no in right part illuminate right part.
				low = mid + 1
			}
		} else {
			// right part is sorted
			if nums[mid] <= target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

/*
Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0   7
Output: 4

we need logN tc so we cant use brute force.
Let's try binary search.
 1. Need to find witch part is sorted  is sorted = a[low]<=a[mid]
 2. if left part is sorted then
    if target in low<= target <= med  illuminate rigth part with -> high=mid-1
 3. else illuminate left part with ->  low = mid+1

 [4,5,6,7,0,1,2]   t=4

  mid=3
  lm=3

*/
