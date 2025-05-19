package daily

/*
3024. Type of Triangle
You are given a 0-indexed integer array nums of size 3 which can form the sides of a triangle.

A triangle is called equilateral if it has all sides of equal length.
A triangle is called isosceles if it has exactly two sides of equal length.
A triangle is called scalene if all its sides are of different lengths.
Return a string representing the type of triangle that can be formed or "none" if it cannot form a triangle.

Example 1:

Input: nums = [3,3,3]
Output: "equilateral"
Explanation: Since all the sides are of equal length, therefore, it will form an equilateral triangle.

Example 2:
Input: nums = [3,4,5]
Output: "scalene"
Explanation:
nums[0] + nums[1] = 3 + 4 = 7, which is greater than nums[2] = 5.
nums[0] + nums[2] = 3 + 5 = 8, which is greater than nums[1] = 4.
nums[1] + nums[2] = 4 + 5 = 9, which is greater than nums[0] = 3.
Since the sum of the two sides is greater than the third side for all three cases, therefore, it can form a triangle.
As all the sides are of different lengths, it will form a scalene triangle.


Constraints:

nums.length == 3
1 <= nums[i] <= 100
*/
/*
  Given nums we need to check:
- can nums may be triangle sides?
- what is the type of triangle? "equilateral", "isosceles", "scalene"
len(nums) == 3 always.
to make easy handle the sides, export them to variables.
a,b,c = a[0],a[1],a[2]
check can we triangle sides?
	if a+b<=c || a+c<=b || b+c <= a:
		return "none"
	if a==b && a==c && b==c:
		return "equal"
	if a!=b && a!=c && b!=c:
		return "isolated"
*/
// def triangleType(self, nums: List[int]) -> str:
//         a,b,c = nums[0],nums[1], nums[2]
//         if a+b<=c or a+c<=b or b+c<=a:
//             return "none"
//         if a==b and a==c and b==c:
//             return "equilateral"
//         if a!=b and a!=c and b!=c:
//             return  "scalene"
//         if (a==b and a!=c) or (a==c and b!=c) or (b==c and a!=b):
//             return "isosceles"
//         return "none"
// solved in python
