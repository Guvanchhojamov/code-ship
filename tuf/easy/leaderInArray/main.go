package main

func main() {

}

/*
Problem statement:
	Given an array, print all the elements witch are leaders.
	A lader is an element that is greater than all
	 of the elements on it's right side in the array.

Example 2:
Input:

 arr = [10, 22, 12, 3, 0, 6]
Output:

 22 12 6
Explanation:

 6 is a leader. In addition to that, 12 is greater than all the elements in its right side (3, 0, 6), also 22 is greater than 12, 3, 0, 6.
*/
/*
  1. The last element is always leader.
	first solution is brute foce 2 llop.
	a. In first loop take each element
 	b. In second loop check a[i] big or not of all elements after that elent, up to end.
	c. If isLeader = true add a[i] in res array.
	Time complexity: ~O(n^2)
	Space complexity: O(1). Not used any extra space. The result array not a extra space.
*/
/*
	2. Second solution is optimal.
		a. Add  last element to the res array. Because last element always a leader.
			and element = a[last]
		b. Start from len(nums)-2.
		c.  if a[i]>element:
			res =  append.a[i]
			element =  a[i]
			else:
			do nothing just continue.
		If order of result is important just reverse res and return.
		TC: O(n)
		SP: O(1). We don't use any extra space here.
*/
func leadersInArray(nums []int) (res []int) {
	// Asumen 1<len(nums)<10^5
	if len(nums) == 1 {
		return nums
	}
	res = append(res, nums[len(nums)-1])
	var maxInRight = nums[len(nums)-1]
	for i := len(nums) - 2; i > 0; i-- {
		if nums[i] > maxInRight {
			res = append(res, nums[i])
			maxInRight = nums[i]
		}
	}
	return
}
