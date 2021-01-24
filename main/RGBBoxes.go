package main

import "fmt"

func main() {
	var r, g, b, n int
	fmt.Scan(&r, &g, &b, &n)
	var ans int = 0

	for i := 0; i <= n; i += r {
		for j := 0; j <= n; j += g {
			if (n-i-j)%b == 0 && n-i-j >= 0 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
