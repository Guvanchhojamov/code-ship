package main

func main (){
    moveZeroes([]nums{0,1,0,3,12})
}

func moveZeroes(nums []int)  {
    i,j :=0,0
    for i < len(nums) {
       if nums[i] != 0 {
        nums[i],nums[j] = nums[j], nums[i]
        j++
       }
       i++
    }    
    fmt.Println(nums)
}
/*
 0 1 0 3 12
 1 0 0 3 12
          

 1 0 1  

*/


/*
 0 1 0 3 12 
 i=0 ; j= 1
 
 for i < len: 
 if n[i]==0: 
    swap n[i],n[j] 
    i++
 j++

 
*/