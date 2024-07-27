from typing import List


def missingNumber(nums: List[int]) -> int:
    i=0
    # Sort the numbers using cyrcle sort.
    while i < len(nums):
        if nums[i]<len(nums) and nums[i]!= nums[nums[i]]: 
            nums[nums[i]],nums[i]=nums[i],nums[nums[i]]
        i+=1

    # Return number as answer if number is not in right index.
    for i in range(len(nums)):
        if i != nums[i]:
            return i 
    return len(nums)

print("missing number is: ", missingNumber([3,0,1]))