package main

import "fmt"

/*
https://atcoder.jp/contests/abc085/tasks/abc085_b

go run 07.go
4 3 8,10,8,6
*/
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	nums := map[int]int{}
	for i := 0; i < n; i++ {
		var tmp int
		_, _ = fmt.Scan(&tmp)
		nums[tmp] = 1
	}
	fmt.Println(len(nums))
}
