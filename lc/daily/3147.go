package daily

/*
3147.

In a mystic dungeon, n magicians are standing in a line.
Each magician has an attribute that gives you energy.
Some magicians can give you negative energy, which means taking energy from you.

You have been cursed in such a way that after absorbing energy from magician i, you will be instantly transported to magician (i + k).
This process will be repeated until you reach the magician where (i + k) does not exist.

In other words, you will choose a starting point and then teleport with k jumps until you reach the end of the magicians' sequence,
absorbing all the energy during the journey.

You are given an array energy and an integer k.
Return the maximum possible energy you can gain.

Note that when you are reach a magician, you must take energy from them, whether it is negative or positive energy.


Example 1:

Input: energy = [5,2,-10,-5,1], k = 3

Output: 3

Explanation: We can gain a total energy of 3 by starting from magician 1 absorbing 2 + 1 = 3.

Example 2:
Input: energy = [-2,-3,-1], k = 2
Output: -1

Explanation: We can gain a total energy of -1 by starting from magician 2.



Constraints:

1 <= energy.length <= 10^5
-1000 <= energy[i] <= 1000
1 <= k <= energy.length - 1

*/
/*
  given n magicans, k - jump steps. Need to gain max possible magican at the end of jump. Usgin all possible jumps.
   - Each magican give energy.
   - Negative magican take enegry.
   - must take magican if we jump to this.
 What is max possible jump position?


Bf:
	do it in 2 loops:
	- i:=0;i<len(n);i++
	  j:=i;j<len(nums);j+k.
	  sum+=arr[j]
	mx = max(mx,sum)
	sum=0
	return mx.
tc: N*(N/K) = N^2/K.
sc: 1

how can we optimize?

Cant solve need to start from back...

*/
