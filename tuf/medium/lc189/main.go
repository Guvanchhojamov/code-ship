package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

// 189. Rotate Array
/*
	for k > 0 ...

	i=0 .. n-1

    1.swap nums[i] = [n-1] i++
*/
func rotate(nums []int, k int) {
	for k > 0 {
		for i := 0; i < len(nums); i++ {
			nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
		}
		k--
	}
	//fmt.Println(k)
}
