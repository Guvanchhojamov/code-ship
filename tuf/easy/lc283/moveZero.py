
def moveZeroes(nums):
    i=0
    j=len(nums)-1
    while i < len(nums):
        if nums[i] == 0:
            nums[i],nums[j] = nums[j], nums[i]
            j-1
        else:
            i+1
    print(nums)


nums = [0,1,0,3,12]
moveZeroes(nums)