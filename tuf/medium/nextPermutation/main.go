package main

import "fmt"

func main() {
	fmt.Println("salam")
}

/*
Problem Statement:
Given an array Arr[] of integers,
 rearrange the numbers of the given array into the lexicographically
 next greater permutation of numbers.

If such an arrangement is not possible,
 it must rearrange to the lowest possible order (i.e., sorted in ascending order).
*/

func nextPermutation(nums []int) {
	var n = len(nums) - 1
	var idx = -1
	for i := n; i > 0; i-- {
		if nums[i-1] < nums[i] {
			idx = i - 1
			break
		}
	}
	// find the nearest biggest from right part and swap with break point.
	if idx != -1 {
		for i := n; i > idx; i-- {
			if nums[i] > nums[idx] {
				nums[i], nums[idx] = nums[idx], nums[i]
				break
			}
		}
	}

	// sort on dec order or reverse rigth part from back, to find nearest permutation.
	i, j := idx+1, n
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

/*
  1 2 3
  1 3 2

*/

/*
 nums = 2 1 5 3 0 0
 output = 2 3 0 0 1 5

 1. starting from back find where
    a[i-1] < a[i]. And make the indexes as needed break point.
 2. On the right part find the nearerest one for needed a[i-1].
    Find the minimum of in all nums for bigger than a[i-1]
    It is the minimum(where num> a[i-1]).
	Starting  from back sind some on grater than a[i-1], but the smallest one.
3. Swap the a[i-1] with this minimum number.
4. Sort or reverse from starting back,  right part. from a[i] -> end.
*/

/*

 n = len(nums)-1
 idx = n-1
  for i:=n;i>0;i--:
	if a[i-1] < a[i]:
	 idx = i-1
	 break;

  2. Find bigger than a[i-1], but smallest one.
	 right part always sorted on decreasing order.
	 so we can start from back and check if a[n]>a[idx] if yes, we must take first occurence.
	 and swap with a[idx], a[n]
	for i:=n;i>idx;i--:
		if a[i]>a[idx]:
			a[i],a[idx] = a[idx],a[i]
			break;
3. Just sort on inc order, or reverse right part from brek point.
   For take next permutation of nums.
   i,j = idx+1,n
   for i<j:
    a[i],a[j] = a[j], a[i]

	return nums.

*/
