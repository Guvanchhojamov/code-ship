package daily

/*
2016. Maximum Difference Between Increasing Elements
Given a 0-indexed integer array nums of size n,
find the maximum difference between nums[i] and nums[j]
(i.e., nums[j] - nums[i]),
- such that 0 <= i < j < n and nums[i] < nums[j].
Return the maximum difference. If no such i and j exists, return -1.

Example 1:

Input: nums = [7,1,5,4]
Output: 4
Explanation:
The maximum difference occurs with i = 1 and j = 2, nums[j] - nums[i] = 5 - 1 = 4.
Note that with i = 1 and j = 0, the difference nums[j] - nums[i] = 7 - 1 = 6, but i > j, so it is not valid.

Example 2:
Input: nums = [9,4,3,2]
Output: -1
Explanation:
There is no i and j such that i < j and nums[i] < nums[j].

Example 3:
Input: nums = [1,5,2,10]
Output: 9
Explanation:
The maximum difference occurs with i = 0 and j = 3, nums[j] - nums[i] = 10 - 1 = 9.

Constraints:
n == nums.length
2 <= n <= 1000
1 <= nums[i] <= 10^9
*/
/*
  ned find max diff between a[j]-a[i]:
- a[i]< a[j]
- i<j
- if no valid pair return -1.
-since i must be less than j, and a[j]>a[i] always  it is good start from back,
-instead start from front.
maxdiff = -1
while j=n->0:
	while i=n-1->0:
	if a[j]>a[j] :
	maxdiff = max(maxdiff,a[j]-a[i])
return maxdiff
tc: n*n
sc:1
/*
is it ok? Yes, but can we optimze n*n, maybe?
waht is the maxdiff, it is when maxEl-minEl=maxDiff, so,
When we find maxEL index, and find minEL until maxElIndex our condition is valid i<j
assume j=maxElindex, find minEl index i<j. If not found return -1
if found minEl, maxEl-MinEl must be max diff, return it.
tc: N+N -> for max, for min, assume max el is the last one.
sc: O(1)
try impliment optimized;
nums = [1,5,2,10];
maxi=3
mini=0 =  10-1=9.
2. [9,4,3,2]
maxi=0
until 0 we dont have min el, so return -1.
*/
/* The second approach was wrong.*/
func maximumDifference(nums []int) int {
	var maxDiff = -1
	for j := len(nums) - 1; j >= 0; j-- {
		for i := j - 1; i >= 0; i-- {
			if nums[j] > nums[i] {
				maxDiff = max(maxDiff, nums[j]-nums[i])
			}
		}
	}
	return maxDiff
}

/*
[7,1,5,4]; res = 4

*/
/*
Approach: Prefix Minimum Value
Intuition
When we fix j,
the chosen index i must satisfy 0≤i<j and nums[i] must be the smallest among those indices.
Therefore, we can iterate over j while maintaining the prefix minimum of nums[0..j−1],denoted as premin.
In this way:
 If nums[j]>premin, we update the answer with nums[j]−premin.
 Otherwise, we update the prefix minimum value premin using nums[j].

Implementation

Complexity Analysis
Time complexity: O(n).

We only need to traverse the array nums once.

Space complexity: O(1).

func maximumDifference(nums []int) int {
	ans := -1
	for i, preMin := 1, nums[0]; i < len(nums); i++ {
		if nums[i] > preMin {
			ans = max(ans, nums[i]-preMin)
		} else {
			preMin = nums[i]
		}
	}
	return ans
}
 - [7,1,5,4]; res = 4
preMin = 7;

*/
