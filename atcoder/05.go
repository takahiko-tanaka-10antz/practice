package main

import "fmt"

/*
https://atcoder.jp/contests/abc083/tasks/abc083_b

go run 05.go
20 2 5
*/
func visit05(x int) int {
	if x == 0 {
		return 0
	}
	return visit05(x/10) + x%10
}

func main() {
	var n, a, b int
	_, _ = fmt.Scan(&n, &a, &b)
	res := 0
	for i := 1; i <= n; i++ {
		sum := visit05(i)
		if a <= sum && sum <= b {
			res += i
		}
	}
	fmt.Println(res)
}
