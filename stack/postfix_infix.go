package stack

/**
 Given postfix string make it prefix.
postfix:AB-DE+F*\
infix:
Solution:
 i=0; st = []string
iterate over s and:
if s[i] == operant:
	st.Add(s[i])
if s[i] == operator:
  Take top 2 elements from stack, add operator between them, add parentheses on begin and end, push to stack.
AB-DE+F*\
A,B
(A+B),DE
(A+B),(D+E),F
(A+B),((D+E)*F)
st= [ ((A+B)/((D+E)*F))] - st[0] is result.
**/
