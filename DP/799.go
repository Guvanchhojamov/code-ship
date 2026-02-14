package dp

func champagneTower(poured int, query_row int, query_glass int) float64 {
	n := query_row
	tower := make([][]float64, n+1)
	for i := range tower {
		tower[i] = make([]float64, n+1)
	}
	tower[0][0] = float64(poured)

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if tower[row][col] > 1.0 {
				halfVolume := (tower[row][col] - 1.0) / 2.0
				tower[row+1][col] += halfVolume
				tower[row+1][col+1] += halfVolume
			}
		}
	}
	// fmt.Println(tower)
	// if query is last and they are > 1 return 1.0
	if tower[query_row][query_glass] > 1.0 {
		return 1.0
	}
	return tower[query_row][query_glass]
}

/*
  Ok we have some glases in pyramid order.
  and we pouring X non negative cups from top.
  and assume 1 unit is glass weight. And if glsss is fuull,
  He puring addition chanpagne to left, right bottom glases.
  We need to find the value in I-th row, J-th columns glass.
  After X puting operations.

input:
 pouring int
 query_row  int
 query_col int
output:
    fullX float32 (0.0000)

*/

/*
 How we can sole this?
 How much is the volume of our glasses?   Its 1.0000 -always.

It looks like, tree problem.
The brute force.
Start from 1..prunes:
 Run BFS

BFS:
 prunex=1
 if curr.Val == 1: // it means our glass is full
 Create newLeftNode
 Create newRightNode
	Call BFS(left, prunex/2) // send half to left.
	Call BFS(right, prunex/2) // send half to the right.
 Else:
  currVal += pruneX

After handling all prunes
Take query_row== as BTreee Level;
Take query_col == Node in query row Level;
Go until query_row level, take query_col’s value
And return this value
TC: poured = N for each poured we walk all BFS nodes. So N^2;
SC: N - for btree.
But hte problem is build correct Btree heer.

Approach-2:
Maybe we can use simpel 2d array;
n=poured count;
N*N array; since we need N glasses to use all the poured.
Then start fill our values
 Ok, looks like the N*N matrix much easier to main tain;
Start filling from (0,0):
For each prouded we take 1 and try to fill curr grass starting form 0,0;
If cuurr is full == 1;
  Arr[i+1][j] = 1/2 - half
  Arr[i][j+1] = 1/2 - half
But the problem how to we know the next glasses not full;
So we need to move to the nex glasses. Until 1 becomes 0;

I think this is similar to recursion problem, and we do same work again and again.
How to we move next? We ned to move diogonaaly from top-left->bottom-right;
Ok, i think its DP.

for the top-down we need to suer recursion form last row and take from paretn and fill current then move to parent;
TC: recursion -> 2^R; memorizaiton; R-memo; R-recursioin;

Can we do bottom-tup tabulation way?
 start(0,0) -> fill all pured -> 100;
 if 100>1; leave 1 in curr glass. take half go
 arr[i+1][j]
 arr[i+1][j+1] and fill them. check them again;
 if > 1 leave 1 adm move;
 reapeat until reach row-reachs end;
 OR prouded value bdecomes 0;
 TC: R*R->we Create 2D array .
 SC: R*R -> we create 2D array;
 The base case all glases are empty so we just crate [R]2D array;
 and strt fill in 2 loops;

*/
