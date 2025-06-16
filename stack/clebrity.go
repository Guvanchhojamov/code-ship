package stack

/*
Given a square matrix mat[][] of size n x n,
such that mat[i][j] = 1 means ith person knows jth person,
the task is to find the celebrity.

A celebrity is a person who is known to all but does not know anyone.
Return the index of the celebrity, if there is no celebrity return -1.

Note: Follow 0-based indexing and mat[i][i] will always be 1.

Examples:

Input: mat[][] = [[1, 1, 0],
 				  [0, 1, 0],
				  [0, 1, 1]
				 ]
Output: 1
Explanation: 0th and 2nd person both know 1. Therefore, 1 is the celebrity.

Input: mat[][] = [[1, 1],
				  [1, 1]]
Output: -1
Explanation: The two people at the party both know each other. None of them is a celebrity.

Input: mat[][] = [[1]]
Output: 0
*/
/*
[1, 1, 0],
[0, 1, 0],
[0, 1, 1]

What we need? n*n shows we have n persons. need return celebrity person.
the result is 1 because
0-th knows 1-th
assume person nums top to bottom.
if a[i][j]==1 ith person knows jth person.
if a[i][j] ==0 ith person don't know jth person.
0=0;0=1;0=0
1=0;1=1;1=0
2=0;2=1;2=1
in there, 0 knows 1;
1 knows 1 only itself.
2 knows 1,2
-we can check take each pair of person(i,j) and check do they know each other?
We need to create 2 arrays with len n.
KnowMe[]
IKnow[]
increase if ith know jth.
iterate over n:
And were knowMe[i]==n-1 && iKnow[i] == 0 then this is our person.

[1, 1, 0]
[0, 1, 0]
[0, 1, 1]   i!=j
iKnow= [1,0,1] = i
knowMe=[0,2,0] = j
celebrity is - every on knows hum, he dont now any one.
0-1
1-0 knows anyone. the 2 person knows 1
2-1
*/
/*
iknow = [n]array
knowme = [n]array
 for i=0 to n :
  for j=0 to n:
  if i!=j:
	if a[i][j] == 1:
		iknow[i]++
		knowme[j]++

for i=0 to n:
	if iknow[i]==0 && knowme[i]==n-1:
	return i - this is our celebrity person.
else return -1
tc: n*n+n
sc: 2n= n
*/
/*
 The other optimal approach is two pointer approach
 what is the max num person can be a celebroty =1
 the min =0
 assume each row is 1 person.
 start from top and bottom.
 i=0; j=n-1
 if i know j, then i cannot be celebrity. inc i++ go to next person.
 if j know i , then j cant be celebrity so dec j-- go to next person.
 do this until i==j
 if i==j then check where i!=j and a[i][j] == 0 j++.
 and check where i!= and a[i][j]==1 i++
 if ok return i or j, it is our ans else return -1 we dont have celebrity.
*/
