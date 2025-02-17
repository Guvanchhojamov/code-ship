package daily

func main() {

}

/*
You are given a string s.

You can perform the following process on s any number of times:

Choose an index i in the string such that there
is at least one character to the left of index i that is equal to s[i],
and at least one character to the right that is also equal to s[i].
	Delete the closest character to the left of index i that is equal to s[i].
	Delete the closest character to the right of index i that is equal to s[i].
Return the minimum length of the final string s that you can achieve.

Example 1:
Input: s = "abaacbcbb"
Output: 5
Explanation:
We do the following operations:

Choose index 2, then remove the characters at indices 0 and 3. The resulting string is s = "bacbcbb".
Choose index 3, then remove the characters at indices 0 and 5. The resulting string is s = "acbcb".

Solution:
   We can genrerate and check all substrings.

    Lets try optimal approach counting chars in string.
  a: 3
  b: 4
  c: 2
    sum :9
 count > 2 we  coint is odd diymek:
         count-dan  2-i ayyryp bilyas diymek, sagdan birini chepden birini.
          ta count > 2 we tak bolan yagdayynda.
 count > 2 we jubut bolanda-da 2-sini ayyryp bilyas.

  diymek jubut bolanda-da, tak bolanda-da count>-den uly bolsa 2-sini ayyryp bilyas.
    her harpy sheydip barlayas hemmesi 2-den az we den bolanson,
result-da galan harplaryn jemi sanyny bereris.
   1 examplde sheyle bolar.
     a:1
     b:2
     c:2
             result = 5 bolar.  we  "acbcb" galar.

   2 example-da "aa"  result =2 bolar 2-den uly hich harp gaytalananok.

   TC:  count hasaplamak uchin: O(n) + her harp uchin sanap ayyryarys jemi harplaryn jemi sany-na den ishlar we yene-de O(n)
       O(n+n) => O(N)
   SC: harplaryn countyny saklamak uchin O(n) hash map ulanyas.


*/

func minimumLength(s string) int {
	var res = 0
	var charsMap = make(map[rune]int, len(s))

	for _, char := range s {
		charsMap[char]++
	}

	for _, count := range charsMap {
		for count > 2 {
			count = count - 2
		}
		res += count
	}
	return res
}
