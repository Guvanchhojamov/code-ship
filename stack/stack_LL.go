package stack

/*
 Implement stack using linked list

Practice:
Pre-requisites: Basic knowledge of stack and operations in the linked list.

Introduction
Stack is a linear data structure in which elements are stored like a pile, i.e. one on top of another.

The following diagram represents a simple stack

Following operations can be performed in the stack:

push(): pushes an element at the top of the stack.
pop(): removes an element from the top of the stack.
size(): returns the size of the stack.
isEmpty(): returns a boolean value for whether the stack is empty or not.
peek()/top(): returns the top element of the stack.
Implementation of stack
The stack can be implemented in two ways:

Statically:  Using arrays
Dynamically: Using a linked list
In this article, we’ll focus more on the implementation of a stack using a linked list.
*/

type Node struct {
	Val  int
	Next *Node
}

type StackLL struct {
	Head *Node
}

func NewStackLL() StackLL {
	return StackLL{}
}

func (s *StackLL) Push(x int) {
	new := &Node{
		Val: x,
	}
	new.Next = s.Head
	s.Head = new // change head, add new element.
}

func (s *StackLL) Pop() int {
	if s.Head == nil {
		panic("pop from empty stack.")
		return -1
	}
	tmp := s.Head
	s.Head = s.Head.Next
	return tmp.Val
}

func (s *StackLL) Top() int {
	if s.Head != nil {
		return s.Head.Val
	}
	return -1
}

func (s *StackLL) Empty() bool {
	return s.Head == nil || s.Head.Val == 0
}

/*
 How about queue with Linked List?
In queue too like similar linked list iteration.
We keep the head and add (link) ne elements from back, instead of changing head itself.
So, this will be first in first out then.
*/

/*
20. Valid Parentheses
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:

- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
- Every close bracket has a corresponding open bracket of the same type.


Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Example 4:
Input: s = "([])"

Output: true

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/
/*
 So, the are can be many solutions using stack, for example.
What is the balanced parantheses?
the open and closed counts must be equal, opened == closed
[]string, using stack we iterate over s, and
if is open '(' add to stack else, remove from stack.
after itration, stack.len ==0 must be zero, if ==0 return true else return flase.
tc: O(n); sc: O(n), we use stack.
2. instead of  keeping paranthesess in stack, we need just count of opened and closed one.
	if `(` balance++ else balance--.
if balance == 0 then return true otherwise return false.
tc: O(n); sc:O(1) we dont use any extra memory.
*/
