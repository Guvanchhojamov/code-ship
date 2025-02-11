package main

import "math"

func main(){

}
/*
 2349. Design a Number Container System
Design a number container system that can do the following:

 - Insert or Replace a number at the given index in the system.
 - Return the smallest index for the given number in the system.

Implement the NumberContainers class:
- NumberContainers() Initializes the number container system.
- void change(int index, int number) Fills the container at index with the number. If there is already a number at that index, replace it.
- int find(int number) Returns the smallest index for the given number, or -1 if there is no index that is filled by number in the system.


Example 1:
Input
["NumberContainers", "find", "change", "change", "change", "change", "find", "change", "find"]
[[], [10], [2, 10], [1, 10], [3, 10], [5, 10], [10], [1, 20], [10]]
Output
[null, -1, null, null, null, null, 1, null, 2]

Explanation
NumberContainers nc = new NumberContainers();
nc.find(10); // There is no index that is filled with number 10. Therefore, we return -1.
nc.change(2, 10); // Your container at index 2 will be filled with number 10.
nc.change(1, 10); // Your container at index 1 will be filled with number 10.
nc.change(3, 10); // Your container at index 3 will be filled with number 10.
nc.change(5, 10); // Your container at index 5 will be filled with number 10.
nc.find(10); // Number 10 is at the indices 1, 2, 3, and 5. Since the smallest index that is filled with 10 is 1, we return 1.
nc.change(1, 20); // Your container at index 1 will be filled with number 20. Note that index 1 was filled with 10 and then replaced with 20.
nc.find(10); // Number 10 is at the indices 2, 3, and 5. The smallest index that is filled with 10 is 2. Therefore, we return 2.


Constraints:
1 <= index, number <= 10^9
At most 10^5 calls will be made in total to change and find. */
/*
 Since range is big square time complexity not acceptable.
     Key points:
        containerNums = []int{} -> to store numbers. (maybe using map is better in there)
        we have find(number) int:
                return min index of number    // we cant use map since numbers are not distinct. But for store min index we can use map.
        change(index, number):
                if containerNums[index] then containerNums[index]=number
                if not exist this index we need add with this index,
                if exists we need replace with given number existing index value.

    What DS we can, need to choose for containerNums:

       Container struct {
           index int
           value int
       }

       numContainers = map[Container.index]Conteiner.value.

*/
type Container struct {
	Index int
	Value int
}

type NumberContainers struct {
	Containers []Container
}


func Constructor() NumberContainers {
	return NumberContainers{
		Containers: make([]Container, 0, 0)
	}
}


func (this *NumberContainers) Change(index int, number int)  {
	for i , container:=range Containers{
		if container.Index == index {
			container[i].Value = number
			return
		}
	}
	this.Containers = append(this.Containers, Container{
		Index:index,
		Value:number,
	})
}


func (this *NumberContainers) Find(number int) int {
	// 1. implement linear search, for to be shure it works
	// 2. implement BS in there for optimal approach.
	var minNumberIndex = math.MaxInt
	var ans = -1
	for i,container:=range this.Containers {
		if container.Value == number {
			minNumberIndex = min(minNumberIndex,container.Index)
			ans = minNumberIndex
		}
	}
	return ans
}


/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */

/*
  Ok it works, but we got time limit exeption,
  What is the TC for this func?
  for change O(n),
  for find O(n)
  For n calls our functions runs N time this is N*n = N^2. Square time. Not acceptable.
  How we can optimize our change() and find functions?
     Key poins:
      In each function we need to find something in given range.
      It Doesn't matter what we do with finded Entity.
        1. Given range
        2. Given target
        3. Need to find?
        4. But our Containers is not sorted? Can we sort it? THen Tc: N*(logN)^2
        Can we use Binary Search then ? No
      Can we use 2 hashMaps?
      1. containersIdxMap[int]int   // we can store there container index with value.
      2. containersNumMap[int]sorted([]int)  // we store in there numbers minimum index

*/
/*
Input
["NumberContainers", "find", "change", "change", "change", "change", "find", "change", "find"]
[[],                  [10],  [2, 10],  [1, 10],  [3, 10],   [5, 10],  [10], [1, 20], [10]]
Output
[null, -1, null, null, null, null, 1, null, 2]
*/
/*
    -1
   cidxMap:
     2:10
     1:10
     3:10
   cNMap:
     10: [2] -> [1,2]
