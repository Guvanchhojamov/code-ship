package bit

/*
260. Single Number III
Given an integer array nums,
in which exactly two elements appear only once and all the other elements appear exactly twice.
Find the two elements that appear only once. You can return the answer in any order.

You must write an algorithm that runs in linear runtime complexity and uses only constant extra space.

Example 1:
Input: nums = [1,2,1,3,2,5]
Output: [3,5]
Explanation:  [5, 3] is also a valid answer.

Example 2:
Input: nums = [-1,0]
Output: [-1,0]

Example 3:
Input: nums = [0,1]
Output: [1,0]

Constraints:

2 <= nums.length <= 3 * 10^4
-2^31 <= nums[i] <= 2^31 - 1
Each integer in nums will appear twice, only two integers will appear once.
*/
/*
  We need return  2 unique element all other elements appears twice.
Simple bf:
	create map, fill with num:frequency:
	return 2 uniq's num:1.
tc: N; sc:N/2+2 = N
So, we must use linear tc and not use any extra space we need to think other solution.
Input: nums = [1,2,1,3,2,5]
Output: [3,5]
  in this related problems bit manipulation is very helpful.
a^a = 0
0^a = a
   (1^1^2^2 = 0 ) ^ (3^5), so what gives us 3^5?
    011 ^ 101 = 110 = 6
From there the nums are uniq minimum the one bit must be different from others.
	Can we check last set bit of each number in nums. Yes.
Since we have 2 uniq nums we need to use buckets, approach.
We take 2 buckets a,b.
   since,
a= 001 001, 010, 010, 011  -> the last set bits index are same for this nums.
b= 101 -> the last set bit index is different.
doing xor of elements a and b, we still a=3, and b = 5, and this is our answer.
	So, what we need to do?
- loop in nums.
- take first  set bit ()    =  110 & 1<<i = i-index of first set bit.
or
 110 & 011 = 010 ^ 110 = 100
-
*/

func singleNumber(nums []int) []int {
	var a, b = 0, 0
	var x = 0
	for _, v := range nums {
		x ^= v // still uniq1^uniq2 only.
	}
	lastSetBitMinNum := x & -x
	for _, num := range nums {
		if lastSetBitMinNum&num > 0 { // if the bit of num is set then
			a ^= num
		} else { //else
			b ^= num
		}
	}
	return []int{a, b}
}
