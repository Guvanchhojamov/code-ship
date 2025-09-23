package daily

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	var v1 = make([]string, 0)
	var v2 = make([]string, 0)
	v1 = strings.Split(version1, ".")
	v2 = strings.Split(version2, ".")
	var mn = min(len(v1), len(v2))
	for i := 0; i < mn; i++ {
		if atoi(v1[i]) < atoi(v2[i]) {
			return -1
		}
		if atoi(v1[i]) > atoi(v2[i]) {
			return 1
		}
	}
	for i := mn; i < len(v1); i++ {
		if atoi(v1[i]) > 0 {
			return 1
		}
	}
	for i := mn; i < len(v2); i++ {
		if atoi(v2[i]) > 0 {
			return -1
		}
	}
	return 0
}

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	fmt.Println(v)
	return int(v)
}

/*
 We need compare versions.
- string, seperated by .
- remove leading zeros.
- if v1<v2 return -1
- if v1>v2 return 1
- if equal return 0.
- if version has less revision, every missing revision sets as 0.
range is 1->500

approach. do simple given steps
-seperate string to arr by .
- take each element from v1 and v2 and compare
- convert to int and remove reading zeros. any time v1 less then v2 return -1.
- if v1 > v2 return 1.
- if all is equal then check len
- if v1 remaining elements are >0 return -1
- if v2 remaining elements are >0 return 1.
- otherwise return 0 .
*/
// optimal version parsing on fly.

func compareVersion2(v1, v2 string) int {
	i, j := 0, 0
	for i < len(v1) || j < len(v2) {
		n1 := 0
		for i < len(v1) && v1[i] != '.' {
			n1 = n1*10 + int(v1[i]-'0')
			i++
		}
		n2 := 0
		for j < len(v2) && v2[j] != '.' {
			n2 = n2*10 + int(v2[j]-'0')
			j++
		}
		if n1 < n2 {
			return -1
		}
		if n1 > n2 {
			return 1
		}
		i++ // skip dot
		j++
	}
	return 0
}
