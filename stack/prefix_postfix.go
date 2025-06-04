package stack

/**
  Given s prefix make it postfix.
s = /-AB*+DEF
Since we given prefix we need to start from back. to make it postfix.
i=0 st = []string
1. If s[i] == operant add to stack.
2. If s[i] == operator pop alst two and add to the back.

s = /-AB*+DEF
F,E,D
F,DE+
DE+F*,B,A
DE+F*,AB-
AB-DE+F*\ - is result.

**/
/*
 Lest write summary in there:
1. infix to postfix, if operand add to ans, if operator and priority<pop.priority add to the stack elese pop from stack add to ans until >.
	return ans
2. postfix to infix, if operant add to stack, if operator, pop 2 elem from stack add () and +- between push back.
3. infix to prefix. reverse, convert to postfix, reverse again.
4. prefix to infix. Start from back to like postfix to infix.
5. prefix to postfix. Start from back, if operant add stack, if operator pop() 2 from stack add operand end of this, push back to stack.
6. postfix to prefix. Start from begin, if operant add to stack,else pop() 2 from stack, add operator to front, push back again.
Great!.
*/
