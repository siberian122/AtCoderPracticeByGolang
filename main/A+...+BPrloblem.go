package main

import "fmt"

func main() {
	var n,a,b int
	fmt.Scan(&n,&a,&b)

	if a > b{
		fmt.Println(0)
	}else if (n==1 && a !=b) {
		fmt.Println(0)
	}else{
		fmt.Println((b-a)*(n-2)+1)
	}
}