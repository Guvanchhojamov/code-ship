package main

import "fmt"

func main() {
	specialDiscount([]int{10, 1, 1, 6})
}

/*
8 4 6 2 3

4 2 4 2 3

8:4
4:2
6:2
2:0
3:0

*/

/*
8 4 6 2 3

stack.push()

	for len(stack)>0 && stack[n]>nums[i+1]:
		discountMap[stack[n]] = stack[n]-nums[i+1]
		stack.pop()

8:4
6:2
4:2
10 1 1 6
*/
func specialDiscount(nums []int) []int {
	stack := []int{nums[0]}
	discountMap := map[int]int{}
	for i := 1; i < len(nums); i++ {
		//discountMap[nums[i]] = 0
		for len(stack) > 0 && stack[len(stack)-1] >= nums[i] {
			discountMap[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	fmt.Println(discountMap)
	for i, val := range nums {
		nums[i] = nums[i] - discountMap[val]
		discountMap[val] = 0
	}
	fmt.Println(nums)
	return nums
}
