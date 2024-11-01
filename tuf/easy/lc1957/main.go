package main

import (
	"fmt"
)

func main() {
	str := makeFancyString("leeetcode")
	fmt.Println(str)
}

func makeFancyString(s string) string {
	var result = []byte{}
	var count = 1
	result = append(result, s[0])
	for i := 1; i < len(s); i++ {
		if len(result) > 0 {
			switch s[i] == s[i-1] {
			case true:
				count++
				if count == 3 {
					count--
					continue
				} else {
					result = append(result, s[i])
				}
			default:
				result = append(result, s[i])
				count = 1
			}
		}
	}
	return string(result)
}

/*
 Input: s = "aaabaaaa"
 Output: "aabaa"

 res = aaa
 count = 3
  charStack = []byte
  for 1 to len(s):
    if s[i]==s[i-1]:
        c++
        if c == 3:
          continue
        res = append(res,s[i])
    else:
    c=1
     res = append(res,s[i])

   return string(stack)
*/

/*


 */
