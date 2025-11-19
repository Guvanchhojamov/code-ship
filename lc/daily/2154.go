package daily

import "sort"

func findFinalValue(nums []int, original int) int {
	var mp = make(map[int]bool)
	for _, num := range nums {
		mp[num] = true
	}

	for mp[original] {
		original = original * 2
	}
	return original
}

func findFinalValueOptimal(nums []int, original int) int {
	sort.Ints(nums)
	for _, num := range nums {
		if num == original {
			original = original * 2
		}
	}
	return original
}

/*
 ok, we need find original number, in nums array.
 and multiply this number by 2.
 then find new number un nums then again, multiply.
 1. find new original number.
 2. multiply by 2.
 3. continue this process until we coudn't find this number.
 4. return final original number after process ends.

So how we can find the num in arr.
Bf:
  Take original number.
  Search linear in loop original number in nums.
  if find, multiply, start loop again with new original number.
  if not find return final original number.
  tc: ~N!. - too slow (need clarify)
  sc: 1.
How can we find number effectively?
   1. sort the array, and do binary search.
 but with new number we need to sort again, so
 it will be same TC with linear approach.
    tc: (NlongN)! - similar with linear.
    sc: 1.
  How can we find number in O(1) time, even it will change?
  We can use HashMap.
  1. create map and add nums to hashmap.
  2. Do process until original in map:
     x = original
  	map[original] = x*2.
	map.delete(x)
  3. after while loop is ends return final original value.
  tc: N + O(1) since unsorted map is O(1) time accessibility.
  sc: O(n) - since, we store N elements in map.
  Ok, lets implement.

  Feedback:
	- After view the solution, I realized theat, I dont fully understand the problem statement.
	- So, i decide delete the perevious element. But it is not necessary at all.
   Second moment is that:
    - Also i dont fully understand requiremeents, so decide us hash set.
	- But it is not necessary, since we only need to check existence of element.
    if we sort once the array, all bigger elements appear after smaller elements.
	 Because if we x * 2 - it is always bigger than x.
	 So we cannot search back, or sort again. We just need to do linear search after sorting once:
	 	- replacing original with x * 2.
		- checking is element equal to new original or not.
	This gives us NLogN + N  Time complexity. But we will win in space complexity. Wich is O(1).
	Outcome:
	 - Read the description more carefully.
	 - Test some custom cases to understand the problem better.
*/
