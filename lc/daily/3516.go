package daily

/*

3516. Find Closest Person
Easy
Topics
premium lock icon
Companies
Hint
You are given three integers x, y, and z, representing the positions of three people on a number line:

x is the position of Person 1.
y is the position of Person 2.
z is the position of Person 3, who does not move.
Both Person 1 and Person 2 move toward Person 3 at the same speed.

Determine which person reaches Person 3 first:

Return 1 if Person 1 arrives first.
Return 2 if Person 2 arrives first.
Return 0 if both arrive at the same time.
Return the result accordingly.



Example 1:

Input: x = 2, y = 7, z = 4

Output: 1

Explanation:

Person 1 is at position 2 and can reach Person 3 (at position 4) in 2 steps.
Person 2 is at position 7 and can reach Person 3 in 3 steps.
Since Person 1 reaches Person 3 first, the output is 1.

Example 2:

Input: x = 2, y = 5, z = 6

Output: 2

Explanation:

Person 1 is at position 2 and can reach Person 3 (at position 6) in 4 steps.
Person 2 is at position 5 and can reach Person 3 in 1 step.
Since Person 2 reaches Person 3 first, the output is 2.

Example 3:

Input: x = 1, y = 5, z = 3

Output: 0

Explanation:

Person 1 is at position 1 and can reach Person 3 (at position 3) in 2 steps.
Person 2 is at position 5 and can reach Person 3 in 2 steps.
Since both Person 1 and Person 2 reach Person 3 at the same time, the output is 0.



Constraints:

1 <= x, y, z <= 100
*/
/*
   have 3 persons.
   p1,p2,p3.
   p1,p2 are movind to p3 with same speed.
   need to find who firt arrive to person3? if both return 0.
- if p1,p2 are miving with same speed.
 - person witch are smaller range arrive first. example:
    p1=3, p2=7, p3 =4.
    4-3 = 1
    4-7 = 3
    this shows us p1 arrive first.

    2 7 4
    d1 = 4-2 =2
    d2 = 4-7 = 3
*/

func findClosest(x int, y int, z int) int {
	d1 := abs(z - x)
	d2 := abs(z - y)
	if d1 > d2 {
		return 2
	} else if d1 < d2 {
		return 1
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return a * (-1)
	}
	return a
}
