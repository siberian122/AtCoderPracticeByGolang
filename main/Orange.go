package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var ans int =0
	fmt.Scan(&n)
	a := make([]int, n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	for i:=0;i<n;i++ {
		cnt := a[i]
		for j:=i;j<n;j++{
			cnt=int(math.Min(float64(cnt),float64(a[j])))
			ans=int(math.Max(float64(cnt*(j-i+1)),float64(ans)))
		}
	}
	fmt.Println(ans)
}
