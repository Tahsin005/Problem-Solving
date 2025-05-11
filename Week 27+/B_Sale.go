package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/34/B

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&prices[i])
	}

	sort.Ints(prices)

	maxProfit := 0
	for i := 0; i < m && i < n; i++ {
		if prices[i] <= 0 {
			maxProfit -= prices[i]
		} else {
			break
		}
	}

	fmt.Println(maxProfit)
}