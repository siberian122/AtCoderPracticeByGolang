package main

import (
	"fmt"
)

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)
	ans := []int{}

	for i := 1; i <= a+b; i++ {
		if (a%i == 0) && (b%i == 0) {
			ans := append(ans, i)
		}
	}
	fmt.Println(ans[-k])
}
