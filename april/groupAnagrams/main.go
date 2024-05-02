package main

import "fmt"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// ["bat"],["nat","tan"],["ate","eat","tea"]]
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) (result [][]string) {
	anagrams := map[uint64][]string{}

	for _, str := range strs {
		if _, ok := anagrams[getUniqKey(str)]; !ok {
			anagrams[getUniqKey(str)] = append(anagrams[getUniqKey(str)], str)
		} else {
			anagrams[getUniqKey(str)] = append(anagrams[getUniqKey(str)], str)
		}

	}
	for _, v := range anagrams {
		result = append(result, v)
	}
	return result
}

func getUniqKey(s string) (uniqKey uint64) {
	var sSum uint64
	var sMultiplier = uint64(1)
	for i := 0; i < len(s); i++ {
		sMultiplier *= uint64(s[i])
		sSum += uint64(s[i])
	}
	return sMultiplier + sSum
}
