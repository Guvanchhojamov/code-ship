package stack

/*
Implement Stack using Array
Problem statement: Implement a stack using an array.
Note: Stack is a data structure that follows the Last In First Out (LIFO) rule.
Example:
Explanation:
push(): Insert the element in the stack.
pop(): Remove and return the topmost element of the stack.
top(): Return the topmost element of the stack
size(): Return the number of remaining elements in the stack.
*/

/*
  in other languages we can use class, in golang we use Struct data type.
  [0,0,0,0,0,0,0] - stack with arrays must be fixed size. (with slices it can be dynamic, in go only.)
[7]int - for example.
*/
/*
push(): Insert the element in the stack.
pop(): Remove and return the topmost element of the stack.
top(): Return the topmost element of the stack
size(): Return the number of remaining elements in the stack.
*/

type Stack struct {
	Nums []int
	Size int
	Idx  int
}

/*
idx = -1 is staring pint.
push add.nums[idx] and idx++, size++
pop idx--, size--
top return nums[idx] just.
size return size value itself.
*/

func NewStack() *Stack {
	return &Stack{
		Idx: -1,
	}
}

func (s *Stack) Push(num int) {
	s.Nums = append(s.Nums, num)
	s.Idx++
	s.Size++
}

func (s *Stack) Pop() int {
	last := s.Nums[s.Idx]
	s.Nums = s.Nums[:len(s.Nums)-1]
	s.Idx--
	s.Size--
	return last
}

func (s *Stack) Top() int {
	return s.Nums[s.Idx]
}

func (s *Stack) Len() int {
	return s.Size
}

//We can optimize it using idiomatic go, just len(s.Nums) instead of index and size variables.
