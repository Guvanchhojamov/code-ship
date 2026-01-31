package dp

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)

	table := make([][]bool, n+1)
	for i := range table {
		table[i] = make([]bool, m+1)
	}
	table[0][0] = true
	for j := 1; j <= m; j++ {
		if p[j-1] == '*' {
			table[0][j] = table[0][j-1]
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				table[i][j] = table[i-1][j-1]
			} else if p[j-1] == '*' {
				table[i][j] = table[i-1][j-1] || table[i-1][j] || table[i][j-1]
			}
		}
	}
	// fmt.Println(table)
	return table[n][m]
}

/*.      c.   *     a     *    b
  [true false true false true false]
a [false false true true true false]
a [false false true true true false]
b [false false true false true true]
*/
/*
 ok we have some cases:
    if they equal = true.
    if this is '?' we take [i-1][j-1]
    if this is '*' we take [i-1][j-1] or [i-1][j] or [i][j-1]

what is the base case:
  if one empty false.
  if both empty true.

*/
