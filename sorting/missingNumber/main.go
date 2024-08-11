package main

import "fmt"

func main() {
	println("hello")
	sortNums([]int{5, 4, 3, 2, 1})
}

func sortNums(nums []int) {
	i := 0
	for i < len(nums) {
		if nums[i]-1 != i {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		} else {
			i++
		}
	}
	fmt.Println(nums)
}
