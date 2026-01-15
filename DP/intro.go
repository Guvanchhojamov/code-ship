package dp

/*
 What is Dynamic programming?
  We can use dp whaere we need dynamically changin values..
  Where subproblems are metters.
  Where we can divide problem to small subproblems, and use subproblem results.

DP usually comes after recursion approach.
If problem goes to recursion, we can us somehow DP on this problem.
 DP is divided some parts:
 - DP in 1D array.
 - DP in 2D array.
 - DP in strings
 etc..
 Where we can store prev results in 1D array by index.
 Where we can store prev result in 2d array by index.
 Where we can store parts of string in the array and use them for other results.

In DP, we need to see some points:
 Overlapping subproblems we can use again agin.
 Optimal subsctructure, The optimal solution which can be done with using their
 optimal solution for their subproblems..

There dp trick defined with 3 steps:
1. Try to represent problem in terms of index..
2. Do all possible staff with this indexes.. according to the problem statement..
3. Got last needed result from them, sum(), count(), min(), max() or other staffs.

If you got is it DP problem, after this is pattern..
1. Define state, start state...
	dp[i] = the max profit in i days or etc..
2. Define state transition..
	dp[i] = max(dp[i-1], dp[i-2]) ...
3. Define base case..
	its like in recursion we has base case..
dp[0] =1; dp[1]=2 - for example..

Since to DP we usually we come after recursion..
 We have 2 types of appraoch, to solve dp problem.
 1. Memoization:
	- Bottom-Up approach:
		In this case we go extreme down, and come back using backtracking.
		It means from f(N) -> to F(0)..
		in this case we go until 0 - base case first, then come back until N.
		Store results in cache as you go or come back..
 2. Tabulation:
	- Top-bottom apprach:
		Start with 0 and use for until N..
		store results in table, in 1D or 2D, and use them with loop index..
		ususally its memory efficient because we dont use recursion..
Common problem style is:
subsequences, change something, max/min sum/count after some operation, palindromes,
longes something, subsequences etc.. Pick not pick problems..
How many ways etc..
 Other notes we will take when we solve problems..

*/
