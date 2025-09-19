package heap

/*
 Heap  - specified completely binary tree.
 If binary tree is satisfied Heap property then the tree is can be heap.
 Heap property is:
	- Max heap
	- Min heap.
- Max - heap is the root node is greater than all of other nodes of tree.
- Min heap - is the roon node is smallest than all othe nodes.
- Root node is the top element of Binary tree.
- complete binary tree:
	- all levels of the binary tree filled from left to right.
	tree level is 0...n
Priority queue:
 ordered by priority, by heap type, max or min heap.
 Heap store elements in array:

 H - is the height of the heap starting from 0.
 S - is the size of array withc stores heap elements.
 The heap node has left nand right nodes.
 lefNodeIndex = 2*i+1
 rightNodeIdex = 2*i+2.

  S = (2^h+1)-1
  Number of nodes in level = 2^h.
  H - log(a)N.
  parentIndex = i/2.
  Time complexity:
make a heap: O(n)
findMax = O(1)
findMin = O(n) (or versa for min heap)
deleteMa = O(n) after deleteion we need reorder heap nodes.
increase-key = O(logn) reorder heap.
insert-node = O(LogN) reorder.

Tasks to do.
- heapify algorythm. Heapify max and min, for given nums.
- Build heap from array.
- Implement max heap.
- implement min heap.
- implement max heap with:
	- insert node
	- delete node
	- increase key.
	- decrease key.
- implement min node too with this methods.

*/
