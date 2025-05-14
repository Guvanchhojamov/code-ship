package daily

/*
There are n dominoes in a line,
and we place each domino vertically upright.
In the beginning, we simultaneously push some of the dominoes either to the left or to the right.

After each second, each domino that is falling to the left pushes the adjacent domino on the left.
Similarly, the dominoes falling to the right push their adjacent dominoes standing on the right.

When a vertical domino has dominoes falling on it from both sides, it stays still due to the balance of the forces.

For the purposes of this question,
we will consider that a falling domino expends
no additional force to a falling or already fallen domino.

You are given a string dominoes representing the initial state where:

dominoes[i] = 'L', if the ith domino has been pushed to the left,
dominoes[i] = 'R', if the ith domino has been pushed to the right, and
dominoes[i] = '.', if the ith domino has not been pushed.
Return a string representing the final state.

Example 1:
Input: dominoes = "RR.L"
Output: "RR.L"
Explanation: The first domino expends no additional force on the second domino.

Example 2:
Input: dominoes = ".L.R...LR..L.."
Output: "LL.RR.LLRRLL.."

Constraints:
n == dominoes.length
1 <= n <= 10^5
dominoes[i] is either 'L', 'R', or '.'.
*/
/*
  L - pushes left just 1 adjacent or fallen to left .
  R - pushes right just 1 adjacent or fallen to right .
  . - not fallen domino
 - any push not affect to already fallen domino.
 - generate string with new state.

What is the simple brute force solution in there?
   if len.s ==1 just return this string.
   We need 3 constraints:
    take a[i]
	if a[i] == '.':
		a[i-1] == 'L' or 'R' and
		a[i+1] == 'L' or 'R'
    just continue this is the balance.
  if
*/
/*
 ".L.R...LR..L.."
LL.R..LLR.LL..
LL.RR.LLR.LL..
Output: "LL.RR.LLRRLL.."
1, 9
L
*/

/*
Approach #1: Adjacent Symbols [Accepted]
Between every group of vertical dominoes ('.'),
we have up to two non-vertical dominoes bordering this group.
Since additional dominoes outside this group do not affect the outcome,
we can analyze these situations individually:
there are 9 of them (as the border could be empty).
Actually, if we border the dominoes by 'L' and 'R',
there are only 4 cases. We'll write new letters between these symbols depending on each case.

Algorithm

Continuing our explanation, we analyze cases:

If we have say "A....B", where A = B, then we should write "AAAAAA".

If we have "R....L", then we will write "RRRLLL", or "RRR.LLL"
if we have an odd number of dots.
If the initial symbols are at positions i and j,
we can check our distance k-i and j-k to decide at position k whether to write 'L', 'R', or '.'.

(If we have "L....R" we don't do anything. We can skip this case.)

Complexity Analysis

Time and Space Complexity: O(N), where N is the length of dominoes.
*/
/*
 Approach #2: Calculate Force [Accepted]
We can calculate the net force applied on every domino.
The forces we care about are how close a domino is to a leftward 'R',
and to a rightward 'L': the closer we are, the stronger the force.

Algorithm

Scanning from left to right,
our force decays by 1 every iteration, and resets to N if we meet an 'R', so that force[i] is higher (than force[j]) if and only if dominoes[i] is closer (looking leftward) to 'R' (than dominoes[j]).

Similarly, scanning from right to left, we can find the force going rightward (closeness to 'L').

For some domino answer[i], if the forces are equal, then the answer is '.'. Otherwise, the answer is implied by whichever force is stronger.

Example

Here is a worked example on the string S = 'R.R...L': We find the force going from left to right is [7, 6, 7, 6, 5, 4, 0]. The force going from right to left is [0, 0, 0, -4, -5, -6, -7]. Combining them (taking their vector addition), the combined force is [7, 6, 7, 2, 0, -2, -7], for a final answer of RRRR.LL.
*/

// I cant understand and cant solve this.
