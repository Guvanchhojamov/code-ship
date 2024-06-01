package main

func main() {
	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
}
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[r], s[l] = s[l], s[r]
		l++
		r--
	}
}
