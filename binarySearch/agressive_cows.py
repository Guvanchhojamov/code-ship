
# Given an array of length ‘N’,
# where each element denotes the position of a stall.

# Now you have ‘N’ stalls and an integer ‘K’ 
# which denotes the number of cows that are aggressive. 
# To prevent the cows from hurting each other, 
# you need to assign the cows to the stalls, 
# such that
#   the minimum distance between any two of them is as large as possible. 
# Return the largest minimum distance.

# Example 1:
# Input Format:
#  N = 6, k = 4, arr[] = {0,3,4,7,10,9}
# Result:
#  3
# Explanation:
#  The maximum possible minimum distance between any two cows will be 3 when 4 cows are placed at positions 
#  {0, 3, 7, 10}.
#  Here the distances between cows are 3, 4, and 3 respectively. 
#  We cannot make the minimum distance greater than 3 in any ways.

#   What we need: 
#    - find minimum for each  stall 
#    - We need to place all  K cows. 
#    - Find max possible in minimums. 
# Approach: 
#     Brute force: 

#   As we need find minimum, we need to start from min valid value, 
#   check for each cow canWe place this cow in this stall or not, 
#   until K 
#      We need to range, 
#      What is the min distance cows can be it is=1 
#      What is the maximum distance between cows can be? 
#      It is the max(arr) - min(arr) this is the max distance can be. 
#      We sure that our max possible value in this range min...to..max()-min(arr)// or max(arr). 
#     We need helper function is possible. 
def agressiveCows(): 
    maxDistance = 1  
    arr= sort(arr) # type: ignore # // We need to sort array. 
    low = 0 
    high = arr[max] - arr[min]
    while  low <= high :
        mid = low+((high-low)/2) 
        if is_possible(arr, mid, k):
            low = mid+1 # // if ok we need to go search maximum possible
        else:
            high = mid-1 
    return high # // high stops in last possible position, is Max possible ans.    
    
def is_possible(arr, distance,k):
    placedCows=1
    lastCow = arr[0] 
    for i in range(0, len(arr)-1):
        if distance <= arr[i]-lastCow:
            lastCow = arr[i]
            placedCows+1
        if placedCows >= k:
                return True
           
    return False 
    # tc:O(n)
    # Sc:O(1)
    # 1 ok 
    # 2 ok 
    # 3 ok 
    # 4 not 
    # 5 not ... 