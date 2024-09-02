package main

import "fmt"

// 2022. Convert 1D Array Into 2D Array.
/*
 m rows
 n cols 
 [1,2,3,4] m=2 n=2   
  1 2 
  3 4 

[1,2,3] m=1 n=3   
 1 2 3 

 0 1 2 

 0 1 2  

0 1 
2 3 

 0 0 
 0 1 
 1 0  
 1 1 
  
i+j +i
0 0  0 
0 1  1 
1 0  2 
1 1  3 

3     1 2 


*/


func main (){
    construct2DArray([]int{1,2,3,4}, 2,2)
}

func construct2DArray(original []int, m int, n int) [][]int {
    if m*n < len(original) || m*n > len(original) {
        return [][]int{}
    }

    var result = make([][]int,0,max(m,n))
    fmt.Println(result)
    for i:=0; i<m; i++{
        rows:=[]int{}
        for j:=0; j<n; j++{
            fmt.Println(original[(i+j)+i])
            rows = append(rows, original[(i+j)+i])
        }
        result = append(result,rows)
    }
    fmt.Println(result)
    return result
}