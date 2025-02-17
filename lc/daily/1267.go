package daily

import "fmt"

// 1267. Count Servers that Communicate

func main() {
	countServers([][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}})
}

func countServers(grid [][]int) int {
	var res = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for l := 0; l < len(grid[i]); l++ {
				for k := 0; k < len(grid[j]); k++ {
					fmt.Print(grid[l][k])
				}
				fmt.Println()
			}
		}
		//	fmt.Println()
	}
	return res
}
