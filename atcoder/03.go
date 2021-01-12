package main

import (
	"fmt"
	"math"
)

/*
https://atcoder.jp/contests/abc081/tasks/abc081_b

go run 03.go
3 16,12,24
*/
func visit03(cnt, a int) int {
	if a%2 == 1 {
		return cnt
	}
	return visit03(cnt+1, a/2)
}

func main() {
	var n, a, cnt int
	var min = math.MaxInt64
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&a)
		cnt = visit03(0, a)
		if cnt < min {
			min = cnt
		}
	}
	fmt.Println(min)
}
