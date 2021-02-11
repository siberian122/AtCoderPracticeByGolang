package main

import (
	"bufio"
	"fmt"
	"os"
	// bufio は　fmt.Scanより高速
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}
	ma := 0
	past := 0
	for i := 0; i < n; i++ {
		ma = max(a[i], ma)
		ans := max(ma*b[i], past)
		fmt.Println(ans)
		past = max(ans, past)
	}

}
