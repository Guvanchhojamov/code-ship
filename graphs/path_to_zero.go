package graphs

type Queue struct {
	row, col int
}

func updateMatrix(mat [][]int) [][]int {
	var q = []Queue{}
	var ans = make([][]int, len(mat))
	for i := range mat {
		ans[i] = make([]int, len(mat[i]))
		for j := range mat[i] {
			if mat[i][j] == 0 {
				q = append(q, Queue{i, j})
				ans[i][j] = 0
			} else {
				ans[i][j] = -1
			}
		}
	}

	dests := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for _, d := range dests {
			nr := node.row + d[0]
			nc := node.col + d[1]
			if (nr >= 0 && nr < len(mat)) &&
				(nc >= 0 && nc < len(mat[0])) &&
				ans[nr][nc] == -1 {
				ans[nr][nc] = ans[node.row][node.col] + 1
				q = append(q, Queue{nr, nc})
			}
		}
	}
	// Or we do it with steps count also, but it is less optimal.
	// step := 1
	// for len(q) > 0 {
	// 	x := len(q)
	// 	for x > 0 {
	// 		node := q[0]
	// 		q = q[1:]
	// 		x--
	// 		for _, d := range dests {
	// 			nr := node.row + d[0]
	// 			nc := node.col + d[1]
	// 			if (nr >= 0 && nr < len(mat)) &&
	// 				(nc >= 0 && nc < len(mat[0])) &&
	// 				ans[nr][nc] == -1 {
	// 				ans[nr][nc] = step
	// 				q = append(q, Queue{nr, nc})
	// 			}
	// 		}
	// 	}
	// 	step++
	// }
	return ans

}

/*
 ok, since we need the distance,
 assume each step is distance 1.
 and assume disstance is level,
 if we look to image, we need to hande elements on each level right, handle elements on each level for this BFS is most optimal.
  but we go from source to destin. we go reversed,
  from destination to the source...
  since it is more efficient way, for example we need to find.
    nearest houses for each hospital in the city.
    we dont search for each hospital, we go from hosppital searching nearest house.
1. Do BFS for each 0.
2. iterate over matrix, take 0 positions add to queue to handle for BFS.
    if it is not zero set -1 for example.
2. each level == distance. if level increased it means distance increased.
    if neighbour is not processed -1. then set to neighbour val to
        curr.val +1, because it is in the next level.
3. at the end return result filled array.
TC: O(n*m)- at worht case we do BFS n*m-1 time.
SC: O(n*m)+(n*M) - response array and queue.
Is there can be egde cases?

*/
