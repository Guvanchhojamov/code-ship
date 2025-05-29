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

nums = 1 7 3 8 5
result = 7,8,8,-1,7
We store elemens in stack increasing order:

1738517385
0123456789

st:8,7,1
5 -1
8 -1
3: 8
7: 8
1: 7
//
st:8,7
5:7
8:-1
3:8
7:8
1:7
output: 7,8,8,-1,7
How to solve?
we 2x the array, so until i==n we just keep stack, using:

	st = nums[last]
	res = -1
	el:=-1

	el = st.top()
	if nums[i%n-1] > stack.top:
		while nums[i%n-1] > stack.top || stack.empty:
			st.pop
		if st.empty:
			el = -1
		else:
		el = st.Top()
		st.add(nums[i%n-1])
	if i <= n-1:
		res.add(el)
	return res.
*/
func nextGreaterElements(nums []int) []int {
	var res = make([]int, 0, len(nums))
	var st = make([]int, 0, len(nums))
	var n = len(nums)

	for i := 2*n - 1; i >= 0; i-- {
		for len(st) > 0 && st[len(st)-1] <= nums[i%n] {
			st = st[:len(st)-1]
		}
		if i < n {
			if len(st) == 0 {
				res = append(res, -1)
			} else {
				res = append(res, st[len(st)-1])
			}
		}
		st = append(st, nums[i%n])
	}
	return reverse(res)
}

func reverse(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
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
Since the nums
*/
/*
/*
 1,2,4,3

 next greater decreasing order
 next smaller is increasign order montohonic stack
 -1, 1, 3, 1

 for previous smaller, or greater element we need to start from beginning.
	and use monotonic stack algorythm.

	// To count the number of next greater elements on the ryght we need to count just, len(st)
	The remaining len stack would be next greater elements for our number.
	if no remaingn elements aka stack is empty we can count it like 0.

*/
