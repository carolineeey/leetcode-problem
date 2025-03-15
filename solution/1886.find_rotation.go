package solution

func FindRotation(mat [][]int, target [][]int) bool {
	n := len(mat)

	// Attempts up to four rotations (0°, 90°, 180°, 270°)
	// After each rotation, it checks if mat matches target
	// If a match is found, it returns true; otherwise, it proceeds to the next rotation
	// If no match is found after all rotations, it returns false
	for k := 0; k < 4; k++ {
		match := true
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if mat[i][j] != target[i][j] {
					match = false
					break
				}
			}

			if !match {
				break
			}
		}

		if match {
			return true
		}

		// transpose matrix
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
			}
		}

		// reverse each row
		for i := 0; i < n; i++ {
			for j := 0; j < n/2; j++ {
				mat[i][j], mat[i][n-j-1] = mat[i][n-j-1], mat[i][j]
			}
		}
	}

	return false
}
