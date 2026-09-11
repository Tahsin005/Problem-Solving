package main

import "fmt"

func main () {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var curSum, sum, index int
	for i := 0; i < k; i++ {
		curSum += a[i]
	}

	sum = curSum
	index = 1
	for i := k; i < n; i++ {
		curSum += a[i] - a[i - k]
		if curSum < sum {
			sum = curSum
			index = i - k + 2 // +2 because we want 1-based index
		}
	}

	fmt.Println(index)
}