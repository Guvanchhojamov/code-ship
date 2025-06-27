

# Given a binary array nums and an integer k,
#  return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

# Example 1:

# Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
# Output: 6
# Explanation: [1,1,1,0,0,1,1,1,1,1,1]
# Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
# Example 2:

# Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
# Output: 10
# Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
# Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
 

# Constraints:

# 1 <= nums.length <= 10^5
# nums[i] is either 0 or 1.
# 0 <= k <= nums.length

# Need find max array with cons 1's and with k 0s witch we can flip. 
# Brute force: 
#     generate all consequtive arrays with given k zeors: 
#     zeros = 0 
#     for i=0 ->n 
#        for j=i->n 
#        if zeros > k {break}
#        if n[i]==0: zeros ++ 
#        mx = max(mx, j-i+1)
# tc: O(n*n)
# scO(1)

 def longestOnes(self, nums: List[int], k: int) -> int:
        zeros = 0 
        mx = 0 
        for i in range(len(nums)):
            zeros = 0
            for j in range(i,len(nums)):
                if nums[j] == 0: zeros+=1
                if zeros > k: break 
                mx = max(mx,j-i+1)
        return mx

# we got TLE error solve optimal one in go.
# optimal one using pointer and sliding window approach. 
#  How can we optimze using sliding window and binary search algos?
#  Input: nums = [1,1,1,0,0,0,1,1,1,1], k = 2
#   Output: 6
# 	take left and right pointers for bundairies of window:
# 	start interate j until zeros>k;
# 	nums[j] == 0 increase zeros.
# 	when zeros > k:
# 	increate left if nums[left]==0 decrease zeros, for nexxt time increse right.
# 	each time check len for maximum mx = max(mx,j-i+1)
# 	continue this while j<n: and return mx.

 def longestOnes(self, nums: List[int], k: int) -> int:
        zeros=0 
        mx = 0 
        left,right = 0,0 
        while right<len(nums):
            if nums[right] == 0: zeros+=1
            if zeros > k: 
                if nums[left] ==0: zeros-=1 
                left+=1
            mx = max(mx,right-left+1)
            right+=1
        return mx