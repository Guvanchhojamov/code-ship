# 1752. Check if Array Is Sorted and Rotated

def check(nums) -> bool:
    rotations = 0
     
    for i, value in enumerate(nums):
        # nums[-1] takes last value of nums.
        if nums[i - 1] > value:
            rotations += 1
        return rotations <= 1



is_sorted = check([3,4,5,1,2])
print("is sorted:", is_sorted)
