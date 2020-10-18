package levenshtein

const increment = 1

func GetLevenshteinDistance(str1, str2 []rune, caseSensitive bool) int {
	rows := len(str2)
	cols := len(str1)

	// If one of the given string is empty, Levenshtein distance is the max length of two string
	if rows == 0 || cols == 0 {
		return max(rows, cols)
	}

	// Minimize space complexity. O(min(str1, str2) instead of O(str1)
	if cols > rows {
		rows, cols = cols, rows
		str1, str2 = str2, str1
	}

	dp := make([]uint16, cols+increment)

	for i := 0; i <= rows; i++ {
		left := uint16(i)

		for j := 1; j <= cols; j++ {
			if i == 0 {
				dp[j], left = uint16(j), dp[j]
			} else {
				diag, top := dp[j-1], dp[j]

				if !isSameChar(str2[i-1], str1[j-1], caseSensitive) {
					diag = min(left, min(diag, top)) + uint16(increment)
				}

				dp[j-1], left = left, diag
			}
		}

		dp[cols] = left
	}

	return int(dp[cols])
}

func isSameChar(c1, c2 rune, caseSensitive bool) bool {
	if c1 == c2 {
		return true
	}

	if caseSensitive {
		return false
	}

	return c1 == c2-32 || c1-32 == c2
}

func min(a, b uint16) uint16 {
	if a <= b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
