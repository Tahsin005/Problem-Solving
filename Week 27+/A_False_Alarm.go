package main

import "fmt"

func solve() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	cki := false
	start, end := 0, 0
	for i := 0; i < n; i++ {
		if !cki && a[i] == 1 {
			cki = true
			start = i
			end = i
		} else if a[i] == 1 {
			end = i
		}
	}

	if end + 1 - start <= k {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		t--
		solve()
	}
}
