package main

import (
	"fmt"
	"math"
)

func abs(a int) int { return int(math.Abs(float64(a))) }

func main() {
	var n int
	fmt.Scan(&n)
	ans, prev := 0, 1
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		ans += abs(prev-a) + 2
		prev = a
	}
	fmt.Println(ans)
}
