package main

// 153. Find Minimum in Rotated Sorted Array
func main() {

}

func findMin(nums []int) int {
	var low, high = 0, len(nums) - 1
	var ans = nums[0]

	for low <= high {
		mid := low + ((high - low) / 2)
		if nums[low] <= nums[mid] { // use <= for when low==mid
			// left part is sorted
			ans = min(ans, nums[low])
			low = mid + 1
		} else {
			// right part issorted. If sorted both it, takes left by default .
			ans = min(ans, nums[mid])
			high = mid - 1
		}
	}

	return ans
}

/*
  1. Brute force
    First what coming to mind is brute force
    Iterate all array and finc minimum number and return it.
    TC: O(n)
    Sc: O(1)
    But this is not acceptable.
  2. We have sorted array and we need to find something in sorted array.
     This notice us we can use binary search there
     With binary search our:
      tc: O(logN) by default.
      So, how can we use binary search in there?

    Since, array is sorted we can use
      Seperatign sorted part and not sorted part,
       1. Find sorted part, do BS in sorted part.
          But minimum can be in not sorted part, so
          nums = [4,5,6,7,0,1,2]
          l=4
          h=6
          m=3->7
          left: 4 5 6 7
          righ: 0 1 2
          each part is sorted, so
          if we go to left we don't find exactly minimum.
          so what we can do?
          we can go to take
          Take a minimum  of sorted part and eliminate this part. Might this min is our minimum.
          and go to another part, make same thing to .
          so stpes.
          1. find sorted part, take minimum, eliminate.
          2. Go to another part do same.
          In sorted array the minimum is the first elemnt so, arr[low]==minimum.
          return last got-ed minimum.
          let's implement.
*/
