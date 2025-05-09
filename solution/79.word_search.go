package solution

func Exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	var dfs func(i, j, k int) bool
	// k is the position in the word we're currently checking.
	// so, word[k] is the current character we want to match with board[i][j].
	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}

		// check bounds if the current cell doesn't match
		if i < 0 || i >= rows || j < 0 || j >= cols || board[i][j] != word[k] {
			return false
		}

		// mark visited cell (temporarily)
		tmp := board[i][j]
		board[i][j] = '#'

		// explore all 4 directions
		// k+1 as we want to check next character in the word
		found := dfs(i+1, j, k+1) || //go down
			dfs(i-1, j, k+1) || //go up
			dfs(i, j+1, k+1) || //go right
			dfs(i, j-1, k+1) //go left

		// restore original character
		board[i][j] = tmp
		return found
	}

	// starting DFS from every cell
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
