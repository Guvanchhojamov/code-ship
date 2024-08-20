package main


func main(){
    check([]int{})
}

// 1752. Check if Array Is Sorted and Rotated

/*
    Solution approach: 
    1. Sorted and rotated array should must have only 1 rotation point. 
    2. The first element of sorted and rotated array must always less than last elemen 
         - nums[len(nums)-1] < nums[0]
    3. If rotations count greater than 1 return false if <= 0 return true.
*/
func check(nums []int) bool {
    var rotation int = 0

    for i,v:=range nums{
    // Note: here we are taking mod, because we have to check last element with first 
    // element. If element is sorted rotated, then first element should always be greater  than the last element
        if v > nums[(i+1) % len(nums)] {
            rotation++
        }
    }
    return rotation <= 1
}
