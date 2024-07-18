package main

func main() {
	reverseSubstr("(oc)el")
}

/*
(u(love)i)

"(ed(et(oc))el)" => leetcode

stack = (ed(et(oc
tmpStr
case v==')':

	for v!='(':
	  tmpStr = stack.pop()
	if v=='(':
	   stack.pop()
	   stack.append(tmpStr)
	   tmpStr=''

default stack=append(v)
*/
func reverseSubstr(s string) string {
	stack, tmpStack := []rune{}, []rune{}
	for _, v := range s {
		switch {
		case v == ')':
			for stack[len(stack)-1] != '(' {
				tmpStack = append(tmpStack, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, tmpStack...)
			tmpStack = []rune{}
		default:
			stack = append(stack, v)
		}
	}
	return string(stack)
}
