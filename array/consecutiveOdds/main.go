package main

import "fmt"

func main() {
	fmt.Println(threeConsecutive([]int{1, 1, 1}))
}

func threeConsecutive(nums []int) bool {
	stack := []int{}
	for _, v := range nums {
		if v%2 != 0 {
			stack = append(stack, v)
		} else {
			stack = []int{}
		}
		if len(stack) == 3 {
			return true
		}
	}
	return false
}
