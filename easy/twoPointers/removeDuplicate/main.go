package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

// my solution
func removeDuplicates(nums []int) int {
	k := 1
	for l, r := 0, 1; r < len(nums); r++ {
		if nums[l] == nums[r] {
			nums = append(nums[:r], nums[r+1:]...)
			r--
		} else {
			l++
			k++
		}
	}
	return k
}

// best solution
func removeDuplicates2(nums []int) int {
	l := 1
	for r := l; r < len(nums); r++ {
		if nums[r] != nums[r-1] {
			nums[l] = nums[r]
			l++
		}
	}
	return l
}
