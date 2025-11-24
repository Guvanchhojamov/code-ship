package daily

func countPalindromicSubsequence(s string) int {
	res := map[string]bool{}
	left := map[rune]int{}
	right := map[rune]int{}
	c := []rune(s)
	for _, ch := range c {
		right[ch]++
	}

	for _, ch := range c {
		right[ch]--
		for v := range left {
			if right[v] > 0 {
				res[string(v)+string(ch)] = true
			}
		}
		left[ch]++
	}
	return len(res)

}

/*
 We are given string.
 We need to find all palindrome subsequences, with lentght 3.
 - if subsecuences are same only counted once.
 - need to return only count, not the subsecuences.
 - s is only lowercase english letters.
 - the len of subsec must be 3, so we skip 1 len chars, even they palindrome.

 example:
  abcde: 0
  aabca:
   aaa - 1
   aca - 1
   aba - 1
   aba - 0
   aca - 0
   since we have only 3 unique palindrome, we return 3.
1. Brute force:
    Generate all subsequences using 3 loops.
    1. first char,second char, 3th char.
    2. merge them and check for palindrome:
    3. to keep uniquenes create map[string]boolean.
    4 if palindrome and not in map add to map.
    5. at the end return map length.
        map contain only unique palindromes with len 3.
    for i=0->n:
    for j = i+1 -> n:
    for k = j+1 -> n:
      [do ceck and adding to map]
    tc: n*n*n
    sc: N
Ok, how can we optimze?
   What is the palindrom rule:
    n == n/2 allways.
 but we nned len 3. so.
 aba -> first and last chars must be same always.
 Since we need only count of palindrom numbers.
 how we can use palindrom rule?
    aabca
 if we collect each number count in map:
    a:3
    b:1
    c:1
if we take 1 from each it is not palindrome.
we need to take:
 1 unique and 2 same.
 3 same
only to create palindrom from them.
so,
if freq >= 3 we count {1}
3.
create map of results map[string]int
    to keep unique palindroms.
create left part in map
create right part map.
- keep and iterate middle element, char.
    and after taking, remove middle element from the right
    and check, all left elements in loop.
        is left element has in right map?
    if has add this combination to res arr.
    add midd element to left map.
*/
