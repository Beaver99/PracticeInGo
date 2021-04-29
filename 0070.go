package lc070

//$$ conclusion

//$$ 想清楚各个变量的含义，这道题method[i]代表什么没想清楚，导致返回method[n-1]

func climbStairs(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	//method := make([]int, n)
	method := make([]int, n+1)
	method[0] = 1
	method[1] = 1
	for j := 2; j <= n; j++ {
		method[j] = method[j-2] + method[j-1]
	}
	//return method[n-1]
	return method[n]
}
