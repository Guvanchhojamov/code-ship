package main

import "fmt"

/*
 if sequence of element on arr bigger than n/2 then majority.
*/

func main() {
	//mel := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	//fmt.Println(mel)
	melOptimal := majElOptimal([]int{10, 9, 9, 9, 10})
	fmt.Println(melOptimal)
}

/*
169. Majority Element
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

*/

/*
	Input: nums = [2,2,1,1,1,2,2]
	Output: 2
*/
// func majorityElement(nums []int) int {
// 	numsMap := map[int]int{}
// 	for _, num := range nums {
// 		if _, ok := numsMap[num]; ok {
// 			numsMap[num]++
// 		} else {
// 			numsMap[num] = 1
// 		}
// 	}

// 	fmt.Println(numsMap)
// 	for number, numberCount := range numsMap {
// 		fmt.Println(number, numberCount)
// 		if numberCount > len(nums)/2 {
// 			return number
// 		}
// 	}

// 	return -1
// }

/*
 The optimal approach for Majority element is than find more sequenced element and
 after that check with, once with other lements in for loop.

 For example:
 [4,4,2,4,3,4,4,3,2,4]

count: 2
el: 4

[10,9,9,9,10]
count: 1
el: 9
*/
func majElOptimal(nums []int) int {
	count := 1
	element := nums[0]

	for i := 1; i < len(nums); i++ {
		fmt.Println(count, element)
		if count == 0 {
			element = nums[i]
			count++
		} else if element == nums[i] {
			count++
		} else {
			count--
		}
		fmt.Println(count, element)
	}
	fmt.Println(element, count)
	majorCount := 0
	for _, val := range nums {
		if element == val && majorCount/2 == 0 {
			return element
		}

	}
	return -1
}
