package heap

/*
 we need to build heap from given numbers.
 Containign heap property.
The heap contains internal nodes and leafs.

internal nodes: 0 -> (n/2)-1
range of leaves: n/2 to n-1.
 How to build a heap?
 Heapify all the indernal nodes.
 using bottom-up logis,

 If we heapify all internal nodes then we build heap.
 because we start from  last internal node, the last internal node is
 parent of last leave node or nodes.
 so, if we start from last leaves, we gurantee that all the above elements are
 be in heap property.

*/

// func main() {
// 	r := build_max_heap([]int{6, 3, 5, 0, 8, 2, 1, 9})
// 	fmt.Println(r)

// 	m := build_min_heap([]int{3, 6, 5, 0, 8, 2, 1, 9})
// 	fmt.Println(m)
// }

func build_max_heap(nums []int) []int {
	for i := (len(nums) / 2) - 1; i >= 0; i-- {
		nums = mx_heapify(nums, i)
	}
	return nums
}

func build_min_heap(nums []int) []int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		nums = mn_heapify(nums, i)
	}
	return nums
}

func mx_heapify(nums []int, i int) []int {
	var left = 2*i + 1
	var right = 2*i + 2
	var largest = i

	if largest >= len(nums) {
		return nums
	}
	if left < len(nums) && nums[left] > nums[largest] {
		largest = left
	}
	if right < len(nums) && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		return mx_heapify(nums, largest)
	}
	return nums
}

func mn_heapify(nums []int, i int) []int {
	var lefti = 2*i + 1
	var righti = 2*i + 2
	var lowesti = i

	if lowesti > len(nums) {
		return nums
	}

	if lefti < len(nums) && nums[lefti] < nums[lowesti] {
		lowesti = lefti
	}
	if righti < len(nums) && nums[righti] < nums[lowesti] {
		lowesti = righti
	}
	if lowesti != i {
		nums[lowesti], nums[i] = nums[i], nums[lowesti]
		return mn_heapify(nums, lowesti)
	}
	// otherwise return nums.
	return nums
}
