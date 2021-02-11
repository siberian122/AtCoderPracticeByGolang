package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var a, b int
	fmt.Fscan(in, &a)
	fmt.Fscan(in, &b)
	X := (a + b) / 2
	Y := a - X
	fmt.Println(X, Y)
}
