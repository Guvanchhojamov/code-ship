package main

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}

/*
[2,0,2,1,1,0]
low=0; mid=0 hight=len-1;
mid==0  swap mid->low; low++; mid++
mid==1  mid++
mid==2  swap mid->hight; hight--;

0,0,2,1,1,2
0,0,2,1,1,2
0 0 2 1 1 2
0 0 2 1 1 2 mid=1;low=1;hight=4; mid==2
0 0 1 1 2 2

This is the DNF algorythm method.

	2 0 1
*/
func sortColors(nums []int) {
	low, i, hight := 0, 0, len(nums)-1
	for i <= hight {
		switch nums[i] {
		case 0:
			nums[low], nums[i] = nums[i], nums[low]
			low++
			i++
		case 1:
			i++
		case 2:
			nums[i], nums[hight] = nums[hight], nums[i]
			hight--
		}
	}
}
