package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}
	var ma = 0
	for i := 0; i < n; i++ {
		ma = max(a[i], ma)
		fmt.Println(ma * b[i])
	}
}
