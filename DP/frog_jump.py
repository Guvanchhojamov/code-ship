from os import *
from sys import *
from collections import *
from math import *

from typing import *

def frogJumpK(n: int, k:int, heights: List[int]) -> int:
    dp = [0] * (n+1)
    for i in range(1,n):
        minCost = abs(heights[i]-heights[i-1])
        for j in range(i+1,len(k)):
            if i-j >= 0:
                cost = dp[i-j] + abs(heights[i] - heights[i-j])
            dp[i] = min(minCost, cost)
    return dp[n-1]
        
def abs(x):
    if x < 0:
        return x*(-1)
    return x
    pass

# k=3; 
# [10 20 30 10] 
#  0,1,2,3 
# [0,0,0,0] 

# def frogJump(n: int, heights: List[int]) -> int:
#     dp = [0] * (n+1)
#     prev_old = 0 
#     prev = abs(heights[1] - heights[0])
#     for i in range(2,n):
#         first =prev + abs(heights[i] - heights[i-1])
#         second=float('inf') 
#         if i > 1:
#             second = prev_old + abs(heights[i] - heights[i-2])
#         curr = min(first, second)
#         prev_old = prev
#         prev = curr
#     return dp[n-1]
        
# def abs(x):
#     if x < 0:
#         return x*(-1)
#     return x
'''
[10 20 30 10] 
 0,1,2,3 
[0,0,0,0] 

for brute force recursive solution we got TLE. 
Let's do memoizaiton on this solution. 
 Storing each result in dp[] array. to quick take and dont 
 process again preoius job.
'''