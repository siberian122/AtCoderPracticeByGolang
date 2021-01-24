package main

import (
	"fmt"
)

func main() {
	var n,int
	var ans int =0
	fmt.Scan(&n)
	var a [n]int
	fmt.Scan(&a)
	for i:=0;i<n;i++ {
		for j:=i;j<n;j++{
			ans=max(min(a[i:j])*(j-i),ans)
		}
	}
	fmt.Println(ans)
}
