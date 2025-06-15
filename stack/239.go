package stack

//Leetcode 239.

/*
	Simple brute force we need chec from. to n-k for each i-k elements max.
	But, we got tle for this solutuon.
	How we can optimize?
	What we need?

-keep window elements always, check max with new element.
to keep window elements we use stack ds, monotonic stack in decreasing order.
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3 Output: [3,3,5,5,6,7]
Algo:

	st decreasing order.[max,..,min] - stack is double ended we can pop,puhs back and front.

for i in range n:
if q==k:

	res:=append(maxel)
	st.pop_front()
	q--

for st.len()>0 and st.last<a[i]:

	st.pop_back()

if len(st) == 0:

	maxel = a[i]

else:

	maxel = st.front()

st.push_back(a[i])
q++
return res.
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3 Output: [3,3,5,5,6,7]
st =6,7
res=3,3,5,5,6
max=6,
q=0
*/
type Deque struct {
	q []int
}

func (d *Deque) Len() int {
	return len(d.q)
}

func (d *Deque) PushBack(v int) {
	d.q = append(d.q, v)
}

func (d *Deque) PopBack() {
	d.q = d.q[:len(d.q)-1]
}

func (d *Deque) Back() int {
	return d.q[len(d.q)-1]
}

func (d *Deque) PopFront() {
	d.q = d.q[1:]
}

func (d *Deque) Front() int {
	return d.q[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	// handel edge cases
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}

	var dq = &Deque{}
	var res = make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		if dq.Len() > 0 && dq.Front() == i-k {
			dq.PopFront()
		}

		for dq.Len() > 0 && nums[dq.Back()] < nums[i] {
			dq.PopBack()
		}
		dq.PushBack(i)

		if i >= k-1 {
			res = append(res, nums[dq.Front()])
		}
	}

	return res
}

/*
return res.
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3 Output: [3,3,5,5,6,7]
7-3=4
*/
