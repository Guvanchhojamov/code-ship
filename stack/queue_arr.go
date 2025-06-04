package stack

/*
Implement Queue Using Array
Problem Statement:
Implement Queue Data Structure using Array with all functions like pop, push, top, size, etc.

Example:
Input: push(4)
       push(14)
       push(24)
       push(34)
       top()
       size()
       pop()

Output:

The element pushed is 4
The element pushed is 14
The element pushed is 24
The element pushed is 34
The peek of the queue before deleting any element 4
The size of the queue before deletion 4
The first element to be deleted 4
The peek of the queue after deleting an element 14
The size of the queue after deleting an element 3
*/

/*
 As we know the stack is fifo workflow.
*/

type Queue struct {
	Nums  [10]int
	Start int
	End   int
	Size  int
}

func NewQueue() *Queue {
	return &Queue{
		Start: -1,
		End:   -1,
	}
}

func (q *Queue) Push(num int) {
	if q.Start < 0 {
		q.Start++
	}
	q.End++
	q.Nums[q.End] = num
}

func (q *Queue) Pop() int {
	first := q.Nums[q.Start]
	q.Nums[q.Start] = 0
	q.Start++
	if q.Start > q.End {
		q.Start = -1
		q.End = -1
	}
	return first
}

func (q *Queue) Top() int {
	return q.Nums[q.Start]
}

func (q *Queue) Len() int {
	return (q.End + 1) - q.Start
}
