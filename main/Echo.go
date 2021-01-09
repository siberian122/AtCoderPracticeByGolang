package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n,&s)
	if s[:n/2] == s[n/2:]{
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}