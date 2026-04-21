package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		
		root := int(math.Sqrt(float64(x)))
		if root*root == x {
			count++
		}
	}

	fmt.Println(count)
}