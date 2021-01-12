package main

import "fmt"

/*
https://atcoder.jp/contests/abc086/tasks/abc086_a

go run 01.go
1 21
*/
func main() {
	var aa, bb int
	_, _ = fmt.Scan(&aa, &bb)

	switch {
	case aa*bb%2 == 0:
		fmt.Println("Even")
	default:
		fmt.Println("Odd")
	}
}
