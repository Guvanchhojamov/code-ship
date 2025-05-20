package daily

/*
3355. Zero Array Transformation I
You are given an integer array nums of length n and a 2D array queries,
where queries[i] = [li, ri].

For each queries[i]:
 - Select a subset of indices within the range [li, ri] in nums.
 - Decrement the values at the selected indices by 1.
 - A Zero Array is an array where all elements are equal to 0.

Return true if it is possible to transform nums
into a Zero Array after processing all the queries sequentially,
otherwise return false.

Example 1:
Input: nums = [1,0,1,1], queries = [[0,2]]
Output: true
Explanation:
For i = 0:
Select the subset of indices as [0, 2] and decrement the values at these indices by 1.
The array will become [0, 0, 0], which is a Zero Array.

Example 2:
Input: nums = [4,3,2,1], queries = [[1,3],[0,2]]
Output: false
Explanation:
For i = 0:
Select the subset of indices as [1, 2, 3] and decrement the values at these indices by 1.
The array will become [4, 2, 1, 0].
For i = 1:
Select the subset of indices as [0, 1, 2] and decrement the values at these indices by 1.
The array will become [3, 1, 0, 0], which is not a Zero Array.

Constraints:
1 <= nums.length <= 10^5
0 <= nums[i] <= 10^5
1 <= queries.length <= 10^5
queries[i].length == 2
0 <= li <= ri < nums.length
*/
/*
  we are given nums, and we need transform zero array using the
	queries[i]=[li,ri] - subset in nums from left-right.
	- decrease nums[li,ri] - by -1 if nums[i]>0.
	- after transform all queries, if nums all elements are 0 return true else return false.
What is simple bf:
	- go throw queries.
	- take nums[li,ri] -> -1 from all elements if > 0 .
	- after iteration go throw nums and count zeros.
	- if zeros = len(nums) return true, else false.
tc: queries*nums = n*n = n^2  - worth case, when nums[1..10^5] and queries=[[1,10^],... 10^5]
sc: O(1)
But we got TLE for this solution.

Example 2:
Input: nums = [4,3,2,1], queries = [[1,3],[0,2]]
Output: false
Explanation:
For i = 0:
Select the subset of indices as [1, 2, 3] and decrement the values at these indices by 1.
The array will become [4, 2, 1, 0].
For i = 1:
Select the subset of indices as [0, 1, 2] and decrement the values at these indices by 1.
The array will become [3, 1, 0, 0], which is not a Zero Array.
 How can we optimize?
- take min and max of the li,ri ->  mini, maxi
- max from nums = mxn
 if queries[i][1] in [mini..maxi] inclusively:
	count++
- go nums:
if i in [mini..maxi] && nums[i] > 0:
	nums[i]-count
- iterate nums if nums[i] > 0 return false else true.

Can we use difference array and prefix sum to check if an index can be made zero?
nums = [4,3,2,1], queries = [[1,3],[0,2]]
px = [4,7,9,10]


Approach: Difference Array
Intuition
We count the maximum number of operations that can be performed at each position using a difference array.
Construct the difference array deltaArray with a length of n + 1
(where n is the length of the array nums),
which is used to record the increment for each query on the number of operations.

For each query interval [left, right],
increment deltaArray[left] by +1,
indicating an increase in the operation count starting from left.
Decrement deltaArray[right + 1] by -1,
indicating that the operation count returns to its original value after right + 1.

Next, perform a prefix sum accumulation on the difference array deltaArray
to obtain the total operation count at each position in the array,
storing these counts in operationCounts.
Traverse the nums array and the operationCounts array,
comparing the actual operation counts (operations) at each position to see if they meet the minimum number
of operations (target) required for zeroing.
If all positions meet operations >= target, return true; otherwise, return false.

Implementation
Complexity Analysis

class Solution:
    def isZeroArray(self, nums: List[int], queries: List[List[int]]) -> bool:
        deltaArray = [0] * (len(nums) + 1)
        for left, right in queries:
            deltaArray[left] += 1
            deltaArray[right + 1] -= 1
        operationCounts = []
        currentOperations = 0
        for delta in deltaArray:
            currentOperations += delta
            operationCounts.append(currentOperations)
        for operations, target in zip(operationCounts, nums):
            if operations < target:
                return False
        return True
*/
/*
nums = [4,3,2,1], queries = [[1,3],[0,2]]
 d = [1,1,0,-1,-1]
opc = [1,2,2,1,0]
curr_op = 0
*/
