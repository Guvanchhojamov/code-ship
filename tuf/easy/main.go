package main

import "fmt"

func main() {
	//fmt.Println(findMax([]int{5, -9, 3, 75, 1, -10, 0}))
	//findSecondMax([]int{3, 1, 6, 5})
	findSecondMaxOptimized([]int{5, -9, 3, 75, 1, -10, 6, 76})
}

func findMax(nums []int) int {
	maximum := nums[0]
	for i := 0; i < len(nums); i++ {
		maximum = max(maximum, nums[i])
	}
	return maximum
}

func findSecondMax(nums []int) {
	maximum := nums[0]
	secondMax := nums[0]
	for _, val := range nums {
		maximum = max(maximum, val)
	}

	for _, v := range nums {
		if v > secondMax && v != maximum {
			secondMax = v
		}
	}

	fmt.Println("max:", maximum, "secondMax:", secondMax)
}

func findSecondMaxOptimized(nums []int) {
	maximum, secondM := nums[0], nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > maximum {
			maximum = nums[i]
		} else if nums[i] > secondM {
			secondM = nums[i]
		}
	}

	fmt.Println("max:", maximum, "secondMax:", secondM)
}

/*

 */

func checkArraySortedRotated(nums []int) {

}
