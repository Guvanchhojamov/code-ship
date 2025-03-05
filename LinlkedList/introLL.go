package linlkedlist

import "fmt"

// 707. Design Linked List
type MyLinkedList struct {
	Data int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{Next: nil}
}

func (this *MyLinkedList) Get(index int) int {
	var tmpHead = this.Next
	var count = 0
	for tmpHead != nil {
		fmt.Println(tmpHead.Data, index)
		if count == index {
			return tmpHead.Data
		}
		count++
		tmpHead = tmpHead.Next
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &MyLinkedList{
		Data: val,
		Next: this.Next,
	}
	this.Next = newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &MyLinkedList{
		Data: val,
		Next: nil,
	}
	var tmpHead = this
	for tmpHead.Next != nil {
		tmpHead = tmpHead.Next
	}
	tmpHead.Next = newNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	newNode := &MyLinkedList{
		Data: val,
		Next: nil,
	}
	var tmpHead = this
	var count = 0
	for tmpHead != nil {
		if count == index {
			newNode.Next = tmpHead.Next
			tmpHead.Next = newNode
			return
		}
		count++
		tmpHead = tmpHead.Next
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}
	var tmpHead = this
	var count = 0
	for tmpHead.Next != nil {
		if count == index {
			tmpHead.Next = tmpHead.Next.Next
			return
		}
		count++
		tmpHead = tmpHead.Next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
/*

 */
func main() {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(1)
	obj.AddAtIndex(1, 2)
	obj.Get(1)
	obj.DeleteAtIndex(1)
	obj.Get(1)
	fmt.Println(obj)
}
