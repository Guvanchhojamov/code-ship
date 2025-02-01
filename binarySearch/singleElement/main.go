package main

// 540. Single Element in a Sorted Array
func singleNonDuplicate(nums []int) int {
	var n = len(nums) - 1
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] != nums[1] {
		return nums[0]
	}
	if nums[n] != nums[n-1] {
		return nums[n]
	}

	var low, high = 1, n - 1
	for low <= high {
		mid := low + ((high - low) / 2)
		if nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
			return nums[mid]
		}
		if (mid%2 == 0 && nums[mid] == nums[mid+1]) ||
			(mid%2 != 0 && nums[mid] == nums[mid-1]) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

/*
  1. Brute force using map.
    tc: O(N)  sc:O(N)
  2. With using XOR logic:
     tc: O(n) sc: O(1)
  3.our target is single element.  Binary search
     Tc: O(longn)
    nums = [1,1,2,2,3,4,4,8,8,8]
    ans = nums[0]
   how we can use index?

   [1,1,2,2,3,4,4,8,8,8]
 i: 0 1 2 3 4 5 6 7 8 9
    a. if a[i] +=a[i-1] and a[i] != a[i+1] this is the single element.
     So how we can find that element?
     In binary search main thing is eliminate left or right part.
     How we can know in witch par our single element?
     How you can see,
     evenI, oddI,evenI,oddI.... singleI, oddI,evenI,oddI,evenI...

     we need to check mid index like:
     if mid==1



*/
