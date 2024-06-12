package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	decodeString("3[a]2[bc]")
}

/*
3[a2[b]]
Hint

1. push elements to the stack until meet ']'
2. pop elements from the stack until meet '[' => contents
3. pop elements digitals from the stack => number

4. repeat "contents" by "number" times and push to the stack
5. repeat step 1 to 4
*/
func decodeString(s string) string {
	var stack []byte
	c := []byte(s)
	var number byte
	var numsStack []string
	numsStack = append(numsStack, regexp.MustCompile("[0-9]+").FindAllString(s, -1)...)
	fmt.Println(numsStack)
	for i := 0; i < len(c); i++ {
		switch c[i] {
		case ']':
			str := []byte{}
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				str = append([]byte{stack[len(stack)-1]}, str...)
				stack = stack[:len(stack)-1]
			}
			fmt.Println("numsStack", len(numsStack), numsStack[len(numsStack)-1])
			fmt.Println(string(stack), string(number), string(str), numsStack[len(numsStack)-1])
			stack = append(stack[:len(stack)-2], stringRepeater(str, numsStack[len(numsStack)-1])...)
			numsStack = numsStack[:len(numsStack)-1]
			fmt.Println(string(stack), string(number))

		default:
			stack = append(stack, c[i])
		}
	}
	fmt.Println(string(stack), string(number))
	return string(stack)
}

func stringRepeater(str []byte, num string) []byte {
	var newStr []byte
	count, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < count; i++ {
		newStr = append(newStr, str...)
	}
	fmt.Println(string(newStr))
	return newStr
}
