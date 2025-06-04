package stack

/*
Problem Statement:
Given an infix expression, Your task is to convert the given infix expression to a postfix expression.

Example 1:
Input: a+b*(c^d-e)^(f+g*h)-i
Output: abcd^e-fgh*+^*+i-
Explanation: Infix to postfix

Example 2:
Input: (p+q)*(m-n)
Output: pq+mn-*
Explanation: Infix to postfix
*/
/*
infix- human; prefix-machine; postfix-machine good/ reverse of prefix.
 The are some rules. infix and postfix we need to keep this rule.
iterate over infix string and generate postfix using rules.
postfix:
-no parentheses
-operations after operants.
Since the topic is stack we need to use stack to store data .

a = a+b*(c^d-e)^(f+g*h)-i
res = abcd^e-fgh*+^*+i-
What we need to do: This problems not in real interview but, we just need to memorise the flow
Infix -> postfix
take 3 variables: i-index, st - stack []string; ans-string.
what is the priority?
	^:3; *,/: 2; -+:1 - you can create priorities map, the power has maximum priority and +- has minimum.
1. iterate string left->right.
2. if s[i] = operant -> add to stack.
3. if s[i] == operator:
	if priority(s[i]) < priority(st.Top()) or st.Empty() || st.Top() == '(':
		add s[i] -> to st.add(s[i])
	else: pop all operators from stack while priority(s[i])>=priority(st.Top()) and add them to ans.add(st.Top())
			after this st.add(s[i]).
3. If s[i]== '(' just add to stack and continue.
4. If s[i] == ')' st.Pop() and res.Add(st.Pop()) until '('. and pop '('.
5. On the end add all remaining elements to the answer and return the ans .

I don't implement this to dont waste the time.
*/
