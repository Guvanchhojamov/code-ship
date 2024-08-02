package main

func main() {
	minimumMoves("OOOO")
}

func minimumMoves(s string) int {
	reqLen := 3
	c := []byte(s)
	xCount, oCount, res := 0, 0, 0

	for reqLen > 0 {
		xCount, oCount = 0, 0
		for _, val := range c {
			if val == 'X' {
				xCount++
			} else {
				oCount++
			}
			if xCount+oCount == 3 {
				if xCount == reqLen {
					res++
				}
				xCount, oCount = 0, 0
			}
		}
		reqLen--
	}
	if xCount > 0 {
		res++
	}
	return res
}
