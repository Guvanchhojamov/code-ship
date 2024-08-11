package main

import "fmt"

func main() {
	nums := []int{5, 4, 3, 2, 1}
	sorted := mergeSort(nums)
	fmt.Println(sorted)
	//readerTest()
}

/*
  n=5     l=0  mid =n/2  r=len-1

  1. 5 4 3    2 1
    n=3 l=0  mid=n/2  r=len-1
    5 4        3
    n=2 l=0   mid=n/2  r=len-1
    5          4

  2.
*/

func mergeSort(nums []int) []int {
	mid := len(nums) / 2
	//right :=  len(nums)-1
	if len(nums) < 2 {
		return nums
	}

	//fmt.Println(nums[l:mid])
	//fmt.Println(nums[mid:r])

	leftSlice := mergeSort(nums[:mid])
	fmt.Println("leftSlice", leftSlice)
	rightSlice := mergeSort(nums[mid:])
	fmt.Println("rightSlice", rightSlice)
	return merge(leftSlice, rightSlice)
}

func merge(leftSlice, rightSlice []int) []int {
	i, j := 0, 0
	mergedSlice := []int{}
	for i < len(leftSlice) && j < len(rightSlice) {
		if leftSlice[i] < rightSlice[j] {
			mergedSlice = append(mergedSlice, leftSlice[i])
			i++
		} else {
			mergedSlice = append(mergedSlice, rightSlice[j])
			j++
		}
	}

	for i < len(leftSlice) {
		mergedSlice = append(mergedSlice, leftSlice[i])
		i++
	}
	for j < len(rightSlice) {
		mergedSlice = append(mergedSlice, rightSlice[j])
		j++
	}
	return mergedSlice
}
