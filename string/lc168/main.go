package main

import "fmt"

// 168. Excel Sheet Column Title.

/*
Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.

For example:

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28 
...
*/

func main(){
 convertToTitle(4)
}

func convertToTitle(columnNumber int) string {
    var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    var lettersMap = make(map[int]string)
    
    for i,val := range letters {
        lettersMap[i+1] = string(val)
    }
    fmt.Println(lettersMap)
    if val,exists:=lettersMap[columnNumber]; exists {
        fmt.Println(val)
        return val
    }
    
    return ""
}