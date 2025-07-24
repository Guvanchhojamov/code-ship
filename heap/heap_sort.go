package heap

/*
 Given array nums we need to sort this nums using heap ds.
 nums = [9,6,8, 2,1,4,3]
 after sort = 9,8,6,4,3,2,1

*/
/*
 what we should do.
  We need to use max heap to sort in desc order or use min-heap for asc order.
  it does not matter, since we sort any order we can easyli read from back or front.
 1. build a max-heap from given arr.
 2. pop root node. store in result arr.
 3. heapify algo again, from index 0.
 4. repeat this until arr is empty.
 tc: n+nlogn
  we build heap on n time, and repeat heapify n times. nlogn.
sc: n for result array.
*/

// func main() {
// 	fmt.Println("Hello, World!")
// 	// nums := Heap{3, 6, 9, 8, 2, 1, 4}
// 	nums := Heap{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	r := heapSort(nums)
// 	fmt.Println(r)
// }

func heapSort(nums Heap) Heap {
	var res = make(Heap, 0, len(nums))

	//build a heap
	for i := len(nums)/2 - 1; i >= 0; i-- {
		nums = heapifyMax(nums, i)
	}

	// pop root element for each node and heapify again
	for len(nums) > 0 {
		res = append(res, nums[0])
		nums[0] = nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		nums = heapifyMax(nums, 0)
	}

	return res
}

func heapifyMax(nums []int, i int) Heap {
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

// Optimal solution.
func sortArrayInPlace(nums Heap) Heap {
	n := len(nums)

	for i := n/2 - 1; i >= 0; i-- {
		heapifyMaxWithSize(nums, i, n)
	}

	for size := n - 1; size > 0; size-- {
		nums[0], nums[size] = nums[size], nums[0]
		heapifyMaxWithSize(nums, 0, size)
	}
	return nums
}

func heapifyMaxWithSize(nums []int, i int, heapSize int) {
	var largest = i
	var left = 2*i + 1
	var right = 2*i + 2

	if left < heapSize && nums[left] > nums[largest] {
		largest = left
	}

	if right < heapSize && nums[right] > nums[largest] {
		largest = right
	}

	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapifyMaxWithSize(nums, largest, heapSize)
	}
}
