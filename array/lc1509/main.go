package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minDiffLowAndHigh([]int{1, 4}))
}

func minDiffLowAndHigh(nums []int) int {
	slices.Sort(nums)
	min := nums[0]
	if len(nums) < 3 {
		return 0 // we can change all 2 values
	}

	nums[len(nums)-1], nums[len(nums)-2], nums[len(nums)-3] = min, min, min

	var max = min // we have numbers less than 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max - min
}
