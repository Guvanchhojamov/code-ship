package stack

/*
146. LRU Cache
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:
- LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
- int get(int key) Return the value of the key if the key exists, otherwise return -1.
- void put(int key, int value) Update the value of the key if the key exists.
- Otherwise, add the key-value pair to the cache.
- If the number of keys exceeds the capacity from this operation, evict the least recently used key.
- The functions get and put must each run in O(1) average time complexity.

Example 1:

Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[       [2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4

Constraints:

1 <= capacity <= 3000
0 <= key <= 10^4
0 <= value <= 10^5
At most 2 * 10^5 calls will be made to get and put.
*/
/*
 Need implement key,val pair cache, LRU.
- Must get,put in O(1).
- If no in key need return -1
- if already exists key, need replace with new val.
- We can store key,val pairs in hashmap.
- if we assume get from hash and put in to hash is in 1 time.
- we can implement is using only hashmap.
- but how to remove least recent used? We don't know it with hashmap.
  To  store,  datas and for quick access we use hashmap. But to remove and preorder, least recently used val .
We need use Doubly linked list.
If the capacity is reached and we need to add new val we, remove from tail.prev.
add to head.next;
*/

type LRUCache struct {
	Data map[int]*DLNode
	Cap  int
	Head *DLNode
	Tail *DLNode
}

func ConstructorCache(capacity int) LRUCache {
	data := make(map[int]*DLNode, 0)
	var dummyHead, dummyTail = NewNode()
	// fmt.Println(dummyHead,dummyTail)

	return LRUCache{
		Data: data,
		Cap:  capacity,
		Head: dummyHead,
		Tail: dummyTail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Data[key]; ok {
		moveToFront(this.Head, node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Data[key]; ok {
		node.Val = value
		moveToFront(this.Head, node)
		return
	}

	if len(this.Data) >= this.Cap {
		lru := this.Tail.Prev
		removeNode(lru)
		delete(this.Data, lru.Key)
	}

	// Add new node
	newNode := &DLNode{
		Key: key,
		Val: value,
	}
	addToFront(this.Head, newNode)
	this.Data[key] = newNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type DLNode struct {
	Key  int
	Val  int
	Prev *DLNode
	Next *DLNode
}

func NewNode() (*DLNode, *DLNode) {
	var dummyHead = &DLNode{
		-1, -1, nil, nil,
	}
	var dummyTail = &DLNode{
		-1, -1, nil, nil,
	}
	dummyHead.Next = dummyTail
	dummyTail.Prev = dummyHead
	return dummyHead, dummyTail
}

func moveToFront(head, node *DLNode) {
	removeNode(node)
	addToFront(head, node)
}

func addToFront(head *DLNode, newNode *DLNode) {
	newNode.Prev = head
	newNode.Next = head.Next
	head.Next.Prev = newNode
	head.Next = newNode
}

func removeNode(node *DLNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
