package main

import (
	"fmt"
	"math"
)

func main () {
	var n int
	fmt.Scan(&n)
	var beg, end string
	fmt.Scan(&beg, &end)

	ans := 0
	for i, _ := range beg {
		diff := int(math.Abs(float64(int(beg[i]) - int(end[i]))))
		if diff < 5 {
			ans += diff
		} else {
			ans += 10 - diff
		}
	}
	fmt.Println(ans)
}