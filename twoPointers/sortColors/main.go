package main

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}

func sortColors(nums []int) {
	low, i, hight := 0, 0, len(nums)-1
	for i <= hight {
		switch nums[i] {
		case 0:
			nums[low], nums[i] = nums[i], nums[low]
			low++
			i++
		case 2:
			nums[i], nums[hight] = nums[hight], nums[i]
			hight--
		default:
			i++
		}
	}
}
