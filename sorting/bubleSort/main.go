package main

import "fmt"

func main() {
	sortNums([]int{5, 4, 3, 2, 1})
	recSort := recursiveSort([]int{5, 4, 3, 2, 1}, 0)
	fmt.Println("recSort", recSort)
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

func recursiveSort(nums []int, i int) []int {
	if i >= len(nums) {
		return nums
	}
	j := i + 1
	for j < len(nums) {
		if nums[i] > nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		}
		j++
	}
	i++
	return recursiveSort(nums, i)
}
