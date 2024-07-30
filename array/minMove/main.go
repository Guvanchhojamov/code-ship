package main 

import "fmt"

func main (){
    minimumMoves("XXXO")
}

func minimumMoves(s string) int {
    reqLen := 3
    xCount,oCount := 0,0
    k:=0
    c := []byte(s)
    for reqLen > 0 {
        for i:=0;i<len(c);i++{
        k++
        if c[i] == 'X' {xCount++}
        if c[i] == 'O' {
                oCount++ 
                if oCount == len(c) {return 0}
            }
            
            if i > 1 {
                if k == 3 {
                    fmt.Println(xCount, reqLen, k, i)
                    if xCount == reqLen {
                            c[i], c[i-1], c[i-2] = 'O','O','O'     
                    }

                }
            }
        }
        reqLen--
        fmt.Println(xCount, string(c))
    }
    return 0 
}