package main

import "fmt"

func main() {
	sortNums([]int{5, 4, 3, 2, 1})
}

func sortNums(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
			fmt.Println(nums)
		}
	}
	fmt.Println(nums)
}
