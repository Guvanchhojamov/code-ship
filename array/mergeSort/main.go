package main 

import "fmt"
import "bufio"
import "os"
import "strings"

func main(){
   // nums := []int{5,4,3,2,1}
   // mergeSort(nums, 0, len(nums)-1)
   readerTest()
}

/*
  n=5     l=0  mid =n/2  r=len-1
  
  1. 5 4 3    2 1 
    n=3 l=0  mid=n/2  r=len-1
    5 4        3 
    n=2 l=0   mid=n/2  r=len-1
    5          4 
    
  2. 
*/

func mergeSort(nums []int,l, r int) []int{
    mid := len(nums)-1 / 2 
    //right :=  len(nums)-1 
    if l >= r {
      return nums
    }

    mergeSort(nums[l:mid],l,mid)
    fmt.Println("nums1",nums)
    //nums =  mergeSort(nums,mid+1, right)
    //fmt.Println("nums2",nums)
    return nums
}


// func readerTest(){
//       reader := bufio.NewReader(os.Stdin)
//        fmt.Println("Running...")

//       // Проверяем, была ли нажата клавиша Enter
//       input, _ := reader.ReadString('\n')

//       input = strings.TrimSpace(input)
//       sum:= 0 
//       numsStr := strings.Split(input, " ")
        
//       fmt.Println("input", input, sum)
// }

