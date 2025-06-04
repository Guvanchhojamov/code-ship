class Solution:
    def triangleType(self, nums: List[int]) -> str:
        a,b,c = nums[0],nums[1], nums[2] 
        if a+b<=c or a+c<=b or b+c<=a:
            return "none"
        if a==b and a==c and b==c:
            return "equilateral"
        if a!=b and a!=c and b!=c:
            return  "scalene"
        if (a==b and a!=c) or (a==c and b!=c) or (b==c and a!=b):
            return "isosceles"
        return "none"



#   Given nums we need to check:
# - can nums may be triangle sides?
# - what is the type of triangle? "equilateral", "isosceles", "scalene"
# len(nums) == 3 always.
# to make easy handle the sides, export them to variables.
# a,b,c = a[0],a[1],a[2]
# check can we triangle sides?
# 	if a+b<=c || a+c<=b || b+c <= a:
# 		return "none"
# 	if a==b && a==c && b==c:
# 		return "equal"
# 	if a!=b && a!=c && b!=c:
# 		return "isolated"