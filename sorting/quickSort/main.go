package main

import "fmt"

func main() {
	nums := []int{10, -2, 30, 90, 40, 50, -1}
	fmt.Println(quickSort(nums))
}

func quickSort(nums []int) []int {
	//var pivot int

	if len(nums) < 2 {
		return nums
	}

	pivot := nums[len(nums)-1]
	i := 0
	fmt.Println(pivot, 0, "nums before modify:", nums)
	for j := i; j < len(nums); j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		// fmt.Println("nums", nums, "pivot", pivot, "i", i, "j", j)
	}

	// swap pivot to it's right position.
	nums[len(nums)-1], nums[i] = nums[i], nums[len(nums)-1]
	fmt.Println("nums after modify", nums, pivot, i)
	// return i as pivot new right position.

	quickSort(nums[:i])
	quickSort(nums[i:])
	//fmt.Println(nums, pivot)
	return nums
}

// we use this function for optimize

// func partition(nums []int, low int) ([]int, int) {
// 	pivot := nums[len(nums)-1]
// 	i := low
// 	fmt.Println(pivot, low, "nums before modify:", nums)
// 	for j := i; j < len(nums); j++ {
// 		if nums[j] < pivot {
// 			nums[i], nums[j] = nums[j], nums[i]
// 			i++
// 		}
// 		// fmt.Println("nums", nums, "pivot", pivot, "i", i, "j", j)
// 	}

// 	// swap pivot to it's right position.
// 	nums[len(nums)-1], nums[i] = nums[i], nums[len(nums)-1]
// 	fmt.Println("nums after modify", nums, pivot, i)
// 	// return i as pivot new right position.
// 	return nums, i
// }
