package main

import (
	"fmt"
	"sort"
)

/*
https://atcoder.jp/contests/abc088/tasks/abc088_b

go run 06.go
3 2,7,4
*/
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	cards := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&cards[i])
	}
	sort.Ints(cards)

	res := 0
	op := 1
	for i := n - 1; i >= 0; i-- {
		res += cards[i] * op
		op *= -1
	}
	fmt.Println(res)
}
