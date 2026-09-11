package main

import (
	"fmt"
	"sort"
)

func main () {
	var tt int
	fmt.Scan(&tt)
	for ; tt > 0; tt-- {
		var n, c int
		fmt.Scan(&n, &c)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}

		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})

		ans := 0
		for _, el := range arr {
			if el * (1<<ans) <= c {
				ans++
			}
		}
		fmt.Println(n - ans)
	}
}