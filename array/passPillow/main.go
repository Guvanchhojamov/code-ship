package main

func main() {
	passThePillow(5, 3)
}
func passThePillow(n int, time int) int {
	// 1 forward
	// -1 backward
	var direction = 1
	var person = 1
	var ticker = 0
	for ticker < time {
		if person == n {
			direction = -1
		} else if person == 1 {
			direction = 1
		}

		if direction == 1 {
			person++
		} else {
			person--
		}
		ticker++
	}
	return person
}
