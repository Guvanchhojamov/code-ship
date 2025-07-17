package heap

import "fmt"

/*
 Heap algorythnms:
 - extract max. (heap.pop())
 - increase key.
 - decrease key.
 - insert key.
 - prerequisite (heapify algorythm)
 -
*/
/*
 extract value.
 1. save root element. root element is always max.
 2. copy last element to root element.
 3. remove last element.
 4. Heapify array. i=0
*/
func main() {
	nums := Heap{9, 8, 7, 5, 4, 3, 2}
	mx := extractMax(nums)
	fmt.Println("maxEl:", mx)

	nums = Heap{9, 8, 7, 5, 4, 3, 2}
	res := increaseNode(nums, 50, 3)
	fmt.Println("with increased node:", res)

	nums = Heap{9, 8, 7, 5, 4, 3, 2}
	res = decreaseNode(nums, 1, 2)
	fmt.Println("with decreased  node:", res)

	nums = Heap{9, 8, 7, 5, 4, 3, 2}
	res = insertElement(nums, 6)
	fmt.Println("with inserted  node:", res)
}

type Heap []int

func extractMax(nums Heap) int {
	var maxel = nums[0]
	var ln = len(nums)
	for i := 0; i < ln; i++ {
		maxel = nums[0]
		nums[0] = nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		nums = maxHeapify(nums, 0)
		fmt.Println(nums, " max: ", maxel)
	}
	return maxel
}

// tc: O(logN-for heapify); // sc: O(logN-for heapify)
func maxHeapify(nums Heap, i int) Heap {
	var left = 2*i + 1
	var right = 2*i + 2
	var largest = i

	if left < len(nums) && nums[left] > nums[largest] {
		largest = left
	}
	if right < len(nums) && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		maxHeapify(nums, largest)
	}
	return nums
}

/*
 increase key
 1. increase the node.
 2. check with parent:
	a.parent>node. stop
	b. node is root. stop
else swipe.
*/

func increaseNode(nums Heap, val, idx int) Heap {
	nums[idx] = val

	for idx > 0 && nums[idx/2] < nums[idx] {
		if nums[idx/2] < nums[idx] {
			nums[idx], nums[idx/2] = nums[idx/2], nums[idx]
		}
		idx = idx / 2
	}
	return nums
}

/*
 tc: O(logn)- worst case.
 sc: O(1)
*/

/*
decrease node.
1. in can only go down.
1. Decrease node value.
2. max Heapify nums from index of changed node.
*/
func decreaseNode(nums Heap, val, i int) Heap {
	if nums[i] < val {
		return nums
	}
	nums[i] = val
	nums = maxHeapify(nums, i)
	return nums
}

/*
 tc: O(logn) heapify
 sc: O(logn) heapify
*/

/*
insert element
1. insert to the element last.
2. go until:
	parent<element or element==0:
tc: O(logn)
sc: O(1)
*/

func insertElement(nums Heap, val int) Heap {
	nums = append(nums, val)
	i := len(nums) - 1
	for nums[i/2] < nums[i] && i > 0 {
		nums[i/2], nums[i] = nums[i], nums[i/2]
		i = i / 2
	}
	return nums
}

/*
 find max = O(1)
 Delete max = O(logn)
 inset  = O(logn)
 increase key = O(logn)
 decrease Key = O(logN)
 findMin = O(n)
 Delete random = O(n)
 Search random = O(n)
*/

/*
 All algorythm for min-heap is also same:
 1. finMax - O(n)
 2. DeleteMax - O(n)
 3. Insert - O(logn):
	insert to last.
	go up checking with parent.
4. increase key - O(logn)
	increase val, go down until end.
5. decrease key - O(logn)
	decrease val, go up until root.
6. Find min - O(1)
7. Delete min - O(logN)
all operations opposite of max-heap.
*/
