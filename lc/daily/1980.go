package daily

import "fmt"

/*
1980. Find Unique Binary String

/*
Given an array of strings 'nums' containing 'n' unique binary strings each of length 'n',
return a binary string of length 'n' that does not appear in 'nums'.
If there are multiple answers, you may return any of them.



Example 1:
Input: nums = ["01","10"]
Output: "11"
Explanation: "11" does not appear in nums. "00" would also be correct.

Example 2:
Input: nums = ["00","01"]
Output: "11"
Explanation: "11" does not appear in nums. "10" would also be correct.
Example 3:

Input: nums = ["111","011","001"]
Output: "101"
Explanation: "101" does not appear in nums. "000", "010", "100", and "110" would also be correct.


Constraints:

n == nums.length
1 <= n <= 16
nums[i].length == n
nums[i] is either '0' or '1'.
All the strings of nums are unique.
*/

/*
   Given nums binary string unique nums, with length=n return binary string with len n with not in nums, any occurence.
     Key points:
      - n = length(nums)
      - nums[i] = n
      - nums[i] is unique.

    The min numbers len is 1 so, the min is [0] or [1]
     if 0 we return 1 else we return 0
     what is the min?
     0 is min
     what is the max?
     Since the n length is 16 then, the max would be 1..1 -> 16 digits, it is `65535` in decimal.
     Algorythm:
        1. Create map with nums key val: since nums are unique
            key = nums[i]
            val = true
         Add fill map with nums.
        2. Create helper function convertToBinary(int num) string:
            witch returns converted to binary result.
        3. Iterate from i=0 to max value = 655335:
                binStr =  convertToBinary(num)
            if  binsMap[binStr] is not exits return binStr
            since we need first occuerence.
        tc: n-fill map, * 0-to max=655335 * convert lon(n) = 655335*n(logN)
        sc: O(n)
        lets try  to solve.
*/

func findDifferentBinaryString(nums []string) string {
	var binsMap = make(map[string]bool, 0)
	var maxDecimal = 655335

	for _, val := range nums {
		binsMap[val] = true
	}

	for i := 0; i <= maxDecimal+1; i++ {
		binStr := fmt.Sprintf("%0*b", len(nums), i)
		if val, exists := binsMap[binStr]; !val && !exists && len(binStr) == len(nums) {
			return binStr
		}
	}
	return ""
}

// nums = ["111","011","001"]
/*
["10","11"]

*/

// func convertToBinaryStr(num int) string{
//     var res []string

//     for n>0 {
//         res = append(res,fmt.Sprintf("%d",n%2))
//         n=n/2
//     }
//     return strings.Join(res,"")
// }
