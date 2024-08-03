package main

func main() {
	minimumMoves("OXOX")
}

func minimumMoves(s string) int {
	c := []byte(s)
	res, i := 0, 0
	for i < len(c) {
		if c[i] == 'X' {
			res++
			i = i + 2
		}
		i++
	}
	return res
}
