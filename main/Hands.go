package main

import (
	"fmt"
	"math"
)

func min(nums ...int) int {
	if len(nums) == 0 {
		panic("function min() requires at least one argument")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}

func main() {
	var a, b, x, y int
	fmt.Scan(&a, &b, &x, &y)
	var use_x, use_y int
	if a < b {
		use_x = x + 2*x*(b-a)
		use_y = x + y*(b-a)
	} else if a == b {
		use_x = x
		use_y = x + y
	} else {
		use_x = x + 2*x*(a-b-1)
		use_y = x + y*(a-b-1)
	}
	fmt.Println(min(use_x, use_y))
}
