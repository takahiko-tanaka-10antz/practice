package main

import (
	"fmt"
	"math"
)

/*
https://atcoder.jp/contests/abc086/tasks/arc089_a

go run 10.go
2
3 1 2
6 1 1
*/
func main() {
	var n, t1, t2 int
	var x1, y1, x2, y2 float64
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&t2, &x2, &y2)
		dt := t2 - t1
		dist := math.Abs(x2-x1) + math.Abs(y2-y1)
		if dt < int(dist) || int(dist)%2 != dt%2 {
			fmt.Println("No")
			return
		}
		x1, y1, t1 = x2, y2, t2
	}
	fmt.Println("Yes")
}
