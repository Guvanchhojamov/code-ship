package main

func main() {

}
func lexicographicallySmallestArray(nums []int, limit int) []int {
	var i = 0
	for i < len(nums) {
		j := i + 1
		for j < len(nums) {
			if nums[j] < nums[i] && nums[i]-nums[j] <= limit {
				nums[i], nums[j] = nums[j], nums[i]
				j = i + 1
			} else {
				j++
			}
		}
		i++
	}
	return nums
}

/*
 2. How can we optimize solution?


 1. Brute force.
   for to be lex. small element. we need maximize
    all smallest elements move to left. To start of array.
    a. Take each element.
    b. iterate elemnt from el.Index -> end.
        for i:=0 range nums:
            for j:=i+1 range nums:
            if a[j] < a[i] :
               try to swap:
                if a[i]-a[j] < limit:
                   swap(a[i],a[j])

                   continue
          return nums;

        nums = [1,7,6,18,2,1], limit = 3
        1. 1 7 6 18 2 1 -> i=0
         i=6
         1 6 7 18 2 1
         i=7
         1 6 7 18 2 1
         i =18
         1 6 7 18 2 1
         i = 2
         1 6 7 18 1 2
         i=1
         1 6 7 18 1 2  ===> result.
     tc: N*N-> O(n^2)
     sc: O(1)
    is it ok? Nope.

*/
