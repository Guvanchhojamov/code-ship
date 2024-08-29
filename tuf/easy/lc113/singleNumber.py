# Given a non-empty array of integers nums, every element appears twice except for one. 
# Find that single one.

# You must implement a solution with a linear runtime complexity and use only constant extra space.


# Example 1:

# Input: nums = [2,2,1]
# Output: 1

# Example 2:
# Input: nums = [4,1,2,1,2]
# Output: 4
# Example 3:

# Input: nums = [1]
# Output: 1

# Hint : XOR of two equal numbers gives 0 :)
# 6^2=4 how? 
#    1  1  0  (binary for 6)
# ^  0  1  0  (binary for 2)
# -------------
#    1  0  0  (result in binary)

def singleNumber(nums) -> int:
    res = 0
    for num in nums:
        print("res", res)
        print("num",num)
        print("res^num", res^num)
        res ^= num    
    return res

unique = singleNumber([4,1,2,1,2])
print(unique)