package main

func main() {
	canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2)
}

func canPlaceFlowers(flowerBed []int, n int) bool {
	place := 0

	for i := 0; i < len(flowerBed); i++ {
		if flowerBed[i] == 0 {
			place++
		}

		if place == 2 {
			n--
			place = 0
		}

		if n <= 0 {
			return true
		}
	}

	return false
}
