package stack

/*
735. Asteroid Collision
We are given an array asteroids of integers representing asteroids in a row.
The indices of the asteriod in the array represent their relative position in space.

For each asteroid,the absolute value represents its size,
and the sign represents its direction
(positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions.
If two asteroids meet, the smaller one will explode.
If both are the same size, both will explode.
Two asteroids moving in the same direction will never meet.

Example 1:

Input: asteroids = [5,10,-5]
Output: [5,10]
Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.

Example 2:
Input: asteroids = [8,-8]
Output: []
Explanation: The 8 and -8 collide exploding each other.

Example 3:
Input: asteroids = [10,2,-5]
Output: [10]
Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.

Constraints:

2 <= asteroids.length <= 10^4
-1000 <= asteroids[i] <= 1000
asteroids[i] != 0
What we need?
- The num[i] - direction and weight of the asteroid.
- Asteroid with height weight will explode smaller one.
- asters with same size explode each other.
- same directions never meet.
nums = [10,2,-5,8,-8,10] =>

Brute force:
st = 10
i = 0-> n
if (st.top > 0 && a[i]<0) || (st.top < 0 && a[i]>0):

	while st.len > 0 && st.top <= abs(a[i]):
		st.top == abs(a[i]) : i++
		st.pop

return st.
tc: O(n)
sc: O(n) - stack.
*/
func asteroidCollision(nums []int) []int {
	var st = make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			st = append(st, nums[i])
		} else {
			for len(st) > 0 && st[len(st)-1] > 0 && st[len(st)-1] < abs(nums[i]) {
				st = st[:len(st)-1]
			}
			if len(st) > 0 && st[len(st)-1] > 0 && st[len(st)-1] == abs(nums[i]) {
				st = st[:len(st)-1]
			} else if len(st) == 0 || st[len(st)-1] < 0 {
				st = append(st, nums[i])
			}
		}
	}
	return st
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
