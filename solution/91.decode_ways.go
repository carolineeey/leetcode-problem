package solution

func NumDecoding(s string) int {
	n := len(s)
	// handle edge cases, if the string is empty or starts with '0', it’s not decodable
	if n == 0 || s[0] == '0' {
		return 0
	}

	// dp[i] will hold the number of ways to decode the substring s[:i]
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	// loop from index 2 to n (inclusive), because we already handled length 0 and 1
	for i := 2; i <= n; i++ {
		// get the current digit at i-1, checking if the last single digit forms a valid character
		oneDigit := s[i-1]
		// get the last two-digit substring ending at index i
		twoDigits := s[i-2 : i]

		// if the current digit is not '0', it's valid (e.g., '1' = A, '2' = B...)
		// so add dp[i-1] to dp[i], meaning you can decode up to i-1 and then decode this last digit
		if oneDigit != '0' {
			dp[i] += dp[i-1]
		}

		// if the last two digits form a valid number (from "10" to "26"),
		// you can also treat them as one character, so add dp[i-2] to dp[i]
		if twoDigits >= "10" && twoDigits <= "26" {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}
