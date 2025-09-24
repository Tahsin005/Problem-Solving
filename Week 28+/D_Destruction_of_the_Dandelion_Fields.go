package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		var ans int
		var odd []int
		var number int
		for j := 0; j < n; j++ {
			fmt.Scan(&number)
			if number&1 == 1 {
				odd = append(odd, number)
			} else {
				ans += number
			}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(odd)))
		ans = min(1, len(odd)) * ans
		for j := 0; j < (len(odd) + 1) / 2; j++ {
			ans += odd[j]
		}
		fmt.Println(ans)
	}
}
