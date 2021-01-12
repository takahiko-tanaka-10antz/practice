package main

import "fmt"

/*
https://atcoder.jp/contests/abc085/tasks/abc085_c

 go run 08.go
20 196000
*/
func main() {
	var n, y int
	_, _ = fmt.Scan(&n, &y)

	for i := 0; i <= n; i++ {
		for j := 0; i+j <= n; j++ {
			if 10000*i+5000*j+1000*(n-i-j) == y {
				fmt.Println(i, j, n-i-j)
				return
			}
		}
	}
	fmt.Println(-1, -1, -1)
}
