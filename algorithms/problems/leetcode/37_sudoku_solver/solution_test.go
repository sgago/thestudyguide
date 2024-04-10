package sudokusolver

import "testing"

/*
	37. Sudoku Solver
	https://leetcode.com/problems/sudoku-solver/

	Solve a sudoku. All rows, columns, and 3x3 sub-grid have all values 1 to 9
	without any duplicate numbers appearing. Fun.

	The sudoku grids are 9x9. Going to call the sub-grids just subs or sub.
	Individual boxes will be called cells or elements. Rows are rows in the matrix;
	Cols are columns in the matrix/grid.

	Sum of 1 to 9 = N*(N+1)/2 = 9*10/2 = 45. Again, all rows, cols, and sub-boxes
	must sum to 45. A sum could be misleading though it doesn't have to be.
	If we sum to 42 and know we've used 8 numbers, then the only conclusion is that
	3 remains. Without the total numbers, we're unsure if the answer is 1 + 2 or 3.

	All boards must have exactly one solution. This helps.

	So, what tools do we have to solve this problem?

	A human typically (somewhat) computes the remaining
	permissible values for each cell. Once only one value remains
	for that cell, we fill it in with the known value.

	This part assumes that we can gracefully solve all cell values
	with 1 remaining, then 2, then 3. Like, the puzzle will, at no point,
	present us with a situation where we simply have to guess a value.
	I'm unsure if we can make this claim.

	We could "solve" the cells with only one possible value,
	reassess the remaining, repeat again, and so on, until the board
	is solved.

	We could do something similar, creating a memo for each cell
	1 to 9 and if it's a solved cell.

	So, again, we can sort of brute force this by tracking all the possible
	values in a cell, finding cells with only a singular value, and repeating
	until solved.

	I guess we could also DFS out a solution, just start guessing, memo'ing,
	backtracking, etc.

	Do we have an optimal substructure + overlapping subproblems? It sort of
	*feels* like it.
*/

const (
	rows = 9
	cols = 9
)

func TestSudokuSolver(t *testing.T) {

}
