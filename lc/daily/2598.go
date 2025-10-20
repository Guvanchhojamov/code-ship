package daily

/*


2598. Smallest Missing Non-negative Integer After Operations
You are given a 0-indexed integer array nums and an integer value.

In one operation, you can add or subtract value from any element of nums.

For example, if nums = [1,2,3] and value = 2, you can choose to subtract value from nums[0] to make nums = [-1,2,3].

The MEX (minimum excluded) of an array is the smallest missing non-negative integer in it.

For example, the MEX of [-1,2,3] is 0 while the MEX of [1,0,3] is 2.
Return the maximum MEX of nums after applying the mentioned operation any number of times.


Example 1:
Input: nums = [1,-10,7,13,6,8], value = 5
Output: 4
Explanation: One can achieve this result by applying the following operations:
- Add value to nums[1] twice to make nums = [1,0,7,13,6,8]
- Subtract value from nums[2] once to make nums = [1,0,2,13,6,8]
- Subtract value from nums[3] twice to make nums = [1,0,2,3,6,8]
The MEX of nums is 4. It can be shown that 4 is the maximum MEX we can achieve.

Example 2:
Input: nums = [1,-10,7,13,6,8], value = 7
			   1,4,0,13, 6, 8
Output: 2
Explanation: One can achieve this result by applying the following operation:
- subtract value from nums[2] once to make nums = [1,-10,0,13,6,8]
The MEX of nums is 2. It can be shown that 2 is the maximum MEX we can achieve.


Constraints:

1 <= nums.length, value <= 10^5
-10^9 <= nums[i] <= 10^9

What we need?
The MEX (minimum excluded) of an array is the smallest missing non-negative integer in it.
  the samllest missing non negative integer in it.
  [-1,2,3] - for this mex is 0.
  [1,0,3] - mex is 2
  [1,2,0] - mex is 3.
  the smallest missing non-negative number.
- we can add or subtract value to any number on each operation, so
- we can do operation any times.
assume val = x.
what we need to do to get minimum mex number?
 - since we need non-negative we may start from 0->10^5 and find the first missing number,
 - for example:
	0, search 0 in nums - ok.
	1, seach 1 in nums - ok.
	2, seach - not ok.  This is Mex.
 it will take n*10^5 time after each operation.
 - ok now how we can know about which nums[i] we need to add or subt value x ?

 our goal is to maximize the mex.
	what we need to do to maximize it?
	we need to find nex one-by one, first try find 0, then 1, then 2 ... etc.
	need to try to make numbers, in asc order, 0,1,2,3 ... adn after each operation compare mex with max_mex.
	if no one can perform .. 4,  for exampe then break loop. our max_mex is 3. we cannot get other
 take mex, iterate over array:

*/
/*
Think about using modular arithmetic.
if x = nums[i] (mod value), then we can make nums[i] equal to x after some number of operations
How does finding the frequency of (nums[i] mod value) help?
*/
/*Input: nums = [1,-10,7,13,6,8], value = 7
			     1,4,0,13, 6, 8
 ok after hint usage we realize we need to create 0,1,2...
 how we can do this and keep uniqueness.
	and we need to mod to nums[i] to value to take modular
	thinking if if num[i] is modular to val, then after some operation we got nums[i] mod val.
		subtract if nums[i]>0 and add if nums[i]<0, UNTIL nums[i] becomes, 0<=nums[i]<=value itself.
	to keep track of mod vals we use hash table: key-nums[i] mod val; value - bool .
	if mod already in map, don't do anything. just skip. because we already has 0,1,2... we need 3, not 1 or 2 or 0.

Flow:
	create map, key-int; val:boolean.
	iterate over nums:
	 if nums[i]<0:
	 	while nums[i] > 0 and < x:
			nums[i]+=x.
	 if nums[i] > x:
		while z > 0 and z < x:  -- since we can to operation unlimited time.
			z = nums[i] mod x;
			if z in map: break.
		add to map.
- after iteration,
sort the array, adn go with,0->len(nums)
if nums[i] != i then: return i-1. -> this is max possible mex. after all iteration.
return len(nums) - it means the nums already in asc order and val is 1.

tc: (n*logx)+nlogn+n.
sc: On. for hash map.
we can avoid, sorting, compare with max on each adding map.
*/
