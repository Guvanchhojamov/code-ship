package main

import "fmt"

func main() {
	fmt.Println(threeConsecutive([]int{1, 2, 34, 3, 4, 5, 7, 23, 12}))
}

func threeConsecutive(nums []int) bool {
	stack := []int{}
	for _, v := range nums {
		if len(stack) == 3 {
			return true
		}

		if v%2 != 0 {
			stack = append(stack, v)
		} else {
			stack = []int{}
		}
	}
	return false
}
