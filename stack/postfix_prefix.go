package stack

/**

Given s postfix make it prefix.
 Since it is the postfix, start from begin.
 i=0 st=[]string  s=AB-DE+F\*

1. start from begin:

if s[i] == operant: add to stack.
if s[i] == operator:
		pop last 2 elements from stack, make operant+pop(1)+pop(2) push back to stack.

s=AB-DE+F*\
AB
-AB,DE
-AB,+DE,F
-AB, *+DEF,
\-AB*+DEF - this is result.
**/
