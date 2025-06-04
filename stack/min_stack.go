package stack

import "math"

/*
155. Min Stack
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

MinStack() initializes the stack object.
 void push(int val) pushes the element val onto the stack.
 void pop() removes the element on the top of the stack.
 int top() gets the top element of the stack.
 int getMin() retrieves the minimum element in the stack.
 You must implement a solution with O(1) time complexity for each function.

Example 1:

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2


Constraints:

-2^31 <= val <= 2^31 - 1
Methods pop, top and getMin operations will always be called on non-empty stacks.
At most 3 * 10^4 calls will be made to push, pop, top, and getMin.
*/
/*
 we use just simple stack implementation from previous solution, to dont waste time and
add MinEl to stack struct to keep track always minimum element under hands, to take element in O(1) time.
*/

type Node struct {
	Val   int
	Next  *Node
	MinEl int
}

type MinStack struct {
	Head  *Node
	MinEl int
}

func Constructor() MinStack {
	return MinStack{
		MinEl: math.MaxInt,
	}
}

func (this *MinStack) Push(val int) {
	newNode := &Node{
		Val:   val,
		MinEl: math.MaxInt,
	}

	if this.Head == nil {
		newNode.MinEl = val
	} else if val < this.Head.MinEl {
		newNode.MinEl = val
	} else {
		newNode.MinEl = this.Head.MinEl
	}
	newNode.Next = this.Head
	this.Head = newNode
}

func (this *MinStack) Pop() {
	this.Head = this.Head.Next
}

func (this *MinStack) Top() int {
	return this.Head.Val
}

func (this *MinStack) GetMin() int {
	return this.Head.MinEl
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
