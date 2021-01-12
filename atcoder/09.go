package main

import (
	"fmt"
	"strings"
)

/*
https://atcoder.jp/contests/abc049/tasks/arc065_a

go run 09.go
erasedream
*/
func main() {
	var str string
	_, _ = fmt.Scan(&str)
	str = strings.Replace(str, "dream", "A", -1)
	str = strings.Replace(str, "erase", "B", -1)
	str = strings.Replace(str, "Aer", "", -1)
	str = strings.Replace(str, "Br", "", -1)
	str = strings.Replace(str, "A", "", -1)
	str = strings.Replace(str, "B", "", -1)

	switch {
	case str == "":
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
