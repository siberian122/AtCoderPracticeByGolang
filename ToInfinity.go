package main

import (
	"fmt"
)

func main() {
	var s string
	var k int
	fmt.Scan(&s, &k)

	for i := 0; i < k; i++ {
		if i >= len(s) {
			break
		}
		
		if string(s[i]) != "1" {
			fmt.Println(string(s[i]))
			return
		}
	}
	fmt.Println(1)
}
