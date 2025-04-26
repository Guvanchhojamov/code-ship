package daily

/*
2145. Count the Hidden Sequences
   You are given a 0-indexed array of n integers differences,
   which describes the differences between each pair of consecutive integers of a hidden sequence of length (n + 1).
   More formally, call the hidden sequence hidden, then we have that differences[i] = hidden[i + 1] - hidden[i].

You are further given two integers lower and upper that
describe the inclusive range of values [lower, upper] that the hidden sequence can contain.

For example, given differences = [1, -3, 4], lower = 1, upper = 6,
the hidden sequence is a sequence of length 4 whose elements are in between 1 and 6 (inclusive).
[3, 4, 1, 5] and [4, 5, 2, 6] are possible hidden sequences.
[5, 6, 3, 7] is not possible since it contains an element greater than 6.
[1, 2, 3, 4] is not possible since the differences are not correct.
Return the number of possible hidden sequences there are.
If there are no possible sequences, return 0.


Example 1:

Input: differences = [1,-3,4], lower = 1, upper = 6
Output: 2
Explanation: The possible hidden sequences are:
- [3, 4, 1, 5]
- [4, 5, 2, 6]
Thus, we return 2.

Example 2:
Input: differences = [3,-4,5,1,-2], lower = -4, upper = 5
Output: 4
Explanation: The possible hidden sequences are:
- [-3, 0, -4, 1, 2, 0]
- [-2, 1, -3, 2, 3, 1]
- [-1, 2, -2, 3, 4, 2]
- [0, 3, -1, 4, 5, 3]
Thus, we return 4.

Example 3:
Input: differences = [4,-7,2], lower = 3, upper = 6
Output: 0
Explanation: There are no possible hidden sequences. Thus, we return 0.


Constraints:

n == differences.length
1 <= n <= 10^5
-10^5 <= differences[i] <= 10^5
-10^5 <= lower <= upper <= 10^5
*/
/*
   we have differences arr witch contains n elements.
	need to find some hidden sequence array,
	dif[i] = h[i+1]-h[i]
   with this formula, the hidden sequence's length always need to be n+1 for valid.
   dif[i] = h[i+1]-h[i]
   dif[i]+h[i] = h[i+1]

*/

// func numberOfArrays(differences []int, lower int, upper int) int {
//     start, mi, mx := math.MaxInt, 0, 0

//     for _,v:=range differences{
//         start = min(start,v)
//     }

//     var h = make([]int,0,len(differences)+1)
//     var res = 0
//     for start <= upper {
//          mi, mx = math.MaxInt, math.MinInt
//         h = []int{}
//         h = append(h,start)
//         for i:=0; i<len(differences);i++{
//             h = append(h, h[i]+differences[i])
//             mi = min(mi, h[len(h)-1])
//             mx = max(mx, h[len(h)-1])
//         }
//        // fmt.Println(h,mi,mx)
//         if mi>=lower && mx<=upper {
//             res++
//         }
//         start++
//     }
//     return res
// }

// solution showed
func numberOfArrays(differences []int, lower int, upper int) int {
	x, mi, mx := 0, 0, 0
	for _, d := range differences {
		x += d
		mi = min(mi, x)
		mx = max(mx, x)
	}
	return max(0, upper-lower-(mx-mi)+1)
}
