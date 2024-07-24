package main

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

	if x > y {
		return processSubString(s, x, y, 'a', 'b')
	}
	return processSubString(s, y, x, 'b', 'a')
}

func processSubString(s string, bigPoint, lessPoint int, firstChar, secondChar rune) int {
	stack, tmpStack, res := []rune{}, []rune{}, 0
	for _, v := range s {
		if len(stack) > 0 && v == secondChar && stack[len(stack)-1] == firstChar {
			stack = stack[:len(stack)-1]
			res += bigPoint
		} else {
			stack = append(stack, v)
		}
	}
	if len(stack) > 0 {
		for _, tmpVal := range stack {
			if len(tmpStack) > 0 && tmpVal == firstChar && tmpStack[len(tmpStack)-1] == secondChar {
				tmpStack = tmpStack[:len(tmpStack)-1]
				res += lessPoint
			} else {
				tmpStack = append(tmpStack, tmpVal)
			}
		}
	}
	return res
}
