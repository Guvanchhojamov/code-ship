package LL

//func main() {
//	fmt.Println("Hello, World!")
//	nums := []int{10, 2, 3, 4, 5, 6}
//	fmt.Println("initial nums:", nums)
//
//	list := arrToList(nums)
//
//	fmt.Println("ListNode head val: ", list.Val)
//
//	nums = ListNodeToArr(list)
//
//	fmt.Println("List to arr:", nums)
//
//	list = insertToHead(list, 15)
//	//list = NewNode(25, nil)
//	fmt.Println("new head: ", list)
//	nums = ListNodeToArr(list)
//
//	fmt.Println("List to arr:", nums)
//
//	list = deleteTailOfList(list)
//	fmt.Println("after tail  deletion:", ListNodeToArr(list))
//	list = deleteHead(list)
//	fmt.Println("after head deletion:", ListNodeToArr(list))
//	list = insertToTail(list, 77)
//	fmt.Println("after tail insertion:", ListNodeToArr(list))
//	list = insertToPosition(list, 88, 6)
//	fmt.Println("after insert position:", ListNodeToArr(list))
//	list = insertToValue(list, 12, 0)
//	fmt.Println("after insert to value:", ListNodeToArr(list))
//	list = deleteIndex(list, 3)
//	fmt.Println("after delete index:", ListNodeToArr(list))
//	list = deleteVal(list, 10)
//	fmt.Println("after delete value:", ListNodeToArr(list))
//
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}

func insertToPosition(head *ListNode, val, position int) *ListNode {
	var newNode = NewNode(val, nil)
	if head == nil {
		return nil
	}
	if position == 0 {
		return &ListNode{
			Val:  val,
			Next: head,
		}
	}
	var cnt = 0
	var tmpHead = head
	for tmpHead.Next != nil {
		if cnt == position-1 {
			newNode.Next = tmpHead.Next
			tmpHead.Next = newNode
			return head
		}
		cnt++
		tmpHead = tmpHead.Next
	}
	// take tail case
	if cnt == position-1 {
		newNode.Next = tmpHead.Next
		tmpHead.Next = newNode
		return head
	}
	tmpHead = tmpHead.Next
	return head
}

func insertToValue(head *ListNode, newVal, val int) *ListNode {
	var newNode = NewNode(newVal, nil)
	if head == nil {
		return nil
	}
	// if position == 0 {
	// 	return &ListNode{
	// 		Val:  newVal,
	// 		Next: head,
	// 	}
	// }
	var tmpHead = head
	for tmpHead != nil {
		if tmpHead.Val == val {
			newNode.Next = tmpHead.Next
			tmpHead.Next = newNode
			return head
		}
		tmpHead = tmpHead.Next
	}
	return head
}

func insertToHead(head *ListNode, num int) *ListNode {
	if head == nil {
		return NewNode(num, nil)
	}
	newHead := NewNode(num, head)
	return newHead
	/*
	   return NewNode(num,head)
	*/
}

func insertToTail(head *ListNode, num int) *ListNode {
	if head == nil {
		return &ListNode{
			Val:  num,
			Next: nil,
		}
	}
	var tmpHead = head
	for tmpHead.Next != nil {
		tmpHead = tmpHead.Next
	}
	var newNode = NewNode(num, nil)
	tmpHead.Next = newNode

	return head
}

func deleteIndex(head *ListNode, index int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	if index == 0 {
		return head.Next
	}
	var tmpHead = head
	var cnt = 0
	for tmpHead.Next != nil {
		if cnt == index-1 {
			tmpHead.Next = tmpHead.Next.Next
			return head
		}
		tmpHead = tmpHead.Next
		cnt++
	}
	return head
}

func deleteVal(head *ListNode, value int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Val == value {
		return head.Next
	}
	var tmpHead = head
	for tmpHead.Next != nil {
		if tmpHead.Next.Val == value {
			tmpHead.Next = tmpHead.Next.Next
			return head
		}
		tmpHead = tmpHead.Next
	}
	return head
}

func deleteHead(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	return head.Next
}

func deleteTailOfList(head *ListNode) *ListNode {
	var inHead = head
	if head == nil || head.Next == nil {
		return nil
	}
	for inHead.Next.Next != nil {
		inHead = inHead.Next
	}
	inHead.Next = nil
	return head
}

func arrToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return NewNode(0, nil)
	}

	var head = NewNode(nums[0], nil)
	var mover = head
	for i := 1; i < len(nums); i++ {
		newNode := NewNode(nums[i], nil)
		mover.Next = newNode
		mover = mover.Next
	}
	return head
}

func ListNodeToArr(head *ListNode) []int {
	var res = []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
