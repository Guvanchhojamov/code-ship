package f3kingLL

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
