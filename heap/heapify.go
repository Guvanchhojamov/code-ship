package heap

/*
give nums and i th node, need to heapify given nums array.

	Heapify means need to keep heap property, maxheap or minheap.
*/

// func main() {
// 	fmt.Println("Hello, World!")
// 	r := max_heapify([]int{1, 14, 10, 8, 7}, 0)
// 	fmt.Println(r)

// 	ll := min_heapify([]int{1, 14, 10, 8, 7}, 0)
// 	fmt.Println(ll)
// }

func max_heapify(nums []int, i int) []int {
	var left = 2*i + 1
	var right = 2*i + 2
	var largest = i
	if largest > len(nums) {
		return nums
	}

	if left < len(nums) && nums[left] > nums[largest] {
		largest = left
	}

	if right < len(nums) && nums[right] > nums[largest] {
		largest = right
	}

	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		return max_heapify(nums, largest)
	}
	return nums
}

func min_heapify(nums []int, i int) []int {
	var left = 2*i + 1
	var right = 2*i + 2
	var lowest = i

	if lowest >= len(nums) {
		return nums
	}

	if left < len(nums) && nums[left] < nums[lowest] {
		lowest = left
	}
	if right < len(nums) && nums[right] < nums[lowest] {
		lowest = right
	}

	if lowest != i {
		nums[i], nums[lowest] = nums[lowest], nums[i]
		return min_heapify(nums, lowest)
	}
	// if lowest is the same node just check next .
	return nums
}

/*
 heapify and build heap is different llgorythms.
 - when we insert some elemtn to already heap arra we need heapify, because the new element
 may violate he heap properyy.
- for heapify the nums must be already on heap property.
heapify:
heapify down
heapify up (bubble up)
insert key, delete key (replace delete root, and last to root)
update key, down or up. increase or decrease key.
tc: LogN
sc: LogN for recursive call, i think.
build heap:
- affects to entire array.
- buld from any array heapifyed array.
- logic is bottom to up.
tc: O(N)
sc: O(n)
*/
