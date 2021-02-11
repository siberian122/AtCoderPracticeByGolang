package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var x int
	fmt.Fscan(in, &x)
	for i := 1; i <= 360; i++ {
		if (i*x)%360 == 0 {
			ans := i
			break
		}
	}
	fmt.Println(ans)
}
