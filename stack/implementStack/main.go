package main

import "fmt"

var commands = []string{"MyQueue", "push", "push", "peek", "pop", "empty"}
var result = make([]MyQueue, 0)

func main() {
	obj := Constructor()
	fmt.Println(obj)
}

type MyQueue struct {
	queue []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyQueue) Pop() int {
	if len(this.queue) > 0 {
		return this.queue[len(this.queue)-1]
	}
	return 0
}

func (this *MyQueue) Peek() int {
	return this.queue[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
