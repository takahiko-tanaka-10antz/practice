package main

import (
	"fmt"
	"strings"
)

/*
https://atcoder.jp/contests/abc081/tasks/abc081_a

go run 02.go
1011
*/
func main() {
	var a string
	_, _ = fmt.Scan(&a)
	fmt.Println(strings.Count(a, "1"))
}
