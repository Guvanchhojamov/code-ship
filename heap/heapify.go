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

/*
 2. Check if an array represents a min-heap or not ?
 to be mni heap, the root must be min element in the array.
 to be this the left and right child must be less then parent node element.
 itreate over arr and check childs, if not less then element, return false, else return true.
 tc: O(logn)
 sc: O(1)
2. Convert min Heap to max Heap
when we build any heap, we start iterating from bottom to top.
and bottom starts from last inner node.
Start from last inner node and go to top and convert min-heap to max recursively.
inner nodes: 0-> n/2-1. 0-indexed array.
while i:=n-2/1;i>=0: do recursive maxheap.
left, right, largest;
default largest is parent or root node.
swap with max(left,right) if they are greater than root node.

*/
