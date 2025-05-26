package stack

/*
503. Next Greater Element II
Given a circular integer array nums (i.e., the next element of nums[nums.length - 1] is nums[0]), return the next greater number for every element in nums.

The next greater number of a number x is the first greater number to its traversing-order next in the array, which means you could search circularly to find its next greater number. If it doesn't exist, return -1 for this number.



Example 1:

Input: nums = [1,2,1]
Output: [2,-1,2]
Explanation: The first 1's next greater number is 2;
The number 2 can't find next greater number.
The second 1's next greater number needs to search circularly, which is also 2.
Example 2:

Input: nums = [1,2,3,4,3]
Output: [2,3,4,-1,4]


Constraints:

1 <= nums.length <= 10^4
-10^9 <= nums[i] <= 10^9
*/

func nextGreaterElements(nums []int) []int {
	
	return []int{}
}

/*
 The nums are circular.
 Brute force:
    for each num[i] check rigt i+1->n
    check left 0->i-1
    if find greater add it to res
    if no greater add -1 to resut.
tc: O(n*n) each time we check whole array.
sc: O(1) we dont use any extra space.
What is the optimal?
Since the nums are cyrcular 
what we can do for cyrcular nums.  We can add to the end sam nums. 
1,2,3,4,3  1,2,3,4,3 
0 1 2 3 4  5 6 7 8 9 
Now we can check from i to i+n. to the right. 
But tc leaves same not changes.
how we can take i correct?   
i % n 
9 % 5 = 4 
8 % 5 = 3 
Can we use stack? 
	start from back. 
res:-1,  
st:3
for i =2*N to i%n-1 ==0:
	if nums[i%n-1]>st.top:
		st.add
if nums[i]<st.top():
	res = add(nums[i%n-1])
	st = add (nums[i%n-1])
else if len(st) == 0:
	res = add(-1)
	st = add(nums[i%n-1])
else
	while len(st)>0 && nums[i%n-1] > st.top():
		st.pop()
	if len(st) ==0:
		res = -1 
		st.add(nums[i%n-1])
	else:
		res = st.top
		st = add(nums[i%n-1])
return res.
1,2,3,4,3  1,2,3,4,3 
0 1 2 3 4  5 6 7 8 9 
res:4,-1,4,3,1    
st:4,3,2,1 

1,2,1  1 2 1 
0 1 2  3 4 5 
st = 2
res = 2,-1, 2

1 7 3 8 5  1 7 3 8 5  
st: 8,7,1
res:8,-1,8,8,7 



*/
