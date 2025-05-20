package bit

/*
 Sieve of Eratosthenes.
204. Count Primes
Given an integer n, return the number of prime numbers that are strictly less than n.

Example 1:
Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.

Example 2:
Input: n = 0
Output: 0

Example 3:
Input: n = 1
Output: 0

Constraints:
0 <= n <= 5 * 10^6
*/
/*
  to get all prime numbers under in optimal tc.
- create array with nums 2 to n+1
- start from p=2, mark all nums[i]/2 ==0 as 0, not prime.
- after p++ p=3 mark all nums[i]/3 ==0 as 0, not prime.
- until sqrt(n).
return numbers in array witch is non zero. or count them and return count, in our case.

*/

func countPrimes(n int) int {
	var count = 0
	var nums = make([]bool, n, n)
	if n <= 2 {
		return 0
	}
	for i := 2; i < n; i++ {
		nums[i] = true
	}
	//fmt.Println(nums)

	for i := 0; i*i < len(nums); i++ {
		if nums[i] {
			for j := i + 1; j < n; j = j + i {
				nums[j] = false
			}
		}
	}

	//fmt.Println(nums)
	for i := range nums {
		if nums[i] {
			count++
		}
	}
	return count
}
