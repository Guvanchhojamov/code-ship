from typing import List


def missingNumber(nums: List[int]) -> int:
    i=0
    # Sort the numbers using cyrcle sort.
    while i < len(nums):
        if nums[i]<len(nums) and nums[i]!= nums[nums[i]]: 
            nums[nums[i]],nums[i]=nums[i],nums[nums[i]]
        else:
            i+=1

    print(nums)
    # Return number as answer if number is not in right index.
    for i in range(len(nums)):
        print("i",i, "nums[i]",nums[i])
        if i != nums[i]:
            return i
    return len(nums)

print("missing number is: ", missingNumber([9,6,4,2,3,5,7,0,1]))