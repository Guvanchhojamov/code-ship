package main

import "fmt"

func main() {
	maximumGain("aabbabkbbbfvybssbtaobaaaabataaadabbbmakgabbaoapbbbbobaabvqhbbzbbkapabaavbbeghacabamdpaaqbqabbjbababmbakbaabajabasaabbwabrbbaabbafubayaazbbbaababbaaha", 1926, 4320)
}

/*
for i in range s:

		if x<y{
			if i==a {
			 for stack.last != b && len.stak > 0 {
			  stack.pop()
			 }
			  if stack.last==b {
			    stack.pop()
				res+=y
			  }
			}
			stack.append(i)
		} else {
	       if i==b {
			 for stack.last != a && len.stak > 0 {
			  stack.pop()
			 }
			  if stack.last==a {
			    stack.pop()
				res+=x
			  }
			}
			stack.append(i)
		}

}
*/
func maximumGain(s string, x int, y int) int {
	//x=ab
	//y=ba
	stack := []rune{}
	tmpStack := []rune{}
	res := 0
	switch x > y {
	case true:
		for _, v := range s {
			if len(stack) > 0 && v == 'b' && stack[len(stack)-1] == 'a' {
				stack = stack[:len(stack)-1]
				res += x
			} else {
				stack = append(stack, v)
			}
		}
		if len(stack) > 0 {
			for _, tmpVal := range stack {
				if len(tmpStack) > 0 && tmpVal == 'a' && tmpStack[len(tmpStack)-1] == 'b' {
					tmpStack = tmpStack[:len(tmpStack)-1]
					res += y
				} else {
					tmpStack = append(tmpStack, tmpVal)
				}
			}
		}
	case false:
		for _, val := range s {
			if len(stack) > 0 && val == 'a' && stack[len(stack)-1] == 'b' {
				stack = stack[:len(stack)-1]
				res += y
			} else {
				stack = append(stack, val)
			}
		}
		fmt.Println(string(stack), res)
		if len(stack) > 0 {
			for _, tmpVal := range stack {
				if len(tmpStack) > 0 && tmpVal == 'b' && tmpStack[len(tmpStack)-1] == 'a' {
					tmpStack = tmpStack[:len(tmpStack)-1]
					res += x
				} else {
					tmpStack = append(tmpStack, tmpVal)
				}
			}
		}
		fmt.Println(string(stack), string(tmpStack), res)
	}
	fmt.Println(res)
	return res
}
