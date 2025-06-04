package stack

/*
 Givens string infix, make them to prefix.
	infix: (A+B)*C-D+F
	prefix:

We write just flow not implement this, because it takes more time. And this related answers not given in real interviews.
Focus on flow only.
1. Reverse the string.
2. Infix to postfix.
3. Reverse answer.
 Reverse:
1.F+D-C*(B+A)
Make postfix:
st = +-*(+
	 remaining: +-*
ans = FDCBA+
ans = FDCBA+*-+
3. +-*+ABCDF - result. (reverse)
If you want try implement.
*/
