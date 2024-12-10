package main

import "fmt"

func main() {
	longestSuccessiveElements([]int{7, 7, 15, 20, 56, 67})
}

/*
Problem Statement:
You are given an array of ‘N’ integers.
You need to find the length of the longest sequence which contains the consecutive elements.
*/

func longestSuccessiveElements(nums []int) int {
	var count int = -1
	var numsMap = make(map[int]bool, 0)
	var incVal int = 1
	var maxSuc = 0
	for _, num := range nums {
		numsMap[num] = true
	}

	for num := range numsMap {
		incVal = 1
		count = -1
		for numsMap[num+incVal] {
			count++
			incVal++
		}
		maxSuc = max(maxSuc, count+1)
	}
	fmt.Println(numsMap, maxSuc)
	return 0
}

/*
Example 1:
Input:
 [100, 200, 1, 3, 2, 4]

Output:
 4

Explanation:
 The longest consecutive subsequence is 1, 2, 3, and 4.

 [3, 8, 3, 5, 7,7,7,7 6,6] out = 5 6 7 8 = 4

 3 3 5 6 6 7 7 7 7 8

 1. sort O(nlogN)
 smallElem = a[0]
 longest = 0
 count = 0
 for i in range nums:
	longest = max(longest,count)
	if arr[i]-1 == smallElem:
		smallElem = a[i]
		count++
	else if a[i]!=smallElem:
		smallElem = a[i]
		count = 0


*/

/*
try to solve with optimal approach.
using hashmap.
*/
