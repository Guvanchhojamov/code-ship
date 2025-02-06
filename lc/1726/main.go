package main

func main() {

}

/*
1726. Tuple with Same Product
Given an array nums
of distinct positive integers,
return the number of tuples (a, b, c, d)
such that a * b = c * d where
a, b, c, and d are elements of nums,
and a != b != c != d.


Example 1:
Input: nums = [2,3,4,6]
Output: 8
Explanation: There are 8 valid tuples:
(2,6,3,4) , (2,6,4,3) , (6,2,3,4) , (6,2,4,3)
(3,4,2,6) , (4,3,2,6) , (3,4,6,2) , (4,3,6,2)

Example 2:
Input: nums = [1,2,4,5,10]
Output: 16
Explanation: There are 16 valid tuples:
(1,10,2,5) , (1,10,5,2) , (10,1,2,5) , (10,1,5,2)
(2,5,1,10) , (2,5,10,1) , (5,2,1,10) , (5,2,10,1)
(2,10,4,5) , (2,10,5,4) , (10,2,4,5) , (10,2,5,4)
(4,5,2,10) , (4,5,10,2) , (5,4,2,10) , (5,4,10,2)


Constraints:

1 <= nums.length <= 1000
1 <= nums[i] <= 10^4
All elements in nums are distinct.
*/

/*
   What is the porblem state?
       key points:
        - we need just 4 elements: a,b,c,d
        - nums are distinct in array.
        - we need return just count of possible combinations and
          so gerate all possible abcd-s not required.
   1. Brute force:
       In 4 loop genereate all possible combinations
        where i!=j!=k!=l
         if i*j == k*l count +
         tc: n^4
         sc: o(1)
       is it ok? Maybe but how can we optimize it?
   lets try to solve brture foce.

*/

func tupleSameProduct(nums []int) int {
	var count = 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				for l := 0; l < len(nums); l++ {
					if (i != j && j != k && k != l && j != l) && (nums[i]*nums[j] == nums[k]*nums[l]) {
						count++
					}
				}
			}
		}
	}
	return count
}

/*
We got time limit exeed. How can we optimize?
*/
func tupleSameProduct(nums []int) int {
	var totalTuplesNumber = 0
	var pairFrequency = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			pairFrequency[nums[i]*nums[j]]++
		}
	}

	for _, frequency := range pairFrequency {
		equalPorodcutPair := (frequency - 1) * frequency / 2
		totalTuplesNumber += 8 * equalPorodcutPair
	}

	return totalTuplesNumber
}

/*
 solved using hint and editioral.
*/
