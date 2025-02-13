package daily

import "sort"

/*
You are given a 0-indexed integer array nums, and an integer k.

In one operation, you will:
- Take the two smallest integers x and y in nums.
- Remove x and y from nums.
- Add min(x, y) * 2 + max(x, y) anywhere in the array.

Note that you can only apply the described operation if nums contains at least two elements.
Return the minimum number of operations needed so that all elements of the array are greater than or equal to k.



Example 1:
Input: nums = [2,11,10,1,3], k = 10
Output: 2
Explanation: In the first operation, we remove elements 1 and 2, then add 1 * 2 + 2 to nums. nums becomes equal to [4, 11, 10, 3].
In the second operation, we remove elements 3 and 4, then add 3 * 2 + 4 to nums. nums becomes equal to [10, 11, 10].
At this stage, all the elements of nums are greater than or equal to 10 so we can stop.
It can be shown that 2 is the minimum number of operations needed so that all elements of the array are greater than or equal to 10.
Example 2:

Input: nums = [1,1,2,4,9], k = 20
Output: 4
Explanation: After one operation, nums becomes equal to [2, 4, 9, 3].
After two operations, nums becomes equal to [7, 4, 9].
After three operations, nums becomes equal to [15, 9].
After four operations, nums becomes equal to [33].
At this stage, all the elements of nums are greater than 20 so we can stop.
It can be shown that 4 is the minimum number of operations needed so that all elements of the array are greater than or equal to 20.

Constraints:

2 <= nums.length <= 2 * 10^5
1 <= nums[i] <= 10^9
1 <= k <= 10^9
The input is generated such that an answer always exists.
That is, there exists some sequence of operations after which all elements of the array are greater than or equal to k.

*/

/*
  What we need?
    - make given operation in optimal way until break point and return operations count.
    Key points:
     - We always need 2 small elements.
     - We need append new value after each operation.
     - We need to check after each operation elems >= k.
     - For do operation we need at least 2 nums in array.
     - TC is big so we need optimal solution. (n*logn near about)
1. Brute force:
    1. find 2 small elements. (N)
    2. Remove them from array. (N)
    3. Add new value to array. (1)
    4. Check all elements >= k. O(n)
    5. Do operations until breakpoint. N for example.
      TC: n*(n*n*n)= N^4 near about.
      SC: O(1)
2. How can we optimize it?
        In sorted array always.
    1. The first 2 elements is in sorted array. a[0], a[1] Always.
    2. add to array min(a[0], a[1]) *2 + max(a[0], a[1]) = a[0]*2+a[1]. Always.
    3. if arr[0] >= k return operations_count. Always. If first is greater others too in sorted array.
    4....
 What DS we can use?
    Priority queue.

*/
Input: nums = [2,11,10,1,3], k = 10
1 2 3 10 11
1*2+2 =5
5 3 10 11
3*2 + 5 = 11
11 10 11
/*
 1.Sort array

  for each element:
 2.Take first 2 and sort.
 3.if a[0]*2+a[1] >= return opCount
 4.else check arr[1]= a[0]*2+a[1]
 5. arr = arr[1:]
*/
/*
  100 101 102 103    k=105
  301 102 103

*/

func minOperations(nums []int, k int) int {
	var opsCount = 0

	// var smallNums = make([]int,2,2)
	for i:=0;i<len(nusm);i++{
		sort.Ints(nums)

		if len(nums)>1 {
			newNum = nums[0]*2 + nums[1]
			if (newNum >= k && nums[1] >=k ) {
				return opsCount+1
			}
			nums[1] = opNum
			nums = nums[1:]
		}else{
			if nums[1]>=k {
				return opsCount+1
			}
		}
		reutrn opsCount
	}

	/*
	   100 101 102 103   k=105
	   opsC =0
	   100 101 102 103
	   301>k 101<
	   301 102 103
	   oc=1
	   102 103 301
	   307>105 and 103> 105
	   oc+1=2
	   307 301
	   301 307

	   [1,1,2,4,9] k=20
	    op=0
	    3 > 20 and 2>20
	    op=1
	    3 2 4 9



	*/




	func minOperations(nums []int, k int) int {
		opsCount := 0
		for {
		sort.Ints(nums)
		if nums[0] >= k { // Check if the smallest element is already >= k
		return opsCount
	}

		if len(nums) == 1 {
		if nums[0] >= k {
		return nums[0]
	}
	}

		newNum := nums[0]*2 + nums[1]
		nums = append([]int{newNum}, nums[2:]...)
		opsCount++
	}
	}


