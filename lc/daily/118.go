package daily

/*

Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:


Example 1:

Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
Example 2:

Input: numRows = 1
Output: [[1]]


Constraints:

1 <= numRows <= 30

*/
/*
 need to generate pascal triangle, for given numrows.
 n = 1 default is 1.
 1
 1 1
 1 2 1
 1 3 3 1

*/

func generate(numRows int) [][]int {
	ans := [][]int{}
	for i := 1; i <= numRows; i++ {
		ans = append(ans, triangleRow(i))
	}
	return ans
}

func triangleRow(n int) []int {
	num := 1
	res := []int{1}
	for i := 1; i < n; i++ {
		num *= n - i
		num /= i
		res = append(res, num)
	}
	return res
}
