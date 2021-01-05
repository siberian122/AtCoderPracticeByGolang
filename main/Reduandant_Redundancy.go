package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for a > 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	ans := 1
	for i := 2; i <= n; i++ {
		ans = ans * i / gcd(ans, i)
	}
	fmt.Println(ans + 1)
}
