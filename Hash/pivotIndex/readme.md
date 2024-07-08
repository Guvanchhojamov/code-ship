# 724. Find Pivot Index

**Difficulty:** Easy

## Description

Given an array of integers `nums`, calculate the pivot index of this array.

The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.

If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.

Return the leftmost pivot index. If no such index exists, return -1.

## Examples

### Example 1

**Input:**

```plaintext
nums = [1,7,3,6,5,6]
Explanation:

The pivot index is 3.

Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11

Right sum = nums[4] + nums[5] = 5 + 6 = 11
```

## Solution
```go
func pivotIndex(nums []int) int {
	pSums := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		pSums = append(pSums, pSums[i-1]+nums[i])
	}

	for i := range nums {
		leftSum := getSum(0, i-1, pSums)
		rightSum := getSum(i+1, len(nums)-1, pSums)
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}

func getSum(l, r int, pSums []int) int {
	switch {
	case l > r:
		return 0
	case l > 0:
		return pSums[r] - pSums[l-1]
	}
	return pSums[r]
}
```