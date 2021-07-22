package lc0150

import "strconv"

func evalRPN(tokens []string) int {
	if len(tokens) != 0 {
		s := make([]int, 0, 128)
		for _, token := range tokens {
			switch token {
			case "+":
				{
					// TODO: DEAL WITH ABNORMAL INPUT
					s[len(s)-2] += s[len(s)-1]
					s = s[:len(s)-1]
				}
			case "-":
				{
					s[len(s)-2] -= s[len(s)-1]
					s = s[:len(s)-1]
				}
			case "*":
				{
					s[len(s)-2] *= s[len(s)-1]
					s = s[:len(s)-1]
				}
			case "/":
				{
					s[len(s)-2] /= s[len(s)-1]
					s = s[:len(s)-1]
				}
			default:
				{
					v, e := strconv.Atoi(token)
					if e == nil {
						s = append(s, v)
					} else {
						panic(e)
					}
				}
			}
		}
		return s[0]
	}
	return 0
}
