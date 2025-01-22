package main

// 81. Search in Rotated Sorted Array II (Duplicates)
func main() {

}

func search(nums []int, target int) bool {
	var low, high = 0, len(nums) - 1

	for low <= high {
		mid := low + ((high - low) / 2)
		if nums[mid] == target {
			return true
		}

		if nums[low] == nums[mid] && nums[mid] == nums[high] {
			low += 1
			high -= 1
			continue
		}
		if nums[low] <= nums[mid] {
			// left part is sorted process this.
			if nums[low] <= target && target <= nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			// right part is sorted process this.
			if nums[mid] <= target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}

/*

  for _,val:=range nums {
    if val == target {
        return true
    }
  }
  return false

  simple solution is brute force.

  lets try solve this with binary search
  Input: nums = [2,5,6,0,0,1,2], target = 1 Output: true

   l=0
   h=6
   m=0
   if a.mid == target  -> return true
   if a.low < a.mid {
    // then left part is sorted process left
      if a.low < target && target{
         high = mid-1
      }else{
         low = mid+1
      }
   }else {
    // right part is sorted. Proceess right

   }
   we have an edge case in there
   if elements in
    low, mid, high are same:
    we cant know witch marp is sorted,so
    if mid element is not target and
    low and high are same as mid then
     we can eliminate low and high because we don't need them hey not same with target.

*/
