package string

// 1903. Largest Odd Number in String
/*
1903. Largest Odd Number in String
You are given a string `num`, representing a large integer.
Return the largest-valued odd integer (as a string) that is a non-empty substring of num,
or an empty string "" if no odd integer exists.
A substring is a contiguous sequence of characters within a string.

Example 1:

Input: num = "52"
Output: "5"
Explanation: The only non-empty substrings are "5", "2", and "52". "5" is the only odd number.
Example 2:

Input: num = "4206"
Output: ""
Explanation: There are no odd numbers in "4206".
Example 3:

Input: num = "35427"
Output: "35427"
Explanation: "35427" is already an odd number.


Constraints:

1 <= num.length <= 10^5
num only consists of digits and does not contain any leading zeros.
*/

/*
   Need to return subscring of nums, with gives us maximum odd number.
      1. Brute force.
        convert all string to int.
        since we need substring, how can we seperate number to  subsctrings?
            with moding to 10 until 0:
        1. Chech is number is odd
        2. Compare with max.
        3. divide num to 10
        35427 %2 !=0 -> max=35427
        3542 %2 ==0 skip
        354 %2 == 0 skip
        35 %2 != 0
            35>35427 not
        3 %2 !=0
            3 > 35427 not
        0  end of the interation
        return maxNum= 35427
        tc: n+logn
        sc: 1
*/
/*
   optimal: Start from back
   if last char is odd retrun
   s[0:i+1] this is our max  subsctring
*/

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if (num[i]-'0')%2 == 1 {
			return num[:i+1]
		}
	}
	return ""
}
