package solution

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int

	for _, t := range tokens {
		switch t {
		case "+", "-", "*", "/":
			second := stack[len(stack)-1]
			first := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var res int
			switch t {
			case "+":
				res = first + second
			case "-":
				res = first - second
			case "*":
				res = first * second
			case "/":
				res = first / second
			}

			stack = append(stack, res)
		default:
			num, _ := strconv.Atoi(t)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
