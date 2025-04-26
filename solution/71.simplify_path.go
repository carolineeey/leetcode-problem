package solution

import "strings"

func SimplifyPath(path string) string {
	parts := strings.Split(path, "/")
	stack := []string{}

	for _, part := range parts {
		if part == "" || part == "." { // means current directory — ignore it
			continue
		}

		if part == ".." { // means go up to the parent — pop from stack if possible
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, part)
		}
	}

	return "/" + strings.Join(stack, "/")
}
