package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 && nums[0]+nums[1] == target {
		return []int{0, 1}
	}
	var targetNums []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				targetNums = append(targetNums, i)
				targetNums = append(targetNums, j)
				return targetNums
			}
		}
	}
	return targetNums
}
