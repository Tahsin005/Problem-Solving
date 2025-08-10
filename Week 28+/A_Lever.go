package main

import "fmt"

func solve() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}

	var ans int = 1
	for i := 0; i < n; i++ {
		if a[i] > b[i] {
			ans += a[i] - b[i]
		}
	}
	fmt.Println(ans)
}

func main() {
	var t int
	fmt.Scan(&t)
	for range t {
		solve()
	}
}
