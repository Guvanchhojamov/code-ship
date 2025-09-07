package daily

/*
1304. Find N Unique Integers Sum up to Zero

Given an integer n,
 return any array containing n unique integers such that they add up to 0.



Example 1:

Input: n = 5
Output: [-7,-1,1,3,4]
Explanation: These arrays also are accepted [-5,-1,1,2,3] , [-3,-1,2,-2,4].

Example 2:
Input: n = 3
Output: [-1,0,1]
Example 3:

Input: n = 1
Output: [0]

Constraints:

1 <= n <= 1000

*/
/*
 Need to create array with length n. and should be sum of all elemments must be 0.
  - elements must be unique.
for example:
	- n =3
	-1 0 1 - example. or reversed version.
*/

/*
 i start from
 add elements to arr.
 if n is odd:
    start with 0.
 if n is even start with 1:
 until len(arr) < n:
    add to arr i and. -i if i > 0.
 return arr .

*/

func sumZero(n int) []int {
	var i = 0
	var arr = make([]int, 0, n)
	if n%2 == 0 {
		i = 1
	}
	for i < n {
		arr = append(arr, i)
		if i > 0 {
			arr = append(arr, i*(-1))
		}
		i++
	}
	return arr
}

func sumZero2(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n/2; i++ {
		arr[i*2] = i + 1      // positive numbers: 1, 2, 3...
		arr[i*2+1] = -(i + 1) // negative numbers: -1, -2, -3...
	}
	return arr
}
