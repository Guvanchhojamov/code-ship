package daily

import (
	"strconv"
	"strings"
)

/*
3484. Design Spreadsheet

A spreadsheet is a grid with 26 columns
(labeled from 'A' to 'Z') and a given number of rows.
Each cell in the spreadsheet can hold an integer value between 0 and 10^5.

Implement the Spreadsheet class:

Spreadsheet(int rows) Initializes a spreadsheet with 26 columns (labeled 'A' to 'Z') and the specified number of rows.
	All cells are initially set to 0.
void setCell(String cell, int value)
	Sets the value of the specified cell.
	The cell reference is provided in the format "AX" (e.g., "A1", "B10"),
	where the letter represents the column (from 'A' to 'Z') and the number represents a 1-indexed row.
void resetCell(String cell) Resets the specified cell to 0.
int getValue(String formula) Evaluates a formula of the form "=X+Y", where X and Y are either cell references or non-negative integers,
	and returns the computed sum.

Note: If getValue references a cell that has not been explicitly set using setCell, its value is considered 0.



Example 1:

Input:
["Spreadsheet", "getValue", "setCell", "getValue", "setCell", "getValue", "resetCell", "getValue"]
[[3], ["=5+7"], ["A1", 10], ["=A1+6"], ["B2", 15], ["=A1+B2"], ["A1"], ["=A1+B2"]]

Output:
[null, 12, null, 16, null, 25, null, 15]

Explanation

Spreadsheet spreadsheet = new Spreadsheet(3); // Initializes a spreadsheet with 3 rows and 26 columns
spreadsheet.getValue("=5+7"); // returns 12 (5+7)
spreadsheet.setCell("A1", 10); // sets A1 to 10
spreadsheet.getValue("=A1+6"); // returns 16 (10+6)
spreadsheet.setCell("B2", 15); // sets B2 to 15
spreadsheet.getValue("=A1+B2"); // returns 25 (10+15)
spreadsheet.resetCell("A1"); // resets A1 to 0
spreadsheet.getValue("=A1+B2"); // returns 15 (0+15)


Constraints:

1 <= rows <= 10^3
0 <= value <= 10^5
The formula is always in the format "=X+Y", where X and Y are either valid cell references or non-negative integers with values less than or equal to 10^5.
Each cell reference consists of a capital letter from 'A' to 'Z' followed by a row number between 1 and rows.
At most 10^4 calls will be made in total to setCell, resetCell, and getValue.
*/

/*
 We need create the sheet and do in this operations,
	- setCell(key, value) void
	- getValue("=a+b") sum int.
	- resetCell(key) void
 - we have 26 cols A-Z
 - we have n rows.
 - initial state all cells are 0.
 - if not found in sheet a and b then getValue(a,b) result are 0. ?

 Ok, How we store keys?
	hash? key-string; value-int
init:
	create hashmap - empty.
setCell:
	set key '[A-Z][1-n]' = value.
getValue:
	-format string.
	-if a - in map then get value from map[a].
	-if b - in map got from map[b]
	-if not in map:
		convert to int.
	return a+b
resetCell:
	delete(mp, cellKey)

TC: init O1(or 26N) + set 1+get 1; delete: ~1 = O1 or O(26N) - if we create whoole map with all possible keys. for n length.
sc O(26N) at max we can store 26 cell for each row, we have n rows.
1. try init with empty map. if its wrong.
2. try implement with whole map.
*/

type Spreadsheet struct {
	Cells map[string]int
}

func Constructor2(rows int) Spreadsheet {
	return Spreadsheet{
		Cells: make(map[string]int),
	}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.Cells[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	delete(this.Cells, cell)
}

// func (this *Spreadsheet) GetValue(formula string) int {
// 	if len(formula) < 1 {
// 		return 0
// 	}
// 	f := []byte(formula)
// 	f = f[1:]

// 	var a, b []byte
// 	i := 0
// 	for i < len(f) && f[i] != '+' {
// 		a = append(a, f[i])
// 		i++
// 	}
// 	i++
// 	for i < len(f) {
// 		b = append(b, f[i])
// 		i++
// 	}
// 	as := string(a)
// 	bs := string(b)

// 	fmt.Println(as, bs)

// 	x, y := 0, 0
// 	if val, ok := this.Cells[as]; ok {
// 		x = val
// 	} else {
// 		x, _ = strconv.Atoi(as)
// 	}

// 	if val, ok := this.Cells[bs]; ok {
// 		y = val
// 	} else {
// 		y, _ = strconv.Atoi(bs)
// 	}
// 	fmt.Println(x, y)

// 	return x + y
// }

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */

/*
["Spreadsheet","getValue","setCell","getValue","setCell","getValue","resetCell","getValue"]
[[3],["=5+7"],["A1",10],["=A1+6"],["B2",15],["=A1+B2"],["A1"],["=A1+B2"]]
Expected
[null,12,null,16,null,25,null,15]

*/
/*
 - it works but we need optimze getValue.
*/

func (this *Spreadsheet) getValue(key string) int {
	if val, ok := this.Cells[key]; ok {
		return val
	}
	x, _ := strconv.Atoi(key)
	return x
}

func (this *Spreadsheet) GetValue(formula string) int {
	if len(formula) < 1 {
		return 0
	}
	formula = formula[1:]

	keys := strings.Split(formula, "+")

	if len(keys) < 2 {
		return 0
	}

	return this.getValue(keys[0]) + this.getValue(keys[1])
}
