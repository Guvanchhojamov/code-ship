package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// nums := Heap{3, 6, 9, 8, 2, 1, 4}
	nums := []int{2, 7, 4, 1, 8, 1}
	r := lastStoneWeight(nums)
	fmt.Println(r)
}

func lastStoneWeight(stones []int) int {
	var n = len(stones)

	for i := n/2 - 1; i >= 0; i-- {
		stones = heapifyMax(stones, i)
	}

	for len(stones) > 1 {
		y := stones[0]
		x := stones[1]
		stones = stones[2:]
		if y != x {
			stones = append(stones, y-x)
		}
		i := len(stones) - 1
		for stones[i/2] < stones[i] && i > 0 {
			stones[i/2], stones[i] = stones[i], stones[i/2]
			i = i / 2
		}
	}
	if len(stones) > 0 {
		return stones[0]
	}
	return 0
}

/*
	we need to take stones and comparet them, and calculate new weight, forstorenes.
	and return last remaining stone.
	x==y destroy both.
	x!=y then y-x=z and add it to respense.

1. brute force.
until n>=0 we do:

	sort arr in asc or desc order.
	take first or last 2 heavy nums.
	compare them:
	if they are equal then remove both from arr.
	if not equal then remove last element.
	replace prelast with new var y-x.

return nums.
tc: n*logn
sc: O(1).
How we can optimze?
What ds we should use to store sorted elements?
assume we store elements in max heap.
and each time we take root, and take max(leftChild, rightChild) - one
of the childs guaraneteed is second heavy numer.
after taking if x==y:

	then we pop root and max(left,right) and. the left child is always
	greater than right child, after heapifying.
	not equal, then we y-x of them.
	remove both, and add this element to the end of arr.
	and heapify to top claculatoin.

we do this operation n times. for each element.

	logn.
	we do n times logn
	tc: nlogn.
	sc: logN.
*/
func heapifyMax(nums []int, i int) []int {
	var largest = i
	var left = 2*i + 1
	var right = 2*i + 2

	if left < len(nums) && nums[left] > nums[largest] {
		largest = left
	}
	if right < len(nums) && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		return heapifyMax(nums, largest)
	}
	return nums
}
