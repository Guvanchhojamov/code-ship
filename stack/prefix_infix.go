package stack

/**
 Given prefix string make it infix.
 s = *+PQ-MN
take i=0 and stack.
Since it is prefix start from back and do steps like postfix->infix.
1. if s[i] == operant add to st
2. if s[i] == operator:
	pop() last 2 from stack add () begin and end add between them operant and push back to stack.
s = *+PQ-MN
st:
NM
(N-M),QP
(N-M),(Q+P),
st = (N-M)*(Q+P). reversed stack[0] is result.
res = (P+Q)*(M-N)
if you want implement it yourself. I don't.)

**/
