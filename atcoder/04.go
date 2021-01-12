package main

import "fmt"

/*
https://atcoder.jp/contests/abc087/tasks/abc087_b

go run 04.go
2 2 2 100
*/
func main() {
	var a, b, c, x int
	_, _ = fmt.Scan(&a, &b, &c, &x)
	cnt := 0

	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if i*500+j*100+k*50 == x {
					cnt++
				}
			}
		}
	}

	fmt.Println(cnt)
}
