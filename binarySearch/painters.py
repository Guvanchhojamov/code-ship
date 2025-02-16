# Problem statement
# Given an array/list of length ‘n’,
# where the array/list represents the boards and each element of the given array/list represents the length of each board.
# Some ‘k’ numbers of painters are available to paint these boards.
# Consider that each unit of a board takes 1 unit of time to paint.

# You are supposed to return the area of the minimum time to get this job done of painting all the ‘n’ boards
# under a constraint that any painter will only paint the continuous sections of boards.

# Example :
# Input: arr = [2, 1, 5, 6, 2, 3], k = 2
# Output: 11
# Explanation:
# First painter can paint boards 1 to 3 in 8 units
# of time and the second painter can paint boards 4-6 in 11 units of time.
# Thus both painters will paint all the boards in max(8,11) = 11 units of time.
# It can be shown that all the boards can't be painted in less than 11 units of time.

# Sample Input 1 :
# 4 2
# 10 20 30 40
# Sample Output 1 :
# 60
# Explanation For Sample Input 1 :
# In this test case, we can divide the first 3 boards for one painter and the last board for the second painter.

#   What we need:
#     Helper function witch returns pointers count.
#     Find range min and max possible answer.
#     from min to max chek with helper function painters_count == k
#     if ok this is our answer and first possible and minimum answer.
# What is the minimum possible answer.
# The max num in  array would be our minimum answer.
# What is the maximum answer is the sum of all nums in array.
#  So we need to check in range min to max.
#  We can call it paintLimit.

def findLargestMinDistance(boards:list, k:int):
    # Write your code here
    # Return an integer

    low,high = max(boards), sum(boards)
    while low<=high:
        mid = low + ((high-low) // 2)
        if paintersCount(boards,mid) <= k:
            high = mid-1
        else:
            low = mid+1
    return low
    pass

def paintersCount(boards: list, paint_limit: int ) -> int:
    painters = 1
    painted_boards = boards[0]

    for i in range(1,len(boards)):
        if painted_boards+boards[i] <= paint_limit:
            painted_boards += boards[i]
        else:
            painters+=1
            painted_boards = boards[i]
    return painters