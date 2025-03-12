package f3kingLL

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

func NewNode(val int, next *ListNode, prev *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
		Prev: prev,
	}
}

//func main() {
//	fmt.Println("Hello, World!")
//	nums := []int{10, 2, 3, 4, 5, 6}
//	fmt.Println("initial nums:", nums)
//	list := arrToListNode(nums)
//	fmt.Println("after conv to list: ", list, list.Next, list.Next.Next.Prev)
//
//	nums = listToArr(list)
//	fmt.Println("after conv to nums: ", nums)
//
//	list = insertToHead(list, 20)
//	fmt.Println("after insert ot head: ", list, list.Next)
//
//	list = insertToTail(list, 30)
//	nums = listToArr(list)
//	fmt.Println("after insert to tail: ", nums)
//
//	list = deleteKth(list, 8)
//	nums = listToArr(list)
//	fmt.Println("after delete k index: ", nums)
//
//	deleteNodeFromList(list.Next.Next)
//	nums = listToArr(list)
//	fmt.Println("after delete node: ", nums)
//}

func deleteKth(head *ListNode, index int) *ListNode {
	if head == nil {
		return nil
	}
	if index == 0 {
		head = head.Next
		head.Prev = nil
		return head
	}
	var cnt = 0
	var tmpHead = head
	for tmpHead != nil {
		if cnt == index {
			break
		}
		cnt++
		tmpHead = tmpHead.Next
	}
	if tmpHead == nil {
		return head
	}
	front := tmpHead.Next
	back := tmpHead.Prev
	if front == nil && back == nil {
		return nil
	} else if front == nil && back != nil {
		// remove tail.
		tmpHead.Prev.Next = nil
	} else if front != nil && back == nil {
		// remove head
		tmpHead = tmpHead.Next
		tmpHead.Prev = nil
	} else {
		tmpHead.Prev.Next = tmpHead.Next
		tmpHead.Next.Prev = tmpHead.Prev
	}
	return head
}

func deleteNodeFromList(node *ListNode) {
	// node!=head always
	if node == nil {
		return
	}
	prev := node.Prev
	next := node.Next

	if prev == nil && next == nil {
		return
	} else if prev != nil && next == nil {
		prev.Next = nil
		node.Prev = nil
	} else {
		prev.Next = node.Next
		next.Prev = node.Prev
		node.Next = nil
		node.Prev = nil
	}
	return
}

func insertToHead(head *ListNode, num int) *ListNode {
	if head == nil {
		return NewNode(num, nil, nil)
	}

	newHead := NewNode(num, head, nil)
	head.Prev = newHead
	return newHead
}

func insertToTail(head *ListNode, num int) *ListNode {
	if head == nil {
		return NewNode(num, nil, nil)
	}
	var tmpHead = head
	for tmpHead.Next != nil {
		tmpHead = tmpHead.Next
	}
	tmpHead.Next = NewNode(num, nil, tmpHead)
	return head
}

func arrToListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var head = NewNode(nums[0], nil, nil)
	var mover = head
	for i := 1; i < len(nums); i++ {
		newNode := NewNode(nums[i], nil, mover)
		mover.Next = newNode
		mover = mover.Next
	}
	return head
}

func listToArr(head *ListNode) []int {
	var res = []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
