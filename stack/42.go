package stack

func trap(height []int) int {
	prefMax := getPrefixMax(height)
	sufMax := getSuffixMax(height)
	total := 0
	for i := range height {
		total += min(prefMax[i], sufMax[i]) - height[i]
	}
	return total
}

func getPrefixMax(nums []int) []int {
	prefixMax := make([]int, len(nums), len(nums))
	prefixMax[0] = nums[0]
	mx := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > mx {
			prefixMax[i] = nums[i]
			mx = nums[i]
		} else {
			prefixMax[i] = mx
		}
	}
	return prefixMax
}

func getSuffixMax(nums []int) []int {
	var suffixMax = make([]int, len(nums), len(nums))
	suffixMax[len(nums)-1] = nums[len(nums)-1]
	mx := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > mx {
			suffixMax[i] = nums[i]
			mx = nums[i]
		} else {
			suffixMax[i] = mx
		}
	}
	return suffixMax
}

/*

Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
 count of filled  cells with water after rain.
 well, cell to be filled.
   the height of cell must be cell-1 > cell < cell+1
We need to count how much water can take each cell,
So how can we do this?
for each element we need to go leftMax, and rightMax.
Take minimum from this min(leftMax,rightMax);
Then we need caclculate height, to do this we need decrease from min(l,r) - nums[i] current heigh.
total = 0
trip = min(leftMax, rightMax)[height] - nums[i] (heigh)
1. How we find leftMax and rightMax?
    To find it we use prefixMax,and suffixMax arrays.
for each element we find leftMax and store it in prefixMax[i]-index. same for suffixMax[i]
iterate 2 times, nums: leftMax, rightMax.
iterate nums for calculate total. with this maxes.
tc: n+n+n = 3n.
sc: n+n = 2N.
*/

/*
To optimize memory we can use another approach.
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]

Taking this example, if we start with left we never see rightmax.
if we start from back we never see the leftMax.
But, we can start from begin and end, and keep leftMax and rightMax.

l,r = for iterate from left, and right.
lmax - keep leftmax element.
rmax - keep rightmax element.
and in there tow pointer approach comes in.
we need always small heigh so,
while l!=r:
  lmax = max(lmax, nums[l])
  rmax = max(rmax, nums[r])
  if l<r:
	val = min(lmax,rmax) - nums[l]
  else:
	val = min(lmax,rmax) - nums[r]
  total+=val
  return total.
*/
