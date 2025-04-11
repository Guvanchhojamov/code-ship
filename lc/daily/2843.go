package daily

/*
You are given two positive integers low and high.

An integer x consisting of 2 * n digits is symmetric

If the sum of the first n digits of x is equal to the sum of the last n digits of x.

Numbers with an odd number of digits are never symmetric.

Return the number of symmetric integers in the range [low, high].

Example 1:
Input: low = 1, high = 100
Output: 9
Explanation: There are 9 symmetric integers between 1 and 100: 11, 22, 33, 44, 55, 66, 77, 88, and 99.

Example 2:
Input: low = 1200, high = 1230
Output: 4
Explanation: There are 4 symmetric integers between 1200 and 1230: 1203, 1212, 1221, and 1230.

Constraints:
1 <= low <= high <= 10^4
*/

/*
 We need to do in there binary search on answers.
1. make number to arr
2. if len(digits)/2 != 0 skip
3. if digits len is even check symmetric or not starting from back and start.
4. if symmetric add increment result.
  tc: n+n/2 = n
  sc: n.

*/

func countSymmetricIntegersOurSolution(low int, high int) int {
	var ans = 0
	for low <= high {
		digits := makeNumsArray(low)
		if len(digits)/2 != 0 {
			low += 1
			continue
		}
		if isSymmetric(digits) {
			ans++
		}
		low += 1
	}
	return ans
}

func makeNumsArray(num int) []int {
	var res = make([]int, 0, 10000)
	for num > 0 {
		res = append(res, num%10)
		num /= 10
	}
	return res
}

func isSymmetric(digits []int) bool {
	var i, j = 0, len(digits) - 1
	leftSum, rightSum := 0, 0
	for i < j {
		leftSum += digits[i]
		rightSum += digits[j]
		i++
		j--
	}
	return leftSum == rightSum
}

/*
 In my solution we got time limit exception.
   so we go to editorial.
*/

/*
Approach 1: Enumeration
Intuition
Enumerate all numbers from low to high:

If it is a two-digit number and is a multiple of 11, then it is a symmetric integer.
If it is a four-digit number, calculate the sum of the thousands and hundreds digits,
as well as the sum of the tens and ones digits.
If they are equal, it is a symmetric (even) integer.
Finally, it returns the number of symmetric integers in the range.
*/

func countSymmetricIntegers(low int, high int) int {
	res := 0
	for a := low; a <= high; a++ {
		if a < 100 && a%11 == 0 {
			res++
		} else if 1000 <= a && a < 10000 {
			left := a/1000 + (a%1000)/100
			right := (a%100)/10 + a%10
			if left == right {
				res++
			}
		}
	}
	return res
}
