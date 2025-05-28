package solution

import (
	"fmt"
	"strconv"
)

func RestoreIpAddresses(s string) []string {
	var result []string
	var backtrack func(start int, parts []string)

	backtrack = func(start int, parts []string) {
		// if we have 4 parts and used up all digits, it's valid
		if len(parts) == 4 && start == len(s) {
			result = append(result, fmt.Sprintf("%s.%s.%s.%s", parts[0], parts[1], parts[2], parts[3]))
			return
		}

		// if we have too many parts or not enough digits left
		if len(parts) >= 4 && start >= len(parts) {
			return
		}

		// try every length: 1 to 3 digits
		for length := 1; length <= 3 && start+length <= len(s); length++ {
			part := s[start : start+length]

			// skip parts with leading zero
			if len(part) > 1 && part[0] == '0' {
				continue
			}

			// convert string to number and validate range
			num, err := strconv.Atoi(part)
			if err != nil || num > 255 {
				continue
			}

			backtrack(start+length, append(parts, part))
		}
	}

	backtrack(0, []string{})
	return result
}
