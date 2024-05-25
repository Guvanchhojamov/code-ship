package main

import "fmt"

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}

func moveZeroes(nums []int) {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] != 0 {
			if nums[l] == 0 {
				nums[l], nums[r] = nums[r], nums[l]
			}
			l++
		}
		r++
	}
	fmt.Println(nums)
}
