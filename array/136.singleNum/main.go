// 136. Single Number

package main
import "fmt"
func main(){    
    nums:=[]int{4,1,2,1,2}
    singleNumber(nums)
}
 
 /* 
    2,2,1

    4,1,2,1,2 = 10
    c

 */ 

 /*
 
  Use Xor/Bit Manipulation
Intuition:
Xor of any two num gives the difference of bit as 1 and same bit as 0.
Thus, using this we get 1 ^ 1 == 0 because the same numbers have same bits.
So, we will always get the single element because all the same ones will evaluate to 0 
and 0^single_number = single_number.
Time: O(n)
Space: O(1)
 */



func singleNumber(nums []int) int {
    result:=0
    fmt.Println(4^5)
    for _,num:=range nums {
        fmt.Println("resBefore:", result)
        
        result = result ^ num    // result^=num  
        
        fmt.Println("resAfter:", result)
    }
    return result
}