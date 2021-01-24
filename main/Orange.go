package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	var cnt float64 = 0
	var ans int = -1
	for i := 1; i <= n; i++ {
		var v, p int
		fmt.Scan(&v, &p)
		cnt += float64(v) * float64(p/100)
		if cnt > float64(x) {
			ans = i + 1
		}
	}
	fmt.Println(ans)
}
